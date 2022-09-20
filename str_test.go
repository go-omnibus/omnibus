package omnibus

import (
	"fmt"
	"testing"
)

func TestStrPadLeft(t *testing.T) {
	s := StrPadLeft("1", 5, "0")

	fmt.Printf("%s", s)
}