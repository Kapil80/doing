package main

import "fmt"

func sub(x, y int, z string) (int, string) {
	return x - y, z
}
func main() {
	fmt.Println(sub(34, 67, "h"))

}
