# 🍇 Seville 🍇
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go](https://github.com/chrisfishbob/seville/actions/workflows/go.yml/badge.svg)](https://github.com/chrisfishbob/seville/actions/workflows/go.yml)  

A C-family bytecode-compiled language in pure Go, no libraries, no parser generators, no nothing!


## Hello World!
Here is a toy program that encapsulates most of seville's functionality.
```
let 🍇 = 8
let arr = [1, 2]
let arr_2 = push(arr, 🍇)
let fib = fn(n) {
    if (n <= 1) {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}
let hashmap = {"foo": "bar"}
hashmap["six"] = fib(6)

if (4 + 4 in arr_2) {
    print("Hello")
}

if (hashmap["s" + "i" + "x"] == arr_2[-1]) {
    print("World!")
} 
```


## Two Sum:
Seville can solve real problems, even without loops (yet, coming real soon).   
Here is an optimal O(n) solution to the Leetcode classic Two Sum question in 100% Seville
```
let two_sum = fn(nums, target) {
    let helper = fn(index, seen_nums) {
        let complement = target - nums[index]
        if (complement in seen_nums) {
            return [seen_nums[complement], index]
        }

        seen_nums[nums[index]] = index
        return helper(index + 1, seen_nums)
    }

    return helper(0, {})
}
```

## Installation
Since Seville is 100% pure Go, running it is as simple as running any typical "Hello, World!" program
in Go. Simply download the code and run:
```
go run seville
```

## Progress Landmarks
### Lexer
:white_check_mark: Single-character operators  
:white_check_mark: Integer literals  
:white_check_mark: String literals  
:white_check_mark: Identifiers  
:white_check_mark: Keywords  
:white_check_mark: Multi-character operators  
:white_check_mark: Unicode / emoji support 🌹  
:white_check_mark: Array literals  
:white_check_mark: Array indices  
:white_check_mark: Hashmap literals  
:white_check_mark: Hashmap indices  
:white_check_mark: In keyword

### Parser
In progress, here are the completed ones ...  
:white_check_mark: Integer literals  
:white_check_mark: Let statements (`let foo = 21`)  
:white_check_mark: Return statements  
:white_check_mark: Prefix operators (`-5`, `!ok`)  
:white_check_mark: Infix operators (`5 > 4 == 3 < 4`)  
:white_check_mark: Boolean literals  
:white_check_mark: Grouped expressions (`2 * (1 + 2)`)  
:white_check_mark: Conditionals (`if ... else ...`)  
:white_check_mark: Function literals (`fn(...) {...}`)  
:white_check_mark: Call expressions (`foo()`)  
:white_check_mark: Strings (`"Hello, World!"`)  
:white_check_mark: String concatendation (`"Hello" + " " + "World!"`)  
:white_check_mark: Array literals (`[1, "hello" + "world", fn(x) {x * 2}]`)  
:white_check_mark: Array indices (`arr[1 * 2]`)  
:white_check_mark: Hashmap literals (`{"chris": "aws", "tim": "apple", "satya": "microsoft"}`)   
:white_check_mark: Hashmap indices (`map["chris"]`)   
:white_check_mark: In keyword (`1 in ["hello", 1, false]`)  
:white_check_mark: Identifier Assignment Expressions (`x = 5`)  
:white_check_mark: Index Assignment Expressions (`arr[5] = 10`)  

### Interpreter
Now evaluating ...  
:white_check_mark: Integer literals (`1, 55`)  
:white_check_mark: Boolean literals (`true, false`)  
:white_check_mark: Bang prefix expression (`!false`)    
:white_check_mark: Minus prefix expression (`-5`)   
:white_check_mark: Infix expressions (`(5 + 5) * 2 == 15 + 5`)   
:white_check_mark: If else expression (`if (...) {...} else {...}`)  
:white_check_mark: Return statements (`return 17;`)   
:white_check_mark: Environments (`let x = 2;`)  
:white_check_mark: First-class functions (`let add_one = fn(n) {n + 1}`)  
:white_check_mark: Funcion calls (`fn(x, y) { return x + y }`)  
:white_check_mark: Recursion (`0, 0, 1, 1, 2, 3, 5, 8, ...`)  
:white_check_mark: Strings (`"Hello, World!"`)  
:white_check_mark: String concatendation (`"Hello" + " " + "World!"`)  
:white_check_mark: Array literals (`[1, "hello", fn(n) {n * 2}]`)  
:white_check_mark: Array indices (`arr[1], arr[2 * 2]`)  
:white_check_mark: Negative array indices(`let arr = [1, 2, 3]; arr[-1] == 3`)  
:white_check_mark: Hashmap literals (`{"chris": "aws", "tim": "apple", "satya": "microsoft"}`)  
:white_check_mark: Hashmap indices (`map["chris"]`)  
:white_check_mark: In keyword (`1 in ["hello", 1, false]`)  
:white_check_mark: Identifier Assignment Expressions (`x = 5`)  
:white_check_mark: Array Index Assignment Expressions (`arr[5] = 10`)  
:white_check_mark: Hashmap Index Assignment Expressions (`name_to_id["chris"] = 24601`)  



## Byte Code Compiler && Virtual Machine: In progress!
### Opcode
:white_check_mark: `OpConstant` represents constant values that are known at compile-time  

### Compiler
:white_check_mark: Integer constants  

## Related Projects
Seviile is not my first language, in fact, it's my third!
Here are the others:  
* JYSS- A lisp with a static type-checker, written in Typed Racket. (Closed-source for academic reasons.)  
* [Spark](https://github.com/chrisfishbob/Spark)- Dynamically-typed lisp written in Python 3.10. (avaialble without install via API hosted on AWS)


## FAQ
Q: Why no loops?  

A: Parsing loops is trivial, a mostly-complete `while` loop exists in a dev branch, but becuase
   I am moving the language away from a tree-walking interpreter to be byte code compiled, I do
   not want to waste the energy to implement them twice. Also, I haven't decided what they should
   look like yet!

Q: Why does this language exist?  

A: Because writing languages is fun!  

## Credits
* *Programming Languages: Application and Interpretation* by Shriram Krishnamurthi  
* *Crafting Interpreters* by Rober Nystrom  
* *Writing an Interpreter* in Go by Thorsten Ball    
* *(How to Write a (Lisp) Interpreter (in Python))* by Peter Norvig  

Thanks all!


