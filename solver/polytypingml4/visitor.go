package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/polytypingml4/parser"
)

type Visitor struct {
	*parser.BasePolyTypingML4Visitor
	p *parser.PolyTypingML4Parser
	l *slog.Logger
}

func NewVisitor(p *parser.PolyTypingML4Parser, l *slog.Logger) *Visitor {
	return &Visitor{
		p: p,
		l: l,
	}
}

func (v *Visitor) Do() (Judgement, error) {
	tree := v.p.Eval()
	if tree == nil {
		return Judgement{}, fmt.Errorf("Visitor.Do: tree is nil: %q", v.LiteralOf(tree))
	}
	return AssertResult[Judgement](tree.Accept(v)).Unwrap()
}

func (v *Visitor) VisitEval(c *parser.EvalContext) any {
	v.l.Debug("VisitEval", "literal", v.LiteralOf(c))

	envRes := AssertResult[Env](c.Env().Accept(v))
	if envRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: visit env: %w", envRes.Err()))
	}

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: visit exp: %w", expRes.Err()))
	}

	typeRes := AssertResult[Type](c.Type_().Accept(v))
	if typeRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: visit type: %w", typeRes.Err()))
	}

	return OK(Judgement{
		Env:  envRes.Val(),
		Exp:  expRes.Val(),
		Type: typeRes.Val(),
	})
}

func (v *Visitor) VisitEnv(c *parser.EnvContext) any {
	v.l.Debug("VisitEnv", "literal", v.LiteralOf(c))

	var env Env
	for _, b := range c.AllBind() {
		bindRes := AssertResult[Bind](b.Accept(v))
		if bindRes.Err() != nil {
			return Err[Env](fmt.Errorf("VisitEnv: visit bind: %w", bindRes.Err()))
		}
		env = append(env, bindRes.Val())
	}
	return OK(env)
}

func (v *Visitor) VisitBind(c *parser.BindContext) any {
	v.l.Debug("VisitBind", "literal", v.LiteralOf(c))

	typeSchemeRes := AssertResult[TypeScheme](c.TyScheme().Accept(v))
	if typeSchemeRes.Err() != nil {
		return Err[Bind](fmt.Errorf("VisitBind: visit type scheme: %w", typeSchemeRes.Err()))
	}

	return OK(Bind{
		Var:        Var(c.Var_().ID().GetText()),
		TypeScheme: typeSchemeRes.Val(),
	})
}

func (v *Visitor) VisitTyScheme(c *parser.TySchemeContext) any {
	v.l.Debug("VisitTyScheme", "literal", v.LiteralOf(c))

	var typeVars []TypeVar
	for _, tv := range c.AllTVar() {
		typeVarRes := AssertResult[TypeVar](tv.Accept(v))
		if typeVarRes.Err() != nil {
			return Err[TypeScheme](fmt.Errorf("VisitTyScheme: visit type var: %w", typeVarRes.Err()))
		}
		typeVars = append(typeVars, typeVarRes.Val())
	}

	typeRes := AssertResult[Type](c.Type_().Accept(v))
	if typeRes.Err() != nil {
		return Err[TypeScheme](fmt.Errorf("VisitTyScheme: visit type: %w", typeRes.Err()))
	}
	return OK(NewTypeScheme(typeVars, typeRes.Val()))
}

func (v *Visitor) VisitTVar(c *parser.TVarContext) any {
	v.l.Debug("VisitTVar", "literal", v.LiteralOf(c))

	tv := TypeVar(c.ID().GetText())
	return OK(tv)
}

func (v *Visitor) VisitTypeVar(c *parser.TypeVarContext) any {
	v.l.Debug("VisitTypeVar", "literal", v.LiteralOf(c))

	typeVarRes := AssertResult[TypeVar](c.TVar().Accept(v))
	if typeVarRes.Err() != nil {
		return Err[Type](fmt.Errorf("VisitTypeVar: visit type var: %w", typeVarRes.Err()))
	}
	return OK[Type](TypeVar(typeVarRes.Val()))
}

func (v *Visitor) VisitBoolType(c *parser.BoolTypeContext) any {
	v.l.Debug("VisitBoolType", "literal", v.LiteralOf(c))

	return OK[Type](BoolType{})
}

func (v *Visitor) VisitIntType(c *parser.IntTypeContext) any {
	v.l.Debug("VisitIntType", "literal", v.LiteralOf(c))

	return OK[Type](IntType{})
}

func (v *Visitor) VisitListType(c *parser.ListTypeContext) any {
	v.l.Debug("VisitListType", "literal", v.LiteralOf(c))

	elemRes := AssertResult[Type](c.Type_().Accept(v))
	if elemRes.Err() != nil {
		return Err[Type](fmt.Errorf("VisitListType: visit element type: %w", elemRes.Err()))
	}
	return OK[Type](ListType{
		Elem: elemRes.Val(),
	})
}

func (v *Visitor) VisitFunType(c *parser.FunTypeContext) any {
	v.l.Debug("VisitFunType", "literal", v.LiteralOf(c))

	paramRes := AssertResult[Type](c.GetParam().Accept(v))
	if paramRes.Err() != nil {
		return Err[Type](fmt.Errorf("VisitFunType: visit parameter type: %w", paramRes.Err()))
	}

	returnRes := AssertResult[Type](c.GetReturn_().Accept(v))
	if returnRes.Err() != nil {
		return Err[Type](fmt.Errorf("VisitFunType: visit return type: %w", returnRes.Err()))
	}
	return OK[Type](FunType{
		Param:  paramRes.Val(),
		Return: returnRes.Val(),
	})
}

func (v *Visitor) VisitParenType(c *parser.ParenTypeContext) any {
	v.l.Debug("VisitParenType", "literal", v.LiteralOf(c))

	return AssertResult[Type](c.Type_().Accept(v))
}

func (v *Visitor) VisitFunExp(c *parser.FunExpContext) any {
	v.l.Debug("VisitFunExp", "literal", v.LiteralOf(c))

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitFunExpr: visit function: %w", funRes.Err()))
	}
	return OK[Exp](FunExp(funRes.Val()))
}

func (v *Visitor) VisitFun(c *parser.FunContext) any {
	v.l.Debug("VisitFun", "literal", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Fun](fmt.Errorf("VisitFun: visit body: %w", bodyRes.Err()))
	}
	return OK(Fun{
		Param: Var(c.GetParam().GetText()),
		Body:  bodyRes.Val(),
	})
}

func (v *Visitor) VisitAppExp(c *parser.AppExpContext) any {
	v.l.Debug("VisitAppExp", "literal", v.LiteralOf(c))

	argRes := AssertResult[Exp](c.GetArg().Accept(v))
	if argRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExpr: visit argument: %w", argRes.Err()))
	}

	funRes := AssertResult[Exp](c.GetFn().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExpr: visit function: %w", funRes.Err()))
	}
	return OK[Exp](AppExp{
		Fun: funRes.Val(),
		Arg: argRes.Val(),
	})
}

func (v *Visitor) VisitMatchExp(c *parser.MatchExpContext) any {
	v.l.Debug("VisitMatchExp", "literal", v.LiteralOf(c))

	argRes := AssertResult[Exp](c.GetArg().Accept(v))
	if argRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExpr: visit arg: %w", argRes.Err()))
	}

	nilRes := AssertResult[Exp](c.GetNilCase().Accept(v))
	if nilRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExpr: visit nil case: %w", nilRes.Err()))
	}
	consRes := AssertResult[Exp](c.GetConsCase().Accept(v))
	if consRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExpr: visit cons case: %w", consRes.Err()))
	}
	return OK[Exp](MatchExp{
		Arg:      argRes.Val(),
		NilCase:  nilRes.Val(),
		HeadVar:  Var(c.GetHead().GetText()),
		TailVar:  Var(c.GetTail().GetText()),
		ConsCase: consRes.Val(),
	})
}

func (v *Visitor) VisitNilExp(c *parser.NilExpContext) any {
	v.l.Debug("VisitNilExp", "literal", v.LiteralOf(c))

	return OK[Exp](NilExp{})
}

func (v *Visitor) VisitConsExp(c *parser.ConsExpContext) any {
	v.l.Debug("VisitConsExp", "literal", v.LiteralOf(c))

	tailRes := AssertResult[Exp](c.GetTail().Accept(v))
	if tailRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitConsExpr: visit tail: %w", tailRes.Err()))
	}

	headRes := AssertResult[Exp](c.GetHead().Accept(v))
	if headRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitConsExpr: visit head: %w", headRes.Err()))
	}
	return OK[Exp](ConsExp{
		Head: headRes.Val(),
		Tail: tailRes.Val(),
	})
}

func (v *Visitor) VisitBinOpExp(c *parser.BinOpExpContext) any {
	v.l.Debug("VisitBinOpExp", "literal", v.LiteralOf(c))

	leftRes := AssertResult[Exp](c.GetLeft().Accept(v))
	if leftRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExpr: visit left: %w", leftRes.Err()))
	}

	rightRes := AssertResult[Exp](c.GetRight().Accept(v))
	if rightRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExpr: visit right: %w", rightRes.Err()))
	}

	op, err := NewOp(c.GetOp())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExpr: new op: %w", err))
	}
	return OK[Exp](BinOpExp{
		Left:  leftRes.Val(),
		Op:    op,
		Right: rightRes.Val(),
	})
}

func (v *Visitor) VisitIfExp(c *parser.IfExpContext) any {
	v.l.Debug("VisitIfExp", "literal", v.LiteralOf(c))

	condRes := AssertResult[Exp](c.GetCond().Accept(v))
	if condRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExpr: visit cond: %w", condRes.Err()))
	}

	thenRes := AssertResult[Exp](c.GetThen().Accept(v))
	if thenRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExpr: visit then: %w", thenRes.Err()))
	}

	elseRes := AssertResult[Exp](c.GetElse_().Accept(v))
	if elseRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExpr: visit else: %w", elseRes.Err()))
	}
	return OK[Exp](IfExp{
		Cond: condRes.Val(),
		Then: thenRes.Val(),
		Else: elseRes.Val(),
	})
}

func (v *Visitor) VisitLetExp(c *parser.LetExpContext) any {
	v.l.Debug("VisitLetExp", "literal", v.LiteralOf(c))

	valRes := AssertResult[Exp](c.GetValue().Accept(v))
	if valRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExpr: visit value exp: %w", valRes.Err()))
	}

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExpr: visit body: %w", bodyRes.Err()))
	}
	return OK[Exp](LetExp{
		Var:  Var(c.Var_().ID().GetText()),
		Bind: valRes.Val(),
		Body: bodyRes.Val(),
	})
}

func (v *Visitor) VisitLetRecExp(c *parser.LetRecExpContext) any {
	v.l.Debug("VisitLetRecExp", "literal", v.LiteralOf(c))

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExpr: visit function: %w", funRes.Err()))
	}

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExpr: visit body: %w", bodyRes.Err()))
	}
	return OK[Exp](LetRecExp{
		Var:  Var(c.Var_().ID().GetText()),
		Bind: funRes.Val(),
		Body: bodyRes.Val(),
	})
}

func (v *Visitor) VisitIntExp(c *parser.IntExpContext) any {
	v.l.Debug("VisitIntExp", "literal", v.LiteralOf(c))

	i, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitIntExpr: parse int: %w", err))
	}
	return OK[Exp](IntExp(i))
}

func (v *Visitor) VisitBoolExp(c *parser.BoolExpContext) any {
	v.l.Debug("VisitBoolExp", "literal", v.LiteralOf(c))

	b, err := strconv.ParseBool(c.BOOL().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBoolExpr: parse bool: %w", err))
	}
	return OK[Exp](BoolExp(b))
}

func (v *Visitor) VisitVarExp(c *parser.VarExpContext) any {
	v.l.Debug("VisitVarExp", "literal", v.LiteralOf(c))

	return OK[Exp](VarExp(Var(c.Var_().ID().GetText())))
}

func (v *Visitor) VisitParenExp(c *parser.ParenExpContext) any {
	v.l.Debug("VisitParenExp", "literal", v.LiteralOf(c))

	return AssertResult[Exp](c.Exp().Accept(v))
}

func (v *Visitor) LiteralOf(st antlr.SyntaxTree) string {
	text := v.p.BaseParser.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
