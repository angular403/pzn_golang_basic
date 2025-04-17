package main

import "fmt"

func main() {

	//name := "Andrew"
	//
	//if name == "Andrew" {
	//	fmt.Println("Hello Andrew")
	//} else if name == "Joko" {
	//	fmt.Println("Hello Joko")
	//} else {
	//	fmt.Println("Data Tidak Ditemukan")
	//}

	name := "Andrew"

	switch name {

	case "Andrew":
		fmt.Println("Hello Andrew")
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello joko")
	default:
		fmt.Println("Data Tidak Ada")
		break
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Ada")
	}

}
