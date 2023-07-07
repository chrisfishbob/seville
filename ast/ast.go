package ast

import (
	"seville/token"
)

type Node interface {
	// TokenLiteral only used for debugging and testing
	TokenLiteral() string
}

type Statement interface {
	Node
	// Dummy methods to help the Go compiler
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (l *Identifier) expressionNode()      {}
func (l *Identifier) TokenLiteral() string { return l.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // The token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
