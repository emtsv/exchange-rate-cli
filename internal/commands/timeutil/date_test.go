package timeutil_test

import (
	"testing"
	"time"

	"github.com/emtsv/exchange-rate-cli/internal/commands/timeutil"
	"github.com/stretchr/testify/require"
)

func TestDate2(t *testing.T) {
	t.Parallel()
	now := time.Now()

	testTable := []struct {
		name         string
		date         string
		expectedDate string
		expectedErr  require.ErrorAssertionFunc
	}{
		{
			name:         "valid date",
			date:         "2000-01-01",
			expectedDate: "2000/01/01",
			expectedErr:  require.NoError,
		},

		{
			name: "to old date",
			date: "1900-01-01",
			expectedErr: func(t require.TestingT, err error, i ...any) {
				require.ErrorIs(t, err, timeutil.ErrDateTooOld)
			},
		},
		{
			date:         "2022-12-",
			expectedDate: "",
			expectedErr: func(t require.TestingT, err error, i ...any) {
				require.ErrorIs(t, err, timeutil.ErrInvalidDate)
			},
		},
		{
			name:         "empty date",
			date:         "",
			expectedDate: now.Format("2006/01/02"),
			expectedErr:  require.NoError,
		},
		{
			name:         "date in future",
			date:         now.AddDate(1, 0, 0).Format("2006-01-02"),
			expectedDate: "",
			expectedErr: func(t require.TestingT, err error, i ...any) {
				require.ErrorIs(t, err, timeutil.ErrDateInFuture)
			},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := timeutil.ParseDate(tt.date)
			tt.expectedErr(t, err)
			require.Equal(t, tt.expectedDate, result)
		})
	}
}
