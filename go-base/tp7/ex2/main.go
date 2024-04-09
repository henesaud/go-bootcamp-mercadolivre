package main

import "fmt"

type product struct {
	name     string
	price    float64
	quantity int
}

type user struct {
	name, surname, email, password string
	age                            int
	products                       []product
}

func New(name string, price float64) product {
	return product{name: name, price: price}
}

func AddProduct(u *user, p product, q int) {
	p.quantity = q
	u.products = append(u.products, p)
}

func DeleteProducts(u *user) {
	u.products = []product{}
}

func main() {
	products := []product{
		{"product1", 10.5, 5},
	}
	products = append(products, New("product2", 20.5))

	u := user{name: "Hene", surname: "SS", email: "hene@ml.com", password: "1234", age: 26, products: products}

	AddProduct(&u, New("product3", 15.5), 3)
	fmt.Println(u)
	DeleteProducts(&u)
	fmt.Println(u)
}
