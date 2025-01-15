package main

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml3/parser"
)

func evalExpr(input string, logger *slog.Logger) (Expr, error) {
	logger.Info("evalExpr", "input", input)
	inputStream := antlr.NewInputStream(input)
	p := parser.NewEvalML3Parser(antlr.NewCommonTokenStream(
		parser.NewEvalML3Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))
	tree := p.Eval()
	l := &Listener{p: p, l: logger}
	l.nextID = 100

	antlr.NewParseTreeWalker().Walk(l, tree)
	root, ok := l.Pop()
	if !ok {
		return nil, fmt.Errorf("no root expr: %v", l.result)
	}
	return root, nil
}

type Listener struct {
	*parser.BaseEvalML3Listener
	p      *parser.EvalML3Parser
	result []Expr
	env    Env
	l      *slog.Logger
	depth  int
	nextID int
}

func (l *Listener) IncrDepth() {
	l.depth++
}

func (l *Listener) DecrDepth() {
	l.depth--
}

func (l *Listener) Empty() bool {
	return len(l.result) == 0
}

func (l *Listener) Push(e Expr) {
	l.result = append(l.result, e)
	// l.logExpr("Push")
}

func (l *Listener) Pop() (Expr, bool) {
	if l.Empty() {
		return nil, false
	}
	e := l.result[len(l.result)-1]
	l.result = slices.Delete(l.result, len(l.result)-1, len(l.result))
	return e, true
}

func (l *Listener) Peek() (Expr, bool) {
	if l.Empty() {
		return nil, false
	}
	return l.result[len(l.result)-1], true
}

func (l *Listener) PopParent() (parent Expr, siblings []Expr, ok bool) {
	for {
		prev := Pop(&l.result)
		if prev.Depth() < l.depth {
			return prev, siblings, true
		}
		siblings = slices.Insert(siblings, 0, prev)
	}
}

func (l *Listener) PeekParent() (parent Expr, siblings []Expr, ok bool) {
	for _, prev := range slices.Backward(l.result) {
		if prev.Depth() < l.depth {
			return prev, siblings, true
		}
		siblings = slices.Insert(siblings, 0, prev)
	}
	return nil, siblings, false
}

func (l *Listener) ParentEnv() Env {
	parent, _, ok := l.PeekParent()
	if ok {
		return parent.Env()
	}
	return l.env
}

func (l *Listener) EnterEveryRule(c antlr.ParserRuleContext) {
	if exprCtx, ok := c.(parser.IExprContext); ok {
		l.EnterExpr(exprCtx)
	}
}

func (l *Listener) EnterDefList(c *parser.DefListContext) {
	l.l.Debug("EnterDefList", "literal", getLiteral(l.p, c))
}

func (l *Listener) EnterDef(c *parser.DefContext) {
	l.l.Debug("EnterDef", "literal", getLiteral(l.p, c))
}

func (l *Listener) EnterValue(c parser.IValueContext) {
	l.l.Debug("EnterValue", "literal", getLiteral(l.p, c))
}

func (l *Listener) EnterExpr(c parser.IExprContext) {
	defer l.IncrDepth()

	// l.l.Debug("EnterExpr", "literal", getLiteral(l.p, c))

	env := slices.Clone(l.env)
	skip := false
	parent, siblings, ok := l.PeekParent()
	if ok {
		env = parent.Env()
		skip = parent.Skip()
	}

	// 1. 親がLetExprで上の兄弟が1つの場合、環境に値の束縛を追加する必要がある
	// 2. 親がLetRecExprで上の兄弟が1つの場合、環境に再帰関数の束縛を追加する必要がある
	// 3. 親がIfExprで上の兄弟がいる場合、無限ループを回避するため無駄な評価を避ける
	//    1. 条件式が評価されていなければthen, elseともに評価しない
	//    2. 条件式がfalseの場合はthenを評価しない
	//    3. 条件式がtrueの場合はelseを評価しない
	switch p := parent.(type) {
	case *LetExpr:
		if len(siblings) != 1 {
			break
		}
		// 親がLetExprで上の兄弟が1つの場合

		env = append(env, Bind{Name: p.Var, Val: siblings[0].EvalTo()})
		l.l.Info("EnterExpr let, body expr", "depth", l.depth, "siblings[0]", siblings[0])
	case *LetRecExpr:
		if len(siblings) != 1 {
			break
		}
		// 親がLetRecExprで上の兄弟が1つの場合

		env = append(env, Bind{
			Name: p.Name,
			Val: RecFunValue{
				RecFun: RecFun{
					Fun:  Fun{Param: p.Param, Body: siblings[0]},
					Name: p.Name,
				},
				Env: env,
			},
		})
		l.l.Info("EnterExpr letrec, body expr", "depth", l.depth, "env", env)
	case *IfExpr:
		if count := len(siblings); count > 0 {
			// 親がIfExprで上の兄弟が1つの場合はthen、2つの場合はelseということになる
			//
			condVal, ok := siblings[0].EvalTo().(BoolValue)
			if !ok {
				skip = true
			}
			// 条件式がfalseの場合はthen、trueの場合はelseを評価しない
			if count == 1 && !condVal || count == 2 && condVal {
				skip = true
			}
			l.l.Info("EnterExpr if, then or else", "skip", skip, "cond", condVal, "literal", getLiteral(c.GetParser(), c))
		}
	default:
	}

	expr := &BaseExpr{
		id:      l.nextID,
		ctx:     c,
		env:     env,
		literal: getLiteral(l.p, c),
		depth:   l.depth,
		skip:    skip,
	}
	l.nextID++
	l.Push(expr)
}

func (l *Listener) EnterParenExpr(c *parser.ParenExprContext) {
	l.l.Info("EnterParenExpr", "literal", getLiteral(l.p, c))
	_, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
}

func (l *Listener) EnterBoolExpr(c *parser.BoolExprContext) {
	l.l.Info("EnterBoolExpr", "literal", getLiteral(l.p, c))
	v, err := NewBoolValue(c.GetParser().GetCurrentToken())
	if err != nil {
		l.l.Error("NewBoolValue", "err", err)
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&BoolExpr{
		BaseExpr: self.(*BaseExpr),
		Val:      v,
	})
}

func (l *Listener) EnterIntExpr(c *parser.IntExprContext) {
	l.l.Info("EnterIntExpr", "literal", getLiteral(l.p, c))
	i, err := NewIntValue(getLiteral(l.p, c))
	if err != nil {
		l.l.Error("NewIntValue", "err", err)
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&IntExpr{
		BaseExpr: self.(*BaseExpr),
		Val:      i,
	})
}

func (l *Listener) EnterNegIntExpr(c *parser.NegIntExprContext) {
	l.l.Info("EnterNegIntExpr", "literal", getLiteral(l.p, c))
	i, err := NewIntValue(c.GetText())
	if err != nil {
		l.l.Error("NewIntValue", "err", err)
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&IntExpr{
		BaseExpr: self.(*BaseExpr),
		Val:      i,
	})
}

func (l *Listener) EnterVarExpr(c *parser.VarExprContext) {
	l.l.Info("EnterVarExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&VarExpr{
		BaseExpr: self.(*BaseExpr),
		Name:     c.VARNAME().GetText(),
	})
}

func (l *Listener) EnterBinOpExpr(c *parser.BinOpExprContext) {
	l.l.Info("EnterBinOpExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}

	l.Push(&BinOpExpr{
		BaseExpr: self.(*BaseExpr),
		Op:       c.GetOp(),
	})
}

func (l *Listener) EnterIfExpr(c *parser.IfExprContext) {
	l.l.Info("EnterIfExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&IfExpr{
		BaseExpr: self.(*BaseExpr),
	})
}

func (l *Listener) EnterLetExpr(c *parser.LetExprContext) {
	l.l.Info("EnterLetExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&LetExpr{
		BaseExpr: self.(*BaseExpr),
		Var:      c.GetBindName().GetText(),
	})
}

func (l *Listener) EnterLetRecExpr(c *parser.LetRecExprContext) {
	l.l.Info("EnterLetRecExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&LetRecExpr{
		BaseExpr: self.(*BaseExpr),
		RecFun: RecFun{
			Name: c.RecFun().GetFunName().GetText(),
			Fun: Fun{
				Param: c.RecFun().Fun().GetParam().GetText(),
			},
		},
	})
}

func (l *Listener) EnterAppExpr(c *parser.AppExprContext) {
	l.l.Info("EnterAppExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&AppExpr{
		BaseExpr: self.(*BaseExpr),
	})
}

func (l *Listener) EnterFunExpr(c *parser.FunExprContext) {
	l.l.Info("EnterFunExpr", "literal", getLiteral(l.p, c))
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	l.Push(&FunExpr{
		BaseExpr: self.(*BaseExpr),
		Fun: Fun{
			Param: c.Fun().GetParam().GetText(),
		},
	})
}

func (l *Listener) ExitEveryRule(c antlr.ParserRuleContext) {
	if ec, ok := c.(parser.IExprContext); ok {
		l.ExitExpr(ec)
	}
}

func (l *Listener) ExitExpr(c parser.IExprContext) {
	l.DecrDepth()
}

func (l *Listener) ExitDef(c *parser.DefContext) {
	switch vc := c.Value().(type) {
	case *parser.IntValueContext:
		i, err := strconv.Atoi(vc.GetText())
		if err != nil {
			l.l.Error("strconv.Atoi", "err", err)
			return
		}
		l.env = append(l.env, Bind{
			Name: c.VARNAME().GetText(),
			Val:  IntValue(i),
		})
	case *parser.BoolValueContext:
		b, err := strconv.ParseBool(vc.GetText())
		if err != nil {
			l.l.Error("strconv.ParseBool", "err", err)
			return
		}
		l.env = append(l.env, Bind{
			Name: c.VARNAME().GetText(),
			Val:  BoolValue(b),
		})
	case *parser.FunValueContext:
		body, ok := l.Pop()
		if !ok {
			l.l.Error("no body found")
			return
		}
		funEnvLen := 0
		if c := vc.DefList(); c != nil {
			funEnvLen = len(c.AllDef())
		}
		funEnv := slices.Clone(l.env[len(l.env)-funEnvLen:])
		l.env = append(slices.Delete(l.env, len(l.env)-funEnvLen, len(l.env)), Bind{
			Name: c.VARNAME().GetText(),
			Val: FunValue{
				Fun: Fun{
					Param: vc.Fun().GetParam().GetText(),
					Body:  body,
				},
				Env: funEnv,
			},
		})
	case *parser.RecFunValueContext:
		body, ok := l.Pop()
		if !ok {
			l.l.Error("no body found")
			return
		}
		funEnvLen := 0
		if c := vc.DefList(); c != nil {
			funEnvLen = len(c.AllDef())
		}
		funEnv := slices.Clone(l.env[len(l.env)-funEnvLen:])
		l.env = append(slices.Delete(l.env, len(l.env)-funEnvLen, len(l.env)), Bind{
			Name: c.VARNAME().GetText(),
			Val: RecFunValue{
				RecFun: RecFun{
					Name: vc.RecFun().GetFunName().GetText(),
					Fun: Fun{
						Param: vc.RecFun().Fun().GetParam().GetText(),
						Body:  body,
					},
				},
				Env: funEnv,
			},
		})
	default:
		l.l.Error("ExitDef: unknown value", "literal", getLiteral(l.p, c))
		return
	}
	bind := l.env[len(l.env)-1]
	l.l.Info("ExitDef: bind value", "literal", getLiteral(l.p, c), "bind", bind)
}

func (l *Listener) ExitParenExpr(c *parser.ParenExprContext) {
	defer l.logExpr("ExitParenExpr")
}

func (l *Listener) ExitBoolExpr(c *parser.BoolExprContext) {
	defer l.logExpr("ExitBoolExpr")
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*BoolExpr)
	if !ok {
		l.l.Error("not a BoolExpr")
		return
	}
	expr.evalTo = expr.Val
	expr.rule = EBool
	l.Push(expr)
}

func (l *Listener) ExitIntExpr(c *parser.IntExprContext) {
	defer l.logExpr("ExitIntExpr")
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*IntExpr)
	if !ok {
		l.l.Error("not a IntExpr")
		return
	}
	expr.evalTo = expr.Val
	expr.rule = EInt
	l.Push(expr)
}

func (l *Listener) ExitNegIntExpr(c *parser.NegIntExprContext) {
	defer l.logExpr("ExitNegIntExpr")
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*IntExpr)
	if !ok {
		l.l.Error("not a IntExpr")
		return
	}
	expr.evalTo = expr.Val
	expr.rule = EInt
	l.Push(expr)
}

func (l *Listener) ExitVarExpr(c *parser.VarExprContext) {
	defer l.logExpr("ExitVarExpr")
	self, ok := l.Pop()
	if !ok {
		l.l.Error("No self expr")
		return
	}
	expr, ok := self.(*VarExpr)
	if !ok {
		l.l.Error("not a VarExpr")
		return
	}

	for _, bind := range self.Env() {
		if bind.Name == expr.Name {
			expr.evalTo = bind.Val
			break
		}
	}
	if expr.evalTo == nil {
		l.Push(expr)
		return
	}

	current := expr
	for env := slices.Clone(expr.Env()); len(env) > 0; env = env[:len(env)-1] {
		current.env = slices.Clone(env)
		if last := env[len(env)-1]; last.Name == current.Name {
			current.rule = EVar1
			break
		}
		current.rule = EVar2
		next := &VarExpr{
			BaseExpr: current.BaseExpr.Clone(),
			Name:     current.Name,
		}
		current.AddChild(next)
		current = next
	}
	l.Push(expr)
}

func (l *Listener) ExitBinOpExpr(c *parser.BinOpExprContext) {
	l.l.Info("ExitBinOpExpr", "literal", getLiteral(l.p, c))
	defer l.logExpr("ExitBinOpExpr")
	right, ok := l.Pop()
	if !ok {
		l.l.Error("No right expr")
		return
	}
	left, ok := l.Pop()
	if !ok {
		l.l.Error("No left expr")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("No self expr")
		return
	}

	expr, ok := self.(*BinOpExpr)
	if !ok {
		l.l.Error("not a BinOpExpr")
		return
	}

	lVal, leftIsInt := left.EvalTo().(IntValue)
	if !leftIsInt {
		l.l.Debug("left is not an IntValue", "left", left.Literal())
	}
	rVal, rightIsInt := right.EvalTo().(IntValue)
	if !rightIsInt {
		l.l.Debug("right is not an IntValue", "right", right.Literal())
	}
	if !(leftIsInt && rightIsInt) {
		expr.Left = left
		expr.Right = right
		l.Push(expr)
		return
	}
	l.l.Info("ExitBinOpExpr", "left", lVal, "op", expr.Op.GetText(), "right", rVal)

	expr.AddChild(left)
	expr.AddChild(right)
	expr.AddChild(NewBExpr(lVal, rVal, expr.Op))
	switch expr.Op.GetTokenType() {
	case parser.EvalML3ParserPLUS:
		expr.evalTo = IntValue(lVal + rVal)
		expr.rule = EPlus
	case parser.EvalML3ParserMINUS:
		expr.evalTo = IntValue(lVal - rVal)
		expr.rule = EMinus
	case parser.EvalML3ParserTIMES:
		expr.evalTo = IntValue(lVal * rVal)
		expr.rule = ETimes
	case parser.EvalML3ParserLT:
		expr.evalTo = BoolValue(lVal < rVal)
		expr.rule = ELt
	default:
		l.l.Error("Unknown op", "op", expr.Op.GetText())
	}
	l.Push(expr)
}

func (l *Listener) ExitIfExpr(c *parser.IfExprContext) {
	defer l.logExpr("ExitIfExpr")

	els, ok := l.Pop()
	if !ok {
		l.l.Error("No else expr")
		return
	}
	then, ok := l.Pop()
	if !ok {
		l.l.Error("No then expr")
		return
	}
	cond, ok := l.Pop()
	if !ok {
		l.l.Error("No condition expr")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("No self expr")
		return
	}
	expr, ok := self.(*IfExpr)
	if !ok {
		l.l.Error("not a IfExpr")
		return
	}

	condValue, ok := cond.EvalTo().(BoolValue)
	if !ok {
		l.l.Debug("condition is not a BoolValue", "cond", cond.EvalTo())
		l.Push(self)
		return
	}

	expr.AddChild(cond)
	if condValue {
		expr.evalTo = then.EvalTo()
		expr.rule = EIfT
		expr.AddChild(then)
	} else {
		expr.evalTo = els.EvalTo()
		expr.rule = EIfF
		expr.AddChild(els)
	}
	l.Push(expr)
}

func (l *Listener) ExitLetExpr(c *parser.LetExprContext) {
	l.l.Debug("ExitLetExpr", "literal", getLiteral(l.p, c))

	body, ok := l.Pop()
	if !ok {
		l.l.Error("no body found")
		return
	}
	bind, ok := l.Pop()
	if !ok {
		l.l.Error("no bind found")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*LetExpr)
	if !ok {
		l.l.Error("not a LetExpr")
		return
	}
	expr.AddChild(bind)
	expr.Bind = bind
	expr.AddChild(body)
	expr.Body = body
	expr.evalTo = body.EvalTo()
	expr.rule = ELet

	l.Push(expr)
}

func (l *Listener) ExitLetRecExpr(c *parser.LetRecExprContext) {
	defer l.logExpr("ExitLetRecExpr")
	body, ok := l.Pop()
	if !ok {
		l.l.Error("no body found")
		return
	}
	bind, ok := l.Pop()
	if !ok {
		l.l.Error("no bind found")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*LetRecExpr)
	if !ok {
		l.l.Error("not a LetRecExpr", "self", self.Literal())
		return
	}

	l.l.Error("not implemented", "expr", expr.Literal(), "body", body.Literal(), "bind", bind.Literal())
	expr.AddChild(body)
	expr.Body = body
	expr.evalTo = body.EvalTo()
	expr.rule = ELetRec

	l.Push(expr)
}

func (l *Listener) ExitAppExpr(c *parser.AppExprContext) {
	defer l.logExpr("ExitAppExpr")
	arg, ok := l.Pop()
	if !ok {
		l.l.Error("no arg found")
		return
	}
	fun, ok := l.Pop()
	if !ok {
		l.l.Error("no fun found")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*AppExpr)
	if !ok {
		l.l.Error("not a AppExpr", "arg", arg, "fun", fun, "self", fmt.Sprintf("%+v", self))
		l.Push(self)
		l.logExpr("ExitAppExpr: not a AppExpr")
		return
	}
	if expr.Skip() {
		l.l.Info("ExitAppExpr: skip", "arg", arg, "fun", fun, "expr", getLiteral(l.p, c))
		l.Push(expr)
		return
	}
	if arg.EvalTo() == nil {
		l.l.Warn("ExitAppExpr: arg.EvalTo() == nil", "arg", arg, "fun", fun, "expr", getLiteral(l.p, c))
		l.Push(expr)
		return
	}

	if recFunVal, recFunEvaled := fun.EvalTo().(RecFunValue); recFunEvaled {
		l.l.Info("ExitAppExpr: recFunEvaled", "arg", arg, "fun", fun, "expr", fmt.Sprintf("%+v", expr), "recFunVal", recFunVal)
		env := append(slices.Clone(recFunVal.Env),
			Bind{
				Name: recFunVal.RecFun.Name,
				Val:  recFunVal,
			},
			Bind{
				Name: recFunVal.RecFun.Fun.Param,
				Val:  arg.EvalTo(),
			},
		)

		e := recFunVal.RecFun.Fun.Body
		s := fmt.Sprintf("%s |- %s", env, e.Literal())
		app, err := evalExpr(s, l.l)
		if err != nil {
			l.l.Error("evalExpr", "err", err)
			return
		}
		expr.evalTo = app.EvalTo()
		expr.AddChild(fun)
		expr.AddChild(arg)
		expr.AddChild(app)
		expr.rule = EAppRec
		l.Push(expr)
		return
	}

	if funVal, funEvaled := fun.EvalTo().(FunValue); funEvaled {
		l.l.Info("ExitAppExpr: funEvaled", "arg", arg, "fun", fun, "expr", fmt.Sprintf("%+v", expr), "funVal", funVal)
		env := append(slices.Clone(funVal.Env), Bind{
			Name: funVal.Fun.Param,
			Val:  arg.EvalTo(),
		})
		e := funVal.Fun.Body
		s := fmt.Sprintf("%s |- %s", env, e.Literal())
		app, err := evalExpr(s, l.l)
		if err != nil {
			l.l.Error("evalExpr", "err", err)
			return
		}
		expr.evalTo = app.EvalTo()
		expr.AddChild(fun)
		expr.AddChild(arg)
		expr.AddChild(app)
		expr.rule = EApp
		l.Push(expr)
		return
	}

	expr.Fun = fun
	expr.Arg = arg
	l.Push(expr)
}

func (l *Listener) ExitFunExpr(c *parser.FunExprContext) {
	defer l.logExpr("ExitFunExpr")

	body, ok := l.Pop()
	if !ok {
		l.l.Error("no body found")
		return
	}
	self, ok := l.Pop()
	if !ok {
		l.l.Error("no expression found")
		return
	}
	expr, ok := self.(*FunExpr)
	if !ok {
		l.l.Error("not a FunExpr", "self", fmt.Sprintf("%+v", self))
		return
	}
	expr.Fun.Body = body
	expr.evalTo = FunValue{
		Fun: Fun{
			Param: expr.Param,
			Body:  body,
		},
		Env: expr.Env(),
	}
	expr.rule = EFun
	l.Push(expr)
}

func (l *Listener) ExitDefList(c *parser.DefListContext) {
	parent, siblings, ok := l.PeekParent()
	l.l.Debug("ExitDefList", "literal", getLiteral(l.p, c), "parent", parent, "siblings", siblings, "ok", ok)
}

func (l *Listener) EnterFunValue(c *parser.FunValueContext) {
	l.l.Info("EnterFunValue", "literal", getLiteral(l.p, c))
}

func (l *Listener) ExitFunValue(c *parser.FunValueContext) {
	l.l.Info("ExitFunValue", "literal", getLiteral(l.p, c))
}

func (l *Listener) logExpr(caller string) {
	if len(l.result) == 0 {
		l.l.Error("logExpr: no expression found", "caller", caller)
		return
	}
	e := l.result[len(l.result)-1]
	l.l.Info(caller,
		"id", e.ID(),
		"type", fmt.Sprintf("%T", e),
		"env", e.Env(),
		"literal", e.Literal(),
		"evalTo", e.EvalTo(),
		"rule", e.Rule(),
		"stackSize", len(l.result),
	)
}

func Push[Slice ~[]E, E any](s *Slice, e E) {
	*s = append(*s, e)
}

func Pop[Slice ~[]E, E any](s *Slice) E {
	if len(*s) == 0 {
		panic("Pop: empty slice")
	}
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func Peek[Slice ~[]E, E any](s *Slice) E {
	if len(*s) == 0 {
		panic("Peek: empty slice")
	}
	return (*s)[len(*s)-1]
}
