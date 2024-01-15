package pkg

import (
	"fmt"
	"os"
	"sort"
	"strings"

	mmap "github.com/edsrzf/mmap-go"
	"github.com/pkg/errors"
)

var filename = "/Users/tylergan/Desktop/VisualStudioCode/Personal/billion_rows/weather_stations.csv"

// Print the results for all aggregators in alphabetical order.
func printResults(aggregators map[string]*Aggregator) {
	// Extract keys from the map and sort them
	var keys []string
	for key := range aggregators {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Iterate over the sorted keys and print the results
	results := make([]string, 0, len(keys))
	for _, key := range keys {
		aggr := aggregators[key]
		mean := aggr.Sum / aggr.Count
		result := fmt.Sprintf("%s=%.1f/%.1f/%.1f", key, aggr.Min, mean, aggr.Max)
		results = append(results, result)

		// Finished with the aggregator. Put it back into the pool.
		ReleaseAggregator(aggr)
	}
	fmt.Printf("{%s}\n", strings.Join(results, ", "))
}

// MapFileIntoMem maps the file into memory so that we can directly access it via memory
// without having to read it from disk.
func MapFileIntoMem() (mmap.MMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Open file failed")
	}
	defer file.Close()

	// Memory-map the file for reading in memory.
	mmapedData, err := mmap.Map(file, mmap.RDONLY, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Memory-map file failed")
	}
	return mmapedData, nil
}
