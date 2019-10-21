package main

import "fmt"

func main() {
	arraySlicing()
}

func arrayDeclarationn() {
	var x [5]int // An array of 5 integers with var declarationn
	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105

	fmt.Printf("x[0] = %d, x[1] = %d, x[2] = %d\n", x[0], x[1], x[2])
	fmt.Println("x = ", x)

	x = [5]int{2, 4, 6, 8, 10} // short declaration with fixed length { array literal}

	//infer the length of the array
	x1 := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println("x1 = ", x1)

	// array is comparable if length and type is same
	x2 := [...]int{9, 89, 7, 9, 11, 13, 9}
	fmt.Println(x1 == x2)

}

func arrayisValueType() {
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)
}

func arrayIteration() {
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	fmt.Printf("Sum of all the elements in array  %v = %f\n", a, sum)

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}
}
func arraySlicing() {
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("%T, %v, %p \n", daysOfWeek, daysOfWeek, &daysOfWeek)

	sliceDays1 := daysOfWeek[2:]    // all from 2
	sliceDays2 := daysOfWeek[:3]    // 0 to 3 => 3-0 = 3 items
	sliceDays3 := daysOfWeek[2:3]   // 3-2 => 1 item
	sliceDays4 := daysOfWeek[:]     // complete copy to slice
	sliceDays5 := daysOfWeek[0:1:4] //set the capacity of slice
	fmt.Printf("%T, %v,%d, %p\n", sliceDays1, sliceDays1, cap(sliceDays1), sliceDays1)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays2, sliceDays2, cap(sliceDays2), sliceDays2)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays3, sliceDays3, cap(sliceDays3), sliceDays3)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays4, sliceDays4, cap(sliceDays4), sliceDays4)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays5, sliceDays5, cap(sliceDays5), sliceDays5)

	//fmt.Println(daysOfWeek[0:3:])
	//var cArray [3]string
	//cArray = daysOfWeek[:3]
}

func variadicArrays(i ...int) {
	fmt.Println(i)
}

func callVariadicArrays() {
	b := []int{1, 2, 3, 4}
	variadicArrays(b...)
}

func multiDimensionArray() {
	a := [2][2]int{
		{3, 5},
		{7, 9}, // This trailing comma is mandatory
	}
	fmt.Println(a)

	// Just like 1D arrays, you don't need to initialize all the elements in a multi-dimensional array.
	// Un-initialized array elements will be assigned the zero value of the array type.
	b := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}

	fmt.Println(b)
}

//go build -gcflags="-d=ssa/check_bce/debug=1" example1.go
func arrayBoundCheck() {

}
