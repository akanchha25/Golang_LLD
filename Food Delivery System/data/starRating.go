package data

type StarRating int

const (
    ONE   StarRating = 1
    TWO   StarRating = 2
    THREE StarRating = 3
    FOUR  StarRating = 4
    FIVE  StarRating = 5
)

func (s StarRating) GetVal() int {
	return int(s)
}