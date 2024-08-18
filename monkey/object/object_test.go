package object

import "testing"


func TestStringHashKey(t *testing.T) {
    hello1 := &String{Value: "Hello World"}
    hello2 := &String{Value: "Hello World"}
    diff1 := &String{Value: "My name is julian"}
    diff2 := &String{Value: "My name is julian"}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("strings with same content have different hash keys")
    }
    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("strings with same content have different hash keys")
    }
    if hello1.HashKey() == diff2.HashKey() {
        t.Errorf("strings with different content have same hash keys")
    }
}

func TestBooleanHashKey(t *testing.T) {
    hello1 := &Boolean{Value: false}
    hello2 := &Boolean{Value: false}
    diff1 := &Boolean{Value: true}
    diff2 := &Boolean{Value: true}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("booleans with same content have different hash keys")
    }
    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("booleans with same content have different hash keys")
    }
    if hello1.HashKey() == diff2.HashKey() {
        t.Errorf("booleans with different content have same hash keys")
    }
}

func TestIntegerHashKey(t *testing.T) {
    hello1 := &Integer{Value: 420}
    hello2 := &Integer{Value: 420}
    diff1 := &Integer{Value: 5}
    diff2 := &Integer{Value: 5}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("integers with same content have different hash keys")
    }
    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("integers with same content have different hash keys")
    }
    if hello1.HashKey() == diff2.HashKey() {
        t.Errorf("integers with different content have same hash keys")
    }
}

