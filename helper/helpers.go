package helper

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"unicode"
)

func IsInt(s string) bool {
	for _, c := range s {
		if c == '-' {
			continue
		}
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// PrimeNumber checks if a number is a prime number and returns true or false
func IsPrimeNumber(number int) bool {
	if number < 0 {
		number = int(math.Abs(float64(number)))
	}
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// PerfecrNumber checks if a number is a Perfect Number and returns true or false
func IsPerfectNumber(number int) bool {
	var count = 0

	if number < 0 {
		number = int(math.Abs(float64(number)))
	}

	for i := 1; i < number; i++ {
		if number%i == 0 {
			count += i
		}
	}
	if count == number {
		return true
	} else {
		return false
	}
}

// IsArmstrong checks if a number is an Armstrong number
func IsArmstrong(number int) bool {
	if number < 0 {
		number = int(math.Abs(float64(number)))
	}
	var numOfDigit int
	var temp = number
	result := 0

	for temp > 0 {
		temp /= 10
		numOfDigit += 1
	}
	temp = number
	for {
		rem := temp % 10
		result += int(math.Pow(float64(rem), float64(numOfDigit)))
		temp /= 10
		if temp == 0 {
			break
		}
	}
	if result == number {
		return true
	} else {
		return false
	}
}

// SumAllDigit sums up all digit in a number
func SumAllDigit(number int) int {
	if number < 0 {
		number = int(math.Abs(float64(number)))
	}
	var sum, digit int

	for number > 0 {
		digit = number % 10
		sum += digit
		number /= 10
	}
	return sum
}

// GetProperties returns a slice of string containing the properties of a number
func GetProperties(number int) []string {
	var properties []string

	if IsArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	return properties
}

func GetFunFact(number int) (string, error) {
	Url := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	resp, err := http.Get(Url)
	if err != nil {
		return "", err
	}
	var data []byte
	if resp.StatusCode == 200 {
		if data, err = io.ReadAll(resp.Body); err != nil {
			return "", err
		}
	}

	return string(data), nil
}
