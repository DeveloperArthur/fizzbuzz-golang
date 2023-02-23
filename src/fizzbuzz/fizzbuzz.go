package fizzbuzz

import ("strconv")

type FizzBuzz struct { 
	QuantidadeDeNumeros int 
}

func (f *FizzBuzz) Execute() []string {
	var result = make([]string, f.QuantidadeDeNumeros)
	for i := 0; i < f.QuantidadeDeNumeros; i++ {
		if IsDisivibleBy(i, 3) && IsDisivibleBy(i, 5) {
			result[i] = "FizzBuzz"
		} else if IsDisivibleBy(i, 3) {
			result[i] = "Fizz"
		} else if IsDisivibleBy(i, 5) {
			result[i] = "Buzz"
		}else {
			result[i] = strconv.Itoa(i)
		}
	}
	return result
}

func IsDisivibleBy(number int, div int) bool {
	return number%div==0
}

func NewFizzBuzz(n int) FizzBuzz {
	return FizzBuzz{n}
}