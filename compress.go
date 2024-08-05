package main

import (
	"huffman/fileutils"
	"huffman/treeUtils"
)

func compress(inputfile string, outputfile string) {
	var m = make(map[rune]int)
	err := fileutils.MapFile(inputfile, m)
	if err != nil {
		panic(err)
	}
	t := treeUtils.BuildHuffmanTree(m)
	table := make(map[rune]string)
	treeUtils.BuildPrefixTable(&t, table)
	fileutils.WriteToFile(m, table, inputfile, outputfile)
}
