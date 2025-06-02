package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	adultYears := getAdultYears(agePointer)
	fmt.Println("age:", age)
	fmt.Println("adultYears:", adultYears)

}

func getAdultYears(age *int) int {
	return *age - 18
}
