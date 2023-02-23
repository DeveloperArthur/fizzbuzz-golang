//TODO: fazer o teste funcionar
package main

import(
	"fmt"
	f "fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T){
	fizzbuzz = f.NewFizzBuzz(16)
	result := fizzbuzz.Execute()

	assertEquals(t, "1", result[0])
	assertEquals(t, "2", result[1])
	assertEquals(t, "Fizz", result[2])
	assertEquals(t, "4", result[3])
	assertEquals(t, "Buzz", result[4])
	assertEquals(t, "Fizz", result[5])
	assertEquals(t, "7", result[6])
	assertEquals(t, "8", result[7])
	assertEquals(t, "Fizz", result[8])
	assertEquals(t, "Buzz", result[9])
	assertEquals(t, "11", result[10])
	assertEquals(t, "Fizz", result[11])
	assertEquals(t, "13", result[12])
	assertEquals(t, "14", result[13])
	assertEquals(t, "FizzBuzz", result[14])
	assertEquals(t, "16", result[15])
}

func TestIsDisivibleBy_shouldReturnTrue(t *testing.T){
	fizzbuzz = f.NewFizzBuzz(0)
	bool result = fizzbuzz.IsDisivibleBy(18, 3)
	assertEqual(result, true)
}

func TestIsDisivibleBy_shouldReturnFalse(t *testing.T){
	fizzbuzz = f.NewFizzBuzz(0)
	bool result = fizzbuzz.IsDisivibleBy(24, 5)
	assertEqual(result, false)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected [%+v] Got: [%+v]", expected, actual)
	}
}