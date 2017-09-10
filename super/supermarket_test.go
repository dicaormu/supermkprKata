package super

import "testing"
import "github.com/stretchr/testify/assert"

func TestComputeSelectionPrice(t *testing.T) {
	product := Product{Price: 10, Name: "lovea nature", Discount: 0}
	productSelection := ProductSelection{Amt: 1, Prod: product}
	price, isDisc := productSelection.computeSelectionPrice()
	if price != 10 {
		t.Error("Expected 10, got ", price)
	}
	if isDisc {
		t.Error("Expected no discount, got ", true)
	}
}

func TestComputeSelectionPriceWithDiscount(t *testing.T) {
	product := Product{Price: 10, Name: "Nivea men", Discount: 5}
	productSelection := ProductSelection{Amt: 1, Prod: product}
	price, isDisc := productSelection.computeSelectionPrice()
	assert.Equal(t, price, float64(5), "Price should be 10")
	assert.Equal(t, isDisc, true, "Price should be with discount")
}

func TestComputeBasketPriceWithoutPromo(t *testing.T) {
	product := Product{Price: 10, Name: "Nivea men", Discount: 0}
	productSelection := ProductSelection{Amt: 1, Prod: product}
	basket := Basket{Contents: []ProductSelection{productSelection}}
	price, withPromo := basket.ComputeBasketPrice()
	assert.Equal(t, price, float64(10), "Price should be 10")
	assert.Equal(t, withPromo, 0, "with promo should be 0")
}

func TestComputeBasketPrice(t *testing.T) {
	product := Product{Price: 10, Name: "Nivea men", Discount: 5}
	productSelection := ProductSelection{Amt: 1, Prod: product}
	product2 := Product{Price: 10, Name: "lovea nature", Discount: 0}
	productSelection2 := ProductSelection{Amt: 1, Prod: product2}
	basket := Basket{Contents: []ProductSelection{productSelection, productSelection2}}
	price, withPromo := basket.ComputeBasketPrice()
	assert.Equal(t, price, float64(15), "Price should be 15")
	assert.Equal(t, withPromo, 1, "with promo should be 1")
}
