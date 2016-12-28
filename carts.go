package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

type TCart []int
type TCarts struct {
	Carts       []TCart `json:"carts"`
	Sorted      bool
	MaxElements int
	MinSupp     float64 `json:"min_support"`
	MinConf     float64 `json:"min_confidence"`

	Els [](map[string]int)
}

func (t TCarts) String() string {
	formatted, err := json.MarshalIndent(t, " ", "\t")
	if err != nil {
		log.Fatalf("could not format json config: %v; ", err)
	}
	return string(formatted)
}

// Sorting the items
func (t *TCarts) Sort() {
	for k, v := range t.Carts {
		sort.Ints(v)
		t.Carts[k] = v
		if len(v) > t.MaxElements {
			t.MaxElements = len(v)
		}
	}
	t.Sorted = true
}

func (t *TCarts) Fill() {

	if t.Els == nil {
		if t.MaxElements < 1 {
			panic("We need max number of items in cart")
		}
		t.Els = make([]map[string]int, t.MaxElements)
	}

	for _, v := range t.Carts {
		size := len(v) - 1
		if t.Els[size] == nil {
			t.Els[size] = map[string]int{}
		}
		t.Els[size][fmt.Sprintf("%v", v)]++
	}

}

// Create permutations of max lot size from param pool.
// For a source of 5 elements, it draws
//    All possible 4 out of 5 then recurses deeper
//    All possible 3 out of 4 then recurses deeper
//    All possible 2 out of 3 then recurses deeper
// => All possible 1 out of 2 (desired lotSize)
//
// Pool is expected to be sorted.
//
// Formula: (Poolsize)(Poolsize-1)(Poolsize-2)...(maxLotSize+1)(maxLotSize)
//          ----------------------------------------------------
//          (1*2*...(maxLotSize-1)*maxLotSize)
// => There are maxLotSize! duplicates
//
// Deduplication is kept in a second stage - for clarity.
// But huge memory allocation waste is incurred.
//
// Improvements:
// We would prune the duplicates instantly.
// A more efficient comparison than sliceToString == sliceToString.
// We could use an array (not a slice) as map key.
// Compare sliceAsMapKey() for a demonstration.
// The supreme solution would be to code the pattern of "increasing duplicity".
//   First permutations are always new, subsequent permutations contain more and more repetition.
func permutateElements(pool []string, maxLotSize int) [][][]string {

	lglc := func(fmt string, vals ...interface{}) {} // log local
	// lglc = log.Printf // enable at will

	gatherSlice := make([]([][]string), len(pool)) // global to all recursions

	var fc func([]string, int) // the actual recursive func

	fc = func(pl []string, lvl int) {
		prefix := strings.Repeat(" ", lvl)
		sz := len(pl)
		// i := sz - 1

		if sz < 2 || sz < maxLotSize+1 {
			return
		}

		// for i := sz - 1; i > 0; i-- {
		lglc("%v%v out of %v\n", prefix, sz-1, sz)

		// We turn the index run upsite down.
		// Thus we get output sorted a-z, such as [A B] [A C] [B C]
		if false {
			for j := 0; j < sz; j++ {
				// original order
			}
			for j := sz - 1; j > -1; j-- {
				// reversed order
			}
		}

		for j := sz - 1; j > -1; j-- {
			// lglc("\t%v %v   %v \n", j, pl[:j], pl[j+1:])
			sg := make([]string, j, sz-1)
			copy(sg, pl[:j])
			sg = append(sg, pl[j+1:]...)
			lglc("%v %v %v\n", prefix, j, sg)
			fc(sg, lvl+4) // recursion
			sort.Strings(sg)
			gatherSlice[sz-1] = append(gatherSlice[sz-1], sg)
		}
		// }

		lglc("\n")
		lglc("\n")

	}

	fc(pool, 0)

	dupes := map[string]bool{}
	deduped := make([]([][]string), len(pool))
	for i, v := range gatherSlice {
		if len(v) > 0 {
			for _, v1 := range v {
				str := fmt.Sprintf("%v", v1)
				if !dupes[str] {
					deduped[i] = append(deduped[i], v1)
					// fmt.Printf("%v ", str)  // fmt.Printf does not append \n - contrary to log.Printf
					dupes[str] = true
				} else {
					// fmt.Printf(".")
				}
			}
		}
		// fmt.Printf("\n")
	}
	// fmt.Printf("\n\n")

	return deduped
}

func sliceAsMapkey() {
	var mp map[[4]int]bool
	mp = map[[4]int]bool{}

	s1 := [4]int{1, 2, 3}
	s2 := [4]int{1, 2, 3}
	s3 := [4]int{3, 3, 3}
	s4 := [4]int{1, 2, 3}

	mp[s1] = true
	mp[s2] = true
	mp[s3] = true
	mp[s4] = true

	for k, v := range mp {
		log.Printf("%#v %v", k, v)
	}
}
