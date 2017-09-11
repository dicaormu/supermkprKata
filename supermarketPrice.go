package main

import (
	"supermarketPrice/super"
	"html"
	"log"
	"net/http"

	"fmt"
)

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

/*This function is just to show how routing works*/
func handleGet() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	app := super.App{}
	app.Initialize()
	app.Run(":8080")
}
