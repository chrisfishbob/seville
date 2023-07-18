package lexer

import (
	"seville/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input :=
		`
	let five = 5;
	let ten = 10;
	
	let add = fn(🌮, y) {
		x + y;
	};
	
	let result = add(five, ten);
	!-*23;
	5 < 123 > 5;

	if (5 < 10) {
		return true;
	} elif (true) {
		return 1
	} else {
		return false;
	}
	21⚽ == 21;
	10 != 23;
	🌹🎶 🇪🇸 
	🍇 =🍇 2🍇
	23 + 🍇 = 1
	1 + 1 == 2
	1 <= 2
	2 >= 12
	5 + 5 ** 12
	"foobar"
	"foo if bar"
	[1, "2"];
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "🌮"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.INT, "23"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "123"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELIF, "elif"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "21"},
		{token.IDENT, "⚽"},
		{token.EQ, "=="},
		{token.INT, "21"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "23"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "🌹🎶"},
		{token.IDENT, "🇪🇸"},
		{token.IDENT, "🍇"},
		{token.ASSIGN, "="},
		{token.IDENT, "🍇"},
		{token.INT, "2"},
		{token.IDENT, "🍇"},
		{token.INT, "23"},
		{token.PLUS, "+"},
		{token.IDENT, "🍇"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "2"},
		{token.INT, "1"},
		{token.LT_OR_EQ, "<="},
		{token.INT, "2"},
		{token.INT, "2"},
		{token.GT_OR_EQ, ">="},
		{token.INT, "12"},
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.INT, "5"},
		{token.EXP, "**"},
		{token.INT, "12"},
		{token.STRING, "foobar"},
		{token.STRING, "foo if bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.STRING, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
