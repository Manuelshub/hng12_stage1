package main

import (
	"encoding/json"

	"github.com/Manuelsub/hng12_stage1/helper"
	"log"
	"net/http"
	"strconv"
)

const Addr = ":8080"

type OnSuccess struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"` //Sum of it's digit
	FunFact    string   `json:"fun_fact"`  // gotten from the number API
}

type OnFailure struct {
	Number string `json:"number"` //The Actual number converted to a string
	Error  bool   `json:"error"`
}

func welcome(w http.ResponseWriter, _ *http.Request) {
	// set necessary headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	message := map[string]string{
		"message": "Welcome to our API",
	}
	json.NewEncoder(w).Encode(message)
	w.WriteHeader(http.StatusOK)
}

func handleNumber(w http.ResponseWriter, r *http.Request) {
	// Set the necessary headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	stringNumber := r.URL.Query().Get("number")
	if stringNumber == "" || !helper.IsInt(stringNumber) {
		w.WriteHeader(http.StatusBadRequest)
		errMessage := OnFailure{
			Number: "alphabet",
			Error:  true,
		}
		json.NewEncoder(w).Encode(errMessage)
		return
	}
	number, _ := strconv.Atoi(stringNumber)

	isPrime := helper.IsPrimeNumber(number)
	isPerfect := helper.IsPerfectNumber(number)
	digitSum := helper.SumAllDigit(number)
	properties := helper.GetProperties(number)
	funFact, err := helper.GetFunFact(number)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMessage := OnFailure{
			Number: "alphabet",
			Error:  true,
		}
		json.NewEncoder(w).Encode(errMessage)
		return
	}

	onSuccess := OnSuccess{
		Number:     number,
		IsPrime:    isPrime,
		IsPerfect:  isPerfect,
		Properties: properties,
		DigitSum:   digitSum,
		FunFact:    funFact,
	}
	json.NewEncoder(w).Encode(onSuccess)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", welcome)
	router.HandleFunc("/api/classify-number", handleNumber)

	loggedRouter := helper.LoggingMiddleware(router)

	log.Println("Listening and Serving on Port", Addr)
	http.ListenAndServe(Addr, loggedRouter)
}
