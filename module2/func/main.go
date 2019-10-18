package main

import (
	"errors"
	"fmt"
	"runtime"
)

var users []user

type user struct {
	name string
	age  int
	id   string
}

func init() {
	users = []user{
		user{name: "john", age: 20, id: "1"},
		user{name: "marshal", age: 20, id: "2"},
	}
}

func main() {
	growTheFunc(100000)
}

/*  1. Function Parameters  */
func withArgs1(name string) (user, error) {
	for _, user := range users {
		if user.name == name {
			return user, nil
		}
	}

	return user{}, errors.New("user is not found")
}

//error - should be the last values to be returned
func withMultipleArgs2(name string, age int) (user, error) {
	for _, user := range users {
		if user.name == name && user.age == age {
			return user, nil
		}
	}

	return user{}, errors.New("user is not found")
}

func withArgsGrouping3(name1, name2 string) (user, error) {
	for _, user := range users {
		if user.name == name1 && user.name == name2 {
			return user, nil
		}
	}

	return user{}, errors.New("user is not found")
}

func withVariadic4(names ...string) {
	fmt.Println(names)
}

/*  2. Function Return Value  */

func withReturnValue() string {

	return "hello"
}

func withMultipleReturnValue() (string, error) {

	return "hello", errors.New("error")
}

func withNamedReturnValue() (message string, err error) {
	message = "hello"
	err = fmt.Errorf("%s", "error")
	return
	//return "hello", errors.New("error")
}

func withNamedReturnValueDefer() (message string, err error) {
	message = "hello"
	err = fmt.Errorf("%s", "error")
	defer func() {
		message = "i am deferred hello"
		err = errors.New("I am deferred error")
	}()

	return
	//return "hello", errors.New("error")
}

/* 3. Function as type  and closure */

type customFunc1 func(string, string) error
type customFunc2 func(string, string) error

func funcAsTypes(c customFunc1) {
	fmt.Println(c("name1", "name2"))
}

func callFunc() {
	f := customFunc2(func(name1, name2 string) error {
		if name1 != name2 {
			return errors.New("not equal")
		}

		return nil
	})
	f1 := customFunc1(f)
	funcAsTypes(f1)
}

/* 5. Advanced */
func stackExample() {
	stackSlice := make([]byte, 10)
	s := runtime.Stack(stackSlice, false)
	fmt.Printf("\n%s", stackSlice[0:s])
}

func first() {
	second()
}

func second() {
	third()
}

func third() {
	for c := 0; c < 10; c++ {
		fmt.Println(runtime.Caller(c))
	}
}

/* 6. Dynamic Statck */

func growTheFunc(bufSize int) {
	b := []byte{}
	for i := 1; i <= bufSize; i++ {
		b = append(b, 1)
	}
}
