package cbr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRateRUB_BasicCases(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name        string
		valute      string
		nominal     int
		expectedVal float64
		expectedErr require.ErrorAssertionFunc
	}{
		{
			name:        "valid parsing",
			valute:      "10,43",
			nominal:     1,
			expectedVal: 10.43,
			expectedErr: require.NoError,
		},
		{
			name:        "comma as separator",
			valute:      "10,49",
			nominal:     1,
			expectedVal: 10.49,
			expectedErr: require.NoError,
		},
		{
			name:    "empty value",
			valute:  "",
			nominal: 1,
			expectedErr: func(t require.TestingT, err error, i ...any) {
				require.ErrorIs(t, err, ErrVal)
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := Valute{
				Value:   tt.valute,
				Nominal: tt.nominal,
			}

			result, err := v.RateRUB()
			tt.expectedErr(t, err)
			if err == nil {
				require.Equal(t, tt.expectedVal, result)
			}
		})
	}
}
