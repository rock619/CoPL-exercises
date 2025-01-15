// Code generated from EvalML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML1
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type EvalML1Parser struct {
	*antlr.BaseParser
}

var EvalML1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalml1ParserInit() {
	staticData := &EvalML1ParserStaticData
	staticData.LiteralNames = []string{
		"", "'evalto'", "'('", "')'", "'if'", "'then'", "'else'", "'+'", "'-'",
		"'*'", "'<'", "", "", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES", "LT", "INT", "BOOL",
		"TRUE", "FALSE", "WS",
	}
	staticData.RuleNames = []string{
		"question", "eval", "expr", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 52, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 32, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 43, 8, 2, 10,
		2, 12, 2, 46, 9, 2, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 3, 0, 1, 4, 4, 0, 2,
		4, 6, 0, 1, 1, 0, 7, 8, 55, 0, 8, 1, 0, 0, 0, 2, 11, 1, 0, 0, 0, 4, 31,
		1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1, 10,
		1, 1, 0, 0, 0, 11, 12, 3, 4, 2, 0, 12, 13, 5, 1, 0, 0, 13, 14, 3, 6, 3,
		0, 14, 3, 1, 0, 0, 0, 15, 16, 6, 2, -1, 0, 16, 32, 5, 12, 0, 0, 17, 32,
		5, 11, 0, 0, 18, 19, 5, 8, 0, 0, 19, 32, 5, 11, 0, 0, 20, 21, 5, 2, 0,
		0, 21, 22, 3, 4, 2, 0, 22, 23, 5, 3, 0, 0, 23, 32, 1, 0, 0, 0, 24, 25,
		5, 4, 0, 0, 25, 26, 3, 4, 2, 0, 26, 27, 5, 5, 0, 0, 27, 28, 3, 4, 2, 0,
		28, 29, 5, 6, 0, 0, 29, 30, 3, 4, 2, 1, 30, 32, 1, 0, 0, 0, 31, 15, 1,
		0, 0, 0, 31, 17, 1, 0, 0, 0, 31, 18, 1, 0, 0, 0, 31, 20, 1, 0, 0, 0, 31,
		24, 1, 0, 0, 0, 32, 44, 1, 0, 0, 0, 33, 34, 10, 4, 0, 0, 34, 35, 5, 9,
		0, 0, 35, 43, 3, 4, 2, 5, 36, 37, 10, 3, 0, 0, 37, 38, 7, 0, 0, 0, 38,
		43, 3, 4, 2, 4, 39, 40, 10, 2, 0, 0, 40, 41, 5, 10, 0, 0, 41, 43, 3, 4,
		2, 3, 42, 33, 1, 0, 0, 0, 42, 36, 1, 0, 0, 0, 42, 39, 1, 0, 0, 0, 43, 46,
		1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 5, 1, 0, 0, 0,
		46, 44, 1, 0, 0, 0, 47, 50, 5, 11, 0, 0, 48, 50, 5, 12, 0, 0, 49, 47, 1,
		0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 4, 31, 42, 44, 49,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EvalML1ParserInit initializes any static state used to implement EvalML1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalML1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalML1ParserInit() {
	staticData := &EvalML1ParserStaticData
	staticData.once.Do(evalml1ParserInit)
}

// NewEvalML1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalML1Parser(input antlr.TokenStream) *EvalML1Parser {
	EvalML1ParserInit()
	this := new(EvalML1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalML1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalML1.g4"

	return this
}

// EvalML1Parser tokens.
const (
	EvalML1ParserEOF   = antlr.TokenEOF
	EvalML1ParserT__0  = 1
	EvalML1ParserT__1  = 2
	EvalML1ParserT__2  = 3
	EvalML1ParserT__3  = 4
	EvalML1ParserT__4  = 5
	EvalML1ParserT__5  = 6
	EvalML1ParserPLUS  = 7
	EvalML1ParserMINUS = 8
	EvalML1ParserTIMES = 9
	EvalML1ParserLT    = 10
	EvalML1ParserINT   = 11
	EvalML1ParserBOOL  = 12
	EvalML1ParserTRUE  = 13
	EvalML1ParserFALSE = 14
	EvalML1ParserWS    = 15
)

// EvalML1Parser rules.
const (
	EvalML1ParserRULE_question = 0
	EvalML1ParserRULE_eval     = 1
	EvalML1ParserRULE_expr     = 2
	EvalML1ParserRULE_value    = 3
)

// IQuestionContext is an interface to support dynamic dispatch.
type IQuestionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Eval() IEvalContext
	EOF() antlr.TerminalNode

	// IsQuestionContext differentiates from other interfaces.
	IsQuestionContext()
}

type QuestionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuestionContext() *QuestionContext {
	var p = new(QuestionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_question
	return p
}

func InitEmptyQuestionContext(p *QuestionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_question
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML1ParserRULE_question

	return p
}

func (s *QuestionContext) GetParser() antlr.Parser { return s.parser }

func (s *QuestionContext) Eval() IEvalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEvalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEvalContext)
}

func (s *QuestionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserEOF, 0)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (s *QuestionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitQuestion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML1Parser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalML1ParserRULE_question)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Eval()
	}
	{
		p.SetState(9)
		p.Match(EvalML1ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEvalContext is an interface to support dynamic dispatch.
type IEvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Value() IValueContext

	// IsEvalContext differentiates from other interfaces.
	IsEvalContext()
}

type EvalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvalContext() *EvalContext {
	var p = new(EvalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML1ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EvalContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML1Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalML1ParserRULE_eval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(11)
		p.expr(0)
	}
	{
		p.SetState(12)
		p.Match(EvalML1ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(13)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML1ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfExprContext struct {
	ExprContext
	cond  IExprContext
	then  IExprContext
	else_ IExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetCond() IExprContext { return s.cond }

func (s *IfExprContext) GetThen() IExprContext { return s.then }

func (s *IfExprContext) GetElse_() IExprContext { return s.else_ }

func (s *IfExprContext) SetCond(v IExprContext) { s.cond = v }

func (s *IfExprContext) SetThen(v IExprContext) { s.then = v }

func (s *IfExprContext) SetElse_(v IExprContext) { s.else_ = v }

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinOpExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewBinOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinOpExprContext {
	var p = new(BinOpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BinOpExprContext) GetOp() antlr.Token { return s.op }

func (s *BinOpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinOpExprContext) GetLeft() IExprContext { return s.left }

func (s *BinOpExprContext) GetRight() IExprContext { return s.right }

func (s *BinOpExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *BinOpExprContext) SetRight(v IExprContext) { s.right = v }

func (s *BinOpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinOpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinOpExprContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserTIMES, 0)
}

func (s *BinOpExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserPLUS, 0)
}

func (s *BinOpExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserMINUS, 0)
}

func (s *BinOpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserLT, 0)
}

func (s *BinOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterBinOpExpr(s)
	}
}

func (s *BinOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitBinOpExpr(s)
	}
}

func (s *BinOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitBinOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegIntExprContext struct {
	ExprContext
}

func NewNegIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegIntExprContext {
	var p = new(NegIntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegIntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegIntExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserMINUS, 0)
}

func (s *NegIntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserINT, 0)
}

func (s *NegIntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterNegIntExpr(s)
	}
}

func (s *NegIntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitNegIntExpr(s)
	}
}

func (s *NegIntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitNegIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML1Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *EvalML1Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, EvalML1ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalML1ParserBOOL:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(16)
			p.Match(EvalML1ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML1ParserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(EvalML1ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML1ParserMINUS:
		localctx = NewNegIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(EvalML1ParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(EvalML1ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML1ParserT__1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(EvalML1ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.expr(0)
		}
		{
			p.SetState(22)
			p.Match(EvalML1ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML1ParserT__3:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(EvalML1ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)

			var _x = p.expr(0)

			localctx.(*IfExprContext).cond = _x
		}
		{
			p.SetState(26)
			p.Match(EvalML1ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)

			var _x = p.expr(0)

			localctx.(*IfExprContext).then = _x
		}
		{
			p.SetState(28)
			p.Match(EvalML1ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)

			var _x = p.expr(1)

			localctx.(*IfExprContext).else_ = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML1ParserRULE_expr)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(34)

					var _m = p.Match(EvalML1ParserTIMES)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(35)

					var _x = p.expr(5)

					localctx.(*BinOpExprContext).right = _x
				}

			case 2:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML1ParserRULE_expr)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(37)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalML1ParserPLUS || _la == EvalML1ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(38)

					var _x = p.expr(4)

					localctx.(*BinOpExprContext).right = _x
				}

			case 3:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML1ParserRULE_expr)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(40)

					var _m = p.Match(EvalML1ParserLT)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(41)

					var _x = p.expr(3)

					localctx.(*BinOpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML1ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML1ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolValueContext struct {
	ValueContext
}

func NewBoolValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolValueContext {
	var p = new(BoolValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitBoolValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntValueContext struct {
	ValueContext
}

func NewIntValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntValueContext {
	var p = new(IntValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) INT() antlr.TerminalNode {
	return s.GetToken(EvalML1ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML1Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML1Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML1Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalML1ParserRULE_value)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalML1ParserINT:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(EvalML1ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML1ParserBOOL:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(EvalML1ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *EvalML1Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvalML1Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
