package main

import "fmt"

func main() {

	num := 6

	if num %2 > 0{
		fmt.Println("ini bilangan ganjil")
		defer cetak("akan berakhir setelah proccess selesai ")
	}

	cetak("oh tentu tidak, defer akan dieksekusi setelah ini")
	
}

func cetak(text string){
	fmt.Println(text)
}

