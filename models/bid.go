package models

import (
	"time"
	_ "encoding/json"
)

// Bid is the main abstraction about auction offer
type Bid struct {
	// ClientID is alphanumeric used to reference the client chosed
	ClientID string	`json:"client_id"`

	// Price is value of bid chosed
	Price float64	`json:"price"`
}

// BidData is the abstraction used in body of bid request
type BidData struct {
	Bid
	// ItemID is alphanumeric used to reference id of item chosed
	ItemID string	`json:"item_id"`
}

// ToBidDB is a easy way to add timestamp on data
func (bd *BidData) ToBidDB() BidDB {
	return BidDB{
		bd.Bid,
		time.Now().Unix(),
	}
}

// BidDB is the abstraction used to store bid
type BidDB struct {
	Bid
	// TimeStamp is the date of made bid in unix format
	TimeStamp int64	`json:"timestamp"`
}
