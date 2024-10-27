package utils

import (
	"errors"
	"math"
)

const DefaultBatchSize = 100

type Batcher struct {
	batchSize int
	count     int
	curCount  int
	times     int
	curTimes  int
}

func NewBatcher(count, batchSize int) (*Batcher, error) {
	mass := new(Batcher)

	if count < 0 {
		return nil, errors.New("count less than 0")
	}
	if batchSize <= 0 {
		return nil, errors.New("batchSize less or equal than 0")
	}

	mass.count = count
	mass.batchSize = batchSize
	mass.times = int(math.Ceil(float64(count) / float64(batchSize)))
	return mass, nil
}

func (i *Batcher) Iter(start, length *int) bool {
	roundSize := i.batchSize

	if i.curTimes != 0 {
		*start += *length
	}
	if *start+roundSize > i.count {
		roundSize = i.count - *start
	}

	i.curCount += roundSize
	*length = roundSize
	i.curTimes++

	return *start < i.count
}
