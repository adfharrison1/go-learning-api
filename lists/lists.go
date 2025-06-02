package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func main() {
	hobbies := [3]string{"coding", "cooking", "climbing"}

	fmt.Println("hobbies: ", hobbies)
	fmt.Println("hobbies[0]: ", hobbies[0])
	fmt.Println("hobbies[1 + 2]: ", hobbies[1:])

	firstHobbies := hobbies[:2]
	fmt.Println("firstHobbies: ", firstHobbies)

	firstHobbies = firstHobbies[1:3]
	fmt.Println("firstHobbies: ", firstHobbies)

	goals := []string{"learn GO", "make an API"}
	fmt.Println("goals: ", goals)

	goals[1] = "make a great API"
	fmt.Println("goals: ", goals)

	goals = append(goals, "go forth and code")
	fmt.Println("goals: ", goals)

	products := []Product{
		{
			name:  "product 1",
			price: 1.50,
		},
		{
			name:  "product 2",
			price: 2.50,
		},
	}

	fmt.Println("products: ", products)

	products = append(products, Product{
		name:  "product 3",
		price: 20.55,
	})

	fmt.Println("products: ", products)

	moreProducts := []Product{
		{
			name:  "product 4",
			price: 1.50,
		},
		{
			name:  "product 5",
			price: 2.50,
		},
	}

	products = append(products, moreProducts...)

	fmt.Println("products: ", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
