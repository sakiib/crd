package main

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

func main() {
	if err := exec.Command("chmod", "+x", "vendor/k8s.io/code-generator/generate-groups.sh").Run(); err != nil {
		panic(errors.Wrapf(err, "chmod"))
	}

	out, err := exec.Command("vendor/k8s.io/code-generator/generate-groups.sh", "all", "github.com/sakiib/crd/pkg/client", "github.com/sakiib/crd/pkg/apis", "book.com:v1alpha1").Output()
	if err != nil {
		panic(errors.Wrapf(err, "run generator"))
	}

	fmt.Println(string(out))
}
