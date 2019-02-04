package rpc

import (
	"fmt"
	"testing"
)

func TestGetPages(t *testing.T) {
	testdata := []struct {
		limit int
		size  int
		pages int
		tail  int
	}{
		{limit: 100, size: 50, pages: 2, tail: 0},
		{limit: 50, size: 50, pages: 1, tail: 0},
		{limit: 51, size: 50, pages: 1, tail: 51},
		{limit: 2, size: 50, pages: 1, tail: 10},
		{limit: 0, size: 50, pages: 0, tail: 0},
		{limit: 123, size: 50, pages: 3, tail: 23},
		{limit: 73, size: 50, pages: 2, tail: 23},
		{limit: 43, size: 50, pages: 1, tail: 43},
		{limit: 1, size: 50, pages: 1, tail: 10},
		{limit: 75, size: 50, pages: 2, tail: 25},
	}

	for i, td := range testdata {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			pages, tail := getPages(td.limit, td.size)
			if pages != td.pages {
				t.Errorf("Pages, expected: %d, got: %d", td.pages, pages)
			}
			if tail != td.tail {
				t.Errorf("Tail, expected: %d, got: %d", td.tail, tail)
			}
		})
	}
}
