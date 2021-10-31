package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

func AverageAge(ages []uint8) float64 {
	var sum uint64
	for _, age := range ages {
		sum += uint64(age)
	}
	return float64(sum) / float64(len(ages))
}

func AveragePaymentAmount(amounts []uint32) float64 {
	// all cents fit in 64 bit int bc 1e14 < max int64
	var sum uint64
	for _, a := range amounts {
		sum += uint64(a)
	}
	return 0.01 * float64(sum) / float64(len(amounts))
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(amounts []uint32) float64 {
	mean := AveragePaymentAmount(amounts)
	squaredDiffs := 0.0
	for _, a := range amounts {
		amount := 0.01 * float64(a)
		diff := amount - mean
		squaredDiffs += diff * diff
	}
	return math.Sqrt(squaredDiffs / float64(len(amounts)))
}

func LoadData() ([]uint8, []uint32) {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	ages := make([]uint8, len(userLines))

	for i, line := range userLines {
		age, _ := strconv.Atoi(line[2])	
		ages[i] = uint8(age)
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	amounts := make([]uint32, len(paymentLines))

	for i, line := range paymentLines {
		paymentCents, _ := strconv.Atoi(line[0])	
		amounts[i] = uint32(paymentCents)
	}

	return ages, amounts
}
