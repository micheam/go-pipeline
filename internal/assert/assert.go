package assert

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func NoError(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return false
	}
	return true
}

func readfile(t *testing.T, filename string) []byte {
	t.Helper()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return b
}

func SameContent(t *testing.T, filename1, filename2 string) bool {
	t.Helper()
	c1 := readfile(t, filename1)
	c2 := readfile(t, filename2)
	if diff := cmp.Diff(c1, c2); diff != "" {
		t.Errorf("file content mismatch (-file1, +file2):%s\n", diff)
		return false
	}
	return true
}
