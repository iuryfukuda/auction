package app

import (
	"time"
	_ "encoding/json"
)

// Bid is the main abstraction about auction offer
type Bid struct {
	// ClientID is alphanumeric used to reference the client chosed
	ClientID string	`json:"item_id"`

	// Price is value of bid chosed
	Price float64	`json:"price"`
}

// BidData is the abstraction used in body of bid request
type BidData struct {
	Bid
	// ClientID is alphanumeric used to reference id of client
	ClientID string	`json:"client_id"`
}

// BidStorage is the abstraction used to store bid
type BidStorage struct {
	Bid
	// TimeStamp is the date of made bid in unix format
	TimeStamp time.Time	`json:"timestamp"`
}