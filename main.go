package main

import (
	"fmt"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
)

func main() {
	fmt.Println("Hello, World!")
	path, err := common.GitRootPath(".")
	fmt.Println(path, err)
}
