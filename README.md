# 🍇 Seville 🍇
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go](https://github.com/chrisfishbob/seville/actions/workflows/go.yml/badge.svg)](https://github.com/chrisfishbob/seville/actions/workflows/go.yml)  

A C-family compiled language in pure Go, no libraries, no parser generators, no nothing!


## Quick Tour
Here is a toy program that capsulates most of Seville's currently supported features:  
```
let x = 2
let y = 10
let newAdder = fn(x) { fn (y) {x + y}}
let min = fn(x, y) {
    if (x < y) {
        x
    } else {
        y
    }
}
let addTwenty = newAdder(20)
min(addTwenty(x), y * 2)    # prints 20
```


## Progress Landmarks
### Lexer
:white_check_mark: Single-character operators  
:white_check_mark: Integer literals  
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
  
### Byte Code Compiler && Virtual Machine
Coming soon!
