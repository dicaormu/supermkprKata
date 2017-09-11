package main

import (
	"supermarketPrice/super"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePromotionsNoPromotion(t *testing.T) {
	product := super.Product{Price: 10, Name: "Nivea men", Discount: 0}
	productSelection := super.ProductSelection{Amt: 1, Prod: product}
	basket := super.Basket{Contents: []super.ProductSelection{productSelection}}
	totalPrice := calculatePromotions(basket)
	assert.Equal(t, totalPrice, float64(10))
}

func TestCalculatePromotionsOnePromotion(t *testing.T) {
	product := super.Product{Price: 10, Name: "Nivea men", Discount: 5}
	productSelection := super.ProductSelection{Amt: 1, Prod: product}
	basket := super.Basket{Contents: []super.ProductSelection{productSelection}}
	totalPrice := calculatePromotions(basket)
	assert.Equal(t, totalPrice, float64(5))
}

func TestCalculatePromotionsTwoPromotion(t *testing.T) {
	product := super.Product{Price: 10, Name: "Nivea men", Discount: 3}
	productSelection := super.ProductSelection{Amt: 1, Prod: product}
	basket := super.Basket{Contents: []super.ProductSelection{productSelection}}
	product2 := super.Product{Price: 5, Name: "lovea natural", Discount: 2}
	productSelection2 := super.ProductSelection{Amt: 1, Prod: product2}
	basket2 := super.Basket{Contents: []super.ProductSelection{productSelection2}}
	totalPrice := calculatePromotions(basket, basket2)
	assert.Equal(t, totalPrice, float64(5))
}

func TestCalculatePromotionsThreePromotion(t *testing.T) {
	product := super.Product{Price: 10, Name: "Nivea men", Discount: 3}
	productSelection := super.ProductSelection{Amt: 1, Prod: product}
	basket := super.Basket{Contents: []super.ProductSelection{productSelection}}
	product2 := super.Product{Price: 5, Name: "lovea natural", Discount: 2}
	productSelection2 := super.ProductSelection{Amt: 1, Prod: product2}
	product3 := super.Product{Price: 15, Name: "t-shirt", Discount: 5}
	productSelection3 := super.ProductSelection{Amt: 1, Prod: product3}
	basket2 := super.Basket{Contents: []super.ProductSelection{productSelection2, productSelection3}}
	totalPrice := calculatePromotions(basket, basket2)
	assert.Equal(t, totalPrice, float64(13))
}

func TestCalculatePromotionsFourPromotion(t *testing.T) {
	product := super.Product{Price: 10, Name: "Nivea men", Discount: 3}
	productSelection := super.ProductSelection{Amt: 1, Prod: product}
	basket := super.Basket{Contents: []super.ProductSelection{productSelection}}
	product2 := super.Product{Price: 5, Name: "lovea natural", Discount: 2}
	productSelection2 := super.ProductSelection{Amt: 1, Prod: product2}
	product3 := super.Product{Price: 15, Name: "Red t-shirt", Discount: 5}
	productSelection3 := super.ProductSelection{Amt: 1, Prod: product3}
	product4 := super.Product{Price: 20, Name: "lindt nougets", Discount: 5}
	productSelection4 := super.ProductSelection{Amt: 1, Prod: product4}
	basket2 := super.Basket{Contents: []super.ProductSelection{productSelection2, productSelection3, productSelection4}}
	totalPrice := calculatePromotions(basket, basket2)
	assert.Equal(t, totalPrice, float64(22))
}
