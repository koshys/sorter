package gosortint

import (
	"sort"
	"sync"
)

// NSlice splits a slice into n and returns them
func NSlice(n int, dat []int) (slices [][]int) {

	dl := len(dat)                // length of big slice
	s := float32(dl) / float32(n) // accurate slice length
	sl := int(s)                  // floor slice length
	if dl < n {
		n = dl
	}

	if (s - float32(sl)) >= 0.5 {
		sl++ // slice length
	}

	for i := 1; i < n; i++ {
		slices = append(slices, dat[(sl*(i-1)):sl*i])
	}
	slices = append(slices, dat[(sl*(n-1)):])

	return
}

// Sort the slice
func Sort(dat []int, wg *sync.WaitGroup) {

	sort.Sort(sort.IntSlice(dat))

	wg.Done()
}

func isEmpty(dat [][]int) bool {

	for _, v := range dat {
		if len(v) > 0 {
			return false
		}
	}

	return true
}

// Merge the sorted slices
func Merge(dat [][]int) (sorted []int) {

	start := false
	type pair struct {
		k int
		v int
	}
	nxt := pair{}
	for !isEmpty(dat) {
		for k, v := range dat {

			if len(v) == 0 {
				continue
			}
			if start {
				if nxt.v > dat[k][0] {
					nxt = pair{k: k, v: dat[k][0]}
				}
			} else {
				nxt = pair{k: k, v: dat[k][0]}
				start = true
			}
		}
		sorted = append(sorted, nxt.v)
		dat[nxt.k] = dat[nxt.k][1:] // pop the queue
		start = false
		nxt = pair{}
	}

	return

}
