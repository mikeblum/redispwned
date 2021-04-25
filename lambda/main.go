package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	config "github.com/mikeblum/redispwned/internal/configs"
	"github.com/mikeblum/redispwned/pkg/crawler"
	"github.com/zmap/zgrab2"
)

type CrawlRequest struct {
	IP   net.IP `json:"ip"`
	Port uint   `json:"port"`
}

type CrawlResponse struct {
	IP      net.IP `json:"ip"`
	Timeout bool   `json:"timeout"`
}

func decodeAPIGatewayRequest(body string, ctx interface{}) error {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return err
	}
	err = json.NewDecoder(strings.NewReader(string(data))).Decode(&ctx)
	return err
}

func HandleLambdaEvent(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := config.NewLog()
	// NOTE - API Gateway request payloads are base64-encoded!!
	var crawlRequest CrawlRequest
	err := decodeAPIGatewayRequest(request.Body, &crawlRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	log.Info(crawlRequest)
	scanner, err := crawler.NewScanner()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	port := crawlRequest.Port
	target := zgrab2.ScanTarget{
		IP:   crawlRequest.IP,
		Port: &port,
	}
	status, info, err := scanner.Scan(target)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	log.Info(status)
	data, err := json.Marshal(info)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	log.Info(string(data))
	crawlResponse := CrawlResponse{
		IP:      crawlRequest.IP,
		Timeout: false,
	}
	apiGatewayResponse, err := json.Marshal(crawlResponse)
	var responseCode int = 200
	if err != nil {
		responseCode = 400
	}
	return events.APIGatewayProxyResponse{Body: string(apiGatewayResponse), StatusCode: responseCode}, err
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
