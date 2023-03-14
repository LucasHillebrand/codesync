package main

import (
	"fmt"
	"os"
	"strings"
)

type fileTree struct {
	directorys []string
	files      []string
}

func listDirTree(startDir string) fileTree {
	if startDir[len(startDir)-1] != '/' {
		startDir += "/"
	}

	directorys := []string{startDir}
	files := make([]string, 0)
	for i := 0; i < len(directorys); i++ {
		dir, _ := os.ReadDir(directorys[i])
		for j := 0; j < len(dir); j++ {
			if dir[j].IsDir() {
				//fmt.Println(directorys[i] + dir[j].Name() + "/")
				directorys = append(directorys, directorys[i]+dir[j].Name()+"/")
			} else {
				//fmt.Println(directorys[i] + dir[j].Name())
				files = append(files, directorys[i]+dir[j].Name())

			}
		}

	}
	//fmt.Println(files)
	//fmt.Println(directorys)

	return fileTree{
		directorys: directorys,
		files:      files,
	}
}

func compress(ftree fileTree) string {
	convfiles := make([]string, 0)
	for i := 0; i < len(ftree.files); i++ {
		file, _ := os.ReadFile(ftree.files[i])
		//fmt.Println(string(file))
		convfiles = append(convfiles, string(hexDump(file)))
	}

	return genCompFile(ftree.directorys, ftree.files, convfiles)
}

func genCompFile(directorys []string, fileNames []string, data []string) string {
	out := make([]string, 0)

	for i := 0; i < len(directorys); i++ {
		out = append(out, "DIR:"+directorys[i])
	}

	for i := 0; i < len(fileNames); i++ {
		out = append(out, "FILE:"+fileNames[i]+":"+data[i])
	}

	str := toString(out)

	return str
}

type compressedStruct struct {
	dirs      []string
	filenames []string
	datahex   []string
}

func loadStream(stream string) compressedStruct {
	file := strings.Split(stream, "\n")
	dirs := make([]string, 0)
	filenames := make([]string, 0)
	data := make([]string, 0)

	for i := 0; i < len(file); i++ {
		ln := strings.Split(file[i], ":")
		if ln[0] == "DIR" {
			dirs = append(dirs, ln[1])
		} else if ln[0] == "FILE" {
			filenames = append(filenames, ln[1])
			data = append(data, ln[2])
		}
	}

	out := compressedStruct{
		dirs:      dirs,
		filenames: filenames,
		datahex:   data,
	}
	return out
}

func decompress(cmpFile compressedStruct) {

	for i := 0; i < len(cmpFile.dirs); i++ {
		os.Mkdir(cmpFile.dirs[i], os.FileMode(0777))
	}

	for i := 0; i < len(cmpFile.filenames); i++ {
		err := os.WriteFile(cmpFile.filenames[i], byteDump([]byte(cmpFile.datahex[i])), os.FileMode(0666))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
