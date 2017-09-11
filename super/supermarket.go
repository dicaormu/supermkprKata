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
	in:=-1
	amt:=-1
	for index, item := range s.Contents {
		if item.Prod.Name==selection.Prod.Name {
			amt = int(math.Min(float64(item.Amt), float64(selection.Amt)))
			amt=item.Amt-amt
			in=index
		}
	}
	fmt.Println(amt)
	if in>-1{
		s.Contents[in].Amt=amt
	}
}

func (ps ProductSelection) computeSelectionPrice() (totalPrice float64, hasDiscount bool) {
	totalPrice = float64(ps.Amt)*ps.Prod.Price - ps.Prod.Discount*float64(ps.Amt)
	if ps.Prod.Discount > 0 {
		hasDiscount = true
	}
	return
}

func (bs Basket) ComputeBasketPrice() (float64, int) {
	basketPrice := float64(0)
	discountedProds := 0
	for _, item := range bs.Contents {
		price, hasDisc := item.computeSelectionPrice()
		basketPrice += price
		if hasDisc {
			discountedProds++
		}
	}
	return basketPrice, discountedProds
}
