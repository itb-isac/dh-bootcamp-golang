package mdl_two

import "fmt"

type product struct {
	name  string
	size  string
	value float64
}

func (p *product) storageTax() float64 {
	if p.size == "m" {
		return p.value * 1.03
	} else if p.size == "g" {
		return (p.value * 1.06) + 2500
	}

	return p.value
}

func (p *product) getName() string {
	return p.name
}

type Product interface {
	getName() string
	storageTax() float64
}

type shop struct {
	productList []Product
}

func (s *shop) Total() float64 {
	lastTotal := 0.
	for _, r := range s.productList {
		lastTotal += r.storageTax()
	}
	return lastTotal
}

func (s *shop) Add(p Product) {
	s.productList = append(s.productList, p)
}

func (s *shop) Cart() []string {
	var cartList []string
	for _, r := range s.productList {
		cartList = append(cartList, r.getName())
	}
	return cartList
}

type Ecommerce interface {
	Total() float64
	Add(p Product)
	Cart() []string
}

func newProduct(name string, size string, value float64) Product {
	return &product{name, size, value}
}

func newShop() Ecommerce {
	return &shop{}
}

func Exercise02() {
	fmt.Println("====== EXERCISE 04 ======")

	Shop := newShop()
	Shop.Add(&product{
		name:  "card",
		size:  "p",
		value: 10000,
	})
	Shop.Add(&product{
		name:  "sofa",
		size:  "g",
		value: 10000,
	})
	Shop.Add(&product{
		name:  "freezer",
		size:  "g",
		value: 10000,
	})
	Shop.Add(&product{
		name:  "notebook",
		size:  "m",
		value: 10000,
	})
	pen := newProduct("pen", "p", 10000)
	Shop.Add(pen)
	fmt.Printf("total ordered %.2f\n", Shop.Total())
	fmt.Println(Shop.Cart())
}
