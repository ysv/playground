package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	reversed := lo.Reverse([]int{1, 2, 3, 4, 5})

	fmt.Printf("reversed: %v\n", reversed)
}
