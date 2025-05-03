// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML1
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalContML1Parser.
type EvalContML1Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalContML1Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#BoolExp.
	VisitBoolExp(ctx *BoolExpContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#IfExp.
	VisitIfExp(ctx *IfExpContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#IntExp.
	VisitIntExp(ctx *IntExpContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#BinOpExp.
	VisitBinOpExp(ctx *BinOpExpContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#UnaryCont.
	VisitUnaryCont(ctx *UnaryContContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#ExpCont.
	VisitExpCont(ctx *ExpContContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#ValueCont.
	VisitValueCont(ctx *ValueContContext) interface{}

	// Visit a parse tree produced by EvalContML1Parser#IfCont.
	VisitIfCont(ctx *IfContContext) interface{}
}
