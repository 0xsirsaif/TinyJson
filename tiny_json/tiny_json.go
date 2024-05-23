package tiny_json

const (
	LBRACE = "{"
	RBRACE = "}"

	INT = "INT"

	DOUBLEQUOTE = "\""

	IDENTIFIER = "IDENT"

	COLON = ":"

	LBRACKET = "["
	RBRACKET = "]"

	TRUE  = "true"
	FALSE = "false"
	NULL  = "null"

	COMMA = ","

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var Keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

type Lexer struct {
	input        string
	ch           byte
	position     int
	nextPosition int
	inString     bool
}

func NewLexer(input string) *Lexer {
	l := Lexer{input: input}
	l.readCharacter()
	return &l
}

func (lexer *Lexer) NextToken() Token {
	var token Token

	lexer.isWhiteSpace()

	switch lexer.ch {
	case '{':
		token = newToken(LBRACE, '{')
	case '}':
		token = newToken(RBRACE, '}')
	case '"':
		token = newToken(DOUBLEQUOTE, '"')
		lexer.inString = !lexer.inString
	case ':':
		token = newToken(COLON, ':')
	case '[':
		token = newToken(LBRACKET, '[')
	case ']':
		token = newToken(RBRACKET, ']')
	case ',':
		token = newToken(COMMA, ',')
	case 0:
		// ASCII NULL
		token.Literal = ""
		token.Type = EOF
	default:
		if isLetter(lexer.ch) {
			if isDigit(lexer.ch) {
				if !lexer.inString {
					token.Literal = lexer.readNumber()
					token.Type = INT
					return token
				}
			}
			token.Literal = lexer.readIdentifier()
			token.Type = lookupIdentifier(token.Literal)
			return token
		} else if isDigit(lexer.ch) {
			token.Literal = lexer.readNumber()
			token.Type = INT
		} else {
			token = newToken(ILLEGAL, lexer.ch)
		}
	}
	lexer.readCharacter()
	return token
}

func (lexer *Lexer) readCharacter() {
	if lexer.nextPosition >= len(lexer.input) {
		// ASCII NUL
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.nextPosition]
	}
	lexer.position = lexer.nextPosition
	lexer.nextPosition++
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.position
	for isDigit(lexer.ch) {
		lexer.readCharacter()
	}
	return lexer.input[startPosition:lexer.position]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readCharacter()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		('0' <= ch && ch <= '9') ||
		(ch == '_') ||
		(ch == '-')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) isWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readCharacter()
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func lookupIdentifier(identString string) TokenType {
	if tokenType, ok := Keywords[identString]; ok {
		return tokenType
	}
	return IDENTIFIER
}
