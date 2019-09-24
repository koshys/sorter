package gosortint

import (
	"testing"
)

func equal(a, b [][]int ) bool {
	if len(a) != len(b) {
		return false
	}

	for i, s := range a {

		for k, v := range s {
			if v != b[i][k] {
				return false
			}
		}		

	}

	return true
}

func TestNSlice ( t *testing.T) {

	cases := []struct{
		in []int
		want [][]int
	}{
		{
			in: []int{9,8,7,6,5,4,3,2,1,0}, 
			want: [][]int{[]int{9,8,7},[]int{6,5,4},[]int{3,2,1},[]int{0}},
		},
		{
			in: []int{9,8,7,6,5,4,3,2,1,0,-1,-2}, 
			want: [][]int{[]int{9,8,7},[]int{6,5,4},[]int{3,2,1},[]int{0,-1,-2}},
		},		
	};

	for _, c := range cases {
		got := NSlice(4,c.in)
		if !equal(got,c.want) {
			t.Errorf("NSlice(4,%v) got %v want %v",c.in,got,c.want)
		}
	}

}

