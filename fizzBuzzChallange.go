package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	var max = 300
	rand.Seed(time.Now().UnixNano())
	var loop = rand.Intn(max)
	
	args := []int{}
    	for i := 0; i < loop; i++ {
		args = append(args, i)
		if (i % 15 == 0) {
			fmt.Printf("FizzBuzz\n")
		} else if (i % 3 == 0) {
			fmt.Printf("Fizz\n")
                } else if (i % 5 == 0) {
                        fmt.Printf("Buzz\n")
		} else {
			fmt.Printf(strconv.Itoa(i) + "\n")
		}
    	}
	fmt.Printf("ActuralMax:" + strconv.Itoa(loop) + " (SettingMax" + strconv.Itoa(max) + ")\n")
}

