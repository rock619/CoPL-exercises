// Code generated from EvalML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML3
import "github.com/antlr4-go/antlr/v4"

type BaseEvalML3Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvalML3Visitor) VisitQuestion(ctx *QuestionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitEval(ctx *EvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitLetRecExpr(ctx *LetRecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitAppExpr(ctx *AppExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitBinOpExpr(ctx *BinOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitNegIntExpr(ctx *NegIntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitFunExpr(ctx *FunExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitDefList(ctx *DefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitDef(ctx *DefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitFun(ctx *FunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitRecFun(ctx *RecFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitFunValue(ctx *FunValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML3Visitor) VisitRecFunValue(ctx *RecFunValueContext) interface{} {
	return v.VisitChildren(ctx)
}
