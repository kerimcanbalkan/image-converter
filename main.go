package main

import (
	"fmt"
	"os"

	"github.com/kerimcanbalkan/image-converter/converter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: converter <inputfile.type> <outputfile.type>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputExt := converter.GetFormat(inputFile)
	outputExt := converter.GetFormat(outputFile)

	// Open the input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error opening input file", err)
		os.Exit(1)
	}
	defer inFile.Close()

	// Get decoder for the input format
	decoder, err := converter.GetDecoder(inputExt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// Decode the input image
	img, err := decoder(inFile)
	if err != nil {
		fmt.Printf("failed to decode input file: %v\n", err)
		os.Exit(1)
	}

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer outFile.Close()

	// Get the encoder for the output format
	encoder, err := converter.GetEncoder(outputExt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// Encode the image in the desired format
	if err := encoder(outFile, img); err != nil {
		fmt.Printf("Failed to encode output file: %v\n", err)
		return
	}

	fmt.Println("Conversion successful:", outputFile)
}
