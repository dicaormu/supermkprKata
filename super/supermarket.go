/**
for more about package see: https://tour.golang.org/basics/1
*/
package super

/**
for more about imports see https://tour.golang.org/basics/2
*/
import (
	"math"
	"time"
	"fmt"
)

/**
for more about types see https://tour.golang.org/moretypes/2
*/
type Product struct {
	Price    float64
	Name     string
	Discount float64
}

type ProductSelection struct {
	Amt  int
	Prod Product
}

type Basket struct {
	Contents []ProductSelection
	Day      time.Weekday
}

//Make necessary changes to see contents from the main
type Stock struct {
	Contents []ProductSelection
}

/**
for more about functions see https://tour.golang.org/basics/4
*/
func (s Stock) SellProduct(selection ProductSelection) {
	in := -1
	amt := -1
	for index, item := range s.Contents {
		if item.Prod.Name == selection.Prod.Name {
			amt = int(math.Min(float64(item.Amt), float64(selection.Amt)))
			amt = item.Amt - amt
			in = index
		}
	}
	fmt.Println(amt)
	if in > -1 {
		s.Contents[in].Amt = amt
	}
}

func (ps ProductSelection) computeSelectionPrice() (totalPrice float64, hasDiscount bool) {
	//todo find total price as the price of the product - discount

	//todo find hasDiscount to see if product has discount (flag)
	return
}

func (bs Basket) ComputeBasketPrice() (float64, int) {
	//todo compute basketprice returning the price of the basket and the amout of products with discount
	return 0, 0
}
