package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"amazon": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["amazon"])

	websites["linkedin"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "linkedin")

	fmt.Println(websites)

	for key, val := range websites {

		fmt.Println(key, val)
	}

	hobbies := [3]string{"coding", "cooking", "climbing"}

	hobbyMap := make(map[string]string, len(hobbies))

	for idx, val := range hobbies {

		fmt.Println(idx, val)
		hobbyMap[fmt.Sprintf("hobby # %v", idx+1)] = val
	}

	fmt.Println(hobbyMap)

}
