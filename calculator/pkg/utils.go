package pkg 

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers and handles division by zero
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func Modulus(a, b float64) float64 {
	return float64(int(a) % int(b))
}

func Power(a, b float64) float64 {
	return float64(int(a) ^ int(b))
}

func Factorial(a float64) float64 {
	if a == 0 {
		return 1
	}
	return a * Factorial(a-1)
} 

func IsPrime(a float64) bool {
	if a <= 1 {
		return false
	}
	for i := 2; i <= int(a/2); i++ {
		if int(a)%i == 0 {
			return false
		}
	}
	return true
}

func IsEven(a float64) bool {
	return int(a)%2 == 0
}

func IsOdd(a float64) bool {
	return int(a)%2 != 0
}

func IsPositive(a float64) bool {
	return a > 0
}

func IsNegative(a float64) bool {
	return a < 0
}

func IsZero(a float64) bool {
	return a == 0
}

