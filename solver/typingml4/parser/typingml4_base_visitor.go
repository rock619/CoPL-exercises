// Code generated from TypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TypingML4
import "github.com/antlr4-go/antlr/v4"

type BaseTypingML4Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTypingML4Visitor) VisitJudgement(ctx *JudgementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitListType(ctx *ListTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitFunType(ctx *FunTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitParenType(ctx *ParenTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitLetRecExpr(ctx *LetRecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitAppExpr(ctx *AppExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitEmptyListExpr(ctx *EmptyListExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitConsExpr(ctx *ConsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitBinOpExpr(ctx *BinOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitFunExpr(ctx *FunExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitEnv(ctx *EnvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitBind(ctx *BindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitEmptyPattern(ctx *EmptyPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitConsPattern(ctx *ConsPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitFun(ctx *FunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypingML4Visitor) VisitRecFun(ctx *RecFunContext) interface{} {
	return v.VisitChildren(ctx)
}
