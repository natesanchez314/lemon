package lexer

import "banana/token"

type Lexer struct {
	input string
	position int
	readPosition int
	currentChar byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}


func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()
	
	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '*':
		tok = newToken(token.MULT, l.currentChar)
	case '/':
		tok = newToken(token.DIV, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '<':
		tok = newToken(token.LT, l.currentChar)
	case '>':
		tok = newToken(token.GT, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case '[':
		tok = newToken(token.LBRACKET, l.currentChar)
	case ']':
		tok = newToken(token.RBRACKET, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case ':':
		tok = newToken(token.COLON, l.currentChar)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpId(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, currentChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currentChar)}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.currentChar == '"' || l.currentChar == 0 {
			break
		}
	}
	return l.input[position: l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position: l.position]
}

func isLetter(currentChar byte) bool {
	return 'a' <= currentChar && currentChar <= 'z' || 'A' <= currentChar && currentChar <= 'Z' || currentChar == '_'
}

func isDigit(currentChar byte) bool {
	return '0' <= currentChar && currentChar <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position: l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	} 
}