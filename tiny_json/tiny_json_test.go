package tiny_json

import (
	"testing"
)

func runTests(t *testing.T, input string, tests []struct {
	expectedType    TokenType
	expectedLiteral string
}) {
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

func TestEmptyInput(t *testing.T) {
	input := ``

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{EOF, ""},
	}

	runTests(t, input, tests)
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

	runTests(t, input, tests)
}

func TestSimpleJSON(t *testing.T) {
	input := `{
		"key": "value",
		"key1": "value",
		"key2": 22222,
		"key_3": 33333,
		"key-4": 44444,
		"11111": 11111,
		"a1b2c3d4000000": true,
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
		{COMMA, ","},

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

	runTests(t, input, tests)
}

func TestMultiKeysJSONWithArray(t *testing.T) {
	input := `{
		"key": "value",
		"key-n": 101,
		"key-o": {},
		"key-l": [],
		"nestedObj": {
			"innerKey1": "innerValue1",
			"innerKey2": 55555,
			"innerArray": [1, 2, 3, {"deepKey": "deepValue"}]
		}
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
		{COMMA, ","},

		// .....
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "nestedObj"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{LBRACE, "{"},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "innerKey1"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "innerValue1"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "innerKey2"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{INT, "55555"},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "innerArray"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{LBRACKET, "["},
		{INT, "1"},
		{COMMA, ","},
		{INT, "2"},
		{COMMA, ","},
		{INT, "3"},
		{COMMA, ","},
		{LBRACE, "{"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "deepKey"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "deepValue"},
		{DOUBLEQUOTE, "\""},

		{RBRACE, "}"},
		{RBRACKET, "]"},
		{RBRACE, "}"},

		{RBRACE, "}"},
	}

	runTests(t, input, tests)
}

func TestEscapeChars(t *testing.T) {
	input := `{
		"quotation-mark": "\" value and \b new value",
		"backslash": "\\ value",
		"solidus": "\/ value",
		"backspace": "\b value",
		"form-feed": "\f value",
		"line-feed": "\n value",
		"carriage-return": "\r value",
		"horizontal-tab": "\t value"
	}
	`

	// a slice of anonymous structs
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "quotation-mark"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\\" value and \\b new value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "backslash"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\\\ value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "solidus"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\/ value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "backspace"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\b value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "form-feed"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\f value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "line-feed"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\n value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "carriage-return"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\r value"},
		{DOUBLEQUOTE, "\""},
		{COMMA, ","},

		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "horizontal-tab"},
		{DOUBLEQUOTE, "\""},
		{COLON, ":"},
		{DOUBLEQUOTE, "\""},
		{IDENTIFIER, "\\t value"},
		{DOUBLEQUOTE, "\""},

		{RBRACE, "}"},
	}

	runTests(t, input, tests)
}
