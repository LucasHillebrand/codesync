package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func initfile(home bool) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("server: ")
	b, _ := buf.ReadBytes('\n')
	if b[len(b)-1] == '\n' {
		b = b[0 : len(b)-1]
	}
	if home {
		p, _ := os.UserHomeDir()
		os.WriteFile(p+"/.syncconf", b, os.FileMode(0666))
	} else {
		os.WriteFile("./.syncconf", b, os.FileMode(0666))
	}
}

func cli() {
	srvb, err := os.ReadFile("./.syncconf")
	if err != nil {
		p, _ := os.UserHomeDir()
		srvb, err = os.ReadFile(p + "/.syncconf")
		if err != nil {
			fmt.Println(".syncconf file not found using default, generate a config file by typing: syncr -ih")
			srvb = []byte("127.0.0.1:1310")
		}
	}

	server := string(srvb)
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(os.Args) >= 4 && os.Args[1] == "push" {
		conn.Write([]byte("push"))
		str := compress(listDirTree(os.Args[2]))
		compressed := os.Args[3] + "\n" + str
		conn.Write(uint64ToBytes(uint64(len(compressed))))
		l := len(compressed)
		for i := 0; i < len(compressed); i += 2048 {
			tmp := l - i
			if i+2048 < l {
				conn.Write([]byte(compressed[i : i+2048]))
			} else if tmp > 0 {
				conn.Write([]byte(compressed[i:]))
			} else {
				return
			}
		}
	} else if len(os.Args) >= 3 && os.Args[1] == "get" {
		conn.Write([]byte("recv"))
		conn.Write(uint64ToBytes(uint64(len(os.Args[2]))))
		conn.Write([]byte(os.Args[2]))
		bsize := make([]byte, 8)
		conn.Read(bsize)
		size := bytesToUint64(bsize)
		str := ""
		for i := uint64(0); i < size; i += 2048 {
			tmp := make([]byte, 2048)
			conn.Read(tmp)
			str += string(tmp)
		}
		str = str[0:size]
		decompress(loadStream(str))
		//fmt.Println(str)
	}

}
