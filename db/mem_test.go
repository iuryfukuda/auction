package db_test

import (
	"fmt"
	"testing"
	"reflect"

	"github.com/iuryfukuda/auction/db"
	"github.com/iuryfukuda/auction/models"
)

type dbTest struct {
	in	models.BidData
	want	models.Stats
}

var dbTests = []dbTest{
	dbTest{
		in: models.BidData{models.Bid{"1", 1.26}, "1"},
		want: models.Stats{TotalBids:1, TotalHits:1, Bids:[]models.Item{models.Item{ID:"1", Hits:1, BestBid:models.BidDB{Bid:models.Bid{ClientID:"1", Price:1.26}, TimeStamp:0}}}},
	},
	dbTest{
		in: models.BidData{models.Bid{"2", 1.29}, "1"},
		want: models.Stats{TotalBids:2, TotalHits:2, Bids:[]models.Item{models.Item{ID:"1", Hits:2, BestBid:models.BidDB{Bid:models.Bid{ClientID:"2", Price:1.29}, TimeStamp:0}}}},
	},
	dbTest{
		in: models.BidData{models.Bid{"3", 1.21}, "1"},
		want: models.Stats{TotalBids:3, TotalHits:2, Bids:[]models.Item{models.Item{ID:"1", Hits:2, BestBid:models.BidDB{Bid:models.Bid{ClientID:"2", Price:1.29}, TimeStamp:0}}}},
	},
	dbTest{
		in: models.BidData{models.Bid{"4", 1.30}, "1"},
		want: models.Stats{TotalBids:4, TotalHits:3, Bids:[]models.Item{models.Item{ID:"1", Hits:3, BestBid:models.BidDB{Bid:models.Bid{ClientID:"4", Price:1.3}, TimeStamp:0}}}},
	},
	dbTest{
		in: models.BidData{models.Bid{"5", 1.12}, "2"},
		want: models.Stats{TotalBids:5, TotalHits:4, Bids:[]models.Item{models.Item{ID:"1", Hits:3, BestBid:models.BidDB{Bid:models.Bid{ClientID:"4", Price:1.3}, TimeStamp:0}}, models.Item{ID:"2", Hits:1, BestBid:models.BidDB{Bid:models.Bid{ClientID:"5", Price:1.12}, TimeStamp:0}}}},
	},
}


func runDbTest(mem *db.Mem, t dbTest) error {
	mem.Save(t.in)

	got := mem.Check()

	// timestamp to zero for can check
	for i := range got.Bids {
		got.Bids[i].BestBid.TimeStamp = 0
	}

	if !reflect.DeepEqual(got, t.want) {
		return fmt.Errorf("got [%#v], want [%#v]", got, t.want)
	}
	return nil
}

func TestDb(t *testing.T) {
	MemDB := db.NewMem()
	for i, test := range dbTests {
		if err := runDbTest(MemDB, test); err != nil {
			t.Fatalf("[%d]: %s", i, err)
		}
	}
}

func BenchmarkDb(b *testing.B) {
	MemDB := db.NewMem()
	for i, test := range dbTests {
		if err := runDbTest(MemDB, test); err != nil {
			b.Fatalf("[%d]: %s", i, err)
		}
	}
}
