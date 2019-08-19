package models

import (
	_ "encoding/json"
)

type Stats struct {
	// TotalBids is the count of all items stored
	TotalBids int	`json:"total_bids"`

 	// TotalHits is the count of hits made in any item
	TotalHits int	`json:"total_hits"`

 	// Bids is a list of all items stored
	Bids []Item	`json:"item_bid"`
} 
