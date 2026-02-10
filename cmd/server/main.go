package main

import (
	"fmt"
	"sort"
)

func main() {

	// range
	ages := map[string]int{
		"Yahya": 30,
		"Sara":  30,
	}

	keys := make([]string, 0, len(ages))
	for k := range ages{
		keys = append(keys, k)
	}

	sort.Strings(keys)


	for _, k := range keys {
		fmt.Println(k,"->", ages[k])
	}



	/*
		nums := []int{10, 20, 30}

		for i := range nums {
			nums[i] = nums[i] *10
		}
		fmt.Println(nums)
	*/

}

/*
func main() {

	names := []string{"Yahya", "Sara"}
	ages := []int{30, 28}

	names = append(names, "Ivar")

	for i := 0; i < len(names); i++ {
		fmt.Printf("%s är %d år gammal. \n", names[i], ages[i])
	}


		names := [2]string{"Yahya", "Sara"}

		ages := [2]int{30, 28}

		fmt.Printf("%s är %d år gammal. \n", names[0], ages[0])

		fmt.Printf("%s är %d år gammal. \n", names[1], ages[1])


}

*/
/*
import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	u := User{ID: 1, Name: "Yahya", Age: 34}
	fmt.Println(u)


	u.print()
}



func (u User) print() {
	fmt.Println("User:", u.ID, u.Name, u.Age)
}
*/

/*

func main() {

	ages := map[string]int{
		"Alice": 30,
		"Yahya": 30,
	}


	ages["Sara"] = 31

	fmt.Println(ages)
}

*/

/*
func main() {

	numbers := []int{10, 20, 30}
	fmt.Println(numbers, "len:", len(numbers), "cap", cap(numbers))


	numbers = append(numbers, 40)
	fmt.Println(numbers, "len:", len(numbers), "cap", cap(numbers))


}


*/
/*
func main() {

	var arr [3]int

	arr[0] = 10

	arr[1] = 20

	arr[2] = 30

	fmt.Println(arr)

	fmt.Println("längd på array: ", len(arr))

}

*/
/*

func main() {

	sum := add(2, 3)
	fmt.Println(sum)

	a, b := divide(10,3)
	fmt.Println("Example a: ", a, "Example b: ", b)


}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}


*/

/*

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	n := 0
	for n < 3 {
		fmt.Println("n =", n)
		n++
	}


	for {
		fmt.Println("Låt det loopa en gång bara")
		break
	}


}
*/
/*

func main() {


	day := "Mon"

	switch day {
	case "Mon":
		fmt.Println("Måndag")

	case "Tue":
		fmt.Println("Tisdag")

	default:
		fmt.Println("Okänd dag")
	}




	temp := 18
	switch {
	case temp < 0:
		fmt.Println("Brrr kallt")

	case temp < 20:
		fmt.Println("Svalt")

	default:
		fmt.Println("Varmt")
	}

}


*/
/*

func main() {

	score := 71

	if x := score - 70; x > 0 {
		fmt.Println("Hello", x)
	} else {
		fmt.Println("Test")
	}

}

*/
/*
func main(){

	const name = "Yahya"
	const age int = 30


	var name string = "Yahya"
	var age int = 30
	var isStudent bool = true
	var heightMeters float64 = 1.30



	name := "Yahya"
	age := 30
	isStudent := true
	heightMeters := 1.30


	fmt.Println(name, age)

}


*/

/*

func main() {

	// Skriver ut text med automatiskt ny rad
	fmt.Println("Hello from Go!")


	// utan ny rad
	fmt.Print("Hello ")

	// Manuell newline
	fmt.Print("Go!\n")


	// Formatted String
	fmt.Printf("Jag heter %s och är %d år \n", "Yahya", 30)


}

*/