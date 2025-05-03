// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML1
import "github.com/antlr4-go/antlr/v4"

type BaseEvalContML1Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvalContML1Visitor) VisitEval(ctx *EvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitIntValue(ctx *IntValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitBoolExp(ctx *BoolExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitIfExp(ctx *IfExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitIntExp(ctx *IntExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitBinOpExp(ctx *BinOpExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitUnaryCont(ctx *UnaryContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitExpCont(ctx *ExpContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitValueCont(ctx *ValueContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvalContML1Visitor) VisitIfCont(ctx *IfContContext) interface{} {
	return v.VisitChildren(ctx)
}
