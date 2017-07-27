package main

import "testing"   
import "fmt"

func TestSquare(t *testing.T) { 
	for input, want := range map[int]int{ 
		2: 4, 
		3: 9,  
	}{
		name := fmt.Sprintf("square(%d)", input)  
		t.Run(name, func(t *testing.T) {
          		if have := square(input); want != have {
            			t.Errorf("want %d, have %d", want, have)
			} 
		})
	} 
}
