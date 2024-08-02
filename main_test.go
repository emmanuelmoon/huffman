package main

import (
	"huffman/fileutils"
	"testing"
)

func TestMain(t *testing.T) {
	var MapGot = make(map[rune]int)
	got := fileutils.MapFile("135-0.txt", MapGot)
	if got != nil {
		t.Errorf(got.Error())
	}

	if MapGot['X'] != 333 {
		t.Errorf("Expected: %v, Got: %v", 333, MapGot['X'])
	}

	if MapGot['t'] != 223000 {
		t.Errorf("Expected: %v, Got: %v", 333, MapGot['X'])
	}
}
