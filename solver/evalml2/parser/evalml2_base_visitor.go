// Code generated from EvalML2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML2
import "github.com/antlr4-go/antlr/v4"

type BaseEvalML2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvalML2Visitor) VisitQuestion(ctx *QuestionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitEval(ctx *EvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitBinOpExpr(ctx *BinOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitNegIntExpr(ctx *NegIntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitDefList(ctx *DefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitDef(ctx *DefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML2Visitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}
