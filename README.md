# Huffman

A command line to compress and decompress files using Huffman coding, written entirely in Golang.

## Installation and running

Install golang using this [link](https://go.dev/doc/install)

Clone the repository

```bash
  git clone https://github.com/emmanuelmoon/huffman.git
```

Build the project

```bash
  go build main.go
```

To compress

```bash
  ./huffman -c <input-file> <output-file>
```

To decompress

```bash
  ./huffman -d <compressed-file> <decompressed-file>
```
