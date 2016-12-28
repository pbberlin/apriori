package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_permutateElements(t *testing.T) {

	type tt struct {
		inp         []string
		placeHolder int
		wnt         [][][]string
	}

	ts := []tt{
		tt{
			[]string{"A", "B", "C"},
			2,
			[][][]string{
				[][]string{
					[]string{"A", "B"}, []string{"A", "C"}, []string{"B", "C"},
				},
				[][]string{
					[]string{"A"}, []string{"B"}, []string{"C"},
				},
			},
		},
		tt{
			[]string{"A", "B", "C", "D"},
			2,
			[][][]string{
				[][]string{
					[]string{"A", "B", "C"},
					[]string{"A", "B", "D"},
					[]string{"A", "C", "D"},
					[]string{"B", "C", "D"},
				},
				[][]string{
					[]string{"A", "B"}, []string{"A", "C"}, []string{"B", "C"},
					[]string{"A", "D"}, []string{"B", "D"},
					[]string{"C", "D"},
				},
				[][]string{
					[]string{"A"}, []string{"B"}, []string{"C"}, []string{"D"},
				},
			},
		},
		tt{
			[]string{"A", "B", "C", "D", "E"},
			2,
			[][][]string{
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"A", "B", "C", "E"},
					[]string{"A", "B", "D", "E"},
					[]string{"A", "C", "D", "E"},
					[]string{"B", "C", "D", "E"},
				},
				[][]string{
					[]string{"A", "B"}, []string{"A", "C"}, []string{"B", "C"},
					[]string{"A", "D"}, []string{"B", "D"},
					[]string{"C", "D"},
				},
			},
		},
	}

	for k, v := range ts {
		gotAll := permutateElements(v.inp, len(v.inp)-len(v.wnt))
		inpStr := fmt.Sprintf("%v", v.inp)
		log.Printf("#%2v \ninp %v", k, inpStr)

		for i := 0; i < len(v.wnt); i++ {
			got := gotAll[len(v.inp)-i-1]
			wntStr := fmt.Sprintf("%v", v.wnt[i])
			gotStr := fmt.Sprintf("%v", got)
			msg := fmt.Sprintf("sz %v\ngot %v\nwnt %v", len(v.inp)-i-1, gotStr, wntStr)
			if gotStr != wntStr {
				t.Errorf("FAIL: %v", msg)
			} else {
				log.Printf("PASS: %v", msg)
			}
		}

		log.Printf("\n\n")

	}

}
