package evaluator

import (
	"seville/lexer"
	"seville/object"
	"seville/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"10 + 10 + 10 + 10 - 5", 35},
		{"3 * 3 * 3 - 2", 25},
		{"6 * 10 + 7", 67},
		{"5 + 2 * 10", 25},
		{"20 - 2 * 10", 0},
		{"20 - 2 ** 3", 12},
		{"20 + -2 ** 3", 12},
		{"100 / 2 * 2 + 10", 110},
		{"2 * (10 + 15)", 50},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 != 2", true},
		{"1 == 2", false},
		{"1 <= 2", true},
		{"1 >= 2", false},
		{"2 >= 2", true},
		{"2 <= 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"false == true", false},
		{"false != true", true},
		{"true != false", true},
		{"(1 < 2) == true", true},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 > 2) != false", false},
		{"1 in [1 + 2, 2 - 1]", true},
		{"0 in [1, 2]", false},
		{"1 == 2 in [1, false]", true},
		{`"hi" == "fdsfd" in [1, false]`, true},
		{`"hi" == "hi" in [1, false]`, false},
		{"1 in [fn(n) {1}, 3]", false},
		{"1 in [1, 2]", true},
		{`"1" in {"1": 2}`, true},
		{`1 in {"1": 2}`, false},
		{`2 in {"1": 2}`, false},
		{`"on" + "e" in {"one": 2}`, true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (+%v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`
		if (10 > 1) {
			if (10 > 1) {
				return 10;
			}
			\return 1;
		}
		`, 10},
		{"return 10;", 10},
		{"return 10; 9", 10},
		{"return 2 * 5; 9", 10},
		{"9; return 2 * 5; 9", 10},
		{`if (2 > 1) {return 1;} else {return 2;}`, 1},
		{`if (2 < 1) {return 1;} else {return 2;}`, 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5 >= true",
			"type mismatch: INTEGER >= BOOLEAN",
		},
		{
			`if (10 > 1) {
				if (10 > 1) {
					return true + false
				}

				return 1;
			}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foobar;", "identifier not found: foobar"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[fn (n) {n + 2}]`, "unusable as hash key: FUNCTION"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		testIntegerObject(t, result, tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; }"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let foo = fn(x) {x}; foo(5)", 5},
		{"let foo = fn(x) {return x}; foo(5);", 5},
		{"let double = fn(x) {x * 2}; double(4)", 8},
		{"let add = fn(x, y) {x + y}; add(5, 6)", 11},
		{"let add = fn(x, y) {x + y}; add(5 + 5, add(5, 5))", 20},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	let newAdder = fn(x) {
		fn(y) {
			x + y;
		}
	}
	
	let addTwo = newAdder(2)
	addTwo(2);`

	testIntegerObject(t, testEval(input), 4)
}

func TestRecursion(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`
	let fibonacci = fn(n) {
		if (n <= 1) {
			n
		} else {
			fibonacci(n - 1) + fibonacci(n - 2)
		}
	}
	
	fibonacci(10)`, 55},
		{`
	let fibonacci = fn(n) {
		if (n <= 1) {
			return n
		}
		fibonacci(n - 1) + fibonacci(n - 2)
	}
	
	fibonacci(10)`, 55},
		{`
	let fibonacci = fn(n) {
		if (n == 0) {
			0
		} else {
			if (n == 1) {
				1
			} else {
				fibonacci(n - 1) + fibonacci(n - 2)
			}
		}
	}
	
	fibonacci(10)`, 55},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello, World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (+%v)", evaluated, evaluated)
	}

	if str.Value != "Hello, World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String after string concatendation. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has the wrong value. got=%q", str.Value)
	}
}

func TestStringComparison(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`"chris" == "chris"`, true},
		{`"chris" != "chris"`, false},
		{`"chris" != "bob"`, true},
		{`"chris" == "bob"`, false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if !testBooleanObject(t, evaluated, tt.expected) {
			return
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("hello")`, 5},
		{`len("🍇🍇hi🍇🍇😂😂")`, 8},
		{`len("hello world")`, 11},
		{`len([1, 2, "hello" + "world"])`, 3},
		{`let arr = [1, 2, 3, 4]; len(arr);`, 4},
		{`let arr = [1, 2, "12345", 4]; len(arr[2]);`, 5},
		{`let arr = [1, 2, 3]; let arr = push(arr, 4); arr[-1]`, 4},
		{`let arr = [1, 2, 3, 4]; let arr = push(arr, "hello"); len(arr[4])`, 5},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q. got=%q", expected, errObj.Message)
			}
		}

	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][-1]",
			3,
		},
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"[1, 2, 3][-0]",
			1,
		},
		{
			"[1, 2, 3][-2]",
			2,
		},
		{
			"[1, 2, 3][-3]",
			1,
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			"array index out of bounds: given index 3, array length is: 3",
		},
		{
			"[1, 2, 3][-4]",
			"array index out of bounds: given index -4, array length is: 3",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch result := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(result))
		case string:
			errorObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not error. got=%T (+%v)", evaluated, evaluated)
			}

			if errorObj.Message != tt.expected {
				t.Errorf("error has wrong message. expected: %s, got: %s", tt.expected, errorObj.Message)
			}
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 23}["foo"]`,
			23,
		},
		{
			`{"foo": 5}["bar"]`,
			"key bar not found in hash map",
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			"key foo not found in hash map",
		},
		{
			`{5:5}[5]`,
			5,
		},
		{
			`{true:5}[true]`,
			5,
		},
		{
			`{false:5}[false]`,
			5,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q. got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestAssignmentExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"x = 2", 2},
		{"x = 2 * 2", 4},
		{"let x = 2; x = 3", 3},
		{"let x = 2; x = 3; x", 3},
		{"let x = 100; let foo = fn(x) {x = 2}; foo(5)", 2},
		{"let x = 100; let foo = fn(x) {x = 2}; foo(5); x", 100},
		{"let arr = [1, 2, 3]; arr[0] = 5", 5},
		{"let arr = [1, 2, 3]; arr[0] = 5; arr[0]", 5},
		{"let arr = []; arr[0] = 1", "Array index out of bounds: given index 0, array length is 0"},
		{"let arr = []; arr[2 * 2] = 1", "Array index out of bounds: given index 4, array length is 0"},
		{`let name_to_age = {}; name_to_age["Charlie"] = 99`, 99},
		{`let name_to_age = {}; name_to_age["Charlie"] = 99; name_to_age["Charlie"]`, 99},
		{`let foo = {"one": 1}; foo["two"] = 7 * 7 -47`, 2},
		{`let foo = {"one": 1}; foo["two"] = 7 * 7 -47; foo["two"]`, 2},
		{`let foo = {"one": 1}; foo["two"] = 7 * 7 -47; foo["one"]`, 1},
		{
			`let foo = {}; foo[fn(n) {n ** 2}] = 1`,
			"Hashmap index must be a hashable type, got type *object.Function",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q. got=%q", expected, errObj.Message)
			}
		}
	}
}
