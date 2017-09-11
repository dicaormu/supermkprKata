/**
for more about package see: https://tour.golang.org/basics/1
*/
package super

/**
for more about imports see https://tour.golang.org/basics/2
*/
import (
	"math"
	"fmt"
)

/**
for more about types see https://tour.golang.org/moretypes/2
*/
type Product struct {
	Price float64
	Name  string
}

type ProductSelection struct {
	Amt  int
	Prod Product
}

type Basket struct {
	contents []ProductSelection
}

//Make necessary changes to see contents from the main
type Stock struct {
	contents []ProductSelection
}

/**
for more about functions see https://tour.golang.org/basics/4
for ranges https://gobyexample.com/range
*/
func (s Stock) sellProduct(selection ProductSelection) {
	in := -1
	amt := -1
	for index, item := range s.contents {
		if item.Prod.Name == selection.Prod.Name {
			amt = int(math.Min(float64(item.Amt), float64(selection.Amt)))
			amt = item.Amt - amt
			in = index
		}
	}
	fmt.Println(amt)
	if in > -1 {
		s.contents[in].Amt = amt
	}

}
