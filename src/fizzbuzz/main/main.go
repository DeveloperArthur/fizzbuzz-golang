package main

import (
	"fmt"
	f "fizzbuzz"
)

func main(){ 
	fizzbuzz := f.NewFizzBuzz(16)

	list := fizzbuzz.Execute()

	for i := 0; i < fizzbuzz.QuantidadeDeNumeros; i++ {
		fmt.Println(list[i])
	}
}