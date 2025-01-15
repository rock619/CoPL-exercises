// Code generated from EvalML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML1
import "github.com/antlr4-go/antlr/v4"

type BaseEvalML1Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvalML1Visitor) VisitQuestion(ctx *QuestionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitEval(ctx *EvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitBinOpExpr(ctx *BinOpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitNegIntExpr(ctx *NegIntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalML1Visitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}
