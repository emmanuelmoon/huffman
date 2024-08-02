package main

import (
	"fmt"
	"huffman/fileutils"
	"huffman/huffman"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Invalid or no argument")
		os.Exit(1)
	}
	filepath := args[1]
	var m = make(map[rune]int)
	err := fileutils.MapFile(filepath, m)
	if err != nil {
		panic(err)
	}
	t := huffman.BuildHuffmanTree(m)
	table := make(map[rune]string)
	huffman.BuildPrefixTable(&t, table)
	fileutils.WriteToFile(m, table, filepath)
}
