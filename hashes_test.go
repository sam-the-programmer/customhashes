package customhashes_test

import (
	"fmt"
	"testing"

	hash "github.com/sam-the-programmer/customhashes"
)

func TestHash1(t *testing.T) {
	runTest(hash.Hash1, t)
}

func runTest(hashFunc hash.HashFuncType, t *testing.T) {
	fmt.Println(hashFunc("Hello World"))
	fmt.Println(hashFunc("Hello World1"))

	hashes := []uint64{}

	for i := 0; i < 10; i++ {
		x := fmt.Sprintf("Hello World%d", i)
		hashes = append(hashes, hashFunc(x))
	}

	// Check for duplicates, so if there are some within 100, it's definitely a bad hash
	for i := 0; i < len(hashes)-1; i++ {
		for j := i + 1; j < len(hashes); j++ {
			if hashes[i] == hashes[j] {
				t.Errorf("Duplicate hash found: %d", hashes[i])
			}
		}
	}
}
