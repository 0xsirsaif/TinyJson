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
- [] Parser
- [] Support Unicode characters
- [] Parse Numbers in Keys, e.g. `{"key1": 1, "key2": 2}`
- [] Escape characters in strings
    - [] `\"` quotation mark
    - [] `\\` backslash
    - [] `\/` solidus
    - [] `\b` backspace
    - [] `\f` form feed
    - [] `\n` line feed
    - [] `\r` carriage return
    - [] `\t` horizontal tab
    - [] `\uXXXX` Unicode character
- [] Support full-fledged number notation
    - [] [+-][0-9]+[.][0-9]+[eE][+-][0-9]+ (e.g. -123.456e+789)
- [] Support comments