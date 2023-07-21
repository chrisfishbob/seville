package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello, World!"}
	hello2 := &String{Value: "Hello, World!"}
	diff1 := &String{Value: "My name is Chris"}
	diff2 := &String{Value: "My name is Chris"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys. string: %s", hello1.Value)
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys. string: %s", hello1.Value)
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with differen content have different hash keys. s1: %s, s2: %s", hello1.Value, diff1.Value)
	}
}

func TestIntegerHashKey(t *testing.T) {
	num1 := &Integer{Value: 1}
	num2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 23}
	diff2 := &Integer{Value: 23}

	if num1.HashKey() != num2.HashKey() {
		t.Errorf("integers with same content have different hash keys. string: %d", num1.Value)
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with same content have different hash keys. string: %d", num1.Value)
	}
	if num1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with differen content have different hash keys. n1: %d, n2: %d", num1.Value, diff1.Value)
	}
}

func TestBooleanHashKey(t *testing.T) {
	num1 := &Boolean{Value: true}
	num2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if num1.HashKey() != num2.HashKey() {
		t.Errorf("booleans with same content have different hash keys. string: %t", num1.Value)
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("booleans with same content have different hash keys. string: %t", num1.Value)
	}
	if num1.HashKey() == diff1.HashKey() {
		t.Errorf("booleans with differen content have different hash keys. n1: %t, n2: %t", num1.Value, diff1.Value)
	}
}
