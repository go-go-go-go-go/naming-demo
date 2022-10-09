package main

import (
	"fmt"

	naming_v1 "github.com/go-go-go-go-go/baby-naming/naming"
	naming_v2 "github.com/go-go-go-go-go/baby-naming/v2/naming"
	naming_v3 "github.com/go-go-go-go-go/baby-naming/v3/naming"
)

func main() {
	name := naming_v1.CreateBabyName()
	fmt.Println("v1:", name)

	name = naming_v2.CreateBabyName(true)
	fmt.Println("v2 (male true):", name)
	name = naming_v2.CreateBabyName(false)
	fmt.Println("v2 (male false):", name)

	name = naming_v3.CreateBabyName(true, 5)
	fmt.Println("v3 (male true):", name)
	name = naming_v3.CreateBabyName(false, 5)
	fmt.Println("v3 (male false):", name)
}
