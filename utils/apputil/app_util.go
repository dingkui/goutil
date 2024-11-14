package apputil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Root() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
