package lexer

import (
	"regexp"
	"seville/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination

}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	var offset int

	if l.readPosition >= len(l.input) {
		l.ch = 0
		offset = 1
	} else {
		r, width := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		offset = width
	}
	l.position = l.readPosition
	l.readPosition += offset
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.LT_OR_EQ, Literal: "<="}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.GT_OR_EQ, Literal: ">="}
		} else {
			tok = newToken(token.GT, l.ch)
		} 
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peakChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return r
	}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || isEmoji(ch)
}

func isEmoji(ch rune) bool {
	emojisAndExtendedPictographicPattern := `
	<a?:.+?:\d{18}>|[\x{1F300}-\x{1F9FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{1F1E6}-\x{1F1FF}]`
	emojiRegex := regexp.MustCompile(emojisAndExtendedPictographicPattern)
	return emojiRegex.MatchString(string(ch))
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
