package file

import (
	"fmt"
	"testing"
)

func TestDoesexist(t *testing.T) {
	ok, err := DoesExist("./tmp/", "")
	fmt.Println(ok, err)
}
