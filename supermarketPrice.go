package main

import "fmt"
import "supermarketPrice/super"


/*
for information about switch:
https://tour.golang.org/flowcontrol/9
*/
func calculatePromotions(baskets ... super.Basket) float64 {
	totalPrice := float64(0)
	//todo compute basket price for all  baskets

	//todo compute value of additional promos
	return totalPrice
}

//you can init an array for example: b := [2]string{"Penn", "Teller"} or b := [...]string{"Penn", "Teller"}
//or you can init an slice, for example: letters := []string{"a", "b", "c", "d"} or s = make([]byte, 5, 5)
func main() {
	fmt.Println("--- Supermarket Price ---")
	product := super.Product{Price: 10, Name: "lovea nature", Discount: 0}
	productSelection := super.ProductSelection{Amt: 10, Prod: product}
	stock := super.Stock{Contents: []super.ProductSelection{productSelection}}
	stock.SellProduct(productSelection)
}
