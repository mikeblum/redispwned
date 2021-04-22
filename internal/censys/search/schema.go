package search

import (
	"net"
	"time"
)

// https://censys.io/api/v1/docs/search
type Request struct {
	// The query to be executed. For example, 80.http.get.headers.server: nginx.
	Query string `json:"query"`
	// The page of the result set to be returned.
	// By default, the API will return the first page of results. One indexed.
	Page int `json:"page"`
	// The fields you would like returned in the result set in "dot notation",
	// e.g. location.country_code.
	Fields []string `json:"fields"`
	// Format of the returned results. Default is flattened
	Flatten bool `json:"flatten"`
}

type Response struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
	Meta    Meta     `json:"metadata"`
}

type Result struct {
	IP       net.IP   `json:"ip"`
	Location Location `json:"location"`
	ASN      ASN      `json:"asn"`
	Redis    Redis    `json:"6379"`
}

type Redis struct {
	Service Service `json:"redis"`
}

type Service struct {
	Banner Banner `json:"banner"`
}

type Banner struct {
	PingResponse  string `json:"ping_response"`
	Version       string `json:"version"`
	Mode          string `json:"mode"`
	UptimeSeconds int    `json:"uptime_in_seconds"`
}

type Location struct {
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type ASN struct {
	ASN         int       `json:"asn"`
	Description string    `json:"description"`
	CIDR        net.IPNet `json:"routed_prefix"`
	CountryCode string    `json:"country_code"`
	Name        string    `json:"name"`
}

type Meta struct {
	Count       int           `json:"count"`
	Query       string        `json:"query"`
	BackendTime time.Duration `json:"backend_time"`
	Page        int           `json:"page"`
	NumPages    int           `json:"pages"`
}
