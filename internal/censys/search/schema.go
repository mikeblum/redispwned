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
	Meta    Meta     `json:"meta"`
}

type Result struct {
	IP net.IP `json:"ip"`
}

type Location struct {
	Country   string  `json:"country"`
	Longitude float64 `json:"longitude"`
}

type Meta struct {
	Count       int           `json:"count"`
	Query       string        `json:"query"`
	BackendTime time.Duration `json:"backend_time"`
	Page        int           `json:"page"`
	NumPages    int           `json:"pages"`
}
