// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML1
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

type EvalContML1Parser struct {
	*antlr.BaseParser
}

var EvalContML1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalcontml1ParserInit() {
	staticData := &EvalContML1ParserStaticData
	staticData.LiteralNames = []string{
		"", "'evalto'", "'if'", "'then'", "'else'", "'('", "')'", "'>>'", "'_'",
		"'{'", "'}'", "'+'", "'-'", "'*'", "'<'", "", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES",
		"LT", "BOOL", "TRUE", "FALSE", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"eval", "value", "exp", "cont",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 85, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 3, 0, 11, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 19, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 46, 8, 2, 10, 2, 12, 2, 49, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 81, 8, 3, 3, 3, 83, 8, 3, 1, 3, 0, 1, 4, 4, 0,
		2, 4, 6, 0, 2, 1, 0, 11, 12, 1, 0, 11, 14, 94, 0, 8, 1, 0, 0, 0, 2, 18,
		1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 82, 1, 0, 0, 0, 8, 10, 3, 4, 2, 0, 9,
		11, 3, 6, 3, 0, 10, 9, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 12, 1, 0, 0,
		0, 12, 13, 5, 1, 0, 0, 13, 14, 3, 2, 1, 0, 14, 15, 5, 0, 0, 1, 15, 1, 1,
		0, 0, 0, 16, 19, 5, 18, 0, 0, 17, 19, 5, 15, 0, 0, 18, 16, 1, 0, 0, 0,
		18, 17, 1, 0, 0, 0, 19, 3, 1, 0, 0, 0, 20, 21, 6, 2, -1, 0, 21, 35, 5,
		18, 0, 0, 22, 35, 5, 15, 0, 0, 23, 24, 5, 2, 0, 0, 24, 25, 3, 4, 2, 0,
		25, 26, 5, 3, 0, 0, 26, 27, 3, 4, 2, 0, 27, 28, 5, 4, 0, 0, 28, 29, 3,
		4, 2, 2, 29, 35, 1, 0, 0, 0, 30, 31, 5, 5, 0, 0, 31, 32, 3, 4, 2, 0, 32,
		33, 5, 6, 0, 0, 33, 35, 1, 0, 0, 0, 34, 20, 1, 0, 0, 0, 34, 22, 1, 0, 0,
		0, 34, 23, 1, 0, 0, 0, 34, 30, 1, 0, 0, 0, 35, 47, 1, 0, 0, 0, 36, 37,
		10, 5, 0, 0, 37, 38, 5, 13, 0, 0, 38, 46, 3, 4, 2, 6, 39, 40, 10, 4, 0,
		0, 40, 41, 7, 0, 0, 0, 41, 46, 3, 4, 2, 5, 42, 43, 10, 3, 0, 0, 43, 44,
		5, 14, 0, 0, 44, 46, 3, 4, 2, 4, 45, 36, 1, 0, 0, 0, 45, 39, 1, 0, 0, 0,
		45, 42, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1,
		0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 51, 5, 7, 0, 0, 51,
		83, 5, 8, 0, 0, 52, 53, 5, 7, 0, 0, 53, 54, 5, 9, 0, 0, 54, 55, 5, 8, 0,
		0, 55, 56, 7, 1, 0, 0, 56, 57, 3, 4, 2, 0, 57, 59, 5, 10, 0, 0, 58, 60,
		3, 6, 3, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 83, 1, 0, 0, 0,
		61, 62, 5, 7, 0, 0, 62, 63, 5, 9, 0, 0, 63, 64, 3, 2, 1, 0, 64, 65, 7,
		1, 0, 0, 65, 66, 5, 8, 0, 0, 66, 68, 5, 10, 0, 0, 67, 69, 3, 6, 3, 0, 68,
		67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 83, 1, 0, 0, 0, 70, 71, 5, 7, 0,
		0, 71, 72, 5, 9, 0, 0, 72, 73, 5, 2, 0, 0, 73, 74, 5, 8, 0, 0, 74, 75,
		5, 3, 0, 0, 75, 76, 3, 4, 2, 0, 76, 77, 5, 4, 0, 0, 77, 78, 3, 4, 2, 0,
		78, 80, 5, 10, 0, 0, 79, 81, 3, 6, 3, 0, 80, 79, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 50, 1, 0, 0, 0, 82, 52, 1, 0, 0, 0, 82,
		61, 1, 0, 0, 0, 82, 70, 1, 0, 0, 0, 83, 7, 1, 0, 0, 0, 9, 10, 18, 34, 45,
		47, 59, 68, 80, 82,
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

// EvalContML1ParserInit initializes any static state used to implement EvalContML1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalContML1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalContML1ParserInit() {
	staticData := &EvalContML1ParserStaticData
	staticData.once.Do(evalcontml1ParserInit)
}

// NewEvalContML1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalContML1Parser(input antlr.TokenStream) *EvalContML1Parser {
	EvalContML1ParserInit()
	this := new(EvalContML1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalContML1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalContML1.g4"

	return this
}

// EvalContML1Parser tokens.
const (
	EvalContML1ParserEOF   = antlr.TokenEOF
	EvalContML1ParserT__0  = 1
	EvalContML1ParserT__1  = 2
	EvalContML1ParserT__2  = 3
	EvalContML1ParserT__3  = 4
	EvalContML1ParserT__4  = 5
	EvalContML1ParserT__5  = 6
	EvalContML1ParserT__6  = 7
	EvalContML1ParserT__7  = 8
	EvalContML1ParserT__8  = 9
	EvalContML1ParserT__9  = 10
	EvalContML1ParserPLUS  = 11
	EvalContML1ParserMINUS = 12
	EvalContML1ParserTIMES = 13
	EvalContML1ParserLT    = 14
	EvalContML1ParserBOOL  = 15
	EvalContML1ParserTRUE  = 16
	EvalContML1ParserFALSE = 17
	EvalContML1ParserINT   = 18
	EvalContML1ParserWS    = 19
)

// EvalContML1Parser rules.
const (
	EvalContML1ParserRULE_eval  = 0
	EvalContML1ParserRULE_value = 1
	EvalContML1ParserRULE_exp   = 2
	EvalContML1ParserRULE_cont  = 3
)

// IEvalContext is an interface to support dynamic dispatch.
type IEvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	Value() IValueContext
	EOF() antlr.TerminalNode
	Cont() IContContext

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
	p.RuleIndex = EvalContML1ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML1ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
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

func (s *EvalContext) EOF() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserEOF, 0)
}

func (s *EvalContext) Cont() IContContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContContext)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML1Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalContML1ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.exp(0)
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalContML1ParserT__6 {
		{
			p.SetState(9)
			p.Cont()
		}

	}
	{
		p.SetState(12)
		p.Match(EvalContML1ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(13)
		p.Value()
	}
	{
		p.SetState(14)
		p.Match(EvalContML1ParserEOF)
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
	p.RuleIndex = EvalContML1ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML1ParserRULE_value

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
	return s.GetToken(EvalContML1ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
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
	return s.GetToken(EvalContML1ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML1Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalContML1ParserRULE_value)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalContML1ParserINT:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(EvalContML1ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalContML1ParserBOOL:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Match(EvalContML1ParserBOOL)
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

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML1ParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyAll(ctx *ExpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExpContext struct {
	ExpContext
}

func NewBoolExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpContext {
	var p = new(BoolExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *BoolExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserBOOL, 0)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (s *BoolExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitBoolExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfExpContext struct {
	ExpContext
	cond  IExpContext
	then  IExpContext
	else_ IExpContext
}

func NewIfExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExpContext {
	var p = new(IfExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *IfExpContext) GetCond() IExpContext { return s.cond }

func (s *IfExpContext) GetThen() IExpContext { return s.then }

func (s *IfExpContext) GetElse_() IExpContext { return s.else_ }

func (s *IfExpContext) SetCond(v IExpContext) { s.cond = v }

func (s *IfExpContext) SetThen(v IExpContext) { s.then = v }

func (s *IfExpContext) SetElse_(v IExpContext) { s.else_ = v }

func (s *IfExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *IfExpContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
}

func (s *IfExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterIfExp(s)
	}
}

func (s *IfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitIfExp(s)
	}
}

func (s *IfExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitIfExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	ExpContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ParenExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExpContext struct {
	ExpContext
}

func NewIntExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExpContext {
	var p = new(IntExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *IntExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExpContext) INT() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserINT, 0)
}

func (s *IntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterIntExp(s)
	}
}

func (s *IntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitIntExp(s)
	}
}

func (s *IntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitIntExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinOpExpContext struct {
	ExpContext
	left  IExpContext
	op    antlr.Token
	right IExpContext
}

func NewBinOpExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinOpExpContext {
	var p = new(BinOpExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *BinOpExpContext) GetOp() antlr.Token { return s.op }

func (s *BinOpExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinOpExpContext) GetLeft() IExpContext { return s.left }

func (s *BinOpExpContext) GetRight() IExpContext { return s.right }

func (s *BinOpExpContext) SetLeft(v IExpContext) { s.left = v }

func (s *BinOpExpContext) SetRight(v IExpContext) { s.right = v }

func (s *BinOpExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOpExpContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *BinOpExpContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
}

func (s *BinOpExpContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserTIMES, 0)
}

func (s *BinOpExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserPLUS, 0)
}

func (s *BinOpExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserMINUS, 0)
}

func (s *BinOpExpContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserLT, 0)
}

func (s *BinOpExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterBinOpExp(s)
	}
}

func (s *BinOpExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitBinOpExp(s)
	}
}

func (s *BinOpExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitBinOpExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML1Parser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *EvalContML1Parser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, EvalContML1ParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case EvalContML1ParserINT:
		localctx = NewIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(21)
			p.Match(EvalContML1ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalContML1ParserBOOL:
		localctx = NewBoolExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(EvalContML1ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case EvalContML1ParserT__1:
		localctx = NewIfExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(EvalContML1ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)

			var _x = p.exp(0)

			localctx.(*IfExpContext).cond = _x
		}
		{
			p.SetState(25)
			p.Match(EvalContML1ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)

			var _x = p.exp(0)

			localctx.(*IfExpContext).then = _x
		}
		{
			p.SetState(27)
			p.Match(EvalContML1ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)

			var _x = p.exp(2)

			localctx.(*IfExpContext).else_ = _x
		}

	case EvalContML1ParserT__4:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(EvalContML1ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.exp(0)
		}
		{
			p.SetState(32)
			p.Match(EvalContML1ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML1ParserRULE_exp)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(37)

					var _m = p.Match(EvalContML1ParserTIMES)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(38)

					var _x = p.exp(6)

					localctx.(*BinOpExpContext).right = _x
				}

			case 2:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML1ParserRULE_exp)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(40)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalContML1ParserPLUS || _la == EvalContML1ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)

					var _x = p.exp(5)

					localctx.(*BinOpExpContext).right = _x
				}

			case 3:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML1ParserRULE_exp)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(43)

					var _m = p.Match(EvalContML1ParserLT)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(44)

					var _x = p.exp(4)

					localctx.(*BinOpExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

// IContContext is an interface to support dynamic dispatch.
type IContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContContext differentiates from other interfaces.
	IsContContext()
}

type ContContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContContext() *ContContext {
	var p = new(ContContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_cont
	return p
}

func InitEmptyContContext(p *ContContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML1ParserRULE_cont
}

func (*ContContext) IsContContext() {}

func NewContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContContext {
	var p = new(ContContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML1ParserRULE_cont

	return p
}

func (s *ContContext) GetParser() antlr.Parser { return s.parser }

func (s *ContContext) CopyAll(ctx *ContContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfContContext struct {
	ContContext
	then  IExpContext
	else_ IExpContext
}

func NewIfContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContContext {
	var p = new(IfContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *IfContContext) GetThen() IExpContext { return s.then }

func (s *IfContContext) GetElse_() IExpContext { return s.else_ }

func (s *IfContContext) SetThen(v IExpContext) { s.then = v }

func (s *IfContContext) SetElse_(v IExpContext) { s.else_ = v }

func (s *IfContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *IfContContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
}

func (s *IfContContext) Cont() IContContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContContext)
}

func (s *IfContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterIfCont(s)
	}
}

func (s *IfContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitIfCont(s)
	}
}

func (s *IfContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitIfCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpContContext struct {
	ContContext
	op antlr.Token
}

func NewExpContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpContContext {
	var p = new(ExpContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *ExpContContext) GetOp() antlr.Token { return s.op }

func (s *ExpContContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserPLUS, 0)
}

func (s *ExpContContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserMINUS, 0)
}

func (s *ExpContContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserTIMES, 0)
}

func (s *ExpContContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserLT, 0)
}

func (s *ExpContContext) Cont() IContContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContContext)
}

func (s *ExpContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterExpCont(s)
	}
}

func (s *ExpContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitExpCont(s)
	}
}

func (s *ExpContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitExpCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContContext struct {
	ContContext
	op antlr.Token
}

func NewValueContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContContext {
	var p = new(ValueContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *ValueContContext) GetOp() antlr.Token { return s.op }

func (s *ValueContContext) SetOp(v antlr.Token) { s.op = v }

func (s *ValueContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContContext) Value() IValueContext {
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

func (s *ValueContContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserPLUS, 0)
}

func (s *ValueContContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserMINUS, 0)
}

func (s *ValueContContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserTIMES, 0)
}

func (s *ValueContContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML1ParserLT, 0)
}

func (s *ValueContContext) Cont() IContContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContContext)
}

func (s *ValueContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterValueCont(s)
	}
}

func (s *ValueContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitValueCont(s)
	}
}

func (s *ValueContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitValueCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContContext struct {
	ContContext
}

func NewUnaryContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContContext {
	var p = new(UnaryContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *UnaryContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.EnterUnaryCont(s)
	}
}

func (s *UnaryContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML1Listener); ok {
		listenerT.ExitUnaryCont(s)
	}
}

func (s *UnaryContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML1Visitor:
		return t.VisitUnaryCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML1Parser) Cont() (localctx IContContext) {
	localctx = NewContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalContML1ParserRULE_cont)
	var _la int

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryContContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(EvalContML1ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(EvalContML1ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpContContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(EvalContML1ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(EvalContML1ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(EvalContML1ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpContContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpContContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(56)
			p.exp(0)
		}
		{
			p.SetState(57)
			p.Match(EvalContML1ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML1ParserT__6 {
			{
				p.SetState(58)
				p.Cont()
			}

		}

	case 3:
		localctx = NewValueContContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Match(EvalContML1ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(EvalContML1ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Value()
		}
		{
			p.SetState(64)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueContContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueContContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(65)
			p.Match(EvalContML1ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(EvalContML1ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML1ParserT__6 {
			{
				p.SetState(67)
				p.Cont()
			}

		}

	case 4:
		localctx = NewIfContContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Match(EvalContML1ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(EvalContML1ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(EvalContML1ParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(EvalContML1ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(EvalContML1ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)

			var _x = p.exp(0)

			localctx.(*IfContContext).then = _x
		}
		{
			p.SetState(76)
			p.Match(EvalContML1ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)

			var _x = p.exp(0)

			localctx.(*IfContContext).else_ = _x
		}
		{
			p.SetState(78)
			p.Match(EvalContML1ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML1ParserT__6 {
			{
				p.SetState(79)
				p.Cont()
			}

		}

	case antlr.ATNInvalidAltNumber:
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

func (p *EvalContML1Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvalContML1Parser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
