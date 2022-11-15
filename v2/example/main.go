package main

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression/v2"
)

func main() {

	r := if_expression.Return(true, "是", "否")
	fmt.Println(r)
	// Output:
	// 是

}
