package model

type Interval struct {
	Max, Min float64
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) IsContains(n float64) bool {
	return n >= i.Min && n <= i.Max
}

func (i Interval) IsSurrouonds(n float64) bool {
	return n > i.Min && n < i.Max
}
