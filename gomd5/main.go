package md5

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagFile = flag.String("file", "", "file used for input")
	flagURL  = flag.String("url", "", "loads data via http get")
)
var input io.Reader

func main() {

	flag.Parse()
	switch {
	case *flagFile != "":
		f, err := os.Open(*flagFile)
		if err != nil {
			fmt.Println("error opening file", *flagFile, err)
			os.Exit(123)
		}
		defer f.Close()
		input = f
	case *flagURL != "":
		resp, err := http.Get(*flagURL)
		if err != nil {
			fmt.Println("error loading file", *flagURL, err)
			os.Exit(456)
		}
		defer resp.Body.Close()
		input = resp.Body
	default:
		input = os.Stdin
	}
	// printMD5(os.Stdin, os.Stdout)
	printMD5(input, os.Stdout)

	/*
		f, err := os.Open("file.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	*/

}

func printMD5(r io.Reader, w io.Writer) {

	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		fmt.Println("error copy data to hash")
		// log.Fatal(err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "%x", h.Sum(nil))

}
