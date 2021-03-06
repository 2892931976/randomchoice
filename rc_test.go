package randomchoice

import (
	"testing"
)

func Benchmark100_3(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		RandomChoice(100, 3)
	}
}

func Benchmark10_3(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		RandomChoice(10, 3)
	}
}

func TestRC(t *testing.T) {
    var x []int
    x = RandomChoice(10, 3)
    if len(x) == 3{
        t.Log("success") 
    } else{
        t.Errorf("error, expect:%v, got:%v\n", 3, len(x)) 
    }

    x = RandomChoice(3, 3)
    if len(x) == 3{
        t.Log("success") 
    } else{
        t.Errorf("error, expect:%v, got:%v\n", 3, len(x)) 
    }

    x = RandomChoice(3, 0)
    if len(x) == 0{
        t.Log("success") 
    } else{
        t.Errorf("error, expect:%v, got:%v\n", 0, len(x)) 
    }

    x = RandomChoice(3, 4)
    if len(x) == 3{
        t.Log("success") 
    } else{
        t.Errorf("error, expect:%v, got:%v\n", 3, len(x)) 
    }
}


