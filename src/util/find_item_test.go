package util

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindItemNotInSlice(t *testing.T) {
	slice1 := []string{"1", "2", "3"}
	slice2 := []string{"1", "2", "3", "4", "5"}
	slice3 := []string{"4", "5"}

	res := FindItemNotInSlice(slice1, slice2)
	log.Println(res)

	require.Equal(t, slice3, res)
}

func TestFindItemNotInSlice2(t *testing.T) {
	slice1 := CSVToSliceStr("../../data/test.csv")
	slice2 := CSVToSliceStr("../../data/ribbon_id_2023-12-01_09-35-09.csv")
	slice3 := []string{"4", "5"}

	res := FindItemNotInSlice(slice1, slice2)
	log.Println(res)

	require.Equal(t, slice3, slice3)
}

func TestFindItemNotInSlice3(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"1", "2", "3", "4", "5"}
	slice3 := []string{"1", "2", "3", "4", "5"}

	res := FindItemNotInSlice(slice1, slice2)
	log.Println(res)

	require.Equal(t, slice3, res)
}
