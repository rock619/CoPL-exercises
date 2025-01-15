package main

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml4/parser"
)

type Visitor struct {
	*parser.BaseEvalML4Visitor
	p   *parser.EvalML4Parser
	env Env
	l   *slog.Logger
}

func NewVisitor(p *parser.EvalML4Parser, l *slog.Logger) *Visitor {
	return &Visitor{
		p: p,
		l: l,
	}
}

func (v *Visitor) Do(tree antlr.ParseTree) (Expr, error) {
	expr, ok := tree.Accept(v).(Expr)
	if !ok {
		return nil, fmt.Errorf("tree.Accept(v) is not Expr")
	}
	return expr, nil
}

func (v *Visitor) getEnv() Env {
	return slices.Clone(v.env)
}

func (v *Visitor) VisitQuestion(c *parser.QuestionContext) interface{} {
	v.l.Info("VisitQuestion", "literal", getLiteral(v.p, c))
	return c.Eval().Accept(v)
}

func (v *Visitor) VisitEval(c *parser.EvalContext) interface{} {
	v.l.Info("VisitEval", "literal", getLiteral(v.p, c))
	var env Env
	if c.DefList() != nil {
		env = c.DefList().Accept(v).(Env)
	}
	defer v.setEnv(env)()

	expr, ok := c.Expr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitEval: expr is not Expr", "expr", getLiteral(v.p, c.Expr()))
		return nil
	}
	return expr
}

func (v *Visitor) VisitDefList(c *parser.DefListContext) interface{} {
	v.l.Info("VisitDefList", "literal", getLiteral(v.p, c), "env", v.env)
	env := make(Env, len(c.AllDef()))
	for i, def := range c.AllDef() {
		bind, ok := def.Accept(v).(Bind)
		if !ok {
			v.l.Error("VisitDefList", "literal", getLiteral(v.p, c), "error", "def is not Bind")
			return nil
		}
		v.l.Info("VisitDefList", "literal", getLiteral(v.p, c), "bind", bind)
		env[i] = bind
	}
	return env
}

func (v *Visitor) VisitDef(c *parser.DefContext) interface{} {
	v.l.Info("VisitDef", "literal", getLiteral(v.p, c), "env", v.env)
	name := c.VARNAME().GetText()
	val, ok := c.Value().Accept(v).(Value)
	if !ok {
		v.l.Error("VisitDef", "literal", getLiteral(v.p, c), "error", "value is not Value")
		return nil
	}
	return Bind{
		Name: name,
		Val:  val,
	}
}

func (v *Visitor) VisitIntValue(c *parser.IntValueContext) interface{} {
	v.l.Info("VisitIntValue", "literal", getLiteral(v.p, c), "env", v.env)
	i, err := NewIntValue(c.INT().GetText())
	if err != nil {
		v.l.Error("VisitIntValue", "literal", getLiteral(v.p, c), "error", err)
		return nil
	}
	return i
}

func (v *Visitor) VisitBoolValue(c *parser.BoolValueContext) interface{} {
	v.l.Info("VisitBoolValue", "literal", getLiteral(v.p, c), "env", v.env)
	b, err := NewBoolValue(c.BOOL().GetText())
	if err != nil {
		v.l.Error("VisitBoolValue", "literal", getLiteral(v.p, c), "error", err)
		return nil
	}
	return BoolValue(b)
}

func (v *Visitor) VisitFunValue(c *parser.FunValueContext) interface{} {
	v.l.Info("VisitFunValue", "literal", getLiteral(v.p, c), "env", v.env)

	var env Env
	if c.DefList() != nil {
		env = c.DefList().Accept(v).(Env)
	}
	defer v.setEnv(env)()

	return FunValue{
		Env: env,
		Fun: Fun{
			Param:       c.Fun().VARNAME().GetText(),
			BodyLiteral: getLiteral(v.p, c.Fun().GetBody()),
			BodyCtx:     c.Fun().GetBody(),
		},
	}
}

func (v *Visitor) VisitRecFunValue(c *parser.RecFunValueContext) interface{} {
	v.l.Info("VisitRecFunValue", "literal", getLiteral(v.p, c), "env", v.env)

	var env Env
	if c.DefList() != nil {
		env = c.DefList().Accept(v).(Env)
	}
	defer v.setEnv(env)()

	return RecFunValue{
		Env: env,
		RecFun: RecFun{
			Name: c.RecFun().VARNAME().GetText(),
			Fun: Fun{
				Param:       c.RecFun().Fun().VARNAME().GetText(),
				BodyLiteral: getLiteral(v.p, c.RecFun().Fun().GetBody()),
				BodyCtx:     c.RecFun().Fun().GetBody(),
			},
		},
	}
}

func (v *Visitor) VisitConsValue(c *parser.ConsValueContext) interface{} {
	v.l.Info("VisitConsValue", "literal", getLiteral(v.p, c), "env", v.env)
	head, ok := c.GetHead().Accept(v).(Value)
	if !ok {
		v.l.Error("VisitConsValue: head is not Value", "head", getLiteral(v.p, c.GetHead()))
		return nil
	}
	tail, ok := c.GetTail().Accept(v).(Value)
	if !ok {
		v.l.Error("VisitConsValue: tail is not Value", "tail", getLiteral(v.p, c.GetTail()))
		return nil
	}
	return ConsValue{
		Head: head,
		Tail: tail,
	}
}

func (v *Visitor) VisitEmptyListValue(c *parser.EmptyListValueContext) interface{} {
	v.l.Info("VisitEmptyListValue", "literal", getLiteral(v.p, c), "env", v.env)
	return EmptyListValue{}
}

// Expr

func (v *Visitor) VisitParenExpr(c *parser.ParenExprContext) interface{} {
	v.l.Info("VisitParenExpr", "literal", getLiteral(v.p, c), "env", v.env)
	return c.Expr().Accept(v)
}

func (v *Visitor) VisitFunExpr(c *parser.FunExprContext) interface{} {
	v.l.Info("VisitFunExpr", "literal", getLiteral(v.p, c), "env", v.env)

	fun := Fun{
		Param:       c.Fun().VARNAME().GetText(),
		BodyLiteral: getLiteral(v.p, c.Fun().GetBody()),
		BodyCtx:     c.Fun().GetBody(),
	}
	evalTo := FunValue{
		Env: v.getEnv(),
		Fun: fun,
	}
	v.l.Info("VisitFunExpr", "literal", getLiteral(v.p, c), "evalTo", evalTo)
	return FunExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
			evalTo:  evalTo,
			rule:    EFun,
		},
		Fun: fun,
	}
}

func (v *Visitor) VisitAppExpr(c *parser.AppExprContext) interface{} {
	v.l.Info("VisitAppExpr", "literal", getLiteral(v.p, c), "env", v.env)
	fun, ok := c.GetFn().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitAppExpr: fn is not Expr", "fn", getLiteral(v.p, c.GetFn()))
		return nil
	}
	arg, ok := c.GetArg().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitAppExpr: arg is not Expr", "arg", getLiteral(v.p, c.GetArg()))
		return nil
	}
	expr := AppExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Fun: fun,
		Arg: arg,
	}

	if arg.EvalTo() == nil {
		return expr
	}

	switch funVal := fun.EvalTo().(type) {
	case nil:
		return expr
	case FunValue:
		env := append(funVal.Env, Bind{
			Name: funVal.Param,
			Val:  arg.EvalTo(),
		})
		defer v.setEnv(env)()
		applied, ok := funVal.BodyCtx.Accept(v).(Expr)
		if !ok {
			v.l.Error("VisitAppExpr: funVal.BodyCtx is not Expr", "funVal.BodyCtx", getLiteral(v.p, funVal.BodyCtx))
			return nil
		}
		expr.Applied = applied
		expr.evalTo = applied.EvalTo()
		expr.rule = EApp
		expr.AddChildren(fun, arg, applied)
		v.l.Info("VisitAppExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
		return expr
	case RecFunValue:
		env := append(funVal.Env, Bind{
			Name: funVal.RecFun.Name,
			Val:  funVal,
		}, Bind{
			Name: funVal.RecFun.Fun.Param,
			Val:  arg.EvalTo(),
		})
		defer v.setEnv(env)()
		applied, ok := funVal.RecFun.Fun.BodyCtx.Accept(v).(Expr)
		if !ok {
			v.l.Error("VisitAppExpr: funVal.RecFun.Fun.BodyCtx is not Expr", "funVal.RecFun.Fun.BodyCtx", getLiteral(v.p, funVal.RecFun.Fun.BodyCtx))
			return nil
		}
		expr.Applied = applied
		expr.evalTo = applied.EvalTo()
		expr.rule = EAppRec
		expr.AddChildren(fun, arg, applied)
		v.l.Info("VisitAppExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
		return expr
	default:
		v.l.Error("VisitAppExpr: fun.evalTo is not FunValue or nil",
			"type", fmt.Sprintf("%T", funVal),
		)
		return nil
	}
}

func (v *Visitor) VisitBinOpExpr(c *parser.BinOpExprContext) interface{} {
	v.l.Info("VisitBinOpExpr", "literal", getLiteral(v.p, c), "env", v.env)
	left, ok := c.GetLeft().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitBinOpExpr: left is not Expr", "left", getLiteral(v.p, c.GetLeft()))
		return nil
	}
	right, ok := c.GetRight().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitBinOpExpr: right is not Expr", "right", getLiteral(v.p, c.GetRight()))
		return nil
	}
	expr := BinOpExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Left:  left,
		Op:    c.GetOp(),
		Right: right,
	}
	if left.EvalTo() == nil || right.EvalTo() == nil {
		return expr
	}
	leftVal, ok := left.EvalTo().(IntValue)
	if !ok {
		v.l.Error("VisitBinOpExpr: left.evalTo is not IntValue", "type", fmt.Sprintf("%T", left.EvalTo()))
		return nil
	}
	rightVal, ok := right.EvalTo().(IntValue)
	if !ok {
		v.l.Error("VisitBinOpExpr: right.evalTo is not IntValue", "type", fmt.Sprintf("%T", right.EvalTo()))
		return nil
	}

	switch c.GetOp().GetTokenType() {
	case parser.EvalML4ParserPLUS:
		expr.evalTo = IntValue(leftVal + rightVal)
		expr.rule = EPlus
	case parser.EvalML4ParserMINUS:
		expr.evalTo = IntValue(leftVal - rightVal)
		expr.rule = EMinus
	case parser.EvalML4ParserTIMES:
		expr.evalTo = IntValue(leftVal * rightVal)
		expr.rule = ETimes
	case parser.EvalML4ParserLT:
		expr.evalTo = BoolValue(leftVal < rightVal)
		expr.rule = ELt
	default:
		v.l.Error("VisitBinOpExpr: unknown op", "op", c.GetOp().GetText())
		return nil
	}
	expr.AddChildren(left, right, NewBExpr(leftVal, rightVal, c.GetOp()))
	v.l.Info("VisitBinOpExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitIfExpr(c *parser.IfExprContext) interface{} {
	v.l.Info("VisitIfExpr", "literal", getLiteral(v.p, c), "env", v.env)
	cond, ok := c.GetCond().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitIfExpr: cond is not Expr", "cond", getLiteral(v.p, c.GetCond()))
		return nil
	}
	expr := IfExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Cond: cond,
	}
	switch condVal := cond.EvalTo().(type) {
	case nil:
		v.l.Info("VisitIfExpr: cond.evalTo is nil", "cond", getLiteral(v.p, c.GetCond()))
		return expr
	case BoolValue:
		if condVal {
			then, ok := c.GetThen().Accept(v).(Expr)
			if !ok {
				v.l.Error("VisitIfExpr: then is not Expr", "then", getLiteral(v.p, c.GetThen()))
				return nil
			}
			expr.Then = then
			expr.evalTo = then.EvalTo()
			expr.rule = EIfT
			expr.AddChildren(cond, then)
			break
		}
		els, ok := c.GetElse_().Accept(v).(Expr)
		if !ok {
			v.l.Error("VisitIfExpr: else is not Expr", "else", getLiteral(v.p, c.GetElse_()))
			return nil
		}
		expr.Else = els
		expr.evalTo = els.EvalTo()
		expr.rule = EIfF
		expr.AddChildren(cond, els)
	default:
		v.l.Error("VisitIfExpr: cond.evalTo is not BoolValue or nil",
			"type", fmt.Sprintf("%T", condVal))
		return nil
	}
	v.l.Info("VisitIfExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitLetExpr(c *parser.LetExprContext) interface{} {
	v.l.Info("VisitLetExpr", "literal", getLiteral(v.p, c), "env", v.env)
	bind, ok := c.GetBindExpr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitLetExpr: bindExpr is not Expr", "bindExpr", getLiteral(v.p, c.GetBindExpr()))
		return nil
	}
	expr := LetExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Var:  c.GetBindName().GetText(),
		Bind: bind,
	}
	env := append(v.getEnv(), Bind{
		Name: expr.Var,
		Val:  bind.EvalTo(),
	})
	defer v.setEnv(env)()
	body, ok := c.GetBody().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitLetExpr: body is not Expr", "body", getLiteral(v.p, c.GetBody()))
		return nil
	}
	expr.Body = body
	expr.evalTo = body.EvalTo()
	expr.rule = ELet
	expr.AddChildren(bind, body)
	v.l.Info("VisitLetExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitLetRecExpr(c *parser.LetRecExprContext) interface{} {
	v.l.Info("VisitLetRecExpr", "literal", getLiteral(v.p, c), "env", v.env)

	expr := LetRecExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
			rule:    ELetRec,
		},
		RecFun: RecFun{
			Name: c.RecFun().GetFunName().GetText(),
			Fun: Fun{
				Param:       c.RecFun().Fun().GetParam().GetText(),
				BodyLiteral: getLiteral(v.p, c.RecFun().Fun().GetBody()),
				BodyCtx:     c.RecFun().Fun().GetBody(),
			},
		},
	}

	env := append(v.getEnv(), Bind{
		Name: c.RecFun().GetFunName().GetText(),
		Val: RecFunValue{
			Env:    v.getEnv(),
			RecFun: expr.RecFun,
		},
	})
	defer v.setEnv(env)()
	body, ok := c.Expr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitLetRecExpr: body is not Expr", "body", getLiteral(v.p, c.Expr()))
		return nil
	}
	expr.evalTo = body.EvalTo()
	expr.AddChildren(body)
	v.l.Info("VisitLetRecExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitMatchExpr(c *parser.MatchExprContext) interface{} {
	v.l.Info("VisitMatchExpr", "literal", getLiteral(v.p, c), "env", v.env)
	matchExpr, ok := c.GetMatchExpr().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitMatchExpr: matchExpr is not Expr", "matchExpr", getLiteral(v.p, c.GetMatchExpr()))
		return nil
	}
	expr := MatchExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Matched: matchExpr,
		HeadVar: c.ConsPattern().GetHeadVar().GetText(),
		TailVar: c.ConsPattern().GetTailVar().GetText(),
	}

	switch matchVal := matchExpr.EvalTo().(type) {
	case nil:
	case EmptyListValue:
		// E-MatchNil
		patternExpr, ok := c.EmptyPattern().Expr().Accept(v).(Expr)
		if !ok {
			v.l.Error("VisitMatchExpr: emptyPattern's expr is not Expr",
				"emptyPattern", getLiteral(v.p, c.EmptyPattern().Expr()))
			return nil
		}
		expr.Empty = patternExpr
		expr.evalTo = patternExpr.EvalTo()
		expr.rule = EMatchNil
		expr.AddChildren(matchExpr, patternExpr)
	case ConsValue:
		// E-MatchCons
		env := append(v.getEnv(),
			Bind{Name: expr.HeadVar, Val: matchVal.Head},
			Bind{Name: expr.TailVar, Val: matchVal.Tail},
		)
		defer v.setEnv(env)()
		patternExpr, ok := c.ConsPattern().Expr().Accept(v).(Expr)
		if !ok {
			v.l.Error("VisitMatchExpr: consPattern's expr is not Expr",
				"consPattern", getLiteral(v.p, c.ConsPattern().Expr()))
			return nil
		}

		expr.Cons = patternExpr
		expr.evalTo = patternExpr.EvalTo()
		expr.rule = EMatchCons
		expr.AddChildren(matchExpr, patternExpr)
	default:
		v.l.Error("VisitMatchExpr: matchExpr.evalTo is not EmptyListValue or ConsValue or nil",
			"literal", getLiteral(v.p, c),
			"type", fmt.Sprintf("%T", matchVal),
		)
		return nil
	}
	v.l.Info("VisitMatchExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitEmptyListExpr(c *parser.EmptyListExprContext) interface{} {
	v.l.Info("VisitEmptyListExpr", "literal", getLiteral(v.p, c), "env", v.env)
	return EmptyListExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
			evalTo:  EmptyListValue{},
			rule:    ENil,
		},
	}
}

func (v *Visitor) VisitConsExpr(c *parser.ConsExprContext) interface{} {
	v.l.Info("VisitConsExpr", "literal", getLiteral(v.p, c), "env", v.env)
	head, ok := c.GetHead().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitConsExpr: head is not Expr", "head", getLiteral(v.p, c.GetHead()))
		return nil
	}
	tail, ok := c.GetTail().Accept(v).(Expr)
	if !ok {
		v.l.Error("VisitConsExpr: tail is not Expr", "head", getLiteral(v.p, c.GetTail()))
		return nil
	}
	expr := ConsExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Head: head,
		Tail: tail,
	}
	if head.EvalTo() == nil || tail.EvalTo() == nil {
		return expr
	}
	expr.evalTo = ConsValue{Head: head.EvalTo(), Tail: tail.EvalTo()}
	expr.rule = ECons
	expr.AddChildren(head, tail)
	v.l.Info("VisitConsExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
	return expr
}

func (v *Visitor) VisitIntExpr(c *parser.IntExprContext) interface{} {
	v.l.Info("VisitIntExpr", "literal", getLiteral(v.p, c), "env", v.env)
	val, err := NewIntValue(c.INT().GetText())
	if err != nil {
		v.l.Error("VisitIntExpr", "literal", getLiteral(v.p, c), "error", err)
		return nil
	}
	return IntExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
			evalTo:  val,
			rule:    EInt,
		},
		Val: val,
	}
}

func (v *Visitor) VisitBoolExpr(c *parser.BoolExprContext) interface{} {
	v.l.Info("VisitBoolExpr", "literal", getLiteral(v.p, c), "env", v.env)
	val, err := NewBoolValue(c.BOOL().GetText())
	if err != nil {
		v.l.Error("VisitBoolExpr", "literal", getLiteral(v.p, c), "error", err)
		return nil
	}
	return BoolExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
			evalTo:  val,
			rule:    EBool,
		},
		Val: BoolValue(val),
	}
}

func (v *Visitor) VisitVarExpr(c *parser.VarExprContext) interface{} {
	v.l.Info("VisitVarExpr", "literal", getLiteral(v.p, c), "env", v.env)
	expr := VarExpr{
		BaseExpr: &BaseExpr{
			env:     v.env,
			literal: getLiteral(v.p, c),
		},
		Name: c.VARNAME().GetText(),
	}
	for _, bind := range slices.Backward(v.env) {
		if bind.Name != expr.Name {
			continue
		}
		expr.evalTo = bind.Val
		expr.rule = EVar
		v.l.Info("VisitVarExpr", "literal", getLiteral(v.p, c), "evalTo", expr.evalTo)
		return expr
	}
	return expr
}

func (v *Visitor) setEnv(env Env) (reset func()) {
	prev := v.getEnv()
	v.env = env
	return func() {
		v.env = prev
	}
}
