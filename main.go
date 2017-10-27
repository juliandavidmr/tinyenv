package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
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
func _getNumCPUs() {
	fmt.Print(runtime.NumCPU())
}

func _getIPs() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}

func run() {
	version := flag.Bool("version", false, "show version")
	listEnviron := flag.String("list", "", "list envoirements")
	cpus := flag.Bool("cpus", false, "number of cpus")
	inter := flag.Bool("inter", false, "network interfaces")
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
	} else if *cpus {
		_getNumCPUs()
	} else if *inter {
		_getIPs()
	} else if *flags {
		fmt.Println("version:", *version)
		fmt.Println("numb:", *listEnviron)
		fmt.Println("tail:", flag.Args())
	}
}
