package main

import (
    "fmt"
//    "math/rand"
//    "time"
)

type Hoge struct{}

func (h Hoge) SeyHoge() {
	fmt.Printf("Hello Hogehoge!\n")
}

func (h Hoge) SeyFuga() {
        fmt.Printf("Hello Fugafuge!\n")
}

func (h Hoge) seyHello(greeting string, user string){
	fmt.Printf(greeting + " " + user + "!\n")
}


func main() {
    hogeInstance := new(Hoge)
    hogeInstance.SeyHoge()
    hogeInstance.seyHello("Hello", "Agepoyo")
}

