package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	encode()
	decode()
}

func encode() {
	// Open the input file
	inputFile, err := os.Open("large_file.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file and a gzip writer
	outputFile, err := os.Create("large_file.txt.gz")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// Set the chunk size (e.g., 1MB)
	chunkSize := 1024 * 1024

	// Create a buffer to store the chunk
	buffer := make([]byte, chunkSize)

	// Process the file in chunks
	for {
		n, err := inputFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading input file:", err)
			return
		}

		_, err = gzipWriter.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	fmt.Println("Gzip encoding complete!")
}

func decode() {
	// Open the input (compressed) file
	inputFile, err := os.Open("large_file.txt.gz")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("decompressed_file.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a gzip reader to decompress the input file
	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gzipReader.Close()

	// Copy the decompressed data from the gzip reader to the output file
	_, err = io.Copy(outputFile, gzipReader)
	if err != nil {
		fmt.Println("Error copying decompressed data:", err)
		return
	}

	fmt.Println("Decompression complete!")
}
