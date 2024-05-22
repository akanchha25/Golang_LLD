package data

type StarRating int

const (
    StarRatingUnset StarRating = iota // Represents an unset or default value
    ONE
    TWO
    THREE
    FOUR
    FIVE
)

func (s StarRating) GetVal() int {
	return int(s)
}
