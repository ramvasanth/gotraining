package main

import "fmt"

type User struct {
	name string
	age  int
	user *User
}

func main() {
	//sliceDeclaration()
	//sliceisReferenceType()
	sliceIteration()
	//sliceSlicing()
	//sliceAppending()
	//variadicSlice()
	//sliceBoundCheck([5]int{1, 2, 3, 4, 5}, 4)
	//multiDimensionSlice()
}

func sliceDeclaration() {
	var x []int // A slice of int and its nil
	//x[0] = 100 will throw error
	fmt.Println(x == nil)

	x = []int{2, 4, 6, 8, 10} // short declaration with fixed length { slice literal}

	// slice is not comparable but assignable if slice value is same
	x1 := []int{9, 89, 7, 9, 11, 13, 9}
	x2 := []int{9, 89, 7, 9, 11, 13, 9}
	x1 = x2
	fmt.Println(x1)
	//	fmt.Println(x1 == x2)

	s := make([]int, 10, 20) // length is 10 and capacity is 20 annd its not nil
	fmt.Printf("%T, %v,%d, %d\n", s, s, cap(s), len(s))

	var s1 []int // it is nil
	fmt.Println(s1 == nil)
	fmt.Printf("%T, %v,%d, %d\n", s1, s1, cap(s1), len(s1))

	sP := new([]int) //  nil but not nil
	fmt.Println(*sP == nil)
	fmt.Println(sP == nil)

	fmt.Printf("%T, %v,%d, %d\n", sP, sP, cap(*sP), len(*sP))

}

func sliceisReferenceType() {
	a1 := []string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // a1 and a2 shares the same memory

	a2[0] = "English 1"

	fmt.Println(a1[0])
}

func sliceIteration() {
	a := []float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	fmt.Printf("Sum of all the elements in slice  %v = %f\n", a, sum)

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}
}

func sliceSlicing() {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("%T, %v, %p \n", daysOfWeek, daysOfWeek, &daysOfWeek)

	sliceDays1 := daysOfWeek[2:]    // all from 2
	sliceDays2 := daysOfWeek[:3]    // 0 to 3 => 3-0 = 3 items
	sliceDays3 := daysOfWeek[2:3]   // 3-2 => 1 item
	sliceDays4 := daysOfWeek[:]     // complete copy
	sliceDays5 := daysOfWeek[0:1:4] //set the capacity of slice
	fmt.Printf("%T, %v,%d, %p\n", sliceDays1, sliceDays1, cap(sliceDays1), sliceDays1)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays2, sliceDays2, cap(sliceDays2), sliceDays2)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays3, sliceDays3, cap(sliceDays3), sliceDays3)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays4, sliceDays4, cap(sliceDays4), sliceDays4)
	fmt.Printf("%T, %v,%d, %p\n", sliceDays5, sliceDays5, cap(sliceDays5), sliceDays5)
}

func sliceAppending() {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("%T, %v,%d, %p\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), daysOfWeek)
	daysOfWeek = append(daysOfWeek, "new day")

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
//compiler optimization called Bounds Check Elimination or BCE.
// The idea behind BCE is to give the compiler hints that index-based memory access is guaranteed to be safe and therefore
// the compiler didn’t have to add extra code to check the memory access at runtime. The safe elimination of these integrity checks can help improve performance
func arrayBoundCheck(s [5]int, i int) {
	_ = s[1] // bounds check
	_ = s[i] //  bounds check
	_ = s[1] //  bounds check eliminated!
	_ = s[0] //  bounds check eliminated!

}
