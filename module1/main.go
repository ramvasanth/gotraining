// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program demonstrating when implicit interface conversions
// are provided by the compiler.
package main

import "fmt"

func main() {

}

// Mover provides support for moving things.
type Mover interface {
	Move()
}

// Locker provides support for locking and unlocking things.
type Locker interface {
	Lock()
	Unlock()
}

// MoveLocker provides support for moving and locking things.
type MoveLocker interface {
	Mover
	Locker
}

// bike represents a concrete type for the example.
type bike struct{}

// Move can change the position of a bike.
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// Lock prevents a bike from moving.
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

// Unlock allows a bike to be moved.
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}
