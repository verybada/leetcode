package simplifypath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "root", input: "/", expected: "/"},
		{name: "canonical path", input: "/home/user", expected: "/home/user"},
		{name: "trailing slash", input: "/home/", expected: "/home"},
		{name: "repeated slashes", input: "/home//foo/", expected: "/home/foo"},
		{name: "only slashes", input: "////", expected: "/"},
		{name: "current directory", input: "/a/./b/./c/.", expected: "/a/b/c"},
		{name: "current directory at root", input: "/./", expected: "/"},
		{name: "parent above root", input: "/../", expected: "/"},
		{name: "parent directory", input: "/home/user/Documents/../Pictures", expected: "/home/user/Pictures"},
		{name: "return to root", input: "/a/b/../..", expected: "/"},
		{name: "multiple parents above root", input: "/a/../../b/../../../c", expected: "/c"},
		{name: "mixed separators and dot segments", input: "/a//b////c/d//././/..", expected: "/a/b/c"},
		{name: "three dots are a directory", input: "/.../a/../b/c/../d/./", expected: "/.../b/d"},
		{name: "dots within directory names", input: "/.hidden/..config/file.txt/....", expected: "/.hidden/..config/file.txt/...."},
		{name: "parent removes dotted directory", input: "/.../../.hidden/..", expected: "/"},
		{name: "preserve directory characters", input: "/AbC_123/foo.bar", expected: "/AbC_123/foo.bar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, simplifyPath(tt.input))
		})
	}
}
