package tiny_json

import (
	"testing"
)

func TestEmptyInput(t *testing.T) {
	input := ``

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestEmptyJSON(t *testing.T) {
	input := `{}`

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},
		{RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestSimpleJSON(t *testing.T) {
	input := `{"key": "value"}`

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "value"},
		{DOUBLEQUOTE, "\""},
		{RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestMultiKeysJSON(t *testing.T) {
	input := `{
  		"key_one": true,
  		"key_two": false,
  		"key_three": null,
  		"key_four": "value",
  		"key_five": 101
	}`

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_one"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{TRUE, "true"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_two"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{FALSE, "false"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_three"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{NULL, "null"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_four"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_five"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "101"},

		{RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestMultiKeysJSONWithArray(t *testing.T) {
	input := `{
		"key": "value",
  		"key-n": 101,
  		"key-o": {},
  		"key-l": []
	}`

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key-n"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "101"},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key-o"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key-l"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{LBRACKET, "["},
		{RBRACKET, "]"},

		{RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
