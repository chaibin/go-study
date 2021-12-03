/**
 * 指针
 *
 */
package main

import "fmt"

func main() {

	fmt.Println(11)
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)

}

func swap(a, b int) (int, int) {
	return b, a
}
