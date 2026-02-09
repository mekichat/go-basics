package main

import (
	"fmt"
	"go-basics/internal/calculator"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	// Skriver ut text med automatiskt ny rad
	fmt.Println("Hello from Go!")

	// utan ny rad
	fmt.Print("Hello ")

	// Manuell newline
	fmt.Print("Go!\n")

	// Formatted String
	fmt.Printf("Jag heter %s och är %d år \n", "Yahya", 30)

	// const declaration
	const name = "Yahya"
	const age int = 30
	const isStudent bool = true
	const heightMeters float64 = 1.30

	fmt.Println("const variables", name, age, isStudent, heightMeters)

	// var declaration
	var name1 string = "Yahya"
	var age1 int = 30
	var isStudent1 bool = true
	var heightMeters1 float64 = 1.30

	fmt.Println("declared with var", name1, age1, isStudent1, heightMeters1)

	// without var declaration
	name2 := "Yahya"
	age2 := 30
	isStudent2 := true
	heightMeters2 := 1.30

	fmt.Println("without var declaration", name2, age2, isStudent2, heightMeters2)

	// declaring of array
	var arr [3]int

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println(arr)
	fmt.Println("längd på array: ", len(arr))

	//array
	names := [2]string{"Yahya", "Sara"}
	ages := [2]int{30, 28}

	fmt.Printf("%s är %d år gammal. \n", names[0], ages[0])
	fmt.Printf("%s är %d år gammal. \n", names[1], ages[1])
	fmt.Print("==================================\n")

	//slice - with no size assigning
	newNames := []string{"Yahya", "Sara"}
	newAges := []int{30, 28}

	newNames = append(newNames, "Ivar")
	newAges = append(newAges, 40)

	for i := 0; i < len(newNames); i++ {
		fmt.Printf("%s är %d år gammal. \n", newNames[i], newAges[i])
	}

	//struct
	u := User{ID: 30, Name: "Mekuria", Age: 38}
	fmt.Println(u)

	print(u)
	u.printContent()

	//map
	agesMap := map[string]int{
		"Alex":  20,
		"Belly": 30,
	}

	//update map
	agesMap["Tommy"] = 45

	fmt.Println(agesMap)

	//map
	schoolMap := map[string]string{
		"Alex":  "Chalmers",
		"Belly": "KTH",
	}

	//update map
	schoolMap["Tommy"] = "STH"

	fmt.Println(schoolMap)

	//slice => len and capacity
	numbers := []int{10, 20, 30}
	fmt.Println(numbers, "len", len(numbers), "cap", cap(numbers))

	// What’s happening here
	// numbers is a slice with 3 elements, Since it’s created with a slice literal,
	// Go allocates just enough capacity for those elements
	// So at this point: len(numbers) = 3 and cap(numbers) = 3

	numbers = append(numbers, 40)
	fmt.Println(numbers, "len", len(numbers), "cap", cap(numbers))

	// You’re appending one more element
	// But the slice’s capacity was already full (cap == len == 3)
	// So Go allocates a new underlying array with a larger capacity and copies the elements over
	// After append: len(numbers) = 4 cap(numbers) = 6 (very commonly doubles,
	// but this growth strategy is implementation-dependent)

	// Key takeaways
	// len = number of elements currently in the slice
	// cap = how many elements the slice can hold before reallocating
	// append: reuses the same array if len < cap, allocates a new array if len == cap

	// If you want to control capacity upfront (for performance), you can do:
	// numbers := make([]int, 0, 10)

	// How capacity grows over time
	// Let’s watch Go grow a slice:
	// Typical output (approx) len= 1 cap= 1 len= 2 cap= 2 len= 3 cap= 4
	// len= 4 cap= 4 len= 5 cap= 8 len= 6 cap= 8 len= 9 cap=16 len=17 cap=32

	s := []int{}

	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Printf("len=%2d cap=%2d\n", len(s), cap(s))
	}

	// addition and division
	sum := add(2, 3)
	fmt.Println("sum of 2 and 3 is : ", sum)

	a, b := divide(10, 3)
	fmt.Println("Example 10 / 3: ", a, " Example 10 % 3: ", b)

	// if condition
	score := 71

	if x := score - 70; x > 0 { // declaration and condition on the same line
		fmt.Println("Hello", x)
	} else {
		fmt.Println("Test")
	}

	// Infinite loop without breakinside for loop
	for {
		fmt.Println("Låt det loopa en gång bara")
		break
	}

	n := 0

	for n < 3 {
		fmt.Println("Hello with Go !!")
		fmt.Println("n =", n)
		n++
	}

	//switch case
	dag := "Mon"

	switch dag {
	case "Mon":
		fmt.Println("Det är Möndag")
	case "Tue":
		fmt.Println("Det är Tisdag")
	default:
		fmt.Println("Okänd dag")
	}

	temp := 18

	switch {
	case temp < 0:
		fmt.Println("Det är bara Kallt")
	case temp < 20:
		fmt.Println("Det är Svalt")
	default:
		fmt.Println("Det är Varmt")

	}

	// using external function defined outside main.go file
	// Name of the function should start with capital letter
	fmt.Println("Addition of Add(2, 3) => ", calculator.Add(2, 3))
	fmt.Println("Subtraction of Subtract(3, 1) => ", calculator.Subtract(3, 1))

}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func print(u User) {
	fmt.Println("======= Using print =============")
	fmt.Printf("Name: %s ID: %d Age: %d", u.Name, u.ID, u.Age)
	fmt.Println("=================================")
}

func (u User) printContent() {
	fmt.Println("======= Using printContent =============")
	fmt.Println("User: ", u.Name, u.ID, u.Age)
	fmt.Println("=================================")
}

/*
func main() {

//
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
