package main

import "testing"

for _, testcase := range testcases {
	name := fmt.Sprintf("divide(%d,%d)", testcase.n, testcase.d)
	pt.Run(name, func(t *testing.T) {
		q, r, err := divide(testcase.n, testcase.d)
		if want, have := testcase.err, err; want != have {
			t.Fatal("err: want %v, have %v, want have)
			}
                if want, have := testcase.q, q; want != have {
                        t.Fatal("err: want %v, have %v, want have)
                	}
                if want, have := testcase.r, r; want != have {
                        t.Fatal("err: want %v, have %v, want have)
        	        }
		})
	}
}

func BenchmarkSquare(b *testing.B) {
	for _, n := range []int{9, 99, 999, 9999} {
		b.Run(fmt.Sprint(n)		
