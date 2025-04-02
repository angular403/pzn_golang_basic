package main

import "fmt"

func main() {
	//var a = 10
	//var b = 20
	//var c = 5
	//var d = 2
	//
	//var result = b/a + c*d
	//
	//result += 50
	//fmt.Println(result)

	//var name1 = "andrew"
	//var name2 = "wiliam"
	//
	//var result bool = name1 != name2
	//fmt.Println(result)

	var nilaiAkhir = 90
	var absensi = 80

	var nilaiLulusAkhir = nilaiAkhir > 80

	var nilaiAbsensi = absensi > 80

	var lulus bool = nilaiLulusAkhir || nilaiAbsensi

	fmt.Println(lulus)
}
