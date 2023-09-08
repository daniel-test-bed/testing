package main

import (
	"fmt"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
	"testing"
)

func TestGitRootPath(t *testing.T) {
	fmt.Printf(common.GitRootPath("."))
}
