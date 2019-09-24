package gosortint

import (
	"sort"
	"sync"
	"fmt"
)


// NSlice splits a slice into n and returns them
func NSlice( n int, dat []int ) ( slices [][]int ) {

	dl := len(dat)  										// length of big slice
	s := float32(dl)/float32(n)								// accurate slice length
	sl := int(s)											// floor slice length
	if dl < n {
		n = dl
	}
	
    
	if ( s - float32(int(s)) ) >= 0.5 {
	  sl++													// slice length
	}

	for  i := 1; i < n; i++ {
		slices = append(slices,dat[ (sl * (i -1 ) ) : sl * i ])
	}
	slices = append(slices,dat[ (sl * (n -1 ) ) :  ])

	return
}

// Sort the slice
func Sort(dat []int, wg *sync.WaitGroup) {

	sort.Sort(sort.IntSlice(dat))

	wg.Done()
}

// Merge the sorted slices
func Merge(dat [][]int) ( sorted []int ) {

	for k , v := range dat {

		fmt.Printf("%d %v\n", k, v )
	}

	return

}
