package main

import (
	"fmt"
	"testing"
)

func TestExtractFields(t *testing.T) {
	var g Generator
	// names := []string{"sqlscan_test.go"}
	g.parsePackageDir("./")
	g.generate("Transaction")
	fmt.Println(g.buf.String())
}
