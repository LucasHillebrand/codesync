package main

import "os"

func NextChars(org string, start, length int) string {
	out := ""
	for i := start; i < start+length && i < len(org); i++ {
		out += string(org[i])
	}
	return out
}

func Count(org, keyword string) int {
	out := 0

	for i := 0; i < len(org); i++ {
		if NextChars(org, i, len(keyword)) == keyword {
			out++
		}
	}

	return out
}

func Split(org, keyword string) []string {
	out := make([]string, Count(org, keyword)+1)

	for i, col := 0, 0; i < len(org); i++ {
		if NextChars(org, i, len(keyword)) == keyword {
			col++
			i += len(keyword) - 1
		} else {
			out[col] += string(org[i])
		}
	}
	return out
}

func ReadFileList(filename string) (out []string, err string) {

	bytes, errstr := os.ReadFile(filename)
	if errstr != nil {
		err = errstr.Error()
	} else {
		str := ""
		for i := 0; i < len(bytes); i++ {
			str += string(bytes[i])
		}
		return Split(str, "\n"), ""
	}

	return
}

func toString(org []string) string {
	out := ""
	for i := 0; i < len(org); i++ {
		out += org[i]
		if i+1 < len(org) {
			out += "\n"
		}
	}
	//fmt.Printf("%s,--\n", out)
	return out
}

func pow(org, power int) int {
	out := 1

	for i := 0; i < power; i++ {
		out *= org
	}

	return out
}

func uint64ToBytes(org uint64) []byte {
	out := make([]byte, 8)

	for i := 0; org > 0; org /= 256 {
		out[i] = byte(org % 256)
		org -= org % 256
		i++
	}

	return out
}

func bytesToUint64(src []byte) uint64 {
	out := uint64(0)

	for i := 0; i < len(src); i++ {
		out += uint64(src[i]) * uint64(pow(256, i))
	}

	return out
}
