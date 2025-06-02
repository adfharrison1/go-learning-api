package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main () {
	investmentAmount := getUserValues("Investment amount: ")
	expectedReturnRate := getUserValues("Expected return rate: ")
	years := getUserValues(("Years: "))

	futureValue, futureRealValue := calculateFutureRates(investmentAmount, years, expectedReturnRate)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future value adjusted for inflation: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func calculateFutureRates(investmentAmount, years, expectedReturnRate float64) (fv float64, frv float64){
	fv = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	frv = fv / math.Pow((1 + inflationRate / 100 ), years)

	return fv, frv
}

func getUserValues(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}