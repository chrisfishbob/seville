# 🍇 Seville 🍇
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go](https://github.com/chrisfishbob/seville/actions/workflows/go.yml/badge.svg)](https://github.com/chrisfishbob/seville/actions/workflows/go.yml)  

A C-family compiled language in pure Go, no libraries, no parser generators, no nothing!


## Quick Tour
Here is a toy program that capsulates most of Seville's currently supported features:  
```
let 🍇 = 0
let y = 23
let newAdder = fn(x) { 
    fn (y) {
        x + y
    }
}
let addTen = newAdder(10)
let fib = fn(n) {
    if (n <= 1) {
        return n
    } 
    fib(n - 1) + fib(n - 2)
}

let name = "Sev" + "ille"
fib(addTen(🍇) - len(name) + 7)       # Outputs 55
```


## Progress Landmarks
### Lexer
:white_check_mark: Single-character operators  
:white_check_mark: Integer literals  
:white_check_mark: Stringf literals  
:white_check_mark: Identifiers  
:white_check_mark: Keywords  
:white_check_mark: Multi-character operators  
:white_check_mark: Unicode / emoji support 🌹  

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
:white_check_mark: Strigns (`"Hello, World!"`)  
:white_check_mark: String concatendation (`"Hello" + " " + "World!"`)  
  
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
  
### Byte Code Compiler && Virtual Machine
Coming soon!
## Related Projects
Seviile is not my first language, in fact, it's my third!
Here are the others:  
* JYSS- A lisp with a static type-checker, written in Typed Racket. (Closed-source for academic reasons.)  
* [Spark](https://github.com/chrisfishbob/Spark)- Dynamically-typed lisp written in Python 3.10. (avaialble without install via API hosted on AWS)


## Credits
* *Programming Languages: Application and Interpretation* by Shriram Krishnamurthi  
* *Crafting Interpreters* by Rober Nystrom  
* *Writing an Interpreter* in Go by Thorsten Ball    
* *(How to Write a (Lisp) Interpreter (in Python))* by Peter Norvig  

Thanks all!


