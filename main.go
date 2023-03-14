package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	srvenabl := false
	port := 1310
	path := "./files"
	inifile := false
	home := false
	stop := false

	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-s":
			srvenabl = true
		case "-p":
			port, _ = strconv.Atoi(os.Args[i+1])
			i++
		case "-e":
			path = os.Args[i+1]
			i++
		case "-i":
			inifile = true
		case "-ih":
			inifile = true
			home = true
		case "--help":
			fmt.Printf("default:\ncodesync push [directory to upload] [name]\ncodesync get [name]\n\noptions:\n-i < create an syncconf file in your current directory\n-ih < create a global syncconf file\n-s < work as an server\n-p [port num] < define a port number for the server\n-e [export dir] < export directory for the .comp files\n")
			stop = true
		}
	}

	if srvenabl {
		srv(port, path)
	} else if inifile {
		initfile(home)
	} else if !stop {
		cli()
	}
}
