package fileutils

import (
	"bufio"
	"encoding/json"
	"os"
)

func MapFile(filepath string, m map[rune]int) error {
	data, err := os.Open(filepath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			m[char] += 1
		}
	}

	return nil
}

func WriteToFile(frequencyTable map[rune]int,
	prefixTable map[rune]string, inputfile string) {
	m, err := json.Marshal(frequencyTable)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("temp.txt")
	if err != nil {
		panic(err)
	}

	f.Write(m)
	f.WriteString("\n")
	defer f.Close()

	f1, err := os.Open(inputfile)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	var buffer byte = 0
	var bits uint8 = 0

	writer := bufio.NewWriter(f)
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		for _, char := range scanner.Text() {
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
					writer.WriteByte(buffer)
					bits = 0
					buffer = 0
				}
			}
		}
	}

	if buffer > 0 {
		writer.WriteByte(buffer)
	}

	writer.Flush()
}
