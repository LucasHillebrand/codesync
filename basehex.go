package main

func hexDump(org []byte) []byte {
	out := make([]byte, len(org)*2)

	for i, j := 0, 0; i < len(org); i++ {
		bytes := byteToHex8(org[i])
		out[j] = bytes[0]
		out[j+1] = bytes[1]
		j += 2
	}

	return out
}

func byteDump(raw []byte) []byte {
	out := make([]byte, len(raw)/2)

	for i, k := 0, 0; i < len(out); i++ {
		out[i] = hex8ToByte([]byte{raw[k], raw[k+1]})
		k += 2
	}
	return out
}

func byteToHex8(num byte) []byte {
	out := ""
	hexChar := "0123456789ABCDEF"
	tmpint := make([]uint8, 2)
	for i := 0; i < 2; i++ {
		tmpint[i] = num % 16
		num -= tmpint[i]
		num /= 16
	}

	for i := 1; i >= 0; i-- {
		out += string(hexChar[tmpint[i]])
	}

	return []byte(out)
}

func hex8ToByte(num []byte) byte {
	out := byte(0)
	hexChar := "0123456789ABCDEF"
	revbytes := make([]byte, 2)

	for i, j := 0, 1; i < len(revbytes); i++ {
		for k := 0; k < len(hexChar); k++ {
			if num[i] == hexChar[k] {
				revbytes[j] = byte(k)
			}
		}
		j--
	}

	for i := 0; i < len(revbytes); i++ {
		out += revbytes[i] * byte(up(16, i))
	}

	return out
}

func up(org, times int) int {
	out := 1

	for i := 0; i < times; i++ {
		out *= org
	}

	return out
}
