/**
for more about package see: https://tour.golang.org/basics/1
*/
package super

/**
for more about imports see https://tour.golang.org/basics/2
*/
import (
	"errors"
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
https://gobyexample.com/range
*/
func (s *Stock) sellProduct(selection ProductSelection) {
	//todo iterate over the selection contents and make the stock amount of each product, the good amout after selling this selection
	errors.New("Not implemented")
}
