package ast

type BlockStatement struct {
	Body []Statement
}

func (s *BlockStatement) stmt() {}

type ExpressionStatement struct {
	Expr Expression
}

func (s *ExpressionStatement) stmt() {}
