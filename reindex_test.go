package reindex

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/Konstantin8105/compare"
)

type info struct{ from, to int }

func TestList(t *testing.T) {
	tcs := [][]info{
		[]info{},
		[]info{
			info{from: 5, to: 1},
		},
		[]info{
			info{from: 1, to: 1},
			info{from: 5, to: 2},
		},
		[]info{
			info{from: 1, to: 1},
			info{from: 5, to: 2},
			info{from: 4, to: 3},
		},
	}
	var buf bytes.Buffer
	for i, ts := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var l List[int]
			for k := range ts {
				l.Add(ts[k].from, ts[k].to)
			}
			for p := range ts {
				a, err := l.Get(ts[p].from)
				if err != nil {
					t.Fatal(err)
				}
				if a != ts[p].to {
					t.Errorf("not same %d != %d", a, ts[p].to)
				}
			}
			fmt.Fprintf(&buf, "%v\n", l.data)
			if _, err := l.Get(10000); err == nil {
				t.Errorf("out of range")
			}
		})
	}
	compare.Test(t, filepath.Join("testdata", ".list"), buf.Bytes())
}

func TestAddPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("negative index is not acceptable")
		}
	}()
	var l List[int]
	l.Add(-1, 0)
}
