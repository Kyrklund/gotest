package main

import "fmt"
import "time"

func main() {
	//Outputs
	fmt.Printf("Hello. The number you supplied is is not divisible by 2\n")
	number := 42
	result := fmt.Sprintf("Woot %d +2 =", number)
	fmt.Println(result)
	fmt.Println("42 modulus =", number%2)

	for j := 7; j <= 9; j++ {
		if j%2 == 0 {
			fmt.Print("The number ", j, " is divisible by 2")
			continue
		}
        fmt.Println(j)
	}
	
	//Working With Switch
	fmt.Println("Today is", time.Now().Weekday())
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend!!!")
    default:
        fmt.Println("It's a weekday :(")
	}
	
	//Arrays
	var a [5]int
	a[3]=50
	fmt.Println(a, "Length of array is", len(a))

	//Slices
	s := make([]string, 3)
	s = append(s, "P")
	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(s,"With it's copy", c)
	
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	//Maps
	m := make(map[string]int)
	m["prod"] = 2
	m["dev"] = 4

	fmt.Println("Printing value for prod:", m["prod"])

}