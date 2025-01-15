// Code generated from EvalML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML3
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

type EvalML3Parser struct {
	*antlr.BaseParser
}

var EvalML3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalml3ParserInit() {
	staticData := &EvalML3ParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'<'", "'if'", "'then'", "'else'", "", "'true'",
		"'false'", "','", "'('", "')'", "'let'", "'in'", "'fun'", "'='", "'rec'",
		"'->'", "'['", "']'", "'|-'", "'evalto'", "'-[0-9]+'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "TIMES", "LT", "IF", "THEN", "ELSE", "BOOL", "TRUE",
		"FALSE", "COMMA", "LPAREN", "RPAREN", "LET", "IN", "FUN", "EQ", "REC",
		"ARROW", "LBRACKET", "RBRACKET", "TURNSTILE", "EVALTO", "NEGINT", "INT",
		"VARNAME", "WS",
	}
	staticData.RuleNames = []string{
		"question", "eval", "expr", "defList", "def", "fun", "recFun", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 3, 1, 23, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		57, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 70, 8, 2, 10, 2, 12, 2, 73, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 78,
		8, 3, 10, 3, 12, 3, 81, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		101, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 110, 8, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 117, 8, 7, 1, 7, 0, 1, 4, 8, 0, 2, 4,
		6, 8, 10, 12, 14, 0, 1, 1, 0, 1, 2, 129, 0, 16, 1, 0, 0, 0, 2, 22, 1, 0,
		0, 0, 4, 56, 1, 0, 0, 0, 6, 74, 1, 0, 0, 0, 8, 82, 1, 0, 0, 0, 10, 86,
		1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 116, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0,
		17, 18, 5, 23, 0, 0, 18, 19, 3, 14, 7, 0, 19, 20, 5, 0, 0, 1, 20, 1, 1,
		0, 0, 0, 21, 23, 3, 6, 3, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23,
		24, 1, 0, 0, 0, 24, 25, 5, 22, 0, 0, 25, 26, 3, 4, 2, 0, 26, 3, 1, 0, 0,
		0, 27, 28, 6, 2, -1, 0, 28, 29, 5, 12, 0, 0, 29, 30, 3, 4, 2, 0, 30, 31,
		5, 13, 0, 0, 31, 57, 1, 0, 0, 0, 32, 57, 3, 10, 5, 0, 33, 34, 5, 5, 0,
		0, 34, 35, 3, 4, 2, 0, 35, 36, 5, 6, 0, 0, 36, 37, 3, 4, 2, 0, 37, 38,
		5, 7, 0, 0, 38, 39, 3, 4, 2, 7, 39, 57, 1, 0, 0, 0, 40, 41, 5, 14, 0, 0,
		41, 42, 5, 26, 0, 0, 42, 43, 5, 17, 0, 0, 43, 44, 3, 4, 2, 0, 44, 45, 5,
		15, 0, 0, 45, 46, 3, 4, 2, 6, 46, 57, 1, 0, 0, 0, 47, 48, 5, 14, 0, 0,
		48, 49, 3, 12, 6, 0, 49, 50, 5, 15, 0, 0, 50, 51, 3, 4, 2, 5, 51, 57, 1,
		0, 0, 0, 52, 57, 5, 8, 0, 0, 53, 57, 5, 25, 0, 0, 54, 57, 5, 24, 0, 0,
		55, 57, 5, 26, 0, 0, 56, 27, 1, 0, 0, 0, 56, 32, 1, 0, 0, 0, 56, 33, 1,
		0, 0, 0, 56, 40, 1, 0, 0, 0, 56, 47, 1, 0, 0, 0, 56, 52, 1, 0, 0, 0, 56,
		53, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 71, 1, 0, 0,
		0, 58, 59, 10, 11, 0, 0, 59, 70, 3, 4, 2, 12, 60, 61, 10, 10, 0, 0, 61,
		62, 5, 3, 0, 0, 62, 70, 3, 4, 2, 11, 63, 64, 10, 9, 0, 0, 64, 65, 7, 0,
		0, 0, 65, 70, 3, 4, 2, 10, 66, 67, 10, 8, 0, 0, 67, 68, 5, 4, 0, 0, 68,
		70, 3, 4, 2, 9, 69, 58, 1, 0, 0, 0, 69, 60, 1, 0, 0, 0, 69, 63, 1, 0, 0,
		0, 69, 66, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72,
		1, 0, 0, 0, 72, 5, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 79, 3, 8, 4, 0,
		75, 76, 5, 11, 0, 0, 76, 78, 3, 8, 4, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 7, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 83, 5, 26, 0, 0, 83, 84, 5, 17, 0, 0, 84, 85, 3, 14,
		7, 0, 85, 9, 1, 0, 0, 0, 86, 87, 5, 16, 0, 0, 87, 88, 5, 26, 0, 0, 88,
		89, 5, 19, 0, 0, 89, 90, 3, 4, 2, 0, 90, 11, 1, 0, 0, 0, 91, 92, 5, 18,
		0, 0, 92, 93, 5, 26, 0, 0, 93, 94, 5, 17, 0, 0, 94, 95, 3, 10, 5, 0, 95,
		13, 1, 0, 0, 0, 96, 117, 5, 25, 0, 0, 97, 117, 5, 8, 0, 0, 98, 100, 5,
		12, 0, 0, 99, 101, 3, 6, 3, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 103, 5, 13, 0, 0, 103, 104, 5, 20, 0, 0,
		104, 105, 3, 10, 5, 0, 105, 106, 5, 21, 0, 0, 106, 117, 1, 0, 0, 0, 107,
		109, 5, 12, 0, 0, 108, 110, 3, 6, 3, 0, 109, 108, 1, 0, 0, 0, 109, 110,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 13, 0, 0, 112, 113, 5, 20,
		0, 0, 113, 114, 3, 12, 6, 0, 114, 115, 5, 21, 0, 0, 115, 117, 1, 0, 0,
		0, 116, 96, 1, 0, 0, 0, 116, 97, 1, 0, 0, 0, 116, 98, 1, 0, 0, 0, 116,
		107, 1, 0, 0, 0, 117, 15, 1, 0, 0, 0, 8, 22, 56, 69, 71, 79, 100, 109,
		116,
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

// EvalML3ParserInit initializes any static state used to implement EvalML3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalML3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalML3ParserInit() {
	staticData := &EvalML3ParserStaticData
	staticData.once.Do(evalml3ParserInit)
}

// NewEvalML3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalML3Parser(input antlr.TokenStream) *EvalML3Parser {
	EvalML3ParserInit()
	this := new(EvalML3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalML3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalML3.g4"

	return this
}

// EvalML3Parser tokens.
const (
	EvalML3ParserEOF       = antlr.TokenEOF
	EvalML3ParserPLUS      = 1
	EvalML3ParserMINUS     = 2
	EvalML3ParserTIMES     = 3
	EvalML3ParserLT        = 4
	EvalML3ParserIF        = 5
	EvalML3ParserTHEN      = 6
	EvalML3ParserELSE      = 7
	EvalML3ParserBOOL      = 8
	EvalML3ParserTRUE      = 9
	EvalML3ParserFALSE     = 10
	EvalML3ParserCOMMA     = 11
	EvalML3ParserLPAREN    = 12
	EvalML3ParserRPAREN    = 13
	EvalML3ParserLET       = 14
	EvalML3ParserIN        = 15
	EvalML3ParserFUN       = 16
	EvalML3ParserEQ        = 17
	EvalML3ParserREC       = 18
	EvalML3ParserARROW     = 19
	EvalML3ParserLBRACKET  = 20
	EvalML3ParserRBRACKET  = 21
	EvalML3ParserTURNSTILE = 22
	EvalML3ParserEVALTO    = 23
	EvalML3ParserNEGINT    = 24
	EvalML3ParserINT       = 25
	EvalML3ParserVARNAME   = 26
	EvalML3ParserWS        = 27
)

// EvalML3Parser rules.
const (
	EvalML3ParserRULE_question = 0
	EvalML3ParserRULE_eval     = 1
	EvalML3ParserRULE_expr     = 2
	EvalML3ParserRULE_defList  = 3
	EvalML3ParserRULE_def      = 4
	EvalML3ParserRULE_fun      = 5
	EvalML3ParserRULE_recFun   = 6
	EvalML3ParserRULE_value    = 7
)

// IQuestionContext is an interface to support dynamic dispatch.
type IQuestionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Eval() IEvalContext
	EVALTO() antlr.TerminalNode
	Value() IValueContext
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
	p.RuleIndex = EvalML3ParserRULE_question
	return p
}

func InitEmptyQuestionContext(p *QuestionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_question
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_question

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

func (s *QuestionContext) EVALTO() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserEVALTO, 0)
}

func (s *QuestionContext) Value() IValueContext {
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

func (s *QuestionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserEOF, 0)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (s *QuestionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitQuestion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalML3ParserRULE_question)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Eval()
	}
	{
		p.SetState(17)
		p.Match(EvalML3ParserEVALTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(18)
		p.Value()
	}
	{
		p.SetState(19)
		p.Match(EvalML3ParserEOF)
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
	p.RuleIndex = EvalML3ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) TURNSTILE() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserTURNSTILE, 0)
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
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalML3ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalML3ParserVARNAME {
		{
			p.SetState(21)
			p.DefList()
		}

	}
	{
		p.SetState(24)
		p.Match(EvalML3ParserTURNSTILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.expr(0)
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
	p.RuleIndex = EvalML3ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_expr

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
	return s.GetToken(EvalML3ParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
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

func (s *IfExprContext) IF() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserIF, 0)
}

func (s *IfExprContext) THEN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserTHEN, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserELSE, 0)
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
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
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

func (s *LetExprContext) LET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLET, 0)
}

func (s *LetExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserEQ, 0)
}

func (s *LetExprContext) IN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserIN, 0)
}

func (s *LetExprContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserVARNAME, 0)
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
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterLetExpr(s)
	}
}

func (s *LetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitLetExpr(s)
	}
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitLetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetRecExprContext struct {
	ExprContext
}

func NewLetRecExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetRecExprContext {
	var p = new(LetRecExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LetRecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetRecExprContext) LET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLET, 0)
}

func (s *LetRecExprContext) RecFun() IRecFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecFunContext)
}

func (s *LetRecExprContext) IN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserIN, 0)
}

func (s *LetRecExprContext) Expr() IExprContext {
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

func (s *LetRecExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterLetRecExpr(s)
	}
}

func (s *LetRecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitLetRecExpr(s)
	}
}

func (s *LetRecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitLetRecExpr(s)

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
	return s.GetToken(EvalML3ParserVARNAME, 0)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitVarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppExprContext struct {
	ExprContext
}

func NewAppExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppExprContext {
	var p = new(AppExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AppExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppExprContext) AllExpr() []IExprContext {
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

func (s *AppExprContext) Expr(i int) IExprContext {
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

func (s *AppExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterAppExpr(s)
	}
}

func (s *AppExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitAppExpr(s)
	}
}

func (s *AppExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitAppExpr(s)

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
	return s.GetToken(EvalML3ParserTIMES, 0)
}

func (s *BinOpExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserPLUS, 0)
}

func (s *BinOpExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserMINUS, 0)
}

func (s *BinOpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLT, 0)
}

func (s *BinOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterBinOpExpr(s)
	}
}

func (s *BinOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitBinOpExpr(s)
	}
}

func (s *BinOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
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
	return s.GetToken(EvalML3ParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
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

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLPAREN, 0)
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

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
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

func (s *NegIntExprContext) NEGINT() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserNEGINT, 0)
}

func (s *NegIntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterNegIntExpr(s)
	}
}

func (s *NegIntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitNegIntExpr(s)
	}
}

func (s *NegIntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitNegIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunExprContext struct {
	ExprContext
}

func NewFunExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunExprContext {
	var p = new(FunExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunExprContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *FunExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterFunExpr(s)
	}
}

func (s *FunExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitFunExpr(s)
	}
}

func (s *FunExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitFunExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *EvalML3Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, EvalML3ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(28)
			p.Match(EvalML3ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.expr(0)
		}
		{
			p.SetState(30)
			p.Match(EvalML3ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Fun()
		}

	case 3:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(EvalML3ParserIF)
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
			p.Match(EvalML3ParserTHEN)
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
			p.Match(EvalML3ParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)

			var _x = p.expr(7)

			localctx.(*IfExprContext).else_ = _x
		}

	case 4:
		localctx = NewLetExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(EvalML3ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)

			var _m = p.Match(EvalML3ParserVARNAME)

			localctx.(*LetExprContext).bindName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(EvalML3ParserEQ)
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
			p.Match(EvalML3ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)

			var _x = p.expr(6)

			localctx.(*LetExprContext).body = _x
		}

	case 5:
		localctx = NewLetRecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(EvalML3ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.RecFun()
		}
		{
			p.SetState(49)
			p.Match(EvalML3ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.expr(5)
		}

	case 6:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(EvalML3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(EvalML3ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewNegIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(EvalML3ParserNEGINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.Match(EvalML3ParserVARNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(71)
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
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, EvalML3ParserRULE_expr)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(59)
					p.expr(12)
				}

			case 2:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML3ParserRULE_expr)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(61)

					var _m = p.Match(EvalML3ParserTIMES)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(62)

					var _x = p.expr(11)

					localctx.(*BinOpExprContext).right = _x
				}

			case 3:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML3ParserRULE_expr)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(64)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalML3ParserPLUS || _la == EvalML3ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(65)

					var _x = p.expr(10)

					localctx.(*BinOpExprContext).right = _x
				}

			case 4:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML3ParserRULE_expr)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(67)

					var _m = p.Match(EvalML3ParserLT)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(68)

					var _x = p.expr(9)

					localctx.(*BinOpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(73)
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
	p.RuleIndex = EvalML3ParserRULE_defList
	return p
}

func InitEmptyDefListContext(p *DefListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_defList
}

func (*DefListContext) IsDefListContext() {}

func NewDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefListContext {
	var p = new(DefListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_defList

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
	return s.GetTokens(EvalML3ParserCOMMA)
}

func (s *DefListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvalML3ParserCOMMA, i)
}

func (s *DefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterDefList(s)
	}
}

func (s *DefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitDefList(s)
	}
}

func (s *DefListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitDefList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) DefList() (localctx IDefListContext) {
	localctx = NewDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalML3ParserRULE_defList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Def()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalML3ParserCOMMA {
		{
			p.SetState(75)
			p.Match(EvalML3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Def()
		}

		p.SetState(81)
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
	EQ() antlr.TerminalNode
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
	p.RuleIndex = EvalML3ParserRULE_def
	return p
}

func InitEmptyDefContext(p *DefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_def
}

func (*DefContext) IsDefContext() {}

func NewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefContext {
	var p = new(DefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_def

	return p
}

func (s *DefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserVARNAME, 0)
}

func (s *DefContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserEQ, 0)
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
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterDef(s)
	}
}

func (s *DefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitDef(s)
	}
}

func (s *DefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Def() (localctx IDefContext) {
	localctx = NewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvalML3ParserRULE_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(EvalML3ParserVARNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(EvalML3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
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

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParam returns the param token.
	GetParam() antlr.Token

	// SetParam sets the param token.
	SetParam(antlr.Token)

	// GetBody returns the body rule contexts.
	GetBody() IExprContext

	// SetBody sets the body rule contexts.
	SetBody(IExprContext)

	// Getter signatures
	FUN() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	VARNAME() antlr.TerminalNode
	Expr() IExprContext

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	param  antlr.Token
	body   IExprContext
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() antlr.Token { return s.param }

func (s *FunContext) SetParam(v antlr.Token) { s.param = v }

func (s *FunContext) GetBody() IExprContext { return s.body }

func (s *FunContext) SetBody(v IExprContext) { s.body = v }

func (s *FunContext) FUN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserFUN, 0)
}

func (s *FunContext) ARROW() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserARROW, 0)
}

func (s *FunContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserVARNAME, 0)
}

func (s *FunContext) Expr() IExprContext {
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

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvalML3ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(EvalML3ParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)

		var _m = p.Match(EvalML3ParserVARNAME)

		localctx.(*FunContext).param = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(EvalML3ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)

		var _x = p.expr(0)

		localctx.(*FunContext).body = _x
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

// IRecFunContext is an interface to support dynamic dispatch.
type IRecFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunName returns the funName token.
	GetFunName() antlr.Token

	// SetFunName sets the funName token.
	SetFunName(antlr.Token)

	// Getter signatures
	REC() antlr.TerminalNode
	EQ() antlr.TerminalNode
	Fun() IFunContext
	VARNAME() antlr.TerminalNode

	// IsRecFunContext differentiates from other interfaces.
	IsRecFunContext()
}

type RecFunContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	funName antlr.Token
}

func NewEmptyRecFunContext() *RecFunContext {
	var p = new(RecFunContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_recFun
	return p
}

func InitEmptyRecFunContext(p *RecFunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_recFun
}

func (*RecFunContext) IsRecFunContext() {}

func NewRecFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecFunContext {
	var p = new(RecFunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_recFun

	return p
}

func (s *RecFunContext) GetParser() antlr.Parser { return s.parser }

func (s *RecFunContext) GetFunName() antlr.Token { return s.funName }

func (s *RecFunContext) SetFunName(v antlr.Token) { s.funName = v }

func (s *RecFunContext) REC() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserREC, 0)
}

func (s *RecFunContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserEQ, 0)
}

func (s *RecFunContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *RecFunContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserVARNAME, 0)
}

func (s *RecFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterRecFun(s)
	}
}

func (s *RecFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitRecFun(s)
	}
}

func (s *RecFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitRecFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) RecFun() (localctx IRecFunContext) {
	localctx = NewRecFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EvalML3ParserRULE_recFun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(EvalML3ParserREC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)

		var _m = p.Match(EvalML3ParserVARNAME)

		localctx.(*RecFunContext).funName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(EvalML3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Fun()
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
	p.RuleIndex = EvalML3ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML3ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML3ParserRULE_value

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
	return s.GetToken(EvalML3ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitBoolValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunValueContext struct {
	ValueContext
}

func NewFunValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunValueContext {
	var p = new(FunValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *FunValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLPAREN, 0)
}

func (s *FunValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserRPAREN, 0)
}

func (s *FunValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLBRACKET, 0)
}

func (s *FunValueContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *FunValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserRBRACKET, 0)
}

func (s *FunValueContext) DefList() IDefListContext {
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

func (s *FunValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterFunValue(s)
	}
}

func (s *FunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitFunValue(s)
	}
}

func (s *FunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitFunValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type RecFunValueContext struct {
	ValueContext
}

func NewRecFunValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecFunValueContext {
	var p = new(RecFunValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *RecFunValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLPAREN, 0)
}

func (s *RecFunValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserRPAREN, 0)
}

func (s *RecFunValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserLBRACKET, 0)
}

func (s *RecFunValueContext) RecFun() IRecFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecFunContext)
}

func (s *RecFunValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML3ParserRBRACKET, 0)
}

func (s *RecFunValueContext) DefList() IDefListContext {
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

func (s *RecFunValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterRecFunValue(s)
	}
}

func (s *RecFunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitRecFunValue(s)
	}
}

func (s *RecFunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitRecFunValue(s)

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
	return s.GetToken(EvalML3ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML3Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML3Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML3Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvalML3ParserRULE_value)
	var _la int

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(EvalML3ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(EvalML3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(EvalML3ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalML3ParserVARNAME {
			{
				p.SetState(99)
				p.DefList()
			}

		}
		{
			p.SetState(102)
			p.Match(EvalML3ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(EvalML3ParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Fun()
		}
		{
			p.SetState(105)
			p.Match(EvalML3ParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewRecFunValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(107)
			p.Match(EvalML3ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalML3ParserVARNAME {
			{
				p.SetState(108)
				p.DefList()
			}

		}
		{
			p.SetState(111)
			p.Match(EvalML3ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(EvalML3ParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.RecFun()
		}
		{
			p.SetState(114)
			p.Match(EvalML3ParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

func (p *EvalML3Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *EvalML3Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
