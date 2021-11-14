package parser

import (
	"fmt"

	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/logger"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseExpr() ast.Expr {
	logger.Debugln("peek =", p.peekToken())

	expr := p.parseExpr_mul()
	for {
		switch p.peekTokenType() {
		case token.ADD, token.SUB:
			tok := p.nextToken()
			expr = &ast.BinaryExpr{
				X:  expr,
				Op: tok,
				Y:  p.parseExpr_mul(),
			}
		default:
			return expr
		}
	}
}

func (p *parser) parseExpr_mul() ast.Expr {
	expr := p.parseExpr_unary()
	for {
		switch p.peekTokenType() {
		case token.MUL, token.QUO:
			tok := p.nextToken()
			expr = &ast.BinaryExpr{
				X:  expr,
				Op: tok,
				Y:  p.parseExpr_unary(),
			}
		default:
			return expr
		}
	}
}

func (p *parser) parseExpr_unary() ast.Expr {
	if _, ok := p.acceptToken(token.ADD); ok {
		return p.parseExpr_primary()
	}
	if _, ok := p.acceptToken(token.SUB); ok {
		return &ast.UnaryExpr{
			X: p.parseExpr_primary(),
		}
	}
	return p.parseExpr_primary()
}

func (p *parser) parseExpr_primary() ast.Expr {
	peek := p.peekToken()

	logger.Debugf("peek = %v\n", peek)

	switch peek.Type {
	case token.IDENT:
		ident := p.nextToken()
		if lparen, ok := p.acceptToken(token.LPAREN); ok {
			var args []ast.Expr
			for {
				if rparen, ok := p.acceptToken(token.RPAREN); ok {
					return &ast.CallExpr{
						Fun: &ast.Ident{
							NamePos: ident.Pos,
							Name:    ident.IdentName(),
						},
						Lparen: lparen.Pos,
						Args:   args,
						Rparen: rparen.Pos,
					}
				}
				args = append(args, p.parseExpr())
				p.acceptToken(token.COMMA)
			}
		}
		return &ast.Ident{
			NamePos: ident.Pos,
			Name:    ident.IdentName(),
		}
	case token.INT:
		tok := p.nextToken()
		return &ast.Number{
			ValuePos: tok.Pos,
			Value:    tok.IntValue(),
			ValueEnd: tok.EndPos(),
		}
	case token.FLOAT:
		tok := p.nextToken()
		return &ast.Number{
			ValuePos: tok.Pos,
			Value:    tok.FloatValue(),
			ValueEnd: tok.EndPos(),
		}

	case token.LPAREN:
		p.nextToken()
		expr := p.parseExpr()
		if _, ok := p.acceptToken(token.RPAREN); !ok {
			p.err = fmt.Errorf("todo")
			panic(p.err)
		}
		p.nextToken()
		return expr
	default:
		p.errorf("todo: peek=%v", peek)
		panic(p.err)
	}
}
