/**
 * 指针
 *
 */
package main

import "fmt"

const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)

func main() {
	fmt.Println(11)
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)
	println(0 == 0, 0 != 0)

}

func swap(a, b int) (int, int) {
	return b, a
}
