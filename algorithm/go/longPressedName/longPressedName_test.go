package longpressedname

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		name     string
		typed    string
		expected bool
	}{
		{"kikcxmvzi", "kiikcxxmmvvzz", false},
		{"pyplrz", "ppyyplr", false},
		{"abc", "abd", false},
		{"ab", "abb", true},
		{"adam", "aa", false},
		{"alex", "aaleex", true},
		{"alex", "aaleexn", false},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
		{"foobar", "barfoo", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t,
				test.expected,
				isLongPressedName(test.name, test.typed),
			)
		})
	}
}
