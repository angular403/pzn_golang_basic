package main

import "fmt"

func main() {
	//names := [...]string{"andrew", "harun", "gusti", "bimo", "adi", "pipah"}
	//
	//slice1 := names[2:5]
	//fmt.Println(slice1)
	//
	//slice2 := names[:6]
	//fmt.Println(slice2)
	//
	//slice3 := names[3:]
	//fmt.Println(slice3)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]

	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice2)

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "Andrew "
	newSlice[1] = "Bobi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Zilong")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
