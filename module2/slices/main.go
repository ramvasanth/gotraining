package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
	user *User
}

func main() {
	//sliceDeclaration()
	//sliceisReferenceType()
	//sliceIteration()
	//sliceSlicing()
	//sliceAppending()
	//sliceMemoryShare1()
	//sliceMemoryShare2()
	//variadicSlices()
	//copySlice()
	//deleteSliceEntry()
	//sliceBoundCheck([5]int{1, 2, 3, 4, 5}, 4)
	//multiDimensionSlice()
	//s := &([]int{1, 2, 3})
	//fmt.Println("cap", cap(s))
	//sliceWithoutPointer(s)
	//fmt.Println(s)
	//fmt.Println("cap", cap(s))
	//sliceWithPointer(*s)
	//fmt.Printf("%p\n", s)

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
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))
	daysOfWeek = append(daysOfWeek, "new day")
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))
	daysOfWeek = append(daysOfWeek, "new day")
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))
}

func sliceMemoryShare1() {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))

	daysOfWeek1 := append(daysOfWeek, "new day")
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek1, daysOfWeek1, cap(daysOfWeek1), len(daysOfWeek1))
	daysOfWeek1[0] = "new Mon"
	fmt.Println(daysOfWeek[0], daysOfWeek1[0])

	daysOfWeek1 = append(daysOfWeek1, "new day 1", "new day 2", "new day  3", "new day 4", "new day 5", "new day 6", "new day 7", "new day 8")
	//fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek1, daysOfWeek1, cap(daysOfWeek1), len(daysOfWeek1))
}

func sliceMemoryShare2() {
	daysOfWeek := make([]string, 7, 10)
	daysOfWeek[0] = "Mon"
	daysOfWeek[1] = "Tue"
	daysOfWeek[2] = "Wed"
	daysOfWeek[3] = "Thu"
	daysOfWeek[4] = "Fri"
	daysOfWeek[5] = "Sat"
	daysOfWeek[6] = "Sun"
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek, daysOfWeek, cap(daysOfWeek), len(daysOfWeek))
	daysOfWeek1 := append(daysOfWeek, "new day")
	daysOfWeek1[0] = "new Mon"
	fmt.Println(daysOfWeek[0], daysOfWeek1[0])
	fmt.Printf("%T, %v, cap-%d, len-%d\n", daysOfWeek1, daysOfWeek1, cap(daysOfWeek1), len(daysOfWeek1))

	fmt.Println(daysOfWeek1[7])
}
func variadicSlices(i ...int) {
	fmt.Println(i)
}

func callVariadicSlices() {
	b := []int{1, 2, 3, 4}
	variadicSlices(b...)
}

func multiDimensionSlice() {
	/* a slice with 5 rows and 2 columns*/
	var a = [][]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// use make to create two dimension slice
	ms := make([][]int, 10, 10)
	for key := range ms {
		ms[key] = make([]int, 10, 10)
	}

}

//go build -gcflags="-d=ssa/check_bce/debug=1" example1.go
//compiler optimization called Bounds Check Elimination or BCE.
// The idea behind BCE is to give the compiler hints that index-based memory access is guaranteed to be safe and therefore
// the compiler didn’t have to add extra code to check the memory access at runtime. The safe elimination of these integrity checks can help improve performance
func sliceBoundCheck(s []int, i int) {
	_ = s[1] // bounds check
	_ = s[i] //  bounds check
	_ = s[1] //  bounds check eliminated!
	_ = s[0] //  bounds check eliminated!

}

//https://github.com/golang/go/wiki/SliceTricks
func copySlice() {
	var s = make([]int, 3)
	n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
	fmt.Println(n)

	s1 := []int{0, 1, 2}
	n1 := copy(s1, s1[1:]) // n == 2, s == []int{1, 2, 2}
	fmt.Println(n1)

	//special case
	var b = make([]byte, 5)
	copy(b, "Hello, world!") // b == []byte("Hello")
}

func deleteSliceEntry() {
	a := []int{0, 1, 2, 3}
	toDelete := 1
	a = append(a[:toDelete], a[toDelete+1:]...)
	fmt.Println(a)
}

func sliceWithoutPointer(s []int) {
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
}

func sliceWithPointer(s *[]int) {

	*s = append(*s, 1, 2, 3, 4, 5)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func sortingSlice() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
}
