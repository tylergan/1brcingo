package pkg

import (
	"math"
	"sync"
)

type Aggregator struct {
	Min, Max, Sum, Count float64
}

// Use a sync.Pool to reuse aggregators, reducing garbage collection overhead.
// Likely not useful since we are only releasing aggregators at the end of the program,
// so no aggregators are reused. But it's a good example of how to use sync.Pool.
var AggregatorPool = sync.Pool{
	New: func() interface{} {
		return &Aggregator{
			Min: math.MaxFloat64,
			Max: -math.MaxFloat64,
		}
	},
}

// NewAggregator returns an Aggregator from the pool.
func NewAggregator() *Aggregator {
	return AggregatorPool.Get().(*Aggregator)
}

// ReleaseAggregator puts an Aggregator back into the pool.
func ReleaseAggregator(aggr *Aggregator) {
	aggr.Min = math.MaxFloat64
	aggr.Max = -math.MaxFloat64
	aggr.Sum = 0
	aggr.Count = 0
	AggregatorPool.Put(aggr)
}

// updateAggregator updates the aggregator for the given station.
func updateAggregator(station string, temp float64, aggregators map[string]*Aggregator) {
	aggr, exists := aggregators[station]
	if !exists {
		aggr = NewAggregator()
		aggregators[station] = aggr
	}

	aggr.Count++
	aggr.Sum += temp
	if temp < aggr.Min {
		aggr.Min = temp
	}
	if temp > aggr.Max {
		aggr.Max = temp
	}
}

// mergeAggregators merges the local aggregators into the final aggregators.
func mergeAggregators(final, local map[string]*Aggregator) {
	for station, localAggr := range local {
		if finalAggr, exists := final[station]; exists {
			// Station already exists in final resuts. merge data
			finalAggr.Min = math.Min(finalAggr.Min, localAggr.Min)
			finalAggr.Max = math.Max(finalAggr.Max, localAggr.Max)
			finalAggr.Sum += localAggr.Sum
			finalAggr.Count += localAggr.Count
		} else {
			// Station does not exist in final results. Add it
			final[station] = &Aggregator{
				Min:   localAggr.Min,
				Max:   localAggr.Max,
				Sum:   localAggr.Sum,
				Count: localAggr.Count,
			}
		}
	}
}
