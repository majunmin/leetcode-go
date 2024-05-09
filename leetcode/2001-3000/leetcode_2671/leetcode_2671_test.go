package leetcode_2671

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	freqTracker := Constructor()
	freqTracker.Add(1)
	freqTracker.Add(1)
	freqTracker.Add(1)
	freqTracker.DeleteOne(2)
	freqTracker.DeleteOne(2)
	freqTracker.Add(3)
	freqTracker.DeleteOne(3)
	freqTracker.DeleteOne(3)

	freqTracker.Add(6)
	freqTracker.Add(6)
	freqTracker.Add(6)
	fmt.Println(freqTracker.HasFrequency(3))
}
