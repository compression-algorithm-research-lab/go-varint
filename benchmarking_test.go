package varint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBenchmarkingValueRatio(t *testing.T) {

	{
		value := uint64(1)
		ratio := BenchmarkingValueRatio(value)
		assert.Equal(t, 12.5, ratio)
	}

}

func TestBenchmarkingValueSliceRatio(t *testing.T) {

	{
		// 12.5% 是上限，如果出现了较大的值压缩率只会比这个低不会比这个高
		values := []uint64{
			1, 2, 3, 4, 5, 999999,
		}
		ratio := BenchmarkingValueSliceRatio(values)
		assert.Equal(t, 16.666666666666668, ratio)
	}

}
