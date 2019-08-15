import (
	"enconnding/json"
)

type Bid struct {
	// ItemID is alphanumeric used to reference the item chosed
	ItemID string	`json:"item_id"`

	// Price is value of bid chosed
	Price float64	`json:"price"`

	// ClientID is alphanumeric used to reference id of client
	ClientID string	`json:"client_id"`
}