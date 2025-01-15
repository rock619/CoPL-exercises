package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/typingml4/parser"
)

type Visitor struct {
	*parser.BaseTypingML4Visitor
	p   *parser.TypingML4Parser
	env *Environment
	l   *slog.Logger
	typ Type
}

func NewVisitor(p *parser.TypingML4Parser, l *slog.Logger) *Visitor {
	return &Visitor{
		p:   p,
		env: NewEnvironment(l),
		l:   l,
	}
}

func (v *Visitor) Do(tree antlr.ParseTree) (Expr, error) {
	expr, ok := tree.Accept(v).(Expr)
	if !ok {
		return nil, fmt.Errorf("tree.Accept(v) is not Expr")
	}
	return expr, nil
}

func (v *Visitor) VisitJudgement(c *parser.JudgementContext) interface{} {
	v.l.Info("VisitJudgement", "literal", getLiteral(v.p, c))

	if c.Env() != nil {
		c.Env().Accept(v)
	}

	typ, ok := c.Type_().Accept(v).(Type)
	if !ok {
		v.l.Error("VisitJudgement: type is not Type", "typ", fmt.Sprintf("%T", typ), "type", getLiteral(v.p, c.Type_()))
		return nil
	}
	defer v.setType(typ)()

	expr, ok := c.Expr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitJudgement: expr is not Expr", "expr", getLiteral(v.p, c.Expr()))
		return nil
	}
	return expr
}

func (v *Visitor) VisitEnv(c *parser.EnvContext) interface{} {
	v.l.Info("VisitEnv", "literal", getLiteral(v.p, c))
	for _, bind := range c.AllBind() {
		bind.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitBind(c *parser.BindContext) interface{} {
	v.l.Info("VisitBind", "literal", getLiteral(v.p, c))
	typ, ok := c.Type_().Accept(v).(Type)
	if !ok {
		v.l.Error("VisitBind: type is not Type", "literal", getLiteral(v.p, c.Type_()))
		return nil
	}
	v.env.AddBind(c.IDENTIFIER().GetText(), typ)
	return nil
}

func (v *Visitor) VisitParenType(c *parser.ParenTypeContext) interface{} {
	v.l.Info("VisitParenType", "literal", getLiteral(v.p, c))
	return c.Type_().Accept(v)
}

func (v *Visitor) VisitIntType(c *parser.IntTypeContext) interface{} {
	return &IntType{}
}

func (v *Visitor) VisitBoolType(c *parser.BoolTypeContext) interface{} {
	return &BoolType{}
}

func (v *Visitor) VisitFunType(c *parser.FunTypeContext) interface{} {
	v.l.Info("VisitFunType", "literal", getLiteral(v.p, c), "param", getLiteral(v.p, c.GetParamType()), "return", getLiteral(v.p, c.GetReturnType()))
	paramType, ok := c.GetParamType().Accept(v).(Type)
	if !ok {
		v.l.Error("VisitFunType: paramType is not Type", "literal", getLiteral(v.p, c.GetParamType()))
		return nil
	}
	retType, ok := c.GetReturnType().Accept(v).(Type)
	if !ok {
		v.l.Error("VisitFunType: returnType is not Type", "literal", getLiteral(v.p, c.GetReturnType()))
		return nil
	}
	return &FunType{
		ParamType:  paramType,
		ReturnType: retType,
	}
}

func (v *Visitor) VisitListType(c *parser.ListTypeContext) interface{} {
	v.l.Info("VisitListType", "literal", getLiteral(v.p, c))
	elemType, ok := c.GetElementType().Accept(v).(Type)
	if !ok {
		v.l.Error("VisitListType: elementType is not Type", "literal", getLiteral(v.p, c.GetElementType()))
		return nil
	}
	return &ListType{
		ElemType: elemType,
	}
}

func (v *Visitor) VisitParenExpr(c *parser.ParenExprContext) interface{} {
	v.logExprContext(c)

	return c.Expr().Accept(v)
}

// VisitFunExpr fun x -> e
//
//	Γ |- fun x -> e : τ1 -> τ2 by T-Fun {
//	  Γ, x : τ1 |- e : τ2
//	}
func (v *Visitor) VisitFunExpr(c *parser.FunExprContext) interface{} {
	v.logExprContext(c)

	expr := FunExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     &NilType{},
		},
		Fun: Fun{
			Param:       c.Fun().IDENTIFIER().GetText(),
			BodyLiteral: getLiteral(v.p, c.Fun().GetBody()),
			BodyCtx:     c.Fun().GetBody(),
		},
	}
	defer v.logExpr(expr, "VisitFunExpr: exit")()

	// 関数の引数と戻り値の型を型変数として置く
	v.env.AddScope()
	defer v.env.RemoveScope()
	paramType := v.env.NewTypeVar()
	v.l.Info("VisitFunExpr: add TypeVar of param", "param", expr.Fun.Param, "typeVar", paramType)
	if givenParamType, ok := v.env.GetType(expr.Fun.Param); ok {
		v.l.Info("VisitFunExpr: AddConstraint", "left", paramType, "right", givenParamType)
		v.env.AddConstraint(paramType, givenParamType)
	}
	v.l.Info("VisitFunExpr: AddBind", "name", expr.Fun.Param, "type", paramType)
	v.env.AddBind(expr.Fun.Param, paramType)
	returnType := v.env.NewTypeVar()
	v.l.Info("VisitFunExpr: add TypeVar of return", "typeVar", returnType)
	expr.typ = &FunType{
		ParamType:  paramType,
		ReturnType: returnType,
	}

	// 関数の型が与えられている場合はそれを使う
	if givenType, ok := v.getType().(*FunType); ok {
		v.l.Info("givenType is FunType", "givenType", givenType, "param", expr.Fun.Param, "paramType", paramType,
			"returnType.Index", returnType.ID, "returnType", returnType)
		v.l.Info("VisitFunExpr: AddConstraint", "left", paramType, "right", givenType.ParamType)
		v.env.AddConstraint(paramType, givenType.ParamType)
		v.l.Info("VisitFunExpr: AddConstraint", "left", returnType, "right", givenType.ReturnType)
		v.env.AddConstraint(returnType, givenType.ReturnType)
		defer v.setType(givenType.ReturnType)()
	} else {
		v.l.Info("VisitFunExpr: givenType is not FunType", "type", v.getType())
		v.l.Info("VisitFunExpr: AddConstraint", "expr.typ", expr.typ, "v.getType()", v.getType())
		v.env.AddConstraint(expr.typ, v.getType())
		defer v.setType(returnType)()
	}

	v.l.Info("VisitFunExpr: visit body expr", "literal", getLiteral(v.p, c.Fun().GetBody()))
	body, ok := c.Fun().GetBody().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitFunExpr: body is not Expr", "literal", getLiteral(v.p, c.Fun().GetBody()))
		return nil
	}
	v.l.Info("VisitFunExpr: visited body expr", "literal", getLiteral(v.p, c.Fun().GetBody()), "body.type", body.Type())
	v.l.Info("VisitFunExpr: AddConstraint", "left", returnType, "right", body.Type())
	v.env.AddConstraint(returnType, body.Type())
	expr.Fun.Body = body
	expr.AddChildren(body)
	expr.rule = TFun
	return expr
}

func (v *Visitor) VisitAppExpr(c *parser.AppExprContext) interface{} {
	v.logExprContext(c)
	v.l.Info("VisitAppExpr", "literal", getLiteral(v.p, c), "fun", getLiteral(v.p, c.GetFn()), "arg", getLiteral(v.p, c.GetArg()), "givenType", v.getType())

	// 引数
	argType := v.env.NewTypeVar()
	v.l.Info("VisitAppExpr: add TypeVar of arg", "typeVar", argType)
	reset := v.setType(argType)
	v.l.Info("VisitAppExpr: visit arg expr", "literal", getLiteral(v.p, c.GetArg()))
	arg, ok := c.GetArg().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitAppExpr: arg is not Expr", "literal", getLiteral(v.p, c.GetArg()))
		return nil
	}
	reset()
	v.l.Info("VisitAppExpr: AddConstraint", "left", argType, "right", arg.Type())
	v.env.AddConstraint(argType, arg.Type())

	returnType := v.env.NewTypeVar()
	v.l.Info("VisitAppExpr: add TypeVar of return", "typeVar", returnType)
	v.l.Info("VisitAppExpr: AddConstraint", "left", returnType, "right", v.getType())
	v.env.AddConstraint(returnType, v.getType())
	funType := &FunType{
		ParamType:  argType,
		ReturnType: returnType,
	}
	reset = v.setType(funType)
	// 関数
	fun, ok := c.GetFn().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitAppExpr: fun is not Expr", "literal", getLiteral(v.p, c.GetFn()))
		return nil
	}
	reset()
	defer v.logExpr(fun, "VisitAppExpr: Fun")()

	expr := AppExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     returnType,
			rule:    TApp,
		},
		Fun: fun,
		Arg: arg,
	}
	expr.AddChildren(fun, arg)
	defer v.logExpr(expr, "exit: VisitAppExpr")()
	return expr
}

func (v *Visitor) VisitMatchExpr(c *parser.MatchExprContext) interface{} {
	v.logExprContext(c)

	v.l.Error("VisitMatchExpr: not implemented", "literal", getLiteral(v.p, c))
	return nil
}

func (v *Visitor) VisitEmptyListExpr(c *parser.EmptyListExprContext) interface{} {
	v.logExprContext(c)

	return EmptyListExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     v.getType(),
			rule:    TNil,
		},
	}
}

func (v *Visitor) VisitConsExpr(c *parser.ConsExprContext) interface{} {
	v.logExprContext(c)

	typ := &ListType{ElemType: v.env.NewTypeVar()}
	v.l.Info("VisitConsExpr: add TypeVar of elem", "typeVar", typ.ElemType)
	v.l.Info("VisitConsExpr: AddConstraint", "left", typ, "right", v.getType())
	v.env.AddConstraint(typ, v.getType())

	reset := v.setType(typ.ElemType)
	head, ok := c.GetHead().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitConsExpr: head is not Expr", "literal", getLiteral(v.p, c.GetHead()))
		return nil
	}
	reset()
	v.l.Info("VisitConsExpr: AddConstraint", "left", typ.ElemType, "right", head.Type())
	v.env.AddConstraint(typ.ElemType, head.Type())

	defer v.setType(typ)()
	tail, ok := c.GetTail().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitConsExpr: tail is not Expr", "literal", getLiteral(v.p, c.GetTail()))
		return nil
	}
	v.l.Info("VisitConsExpr: AddConstraint", "left", typ, "right", tail.Type())
	v.env.AddConstraint(typ, tail.Type())

	expr := ConsExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     typ,
			rule:    TCons,
		},
		Head: head,
		Tail: tail,
	}
	expr.AddChildren(head, tail)
	return expr
}

// VisitBinOpExpr 2項演算 e (+ | - | * | <) e op
//
//	Γ |- e1 (+ | - | *) e2 : int by (T-Plus | T-Minus | T-Times) {
//	  Γ |- e1 : int
//	  Γ |- e2 : int
//	}
//	or
//	Γ |- e1 < e2 : bool by T-Lt {
//	  Γ |- e1 : int
//	  Γ |- e2 : int
//	}
func (v *Visitor) VisitBinOpExpr(c *parser.BinOpExprContext) interface{} {
	v.logExprContext(c)

	// 2項演算子の左右は必ずintに確定する
	reset := v.setType(&IntType{})
	v.l.Info("VisitBinOpExpr: visit left expr", "literal", getLiteral(v.p, c.GetLeft()))
	left, ok := c.GetLeft().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitBinOpExpr: left is not Expr", "left", getLiteral(v.p, c.GetLeft()))
		return nil
	}
	v.l.Info("VisitBinOpExpr: visit right expr", "literal", getLiteral(v.p, c.GetRight()))
	right, ok := c.GetRight().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitBinOpExpr: right is not Expr", "right", getLiteral(v.p, c.GetRight()))
		return nil
	}
	reset()

	expr := BinOpExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
		},
		Left:  left,
		Op:    c.GetOp(),
		Right: right,
	}
	expr.AddChildren(left, right)
	defer v.logExpr(expr, "VisitBinOpExpr: exit")()

	switch expr.Op.GetTokenType() {
	case parser.TypingML4ParserPLUS:
		expr.typ = &IntType{}
		expr.rule = TPlus
	case parser.TypingML4ParserMINUS:
		expr.typ = &IntType{}
		expr.rule = TMinus
	case parser.TypingML4ParserTIMES:
		expr.typ = &IntType{}
		expr.rule = TTimes
	case parser.TypingML4ParserLT:
		// 2項演算子のうち '<' だけは型がboolになる
		expr.typ = &BoolType{}
		expr.rule = TLt
	default:
		v.l.Error("VisitBinOpExpr: unknown operator", "literal", getLiteral(v.p, c), "op", expr.Op.GetText())
		return nil
	}
	return expr
}

func (v *Visitor) VisitIfExpr(c *parser.IfExprContext) interface{} {
	v.logExprContext(c)

	v.env.AddScope()
	reset := v.setType(&BoolType{})
	v.l.Info("VisitIfExpr: visit cond expr", "literal", getLiteral(v.p, c.GetCond()))
	cond, ok := c.GetCond().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitIfExpr: cond is not Expr", "cond", getLiteral(v.p, c.GetCond()))
		return nil
	}
	v.logExpr(cond, "VisitIfExpr: cond")()
	reset()
	v.env.RemoveScope()

	v.env.AddScope()
	v.l.Info("VisitIfExpr: visit then expr", "literal", getLiteral(v.p, c.GetThen()))
	then, ok := c.GetThen().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitIfExpr: then is not Expr", "then", getLiteral(v.p, c.GetThen()))
		return nil
	}
	v.env.RemoveScope()
	v.env.AddScope()
	v.l.Info("VisitIfExpr: visit else expr", "literal", getLiteral(v.p, c.GetElse_()))
	els, ok := c.GetElse_().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitIfExpr: else is not Expr", "else", getLiteral(v.p, c.GetElse_()))
		return nil
	}
	v.env.RemoveScope()
	expr := IfExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     then.Type(),
		},
		Cond: cond,
		Then: then,
		Else: els,
	}
	v.logExpr(cond, "VisitIfExpr: cond")()
	v.logExpr(then, "VisitIfExpr: then")()
	v.logExpr(els, "VisitIfExpr: else")()
	defer v.logExpr(expr, "VisitIfExpr: exit")()

	expr.rule = TIf
	expr.AddChildren(cond, then, els)
	return expr
}

func (v *Visitor) VisitLetExpr(c *parser.LetExprContext) interface{} {
	v.logExprContext(c)
	v.l.Info("VisitLetExpr", "bindExpr", getLiteral(v.p, c.GetBindExpr()), "body", getLiteral(v.p, c.GetBody()), "givenType", v.getType())

	bindType := v.env.NewTypeVar()
	v.l.Info("VisitLetExpr: add TypeVar of bind", "typeVar", bindType)
	reset := v.setType(bindType)
	v.l.Info("VisitLetExpr: visit bindExpr", "bindExpr", getLiteral(v.p, c.GetBindExpr()))
	bindExpr, ok := c.GetBindExpr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitLetExpr: bindExpr is not Expr", "bindExpr", getLiteral(v.p, c.GetBindExpr()))
		return nil
	}
	reset()

	expr := LetExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     v.env.NewTypeVar(),
		},
		Var:  c.IDENTIFIER().GetText(),
		Bind: bindExpr,
	}
	v.l.Info("VisitLetExpr: add typeVar of let body", "body", getLiteral(v.p, c.GetBody()), "typeVar", expr.typ)

	v.env.AddScope()
	defer v.env.RemoveScope()
	v.l.Info("VisitLetExpr: add var to env", "var", expr.Var, "type", fmt.Sprintf("%#v", bindExpr.Type()))
	v.env.AddBind(expr.Var, bindExpr.Type())

	v.l.Info("VisitLetExpr: AddConstraint", "left", expr.typ, "right", v.getType())
	v.env.AddConstraint(expr.typ, v.getType())
	v.l.Info("VisitLetExpr: visit body", "body", getLiteral(v.p, c.GetBody()), "givenType", v.getType())
	body := c.GetBody().Accept(v)
	bodyExpr, ok := body.(Expr)
	if !ok {
		v.l.Error("VisitLetExpr: bodyExpr is not Expr", "bodyExpr", getLiteral(v.p, c.GetBody()), "body", body)
		return nil
	}
	expr.Body = bodyExpr
	v.l.Info("VisitLetExpr: AddConstraint", "left", expr.typ, "right", bodyExpr.Type())
	v.env.AddConstraint(expr.typ, bodyExpr.Type())
	expr.rule = TLet
	expr.AddChildren(bindExpr, bodyExpr)
	return expr
}

func (v *Visitor) VisitLetRecExpr(c *parser.LetRecExprContext) interface{} {
	v.logExprContext(c)

	v.l.Error("VisitLetRecExpr: not implemented", "literal", getLiteral(v.p, c))
	return nil
}

func (v *Visitor) VisitIntExpr(c *parser.IntExprContext) interface{} {
	v.logExprContext(c)

	return IntExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     &IntType{},
			rule:    TInt,
		},
	}
}

func (v *Visitor) VisitBoolExpr(c *parser.BoolExprContext) interface{} {
	v.logExprContext(c)

	return BoolExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			typ:     &BoolType{},
			rule:    TBool,
		},
	}
}

func (v *Visitor) VisitVarExpr(c *parser.VarExprContext) interface{} {
	v.logExprContext(c)

	expr := VarExpr{
		BaseExpr: &BaseExpr{
			env:     v.env.Clone(),
			literal: getLiteral(v.p, c),
			rule:    TVar,
		},
		Name: c.IDENTIFIER().GetText(),
	}
	defer v.logExpr(expr, "VisitVarExpr: exit")()

	expr.BaseExpr.typ = v.env.NewTypeVar()

	if t, ok := v.env.GetType(expr.Name); ok {
		v.l.Info("VisitVarExpr: AddConstraint", "left", expr.BaseExpr.typ, "right", t)
		v.env.AddConstraint(expr.BaseExpr.typ, t)
	}
	v.l.Info("VisitVarExpr: AddConstraint", "left", expr.BaseExpr.typ, "right", v.getType())
	v.env.AddConstraint(expr.BaseExpr.typ, v.getType())

	return expr
}

func (v *Visitor) getType() Type {
	return v.typ
}

func (v *Visitor) setType(t Type) (reset func()) {
	prev := v.typ
	v.typ = t
	return func() {
		v.typ = prev
	}
}

func (v *Visitor) logExprContext(c parser.IExprContext, args ...any) {
	v.l.Info(fmt.Sprintf("%T", c),
		append([]any{
			"literal", getLiteral(c.GetParser(), c.GetRuleContext()),
			"pos", getInterval(c),
			"env", v.env,
			"type", v.getType(),
		}, args...)...,
	)
}

func (v *Visitor) logExpr(e Expr, msg string) func() {
	return func() {
		v.l.Info(msg, "literal", e.Literal(), "env", e.Env(), "type", e.Type(), "rule", e.Rule())
	}
}

func getLiteral(p antlr.Parser, st antlr.SyntaxTree) string {
	text := p.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}

func getInterval(st antlr.SyntaxTree) string {
	i := st.GetSourceInterval()
	return fmt.Sprintf("[%d:%d]", i.Start, i.Stop)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
