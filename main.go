package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "huffman",
		Description: "encode and decode files using huffman coding",
		Usage:       "a compression tool",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "compress", Aliases: []string{"c"}, Usage: "compress a file (to be followed by input and output filenames)", DefaultText: ""},
			&cli.BoolFlag{Name: "decompress", Aliases: []string{"d"}, Usage: "decompress a file", DefaultText: ""},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("compress") && cCtx.Bool("decompress") {
				return errors.New("provide only one flag")
			}
			if cCtx.NArg() < 2 {
				return errors.New("provide both input and output file names")
			}
			input := cCtx.Args().Get(0)
			output := cCtx.Args().Get(1)
			if cCtx.Bool("compress") {
				compress(input, output)
			} else if cCtx.Bool("decompress") {
				decompressFile(input, output)
			} else {
				log.Fatal("flag is required")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
