package data


// SpotSelection represents the selection strategy for a parking spot.
type SpotSelection int

const (
    // RANDOM represents a random selection strategy.
    RANDOM SpotSelection = iota
    // NEAREST represents the nearest selection strategy.
    NEAREST
)

// String provides a string representation for SpotSelection values.
func (s SpotSelection) String() string {
    return [...]string{"RANDOM", "NEAREST"}[s]
}
