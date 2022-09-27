package helpers

// _ProcessCounterStruct is a struct related to process counting with three fields, requested, processed, and error count.
type _ProcessCounterStruct struct{ RequestedCount, ProcessedCount, ErrorCount uint64 }

// Returns a _ProcessCounterStruct, a function to increment the processed count, and a function to
// increment the error count.
func UseProcessCounter() (counter *_ProcessCounterStruct, incrementProcessed func(), incrementError func()) {
	counter = new(_ProcessCounterStruct)
	incrementProcessed = func() {
		counter.RequestedCount++
		counter.ProcessedCount++
	}
	incrementError = func() {
		counter.RequestedCount++
		counter.ErrorCount++
	}
	return counter, incrementProcessed, incrementError
}
