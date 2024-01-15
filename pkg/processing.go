package pkg

import (
	"bytes"
	"log"
	"runtime"
	"strconv"

	"github.com/edsrzf/mmap-go"
)

// parse each individual line and update the aggregator for the station
// we deal with each line in bytes to enhance index finding performance
func parseLine(line []byte, aggregators map[string]*Aggregator) {
	sepIndex := bytes.IndexByte(line, ';')
	if sepIndex == -1 {
		log.Println("Separator not found in line:", string(line))
		return
	}

	station := string(line[:sepIndex])
	tempStr := string(line[sepIndex+1:])
	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		log.Println("Unexpected line format:", string(line))
		return
	}

	// Update aggregator for the station
	updateAggregator(station, temp, aggregators)
}

// processChunk processes a chunk of data and updates the local aggregator.
func processChunk(data []byte, results chan<- map[string]*Aggregator) {
	localAggr := make(map[string]*Aggregator)
	lines := bytes.Split(data, []byte("\n")) // split the chunk into lines
	for _, line := range lines {
		if len(line) > 0 {
			parseLine(line, localAggr)
		}
	}
	results <- localAggr
}

func ProcessData(data mmap.MMap) {
	processorCnt := runtime.NumCPU()
	chunkSize := len(data) / processorCnt
	results := make(chan map[string]*Aggregator, processorCnt)

	var start, end int
	for i := 0; i < processorCnt; i++ {
		if i == 0 { // first chunk
			start = 0
			end = chunkSize
		} else { // subsequent chunks
			start = end + 1 // since we ensure that "end" is always a newline, just set to + 1
			end = start + chunkSize
		}

		if i == processorCnt-1 {
			end = len(data)
		}
		for end < len(data) && data[end] != '\n' { // ensure that "end" is always a newline
			end++
		}
		go processChunk(data[start:end], results)
	}

	// Merge results from all chunks
	finalRes := make(map[string]*Aggregator)
	for i := 0; i < processorCnt; i++ {
		localAggr := <-results
		mergeAggregators(finalRes, localAggr)
	}

	// Sort and print results
	printResults(finalRes)
}
