/**
こちらを参考に色々確認。
https://postd.cc/is-go-object-oriented/
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func main() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gotham",
			State:  "NY",
			Zip:    "01313",
		},
	}

	p.Talk()
	p.Location()

	/**
	コピペは正常に動いたので。。
	 */
	expand()
}

type Product struct {
	Id int
	Type int
	TypeName string
	Name string
}

const (
	ProductTypeCD = 0
	ProductTypeDVD = 1
)

var ProductTypeDisplay = [2]string{"CD", "DVD"}

func expand () {
	fmt.Println(ProductTypeDisplay[0])
	rand.Seed(time.Now().UnixNano())

	var randNum = rand.Intn(10)
	product := Product{
		Id: randNum,
		Name: "女々しくて。",
	}

	if (product.Id % 2 == 1) {
		product.TypeName = ProductTypeDisplay[ProductTypeCD]
	} else {
		product.TypeName = ProductTypeDisplay[ProductTypeDVD]
	}

	fmt.Println("product information:", product.Name, product.TypeName, randNum)
}