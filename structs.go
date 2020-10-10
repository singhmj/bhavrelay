package main

import "time"

// Bhav :
type Bhav struct {
	Symbol      string    `json:"symbol"`
	Series      string    `json:"series"`
	Open        float64   `json:"open"` // we can loose value
	High        float64   `json:"high"`
	Low         float64   `json:"low"`
	Close       float64   `json:"close"`
	Last        float64   `json:"last"`
	PrevClose   float64   `json:"prev_close"`
	TotalRDQTY  int32     `json:"total_rd_qty"`
	TotalRDVal  float64   `json:"total_rd_val"`
	TotalTrades int32     `json:"total_trades"`
	ISIN        string    `json:"isin"`
	Timestamp   time.Time `json:"timestamp"`
}

// SystemConfig :
type SystemConfig struct {
	Dumper DumperConfig `json:"dumper"`
	Bhav   BhavConfig   `json:"bhav"`
}

// DumperConfig :
type DumperConfig struct {
	WebServer WebServerConfig `json:"web-server"`
	Scheduler SchedulerConfig `json:"scheduler"`
}

// WebServerConfig :
type WebServerConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// SchedulerConfig :
type SchedulerConfig struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

// BhavConfig :
type BhavConfig struct {
	URL      string `json:"url"`
	DataType string `json:"data_type"`
}
