# Tiny JSON Parser

This project is a simple JSON parser written in Go.

## Features

- Tokenizes basic JSON structures including objects, arrays, strings, and numbers.
- Handles whitespace and special characters.
- Provides a simple API for tokenizing strings.

## Usage

To use the lexer, create a new instance with the input string and call the `NextToken` method to retrieve the next token.

```go
lexer := NewLexer(input)
token := lexer.NextToken()
```

## TODO:
- [x] Parse Numbers in Keys, e.x. `{"key1": 1, "key2": 2, "3", 3, "a123": true}`
- [] Parser
- [] Support Unicode characters
- [] Escape characters in strings
    - [x] `\"` quotation mark
    - [x] `\\` backslash
    - [x] `\/` solidus
    - [x] `\b` backspace
    - [x] `\f` form feed
    - [x] `\n` line feed
    - [x] `\r` carriage return
    - [x] `\t` horizontal tab
    - [] `\uXXXX` Unicode character
- [] Support full-fledged number notation
    - [] [+-][0-9]+[.][0-9]+[eE][+-][0-9]+ (e.g. -123.456e+789)
- [] Support comments