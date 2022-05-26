package main

import "fmt"

func main() {
	var brand = "Google"
	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)
	ops, ok, fail := 2350, 543, 433

	fmt.Println("hi\nhi\\h")
	fmt.Println(
		"total:", ops, "- success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)

	var speed int
	var heat float64
	var off bool
	fmt.Printf("speed type : %T \n", speed)
	fmt.Printf("heat type : %T \n", heat)
	fmt.Printf("off type : %T \n", off)
}
