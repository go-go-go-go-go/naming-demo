package main

import (
	"fmt"

	"gitlab.eng.vmware.com/chuanhaoq/baby-naming/naming"
)

func main() {
	name := naming.CreateBabyName()
	fmt.Println(name)
}
