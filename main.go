package main

import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"sorter/gosort"
	"sync"
)

// a function that can read a line from stdin
func readLine() ( l string, e error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func main() {


	var wg sync.WaitGroup

	fmt.Print("Enter a sequence of integers separated by space for sorting\n")
	in , e := readLine()
	if e != nil {
		fmt.Printf("An Error has occured %v\n",e)
		os.Exit(1)
	}	

	s := strings.Fields(in)
	var unsorted []int
	var n int

	for i := 0; i < len(s) ; i++ {
		n, e = strconv.Atoi(s[i])
		if e != nil {
			fmt.Printf("An Error has occured %v\n",e)
			os.Exit(1)			
		}
		unsorted = append(unsorted, n)
	}	


	work := gosortint.NSlice(4,unsorted)
	for _, w := range work {

		wg.Add(1)
		go gosortint.Sort(w,&wg)

	}
	wg.Wait()
	fmt.Printf("%v\n",work)
	gosortint.Merge(work)


}