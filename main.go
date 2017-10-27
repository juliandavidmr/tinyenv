package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const VERSION = "1.0"

func main() {
	// fmt.Println("FOO:", os.Getenv("ANDROID_HOME"))
	run()
}

func _getListEnviron(pos int) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[pos])
	}
}

func run() {
	version := flag.Bool("version", false, "show version")
	listEnviron := flag.String("list", "*", "list envoirements")
	flags := flag.Bool("flags", false, "list envoirements")

	flag.Parse()

	if *version { // Version
		fmt.Print("v" + VERSION)
	} else if len(*listEnviron) > 0 { // List envoirements
		switch *listEnviron {
		case "keys":
			_getListEnviron(0)
		case "values":
			_getListEnviron(1)
		}
	} else if *flags {
		fmt.Println("version:", *version)
		fmt.Println("numb:", *listEnviron)
		fmt.Println("tail:", flag.Args())
	}
}
