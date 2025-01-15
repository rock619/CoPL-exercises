// Code generated from EvalML2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML2
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

type EvalML2Parser struct {
	*antlr.BaseParser
}

var EvalML2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalml2ParserInit() {
	staticData := &EvalML2ParserStaticData
	staticData.LiteralNames = []string{
		"", "'evalto'", "'('", "')'", "'if'", "'then'", "'else'", "'let'", "'='",
		"'in'", "'+'", "'-'", "'*'", "'<'", "", "", "'true'", "'false'", "",
		"','", "'|-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES", "LT",
		"INT", "BOOL", "TRUE", "FALSE", "VARNAME", "COMMA", "TURNSTILE", "WS",
	}
	staticData.RuleNames = []string{
		"question", "eval", "expr", "defList", "def", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 80, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 17, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 59, 8, 2, 10, 2, 12, 2, 62, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3,
		67, 8, 3, 10, 3, 12, 3, 70, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3,
		5, 78, 8, 5, 1, 5, 0, 1, 4, 6, 0, 2, 4, 6, 8, 10, 0, 1, 1, 0, 10, 11, 85,
		0, 12, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 63, 1, 0, 0,
		0, 8, 71, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5,
		0, 0, 1, 14, 1, 1, 0, 0, 0, 15, 17, 3, 6, 3, 0, 16, 15, 1, 0, 0, 0, 16,
		17, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 19, 5, 20, 0, 0, 19, 20, 3, 4,
		2, 0, 20, 21, 5, 1, 0, 0, 21, 22, 3, 10, 5, 0, 22, 3, 1, 0, 0, 0, 23, 24,
		6, 2, -1, 0, 24, 48, 5, 15, 0, 0, 25, 48, 5, 14, 0, 0, 26, 27, 5, 11, 0,
		0, 27, 48, 5, 14, 0, 0, 28, 48, 5, 18, 0, 0, 29, 30, 5, 2, 0, 0, 30, 31,
		3, 4, 2, 0, 31, 32, 5, 3, 0, 0, 32, 48, 1, 0, 0, 0, 33, 34, 5, 4, 0, 0,
		34, 35, 3, 4, 2, 0, 35, 36, 5, 5, 0, 0, 36, 37, 3, 4, 2, 0, 37, 38, 5,
		6, 0, 0, 38, 39, 3, 4, 2, 2, 39, 48, 1, 0, 0, 0, 40, 41, 5, 7, 0, 0, 41,
		42, 5, 18, 0, 0, 42, 43, 5, 8, 0, 0, 43, 44, 3, 4, 2, 0, 44, 45, 5, 9,
		0, 0, 45, 46, 3, 4, 2, 1, 46, 48, 1, 0, 0, 0, 47, 23, 1, 0, 0, 0, 47, 25,
		1, 0, 0, 0, 47, 26, 1, 0, 0, 0, 47, 28, 1, 0, 0, 0, 47, 29, 1, 0, 0, 0,
		47, 33, 1, 0, 0, 0, 47, 40, 1, 0, 0, 0, 48, 60, 1, 0, 0, 0, 49, 50, 10,
		5, 0, 0, 50, 51, 5, 12, 0, 0, 51, 59, 3, 4, 2, 6, 52, 53, 10, 4, 0, 0,
		53, 54, 7, 0, 0, 0, 54, 59, 3, 4, 2, 5, 55, 56, 10, 3, 0, 0, 56, 57, 5,
		13, 0, 0, 57, 59, 3, 4, 2, 4, 58, 49, 1, 0, 0, 0, 58, 52, 1, 0, 0, 0, 58,
		55, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0,
		0, 61, 5, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 68, 3, 8, 4, 0, 64, 65, 5,
		19, 0, 0, 65, 67, 3, 8, 4, 0, 66, 64, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68,
		66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 7, 1, 0, 0, 0, 70, 68, 1, 0, 0,
		0, 71, 72, 5, 18, 0, 0, 72, 73, 5, 8, 0, 0, 73, 74, 3, 10, 5, 0, 74, 9,
		1, 0, 0, 0, 75, 78, 5, 14, 0, 0, 76, 78, 5, 15, 0, 0, 77, 75, 1, 0, 0,
		0, 77, 76, 1, 0, 0, 0, 78, 11, 1, 0, 0, 0, 6, 16, 47, 58, 60, 68, 77,
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

// EvalML2ParserInit initializes any static state used to implement EvalML2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalML2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalML2ParserInit() {
	staticData := &EvalML2ParserStaticData
	staticData.once.Do(evalml2ParserInit)
}

// NewEvalML2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalML2Parser(input antlr.TokenStream) *EvalML2Parser {
	EvalML2ParserInit()
	this := new(EvalML2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalML2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalML2.g4"

	return this
}

// EvalML2Parser tokens.
const (
	EvalML2ParserEOF       = antlr.TokenEOF
	EvalML2ParserT__0      = 1
	EvalML2ParserT__1      = 2
	EvalML2ParserT__2      = 3
	EvalML2ParserT__3      = 4
	EvalML2ParserT__4      = 5
	EvalML2ParserT__5      = 6
	EvalML2ParserT__6      = 7
	EvalML2ParserT__7      = 8
	EvalML2ParserT__8      = 9
	EvalML2ParserPLUS      = 10
	EvalML2ParserMINUS     = 11
	EvalML2ParserTIMES     = 12
	EvalML2ParserLT        = 13
	EvalML2ParserINT       = 14
	EvalML2ParserBOOL      = 15
	EvalML2ParserTRUE      = 16
	EvalML2ParserFALSE     = 17
	EvalML2ParserVARNAME   = 18
	EvalML2ParserCOMMA     = 19
	EvalML2ParserTURNSTILE = 20
	EvalML2ParserWS        = 21
)

// EvalML2Parser rules.
const (
	EvalML2ParserRULE_question = 0
	EvalML2ParserRULE_eval     = 1
	EvalML2ParserRULE_expr     = 2
	EvalML2ParserRULE_defList  = 3
	EvalML2ParserRULE_def      = 4
	EvalML2ParserRULE_value    = 5
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
	p.RuleIndex = EvalML2ParserRULE_question
	return p
}

func InitEmptyQuestionContext(p *QuestionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_question
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_question

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
	return s.GetToken(EvalML2ParserEOF, 0)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (s *QuestionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitQuestion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalML2ParserRULE_question)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Eval()
	}
	{
		p.SetState(13)
		p.Match(EvalML2ParserEOF)
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
	TURNSTILE() antlr.TerminalNode
	Expr() IExprContext
	Value() IValueContext
	DefList() IDefListContext

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
	p.RuleIndex = EvalML2ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) TURNSTILE() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserTURNSTILE, 0)
}

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

func (s *EvalContext) DefList() IDefListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefListContext)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalML2ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalML2ParserVARNAME {
		{
			p.SetState(15)
			p.DefList()
		}

	}
	{
		p.SetState(18)
		p.Match(EvalML2ParserTURNSTILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
		p.expr(0)
	}
	{
		p.SetState(20)
		p.Match(EvalML2ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
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
	p.RuleIndex = EvalML2ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_expr

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
	return s.GetToken(EvalML2ParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
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
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExprContext struct {
	ExprContext
	bindName antlr.Token
	bindExpr IExprContext
	body     IExprContext
}

func NewLetExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetExprContext {
	var p = new(LetExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LetExprContext) GetBindName() antlr.Token { return s.bindName }

func (s *LetExprContext) SetBindName(v antlr.Token) { s.bindName = v }

func (s *LetExprContext) GetBindExpr() IExprContext { return s.bindExpr }

func (s *LetExprContext) GetBody() IExprContext { return s.body }

func (s *LetExprContext) SetBindExpr(v IExprContext) { s.bindExpr = v }

func (s *LetExprContext) SetBody(v IExprContext) { s.body = v }

func (s *LetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExprContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserVARNAME, 0)
}

func (s *LetExprContext) AllExpr() []IExprContext {
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

func (s *LetExprContext) Expr(i int) IExprContext {
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

func (s *LetExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterLetExpr(s)
	}
}

func (s *LetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitLetExpr(s)
	}
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitLetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExprContext struct {
	ExprContext
}

func NewVarExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprContext {
	var p = new(VarExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserVARNAME, 0)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitVarExpr(s)

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
	return s.GetToken(EvalML2ParserTIMES, 0)
}

func (s *BinOpExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserPLUS, 0)
}

func (s *BinOpExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserMINUS, 0)
}

func (s *BinOpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserLT, 0)
}

func (s *BinOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterBinOpExpr(s)
	}
}

func (s *BinOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitBinOpExpr(s)
	}
}

func (s *BinOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
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
	return s.GetToken(EvalML2ParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
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
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
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
	return s.GetToken(EvalML2ParserMINUS, 0)
}

func (s *NegIntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserINT, 0)
}

func (s *NegIntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterNegIntExpr(s)
	}
}

func (s *NegIntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitNegIntExpr(s)
	}
}

func (s *NegIntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitNegIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *EvalML2Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, EvalML2ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalML2ParserBOOL:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(24)
			p.Match(EvalML2ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(EvalML2ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserMINUS:
		localctx = NewNegIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(EvalML2ParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Match(EvalML2ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserVARNAME:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(28)
			p.Match(EvalML2ParserVARNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserT__1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(EvalML2ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.expr(0)
		}
		{
			p.SetState(31)
			p.Match(EvalML2ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserT__3:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(EvalML2ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)

			var _x = p.expr(0)

			localctx.(*IfExprContext).cond = _x
		}
		{
			p.SetState(35)
			p.Match(EvalML2ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)

			var _x = p.expr(0)

			localctx.(*IfExprContext).then = _x
		}
		{
			p.SetState(37)
			p.Match(EvalML2ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)

			var _x = p.expr(2)

			localctx.(*IfExprContext).else_ = _x
		}

	case EvalML2ParserT__6:
		localctx = NewLetExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(EvalML2ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)

			var _m = p.Match(EvalML2ParserVARNAME)

			localctx.(*LetExprContext).bindName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(EvalML2ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)

			var _x = p.expr(0)

			localctx.(*LetExprContext).bindExpr = _x
		}
		{
			p.SetState(44)
			p.Match(EvalML2ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)

			var _x = p.expr(1)

			localctx.(*LetExprContext).body = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML2ParserRULE_expr)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(50)

					var _m = p.Match(EvalML2ParserTIMES)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(51)

					var _x = p.expr(6)

					localctx.(*BinOpExprContext).right = _x
				}

			case 2:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML2ParserRULE_expr)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(53)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalML2ParserPLUS || _la == EvalML2ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(54)

					var _x = p.expr(5)

					localctx.(*BinOpExprContext).right = _x
				}

			case 3:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML2ParserRULE_expr)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(56)

					var _m = p.Match(EvalML2ParserLT)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(57)

					var _x = p.expr(4)

					localctx.(*BinOpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

// IDefListContext is an interface to support dynamic dispatch.
type IDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDef() []IDefContext
	Def(i int) IDefContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDefListContext differentiates from other interfaces.
	IsDefListContext()
}

type DefListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefListContext() *DefListContext {
	var p = new(DefListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_defList
	return p
}

func InitEmptyDefListContext(p *DefListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_defList
}

func (*DefListContext) IsDefListContext() {}

func NewDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefListContext {
	var p = new(DefListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_defList

	return p
}

func (s *DefListContext) GetParser() antlr.Parser { return s.parser }

func (s *DefListContext) AllDef() []IDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefContext); ok {
			len++
		}
	}

	tst := make([]IDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefContext); ok {
			tst[i] = t.(IDefContext)
			i++
		}
	}

	return tst
}

func (s *DefListContext) Def(i int) IDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
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

	return t.(IDefContext)
}

func (s *DefListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EvalML2ParserCOMMA)
}

func (s *DefListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvalML2ParserCOMMA, i)
}

func (s *DefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterDefList(s)
	}
}

func (s *DefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitDefList(s)
	}
}

func (s *DefListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitDefList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) DefList() (localctx IDefListContext) {
	localctx = NewDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalML2ParserRULE_defList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Def()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalML2ParserCOMMA {
		{
			p.SetState(64)
			p.Match(EvalML2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Def()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IDefContext is an interface to support dynamic dispatch.
type IDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARNAME() antlr.TerminalNode
	Value() IValueContext

	// IsDefContext differentiates from other interfaces.
	IsDefContext()
}

type DefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefContext() *DefContext {
	var p = new(DefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_def
	return p
}

func InitEmptyDefContext(p *DefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_def
}

func (*DefContext) IsDefContext() {}

func NewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefContext {
	var p = new(DefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_def

	return p
}

func (s *DefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML2ParserVARNAME, 0)
}

func (s *DefContext) Value() IValueContext {
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

func (s *DefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterDef(s)
	}
}

func (s *DefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitDef(s)
	}
}

func (s *DefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) Def() (localctx IDefContext) {
	localctx = NewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvalML2ParserRULE_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(EvalML2ParserVARNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(EvalML2ParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
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
	p.RuleIndex = EvalML2ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML2ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML2ParserRULE_value

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
	return s.GetToken(EvalML2ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
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
	return s.GetToken(EvalML2ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML2Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML2Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML2Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvalML2ParserRULE_value)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalML2ParserINT:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(EvalML2ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalML2ParserBOOL:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(EvalML2ParserBOOL)
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

func (p *EvalML2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *EvalML2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
