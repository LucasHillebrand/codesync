package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func srv(port int, path string) {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	os.Mkdir(path, os.FileMode(0777))

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	for {
		conn, err := ln.Accept()
		if err == nil {
			go func(conn net.Conn) {
				mode := make([]byte, 4)
				conn.Read(mode)

				if string(mode) == "push" {
					bsize := make([]byte, 8)
					conn.Read(bsize)
					size := bytesToUint64(bsize)
					fmt.Printf("client: %s sends %d bytes\n", conn.RemoteAddr().String(), size)
					str := ""
					for i := 0; i < int(size); i += 2048 {
						tmp := make([]byte, 2048)
						conn.Read(tmp)
						str += string(tmp)
						//fmt.Printf("recieved %d from %d bytes, %d bytes left\n", i, size, int(size)-i)
					}

					lines := strings.Split(str, "\n")
					os.WriteFile(path+lines[0]+".comp", []byte(str[0:size]), os.FileMode(0666))
				} else if string(mode) == "recv" {
					bsize := make([]byte, 8)
					conn.Read(bsize)
					size := bytesToUint64(bsize)
					name := make([]byte, size)
					conn.Read(name)
					file, err := os.ReadFile(path + string(name) + ".comp")
					if err != nil {
						conn.Write(uint64ToBytes(0))
					} else {
						fmt.Printf("client: %s recieves %d bytes\n", conn.RemoteAddr().String(), len(file))
						conn.Write(uint64ToBytes(uint64(len(file))))
						for i := 0; i < len(file); i += 2048 {
							if i+2048 < len(file) {
								conn.Write(file[i : i+2048])
							} else {
								conn.Write(file[i:])
							}
						}
					}
				} else if string(mode) == "list" {
					dirs, _ := os.ReadDir(path)
					files := make([]string, 0)
					for i := 0; i < len(dirs); i++ {
						name := dirs[i].Name()
						files = append(files, name[0:len(name)-len(".comp")])
					}
					str := toString(files)
					conn.Write(uint64ToBytes(uint64(len(str))))
					conn.Write([]byte(str))
				}
			}(conn)
		}
	}
}
