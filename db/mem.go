package db

import (
	"encoding/json"

	"github.com/zbioe/auction/models"
)

type BidDBN struct{
	NHit	int
	BidDB	models.BidDB
}

type MItem map[string]BidDBN

type Mem struct {
	MI	MItem
	NBid	int
	NHit	int
}

func NewMem() *Mem {
	return &Mem{make(MItem), 0, 0}
}

func (m *Mem) Save(bd models.BidData) {
	m.NBid++
	item, ok := m.MI[bd.ItemID]

	if !ok {
		m.MI[bd.ItemID] = BidDBN{1, bd.ToBidDB()}
		m.NHit++
		return
	}

	if bd.Price > item.BidDB.Price {
		m.NHit++
		item.NHit++
		item.BidDB = bd.ToBidDB()
		m.MI[bd.ItemID] = item
	}
}

func (m *Mem) Check() models.Stats {
	var bids = make([]models.Item, 0)
	for k, bn := range m.MI {
		bids = append(bids, models.Item{k, bn.NHit, bn.BidDB})
	}
	return models.Stats{
		TotalBids: m.NBid,
		TotalHits: m.NHit,
		Bids: bids,
	}
}

func (m *Mem) ToJSON() ([]byte, error) {
	b, err := json.Marshal(m)
	return b, err
}

func MemFromJSON(b []byte) (*Mem, error) {
	var mem Mem
	err := json.Unmarshal(b, &mem)
	if err != nil {
		return nil, err
	}
	return &mem, nil
}
