package main

import (
	"log"

	"github.com/tylergan/billion_rows/pkg"
)

func main() {
	// Load data into memory
	mappedData, err := pkg.MapFileIntoMem()
	if err != nil {
		log.Fatal(err)
	}
	defer mappedData.Unmap()

	mmapData, error := pkg.MapFileIntoMem()
	if error != nil {
		log.Fatal(error)
	}

	pkg.ProcessData(mmapData)
}
