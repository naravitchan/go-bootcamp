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

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Printf("Planet : %v \n", planet)
	fmt.Printf("distance : %v \n", distance)
	fmt.Printf("orbital : %v \n", orbital)
	fmt.Printf("Does %v have life? %v \n", planet, hasLife)

	fmt.Printf("---- cast type safety ----\n")
	fmt.Printf("Planet : %s \n", planet)
	fmt.Printf("distance : %d \n", distance)
	fmt.Printf("orbital : %f \n", orbital)
	fmt.Printf("orbital : %.0f \n", orbital)
	fmt.Printf("orbital : %.10f \n", orbital)
	fmt.Printf("Does %s have life? %t \n", planet, hasLife)
}
