package gocat

import (
	"io"
	"log"
	"os"
)

func main() {
	// fmt.Println(os.Args[1:])

	if len(os.Args) == 1 {
		log.Println("minimum one file is needed")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		fd, err := os.Open(arg)
		if err != nil {
			log.Println("error open file ", arg)
			os.Exit(2)
		}

		_, err = io.Copy(os.Stdout, fd)
		if err != nil {
			log.Println("error with io.Copy[]", arg)
		}

		fd.Close()
		// fmt.Println(arg)
	}

}
