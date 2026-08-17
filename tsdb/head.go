package tsdb

// memSeries represents an in-memory time series.
type memSeries struct {
	// ... existing fields
	lastMmappedTime int64
}

// append handles adding samples during WAL replay, ensuring they don't conflict with mmapped chunks.
func (s *memSeries) append(t int64, v float64) error {
	if t <= s.lastMmappedTime {
		return nil // Gracefully ignore samples already in mmapped chunks
	}
	// ... existing append logic
	return nil
}

// Head.Init logic would call loadMmappedChunks before WAL replay.
func (h *Head) Init() error {
	if err := h.loadMmappedChunks(); err != nil {
		return err
	}
	// Proceed with WAL replay
	return nil
}