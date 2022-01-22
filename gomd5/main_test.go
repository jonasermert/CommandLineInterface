package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("go")
	out := &bytes.Buffer()
	printMD5(in, out)
	want := "YOUR HASHCODE\n" // You can generate it with md5-generator.de
	got := out.String()
	if got != want {
		t.Errorf("got wrong md5 hash")
	}
}

// Negativ Test

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("go2")
	out := &bytes.Buffer()
	printMD5(in, out)
	want := "YOUR HASHCODE\n" // You can generate it with md5-generator.de
	got := out.String()
	if got != want {
		t.Errorf("got wrong md5 hash")
	}
}
