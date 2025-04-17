package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Andrew",
		"Alamat": "Bekasi",
	}

	fmt.Println(person["name"])
	fmt.Println(person["alamat"])
	fmt.Println(person)
}
