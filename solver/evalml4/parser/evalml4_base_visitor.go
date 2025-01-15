// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML4
import "github.com/antlr4-go/antlr/v4"

type BaseEvalML4Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvalML4Visitor) VisitQuestion(ctx *QuestionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitEval(ctx *EvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitLetRecExpr(ctx *LetRecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitAppExpr(ctx *AppExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitEmptyListExpr(ctx *EmptyListExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitConsExpr(ctx *ConsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitBinOpExpr(ctx *BinOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitFunExpr(ctx *FunExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitDefList(ctx *DefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitDef(ctx *DefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitEmptyPattern(ctx *EmptyPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitConsPattern(ctx *ConsPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitFun(ctx *FunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitRecFun(ctx *RecFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitEmptyListValue(ctx *EmptyListValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitFunValue(ctx *FunValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitRecFunValue(ctx *RecFunValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML4Visitor) VisitConsValue(ctx *ConsValueContext) interface{} {
	return v.VisitChildren(ctx)
}
