package hotlist_test

import (
	"testing"

	"github.com/91rmillere/hotlist"
)

func TestWhitelist_Search(t *testing.T) {

	wl := hotlist.NewWhitelist()

	wl.Items["ABC1234"] = &hotlist.Entry{
		PlateNumber:      "ABC1234",
		MatchingStrategy: hotlist.MatchingStrategyExact,
		Description:      "Allowed To be there",
	}

	wl.Items["QWERTY"] = &hotlist.Entry{
		PlateNumber:      "QWERTY",
		MatchingStrategy: hotlist.MatchingStrategyLenient,
		Description:      "Allowed to be there",
	}

	tt := []struct {
		name       string
		cantidates []string
		expected   hotlist.Result
	}{
		{
			name:       "exact plate exists",
			cantidates: []string{"ABC1234"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "ABC1234",
					MatchingStrategy: hotlist.MatchingStrategyExact,
					Description:      "Allowed To be there",
				},
				IsAlert: false,
				IsBest:  true,
			},
		},
		{
			name:       "exact plate does not exist",
			cantidates: []string{"1234ABC"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "1234ABC",
					MatchingStrategy: hotlist.MatchingStrategyUndefined,
					Description:      "",
				},
				IsAlert: true,
				IsBest:  false,
			},
		},
		{

			name:       "lenient plate exists best match",
			cantidates: []string{"QWERTY", "OWERTY"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "QWERTY",
					MatchingStrategy: hotlist.MatchingStrategyLenient,
					Description:      "Allowed to be there",
				},
				IsAlert: false,
				IsBest:  true,
			},
		},
		{

			name:       "lenient plate exists not best match",
			cantidates: []string{"OWERTY", "QWERTY"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "QWERTY",
					MatchingStrategy: hotlist.MatchingStrategyLenient,
					Description:      "Allowed to be there",
				},
				IsAlert: false,
				IsBest:  false,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := wl.Search(tc.cantidates...)
			if actual != tc.expected {
				t.Errorf("expected %v got %v", tc.expected, actual)
			}
		})
	}

}

func TestBlacklist_Search(t *testing.T) {

	bl := hotlist.NewBlacklist()

	bl.Items["ABC1234"] = &hotlist.Entry{
		PlateNumber:      "ABC1234",
		MatchingStrategy: hotlist.MatchingStrategyExact,
		Description:      "Blue Dodge Charger",
	}

	bl.Items["QWERTY"] = &hotlist.Entry{
		PlateNumber:      "QWERTY",
		MatchingStrategy: hotlist.MatchingStrategyLenient,
		Description:      "Red Ford F-150",
	}

	tt := []struct {
		name       string
		cantidates []string
		expected   hotlist.Result
	}{
		{
			name:       "exact plate exists",
			cantidates: []string{"ABC1234"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "ABC1234",
					MatchingStrategy: hotlist.MatchingStrategyExact,
					Description:      "Blue Dodge Charger",
				},
				IsAlert: true,
				IsBest:  true,
			},
		}, {
			name:       "exact plate does not exist",
			cantidates: []string{"ABC12345"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "ABC12345",
					MatchingStrategy: hotlist.MatchingStrategyUndefined,
					Description:      "",
				},
				IsAlert: false,
				IsBest:  false,
			},
		},
		{
			name:       "lenient plate exists is best",
			cantidates: []string{"QWERTY", "OWERTY"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "QWERTY",
					MatchingStrategy: hotlist.MatchingStrategyLenient,
					Description:      "Red Ford F-150",
				},
				IsAlert: true,
				IsBest:  true,
			},
		},
		{
			name:       "lenient plate exists is not best",
			cantidates: []string{"OWERTY", "QWERTY"},
			expected: hotlist.Result{
				Entry: hotlist.Entry{
					PlateNumber:      "QWERTY",
					MatchingStrategy: hotlist.MatchingStrategyLenient,
					Description:      "Red Ford F-150",
				},
				IsAlert: true,
				IsBest:  false,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := bl.Search(tc.cantidates...)
			if actual != tc.expected {
				t.Errorf("expected %v got %v", tc.expected, actual)
			}
		})
	}
}
