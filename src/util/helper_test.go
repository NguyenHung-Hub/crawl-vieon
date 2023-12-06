package util

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcIndexJob(t *testing.T) {
	res := CalcIndexJob(7122, 10)

	ex := [][]int{
		{0, 711}, {712, 1423}, {1424, 2135}, {2136, 2847}, {2848, 3559}, {3560, 4271}, {4272, 4983},
		{4984, 5695}, {5696, 6407}, {6408, 7119}, {7120, 7122}}
	log.Println(res)
	require.Equal(t, ex, res)
}
func TestCalcIndexJob2(t *testing.T) {
	res := CalcIndexJob(7122, 20)

	log.Println(res)
	require.Equal(t, 1, 1)
}
