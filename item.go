package main

import (
	"enconnding/json"
)

type Item struct {
	// ID is the id of item
	ID string	`json:"item_id"`

	// Hits is used to count the hits made in item
	Hits int	`json:"hits"`

	// BestBid is used to store the best bid made so far
	BestBid Bid	`json:"best_bid"`
}
