package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCSVToSliceStr(t *testing.T) {

	s1 := CSVToSliceStr("../csv/test/id.csv")
	s2 := []string{"bff041f2-1da7-4c77-abff-e1a6e9e564f5", "76a50f23-d58c-4ece-994a-606c53a78707"}
	require.Equal(t, s1, s2)
}

func TestWriteSliceStrToCSV(t *testing.T) {

	ids := []string{"bff041f2-1da7-4c77-abff-e1a6e9e564f5", "76a50f23-d58c-4ece-994a-606c53a78707"}

	err := WriteSliceStrToCSV("../csv/test/", "id_2", ids, true)
	require.NoError(t, err)

}
