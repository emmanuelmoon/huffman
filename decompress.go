package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"huffman/huffman"
	"io"
	"os"
	"strings"
)

func decompress(filename string) {
	m := make(map[rune]int)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	l, _, err := reader.ReadLine()
	if err != nil {
		os.Exit(1)
	}
	json.Unmarshal(l, &m)

	prefixTable := make(map[rune]string)

	tree := huffman.BuildHuffmanTree(m)
	huffman.BuildPrefixTable(&tree, prefixTable)

	reverseTable := make(map[string]rune)

	for key, value := range prefixTable {
		reverseTable[value] = key
	}

	f, err = os.Create("decoded.txt")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	var i uint8 = 0
	builder := strings.Builder{}
	for {
		x, err := reader.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			os.Exit(1)
		}
		if err != nil {
			break
		}
		for i = 0; i < 8; i++ {
			bit := (x & (1 << (7 - i)) >> (7 - i))
			if bit == 1 {
				builder.WriteString("1")
			} else {
				builder.WriteString("0")
			}
			val, ok := reverseTable[builder.String()]
			if ok {
				writer.WriteRune(val)
				builder.Reset()
			}
		}
	}

	writer.Flush()
}
