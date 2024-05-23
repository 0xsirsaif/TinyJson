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
			t.Fatalf("test_samples[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test_samples[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
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
			t.Fatalf("test_samples[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test_samples[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestSimpleJSON(t *testing.T) {
	input := `{
		"key": "value",
		"key1": "value",
		"key2": 22222,
		"key_3": 33333,
		"key-4": 44444,
		"11111": 11111,
		"a1b2c3d4000000": true
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
		{IDENTIFIER, "key1"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key2"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "22222"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key_3"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "33333"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "key-4"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "44444"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "11111"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "11111"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "a1b2c3d4000000"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{TRUE, "true"},

		{RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("test_samples[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test_samples[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
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
			t.Fatalf("test_samples[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test_samples[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
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
		{COMMA, ","},

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
			t.Fatalf("test_samples[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test_samples[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
