package helpers

// _ProcessCounterStruct is a struct with three fields, each of which is an unsigned 64-bit integer related to process counting.
type _ProcessCounterStruct struct{ RequestedCount, ProcessedCount, ErrorCount uint64 }

// UseProcessCounter returns a _CounterStruct data, a function to increment the processed count, and a function to
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
