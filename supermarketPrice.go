package main

import "fmt"
import "supermarketPrice/super"

func calculatePromotions(baskets ... super.Basket) float64 {
	totalPrice := float64(0)
	promo := 0
	for _, basket := range baskets {
		price, inPromo := basket.ComputeBasketPrice()
		totalPrice = totalPrice + price
		promo += inPromo
	}
	switch promo {
	case 0, 1:
	case 2:
		totalPrice = totalPrice - 5
	case 3:
		totalPrice = totalPrice - totalPrice*0.1 - 5
	default:
		totalPrice = totalPrice - 5 - float64(2*promo)
	}
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
