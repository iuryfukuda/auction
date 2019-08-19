package db

import (
	_ "fmt"

	"github.com/iuryfukuda/auction/models"
)

type bidDBN struct{
	nHit	int
	bidDB	models.BidDB
}

type mItem map[string]bidDBN

type Mem struct {
	mI	mItem
	nBid	int
	nHit	int
}

func NewMem() *Mem {
	return &Mem{make(mItem), 0, 0}
}

func (m *Mem) Save(bd models.BidData) {
	m.nBid++
	item, ok := m.mI[bd.ItemID]

	if !ok {
		m.mI[bd.ItemID] = bidDBN{1, bd.ToBidDB()}
		m.nHit++
		return
	}

	if bd.Price > item.bidDB.Price {
		m.nHit++
		item.nHit++
		item.bidDB = bd.ToBidDB()
		m.mI[bd.ItemID] = item
	}
}

func (m *Mem) Check() models.Stats {
	var bids = make([]models.Item, 0)
	for k, bn := range m.mI {
		bids = append(bids, models.Item{k, bn.nHit, bn.bidDB})
	}
	return models.Stats{
		TotalBids: m.nBid,
		TotalHits: m.nHit,
		Bids: bids,
	}
}
