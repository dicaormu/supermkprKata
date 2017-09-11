package super

import "testing"


func initSupermarket() Stock {
	return Stock{contents: []ProductSelection{
		{Prod: Product{Name: "eau", Price: 1}, Amt: 50},
		{Prod: Product{Name: "gatorade", Price: 2}, Amt: 10},
		{Prod: Product{Name: "stickers", Price: 10}, Amt: 50},
		{Prod: Product{Name: "chocolat", Price: 15}, Amt: 10},
		{Prod: Product{Name: "crepes", Price: 1}, Amt: 100},
		{Prod: Product{Name: "gelatine", Price: 3}, Amt: 30},
		{Prod: Product{Name: "piles", Price: 10}, Amt: 10},
		{Prod: Product{Name: "nivea", Price: 20}, Amt: 50}}}
}

func TestSellProduct(t *testing.T) {
	supermarket := initSupermarket()
	supermarket.sellProduct(ProductSelection{Prod: Product{Name: "eau"}, Amt: 5})

	if supermarket.contents[0].Amt != 45 {
		t.Error("Expected 45, got ", supermarket.contents[0].Amt)
	}
}