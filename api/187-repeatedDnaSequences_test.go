package api

import (
	"fmt"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
