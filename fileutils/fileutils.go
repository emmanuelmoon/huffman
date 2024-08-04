package fileutils

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

func MapFile(filepath string, m map[rune]int) error {
	file, err := os.Open("135-0.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		m[r] += 1
	}
	return nil
}

func WriteToFile(frequencyTable map[rune]int,
	prefixTable map[rune]string, inputfile string, outputfile string) {
	m, err := json.Marshal(frequencyTable)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(m)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString("\n")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f1, err := os.Open(inputfile)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	var buffer byte = 0
	var bits uint8 = 0

	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(f1)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		value, ok := prefixTable[char]
		if !ok {
			os.Exit(1)
		}
		for _, b := range value {
			if b == '1' {
				buffer |= 1 << (7 - bits)
			}
			bits++
			if bits == 8 {
				err := writer.WriteByte(buffer)
				if err != nil {
					panic(err)
				}
				bits = 0
				buffer = 0
			}
		}
	}

	if buffer > 0 {
		err := writer.WriteByte(buffer)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}
