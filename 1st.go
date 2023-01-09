package main

import "fmt"

func add(x, y int, z string) (int, string) {
	return x + y, z
}
func main() {
	fmt.Println(add(34, 67, "h"))
	fmt.Println(add(34, 67, "h"))

}
