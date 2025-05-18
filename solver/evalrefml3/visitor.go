package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalrefml3/parser"
)

type Visitor struct {
	*parser.BaseEvalRefML3Visitor
	p *parser.EvalRefML3Parser
	l *slog.Logger
}

func NewVisitor(p *parser.EvalRefML3Parser, l *slog.Logger) *Visitor {
	return &Visitor{
		p: p,
		l: l,
	}
}

func (v *Visitor) Do() (Judgement, error) {
	tree := v.p.Eval()
	if tree == nil {
		return nil, fmt.Errorf("tree is nil: %q", v.LiteralOf(tree))
	}

	res := AssertResult[Judgement](tree.Accept(v))
	return res.Val(), res.Err()
}

func (v *Visitor) VisitEval(c *parser.EvalContext) any {
	v.l.Debug("VisitEval", "eval", v.LiteralOf(c))

	var storeIn Store
	if c.GetStoreIn() != nil {
		storeInRes := AssertResult[Store](c.GetStoreIn().Accept(v))
		if storeInRes.Err() != nil {
			return Err[Judgement](fmt.Errorf("VisitEval: %w", storeInRes.Err()))
		}
		storeIn = storeInRes.Val()
	}

	var env Env
	if c.Env() != nil {
		envRes := AssertResult[Env](c.Env().Accept(v))
		if envRes.Err() != nil {
			return Err[Judgement](fmt.Errorf("VisitEval: %w", envRes.Err()))
		}
		env = envRes.Val()
	}

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", expRes.Err()))
	}

	evalToRes := AssertResult[Value](c.Value().Accept(v))
	if evalToRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", evalToRes.Err()))
	}

	var storeOut Store
	if c.GetStoreOut() != nil {
		storeOutRes := AssertResult[Store](c.GetStoreOut().Accept(v))
		if storeOutRes.Err() != nil {
			return Err[Judgement](fmt.Errorf("VisitEval: %w", storeOutRes.Err()))
		}
		storeOut = storeOutRes.Val()
	}

	return OK[Judgement](NewEvaledExpJudgement(storeIn, env, expRes.Val(), evalToRes.Val(), storeOut))
}

func (v *Visitor) VisitStore(c *parser.StoreContext) any {
	v.l.Debug("VisitStore", "store", v.LiteralOf(c), "c.AllAssign()", len(c.AllAssign()))

	store := make(Store, 0)
	for _, a := range c.AllAssign() {
		res := AssertResult[Assign](a.Accept(v))
		if res.Err() != nil {
			return Err[Store](fmt.Errorf("VisitStore: %w", res.Err()))
		}
		store.Set(res.Val())
	}

	return OK(store)
}

func (v *Visitor) VisitAssign(c *parser.AssignContext) any {
	v.l.Debug("VisitAssign", "assign", v.LiteralOf(c), "loc", c.LOC().GetText())

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Assign](fmt.Errorf("VisitAssign: %w", valRes.Err()))
	}

	return OK(Assign{
		Loc:   LocValue(c.LOC().GetText()),
		Value: valRes.Val(),
	})
}

func (v *Visitor) VisitEnv(c *parser.EnvContext) any {
	v.l.Debug("VisitEnv", "env", v.LiteralOf(c))
	env := make(Env, 0)
	for _, b := range c.AllBind() {
		res := AssertResult[Bind](b.Accept(v))
		if res.Err() != nil {
			return Err[Env](fmt.Errorf("VisitEnv: %w", res.Err()))
		}
		env.Set(res.Val())
	}
	return OK(env)
}

func (v *Visitor) VisitBind(c *parser.BindContext) any {
	v.l.Debug("VisitBind", "bind", v.LiteralOf(c))
	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Bind](fmt.Errorf("VisitBind: %w", valRes.Err()))
	}
	return OK(Bind{
		Var:   Var(c.IDENTIFIER().GetText()),
		Value: valRes.Val(),
	})
}

func (v *Visitor) VisitIntValue(c *parser.IntValueContext) any {
	v.l.Debug("VisitIntValue", "intValue", v.LiteralOf(c))

	i, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		return Err[Value](fmt.Errorf("VisitIntValue: invalid int value: %s", c.INT().GetText()))
	}
	return OK[Value](IntValue(i))
}

func (v *Visitor) VisitBoolValue(c *parser.BoolValueContext) any {
	v.l.Debug("VisitBoolValue", "boolValue", v.LiteralOf(c))

	b, err := strconv.ParseBool(c.BOOL().GetText())
	if err != nil {
		return Err[Value](fmt.Errorf("VisitBoolValue: invalid bool value: %s", c.BOOL().GetText()))
	}
	return OK[Value](BoolValue(b))
}

func (v *Visitor) VisitLocValue(c *parser.LocValueContext) any {
	v.l.Debug("VisitLocValue", "locValue", v.LiteralOf(c))

	return OK[Value](LocValue(c.LOC().GetText()))
}

func (v *Visitor) VisitFunValue(c *parser.FunValueContext) any {
	v.l.Debug("VisitFunValue", "funValue", v.LiteralOf(c))

	var env Env
	if c.Env() != nil {
		envRes := AssertResult[Env](c.Env().Accept(v))
		if envRes.Err() != nil {
			return Err[Value](fmt.Errorf("VisitFunValue: %w", envRes.Err()))
		}
		env = envRes.Val()
	}

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitFunValue: %w", funRes.Err()))
	}
	return OK[Value](FunValue{
		Env: env,
		Fun: funRes.Val(),
	})
}

func (v *Visitor) VisitFun(c *parser.FunContext) any {
	v.l.Debug("VisitFun", "fun", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitFun: %w", bodyRes.Err()))
	}
	return OK(Fun{
		Param: Var(c.GetParam().GetText()),
		Body:  bodyRes.Val(),
	})
}

func (v *Visitor) VisitRecFunValue(c *parser.RecFunValueContext) any {
	v.l.Debug("VisitRecFunValue", "recFunValue", v.LiteralOf(c))

	var env Env
	if c.Env() != nil {
		envRes := AssertResult[Env](c.Env().Accept(v))
		if envRes.Err() != nil {
			return Err[Value](fmt.Errorf("VisitRecFunValue: %w", envRes.Err()))
		}
		env = envRes.Val()
	}

	funRes := AssertResult[RecFun](c.RecFun().Accept(v))
	if funRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitRecFunValue: %w", funRes.Err()))
	}
	return OK[Value](RecFunValue{
		Env: env,
		Fun: funRes.Val(),
	})
}

func (v *Visitor) VisitRecFun(c *parser.RecFunContext) any {
	v.l.Debug("VisitRecFun", "recFun", v.LiteralOf(c))

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitRecFun: %w", funRes.Err()))
	}
	return OK(RecFun{
		Name: Var(c.GetFunName().GetText()),
		Fun:  funRes.Val(),
	})
}

func (v *Visitor) VisitParenExp(c *parser.ParenExpContext) any {
	v.l.Debug("VisitParenExp", "parenExp", v.LiteralOf(c))

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitParenExp: %w", expRes.Err()))
	}
	return OK[Exp](ParenExp{
		Inner: expRes.Val(),
	})
}

func (v *Visitor) VisitRefExp(c *parser.RefExpContext) any {
	v.l.Debug("VisitRefExp", "refExp", v.LiteralOf(c))

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitRefExp: %w", expRes.Err()))
	}

	return OK[Exp](RefExp{
		Exp: expRes.Val(),
	})
}

func (v *Visitor) VisitDerefExp(c *parser.DerefExpContext) any {
	v.l.Debug("VisitDerefExp", "derefExp", v.LiteralOf(c))

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitDerefExp: %w", expRes.Err()))
	}

	return OK[Exp](DerefExp{
		Exp: expRes.Val(),
	})
}

func (v *Visitor) VisitFunExp(c *parser.FunExpContext) any {
	v.l.Debug("VisitFunExp", "funExp", v.LiteralOf(c))

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitFunExp: %w", funRes.Err()))
	}
	return OK[Exp](FunExp{
		Fun: funRes.Val(),
	})
}

func (v *Visitor) VisitAppExp(c *parser.AppExpContext) any {
	v.l.Debug("VisitAppExp", "appExp", v.LiteralOf(c))

	funRes := AssertResult[Exp](c.GetFn().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExp: %w", funRes.Err()))
	}

	argRes := AssertResult[Exp](c.GetArg().Accept(v))
	if argRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExp: %w", argRes.Err()))
	}

	return OK[Exp](AppExp{
		Fun: funRes.Val(),
		Arg: argRes.Val(),
	})
}

func (v *Visitor) VisitBinOpExp(c *parser.BinOpExpContext) any {
	v.l.Debug("VisitBinOpExp", "binOpExp", v.LiteralOf(c))

	leftRes := AssertResult[Exp](c.GetLeft().Accept(v))
	if leftRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", leftRes.Err()))
	}

	op, err := NewOp(c.GetOp())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", err))
	}

	rightRes := AssertResult[Exp](c.GetRight().Accept(v))
	if rightRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", rightRes.Err()))
	}
	return OK[Exp](BinOpExp{
		Left:  leftRes.Val(),
		Op:    op,
		Right: rightRes.Val(),
	})
}

func (v *Visitor) VisitAssignExp(c *parser.AssignExpContext) any {
	v.l.Debug("VisitAssignExp", "assignExp", v.LiteralOf(c))

	rightRes := AssertResult[Exp](c.GetRight().Accept(v))
	if rightRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAssignExp: %w", rightRes.Err()))
	}

	leftRes := AssertResult[Exp](c.GetLeft().Accept(v))
	if leftRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAssignExp: %w", leftRes.Err()))
	}

	return OK[Exp](AssignExp{
		Left:  leftRes.Val(),
		Right: rightRes.Val(),
	})
}

func (v *Visitor) VisitIfExp(c *parser.IfExpContext) any {
	v.l.Debug("VisitIfExp", "ifExp", v.LiteralOf(c))

	condRes := AssertResult[Exp](c.GetCond().Accept(v))
	if condRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", condRes.Err()))
	}

	thenRes := AssertResult[Exp](c.GetThen().Accept(v))
	if thenRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", thenRes.Err()))
	}

	elseRes := AssertResult[Exp](c.GetElse_().Accept(v))
	if elseRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", elseRes.Err()))
	}

	return OK[Exp](IfExp{
		Cond: condRes.Val(),
		Then: thenRes.Val(),
		Else: elseRes.Val(),
	})
}

func (v *Visitor) VisitLetExp(c *parser.LetExpContext) any {
	v.l.Debug("VisitLetExp", "letExp", v.LiteralOf(c))

	bindRes := AssertResult[Exp](c.GetBindExp().Accept(v))
	if bindRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExp: %w", bindRes.Err()))
	}

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExp: %w", bodyRes.Err()))
	}

	return OK[Exp](LetExp{
		Var:  Var(c.GetVar_().GetText()),
		Bind: bindRes.Val(),
		Body: bodyRes.Val(),
	})
}

func (v *Visitor) VisitLetRecExp(c *parser.LetRecExpContext) any {
	v.l.Debug("VisitLetRecExp", "letRecExp", v.LiteralOf(c))

	recFunRes := AssertResult[RecFun](c.RecFun().Accept(v))
	if recFunRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExp: %w", recFunRes.Err()))
	}

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExp: %w", bodyRes.Err()))
	}

	return OK[Exp](LetRecExp{
		RecFun: recFunRes.Val(),
		Body:   bodyRes.Val(),
	})
}

func (v *Visitor) VisitIntExp(c *parser.IntExpContext) any {
	v.l.Debug("VisitIntExp", "intExp", v.LiteralOf(c))

	i, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitIntExp: invalid int value: %s", c.INT().GetText()))
	}
	return OK[Exp](IntExp(i))
}

func (v *Visitor) VisitBoolExp(c *parser.BoolExpContext) any {
	v.l.Debug("VisitBoolExp", "boolExp", v.LiteralOf(c))

	b, err := strconv.ParseBool(c.BOOL().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBoolExp: invalid bool value: %s", c.BOOL().GetText()))
	}
	return OK[Exp](BoolExp(b))
}

func (v *Visitor) VisitVarExp(c *parser.VarExpContext) any {
	v.l.Debug("VisitVarExp", "varExp", v.LiteralOf(c))

	return OK[Exp](VarExp(c.IDENTIFIER().GetText()))
}

func (v *Visitor) LiteralOf(st antlr.SyntaxTree) string {
	text := v.p.BaseParser.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
