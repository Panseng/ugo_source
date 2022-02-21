package parser

import (
	"github.com/wa-lang/ugo/ast"
	"github.com/wa-lang/ugo/token"
)

// parseFile 由词法数组，转换为语法树
func (p *Parser) parseFile() {
	p.file = &ast.File{
		Filename: p.Filename(),
		Source:   p.Source(),
	}

	// package xxx
	p.file.Pkg = p.parsePackage() // 解析 引入的依赖 pos会前移

	for {
		switch tok := p.PeekToken(); tok.Type { // 获取下一个token，注意，PeekToken 读取 token 后，会回退
		case token.EOF:
			return
		case token.ERROR:
			panic(tok)
		case token.SEMICOLON:
			p.AcceptTokenList(token.SEMICOLON)

		case token.VAR:
			p.file.Globals = append(p.file.Globals, p.parseStmt_var())
		case token.FUNC:
			p.file.Funcs = append(p.file.Funcs, p.parseFunc())

		default:
			p.errorf(tok.Pos, "unknown token: %v", tok)
		}
	}
}

func (p *Parser) parsePackage() *ast.PackageSpec {
	tokPkg := p.MustAcceptToken(token.PACKAGE)
	tokPkgIdent := p.MustAcceptToken(token.IDENT)

	return &ast.PackageSpec{
		PkgPos:  tokPkg.Pos,
		NamePos: tokPkgIdent.Pos,
		Name:    tokPkgIdent.Literal,
	}
}
