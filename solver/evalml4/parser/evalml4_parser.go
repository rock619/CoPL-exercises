// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML4
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

type EvalML4Parser struct {
	*antlr.BaseParser
}

var EvalML4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalml4ParserInit() {
	staticData := &EvalML4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'|'", "", "'::'", "'match'", "'with'", "'+'", "'-'", "'*'", "'<'",
		"'if'", "'then'", "'else'", "", "'true'", "'false'", "','", "'('", "')'",
		"'let'", "'in'", "'fun'", "'='", "'rec'", "'->'", "'['", "']'", "'|-'",
		"'evalto'",
	}
	staticData.SymbolicNames = []string{
		"", "OR", "EMPTYLIST", "CONS", "MATCH", "WITH", "PLUS", "MINUS", "TIMES",
		"LT", "IF", "THEN", "ELSE", "BOOL", "TRUE", "FALSE", "COMMA", "LPAREN",
		"RPAREN", "LET", "IN", "FUN", "EQ", "REC", "ARROW", "LBRACKET", "RBRACKET",
		"TURNSTILE", "EVALTO", "INT", "VARNAME", "WS",
	}
	staticData.RuleNames = []string{
		"question", "eval", "expr", "defList", "def", "emptyPattern", "consPattern",
		"fun", "recFun", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 153, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 27, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 68, 8,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 84, 8, 2, 10, 2, 12, 2, 87, 9, 2, 1, 3, 1, 3, 1, 3,
		5, 3, 92, 8, 3, 10, 3, 12, 3, 95, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 126, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 135, 8, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 143, 8, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 148, 8, 9, 10, 9, 12, 9, 151, 9, 9, 1, 9, 0, 2, 4, 18, 10, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 0, 1, 1, 0, 6, 7, 165, 0, 20, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 88, 1, 0, 0, 0, 8, 96, 1, 0, 0,
		0, 10, 100, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0, 14, 110, 1, 0, 0, 0, 16, 115,
		1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 28, 0,
		0, 22, 23, 3, 18, 9, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 27,
		3, 6, 3, 0, 26, 25, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0,
		28, 29, 5, 27, 0, 0, 29, 30, 3, 4, 2, 0, 30, 3, 1, 0, 0, 0, 31, 32, 6,
		2, -1, 0, 32, 33, 5, 17, 0, 0, 33, 34, 3, 4, 2, 0, 34, 35, 5, 18, 0, 0,
		35, 68, 1, 0, 0, 0, 36, 68, 3, 14, 7, 0, 37, 38, 5, 4, 0, 0, 38, 39, 3,
		4, 2, 0, 39, 40, 5, 5, 0, 0, 40, 41, 3, 10, 5, 0, 41, 42, 5, 1, 0, 0, 42,
		43, 3, 12, 6, 0, 43, 68, 1, 0, 0, 0, 44, 68, 5, 2, 0, 0, 45, 46, 5, 10,
		0, 0, 46, 47, 3, 4, 2, 0, 47, 48, 5, 11, 0, 0, 48, 49, 3, 4, 2, 0, 49,
		50, 5, 12, 0, 0, 50, 51, 3, 4, 2, 6, 51, 68, 1, 0, 0, 0, 52, 53, 5, 19,
		0, 0, 53, 54, 5, 30, 0, 0, 54, 55, 5, 22, 0, 0, 55, 56, 3, 4, 2, 0, 56,
		57, 5, 20, 0, 0, 57, 58, 3, 4, 2, 5, 58, 68, 1, 0, 0, 0, 59, 60, 5, 19,
		0, 0, 60, 61, 3, 16, 8, 0, 61, 62, 5, 20, 0, 0, 62, 63, 3, 4, 2, 4, 63,
		68, 1, 0, 0, 0, 64, 68, 5, 29, 0, 0, 65, 68, 5, 13, 0, 0, 66, 68, 5, 30,
		0, 0, 67, 31, 1, 0, 0, 0, 67, 36, 1, 0, 0, 0, 67, 37, 1, 0, 0, 0, 67, 44,
		1, 0, 0, 0, 67, 45, 1, 0, 0, 0, 67, 52, 1, 0, 0, 0, 67, 59, 1, 0, 0, 0,
		67, 64, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 85, 1,
		0, 0, 0, 69, 70, 10, 13, 0, 0, 70, 84, 3, 4, 2, 14, 71, 72, 10, 10, 0,
		0, 72, 73, 5, 3, 0, 0, 73, 84, 3, 4, 2, 10, 74, 75, 10, 9, 0, 0, 75, 76,
		5, 8, 0, 0, 76, 84, 3, 4, 2, 10, 77, 78, 10, 8, 0, 0, 78, 79, 7, 0, 0,
		0, 79, 84, 3, 4, 2, 9, 80, 81, 10, 7, 0, 0, 81, 82, 5, 9, 0, 0, 82, 84,
		3, 4, 2, 8, 83, 69, 1, 0, 0, 0, 83, 71, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0,
		83, 77, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 5, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88,
		93, 3, 8, 4, 0, 89, 90, 5, 16, 0, 0, 90, 92, 3, 8, 4, 0, 91, 89, 1, 0,
		0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 7,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 30, 0, 0, 97, 98, 5, 22, 0,
		0, 98, 99, 3, 18, 9, 0, 99, 9, 1, 0, 0, 0, 100, 101, 5, 2, 0, 0, 101, 102,
		5, 24, 0, 0, 102, 103, 3, 4, 2, 0, 103, 11, 1, 0, 0, 0, 104, 105, 5, 30,
		0, 0, 105, 106, 5, 3, 0, 0, 106, 107, 5, 30, 0, 0, 107, 108, 5, 24, 0,
		0, 108, 109, 3, 4, 2, 0, 109, 13, 1, 0, 0, 0, 110, 111, 5, 21, 0, 0, 111,
		112, 5, 30, 0, 0, 112, 113, 5, 24, 0, 0, 113, 114, 3, 4, 2, 0, 114, 15,
		1, 0, 0, 0, 115, 116, 5, 23, 0, 0, 116, 117, 5, 30, 0, 0, 117, 118, 5,
		22, 0, 0, 118, 119, 3, 14, 7, 0, 119, 17, 1, 0, 0, 0, 120, 121, 6, 9, -1,
		0, 121, 143, 5, 29, 0, 0, 122, 143, 5, 13, 0, 0, 123, 125, 5, 17, 0, 0,
		124, 126, 3, 6, 3, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		127, 1, 0, 0, 0, 127, 128, 5, 18, 0, 0, 128, 129, 5, 25, 0, 0, 129, 130,
		3, 14, 7, 0, 130, 131, 5, 26, 0, 0, 131, 143, 1, 0, 0, 0, 132, 134, 5,
		17, 0, 0, 133, 135, 3, 6, 3, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0,
		0, 135, 136, 1, 0, 0, 0, 136, 137, 5, 18, 0, 0, 137, 138, 5, 25, 0, 0,
		138, 139, 3, 16, 8, 0, 139, 140, 5, 26, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		143, 5, 2, 0, 0, 142, 120, 1, 0, 0, 0, 142, 122, 1, 0, 0, 0, 142, 123,
		1, 0, 0, 0, 142, 132, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 149, 1, 0,
		0, 0, 144, 145, 10, 1, 0, 0, 145, 146, 5, 3, 0, 0, 146, 148, 3, 18, 9,
		1, 147, 144, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 19, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 9, 26, 67,
		83, 85, 93, 125, 134, 142, 149,
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

// EvalML4ParserInit initializes any static state used to implement EvalML4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalML4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalML4ParserInit() {
	staticData := &EvalML4ParserStaticData
	staticData.once.Do(evalml4ParserInit)
}

// NewEvalML4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalML4Parser(input antlr.TokenStream) *EvalML4Parser {
	EvalML4ParserInit()
	this := new(EvalML4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalML4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalML4.g4"

	return this
}

// EvalML4Parser tokens.
const (
	EvalML4ParserEOF       = antlr.TokenEOF
	EvalML4ParserOR        = 1
	EvalML4ParserEMPTYLIST = 2
	EvalML4ParserCONS      = 3
	EvalML4ParserMATCH     = 4
	EvalML4ParserWITH      = 5
	EvalML4ParserPLUS      = 6
	EvalML4ParserMINUS     = 7
	EvalML4ParserTIMES     = 8
	EvalML4ParserLT        = 9
	EvalML4ParserIF        = 10
	EvalML4ParserTHEN      = 11
	EvalML4ParserELSE      = 12
	EvalML4ParserBOOL      = 13
	EvalML4ParserTRUE      = 14
	EvalML4ParserFALSE     = 15
	EvalML4ParserCOMMA     = 16
	EvalML4ParserLPAREN    = 17
	EvalML4ParserRPAREN    = 18
	EvalML4ParserLET       = 19
	EvalML4ParserIN        = 20
	EvalML4ParserFUN       = 21
	EvalML4ParserEQ        = 22
	EvalML4ParserREC       = 23
	EvalML4ParserARROW     = 24
	EvalML4ParserLBRACKET  = 25
	EvalML4ParserRBRACKET  = 26
	EvalML4ParserTURNSTILE = 27
	EvalML4ParserEVALTO    = 28
	EvalML4ParserINT       = 29
	EvalML4ParserVARNAME   = 30
	EvalML4ParserWS        = 31
)

// EvalML4Parser rules.
const (
	EvalML4ParserRULE_question     = 0
	EvalML4ParserRULE_eval         = 1
	EvalML4ParserRULE_expr         = 2
	EvalML4ParserRULE_defList      = 3
	EvalML4ParserRULE_def          = 4
	EvalML4ParserRULE_emptyPattern = 5
	EvalML4ParserRULE_consPattern  = 6
	EvalML4ParserRULE_fun          = 7
	EvalML4ParserRULE_recFun       = 8
	EvalML4ParserRULE_value        = 9
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
	p.RuleIndex = EvalML4ParserRULE_question
	return p
}

func InitEmptyQuestionContext(p *QuestionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_question
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_question

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
	return s.GetToken(EvalML4ParserEVALTO, 0)
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
	return s.GetToken(EvalML4ParserEOF, 0)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (s *QuestionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitQuestion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalML4ParserRULE_question)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Eval()
	}
	{
		p.SetState(21)
		p.Match(EvalML4ParserEVALTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.value(0)
	}
	{
		p.SetState(23)
		p.Match(EvalML4ParserEOF)
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
	p.RuleIndex = EvalML4ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) TURNSTILE() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserTURNSTILE, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalML4ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalML4ParserVARNAME {
		{
			p.SetState(25)
			p.DefList()
		}

	}
	{
		p.SetState(28)
		p.Match(EvalML4ParserTURNSTILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
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
	p.RuleIndex = EvalML4ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_expr

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
	return s.GetToken(EvalML4ParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserIF, 0)
}

func (s *IfExprContext) THEN() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserTHEN, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserELSE, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserLET, 0)
}

func (s *LetExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEQ, 0)
}

func (s *LetExprContext) IN() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserIN, 0)
}

func (s *LetExprContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserVARNAME, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterLetExpr(s)
	}
}

func (s *LetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitLetExpr(s)
	}
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserLET, 0)
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
	return s.GetToken(EvalML4ParserIN, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterLetRecExpr(s)
	}
}

func (s *LetRecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitLetRecExpr(s)
	}
}

func (s *LetRecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitLetRecExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppExprContext struct {
	ExprContext
	fn  IExprContext
	arg IExprContext
}

func NewAppExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppExprContext {
	var p = new(AppExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AppExprContext) GetFn() IExprContext { return s.fn }

func (s *AppExprContext) GetArg() IExprContext { return s.arg }

func (s *AppExprContext) SetFn(v IExprContext) { s.fn = v }

func (s *AppExprContext) SetArg(v IExprContext) { s.arg = v }

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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterAppExpr(s)
	}
}

func (s *AppExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitAppExpr(s)
	}
}

func (s *AppExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitAppExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyListExprContext struct {
	ExprContext
}

func NewEmptyListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyListExprContext {
	var p = new(EmptyListExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EmptyListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyListExprContext) EMPTYLIST() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEMPTYLIST, 0)
}

func (s *EmptyListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterEmptyListExpr(s)
	}
}

func (s *EmptyListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitEmptyListExpr(s)
	}
}

func (s *EmptyListExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitEmptyListExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConsExprContext struct {
	ExprContext
	head IExprContext
	tail IExprContext
}

func NewConsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsExprContext {
	var p = new(ConsExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ConsExprContext) GetHead() IExprContext { return s.head }

func (s *ConsExprContext) GetTail() IExprContext { return s.tail }

func (s *ConsExprContext) SetHead(v IExprContext) { s.head = v }

func (s *ConsExprContext) SetTail(v IExprContext) { s.tail = v }

func (s *ConsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsExprContext) CONS() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserCONS, 0)
}

func (s *ConsExprContext) AllExpr() []IExprContext {
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

func (s *ConsExprContext) Expr(i int) IExprContext {
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

func (s *ConsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterConsExpr(s)
	}
}

func (s *ConsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitConsExpr(s)
	}
}

func (s *ConsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitConsExpr(s)

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
	return s.GetToken(EvalML4ParserVARNAME, 0)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserTIMES, 0)
}

func (s *BinOpExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserPLUS, 0)
}

func (s *BinOpExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserMINUS, 0)
}

func (s *BinOpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserLT, 0)
}

func (s *BinOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterBinOpExpr(s)
	}
}

func (s *BinOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitBinOpExpr(s)
	}
}

func (s *BinOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserLPAREN, 0)
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
	return s.GetToken(EvalML4ParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitParenExpr(s)

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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterFunExpr(s)
	}
}

func (s *FunExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitFunExpr(s)
	}
}

func (s *FunExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitFunExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExprContext struct {
	ExprContext
	matchExpr IExprContext
}

func NewMatchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExprContext {
	var p = new(MatchExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MatchExprContext) GetMatchExpr() IExprContext { return s.matchExpr }

func (s *MatchExprContext) SetMatchExpr(v IExprContext) { s.matchExpr = v }

func (s *MatchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExprContext) MATCH() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserMATCH, 0)
}

func (s *MatchExprContext) WITH() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserWITH, 0)
}

func (s *MatchExprContext) EmptyPattern() IEmptyPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyPatternContext)
}

func (s *MatchExprContext) OR() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserOR, 0)
}

func (s *MatchExprContext) ConsPattern() IConsPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsPatternContext)
}

func (s *MatchExprContext) Expr() IExprContext {
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

func (s *MatchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterMatchExpr(s)
	}
}

func (s *MatchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitMatchExpr(s)
	}
}

func (s *MatchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitMatchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *EvalML4Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, EvalML4ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
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
			p.SetState(32)
			p.Match(EvalML4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.expr(0)
		}
		{
			p.SetState(34)
			p.Match(EvalML4ParserRPAREN)
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
			p.SetState(36)
			p.Fun()
		}

	case 3:
		localctx = NewMatchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(EvalML4ParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)

			var _x = p.expr(0)

			localctx.(*MatchExprContext).matchExpr = _x
		}
		{
			p.SetState(39)
			p.Match(EvalML4ParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.EmptyPattern()
		}
		{
			p.SetState(41)
			p.Match(EvalML4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.ConsPattern()
		}

	case 4:
		localctx = NewEmptyListExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(EvalML4ParserEMPTYLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(EvalML4ParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)

			var _x = p.expr(0)

			localctx.(*IfExprContext).cond = _x
		}
		{
			p.SetState(47)
			p.Match(EvalML4ParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)

			var _x = p.expr(0)

			localctx.(*IfExprContext).then = _x
		}
		{
			p.SetState(49)
			p.Match(EvalML4ParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)

			var _x = p.expr(6)

			localctx.(*IfExprContext).else_ = _x
		}

	case 6:
		localctx = NewLetExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(EvalML4ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)

			var _m = p.Match(EvalML4ParserVARNAME)

			localctx.(*LetExprContext).bindName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(EvalML4ParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)

			var _x = p.expr(0)

			localctx.(*LetExprContext).bindExpr = _x
		}
		{
			p.SetState(56)
			p.Match(EvalML4ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)

			var _x = p.expr(5)

			localctx.(*LetExprContext).body = _x
		}

	case 7:
		localctx = NewLetRecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(EvalML4ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.RecFun()
		}
		{
			p.SetState(61)
			p.Match(EvalML4ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.expr(4)
		}

	case 8:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(EvalML4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(EvalML4ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(EvalML4ParserVARNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(85)
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
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AppExprContext).fn = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_expr)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(70)

					var _x = p.expr(14)

					localctx.(*AppExprContext).arg = _x
				}

			case 2:
				localctx = NewConsExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ConsExprContext).head = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_expr)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(72)
					p.Match(EvalML4ParserCONS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)

					var _x = p.expr(10)

					localctx.(*ConsExprContext).tail = _x
				}

			case 3:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _m = p.Match(EvalML4ParserTIMES)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)

					var _x = p.expr(10)

					localctx.(*BinOpExprContext).right = _x
				}

			case 4:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_expr)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalML4ParserPLUS || _la == EvalML4ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)

					var _x = p.expr(9)

					localctx.(*BinOpExprContext).right = _x
				}

			case 5:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_expr)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(81)

					var _m = p.Match(EvalML4ParserLT)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)

					var _x = p.expr(8)

					localctx.(*BinOpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(87)
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
	p.RuleIndex = EvalML4ParserRULE_defList
	return p
}

func InitEmptyDefListContext(p *DefListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_defList
}

func (*DefListContext) IsDefListContext() {}

func NewDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefListContext {
	var p = new(DefListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_defList

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
	return s.GetTokens(EvalML4ParserCOMMA)
}

func (s *DefListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EvalML4ParserCOMMA, i)
}

func (s *DefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterDefList(s)
	}
}

func (s *DefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitDefList(s)
	}
}

func (s *DefListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitDefList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) DefList() (localctx IDefListContext) {
	localctx = NewDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalML4ParserRULE_defList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Def()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalML4ParserCOMMA {
		{
			p.SetState(89)
			p.Match(EvalML4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Def()
		}

		p.SetState(95)
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
	p.RuleIndex = EvalML4ParserRULE_def
	return p
}

func InitEmptyDefContext(p *DefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_def
}

func (*DefContext) IsDefContext() {}

func NewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefContext {
	var p = new(DefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_def

	return p
}

func (s *DefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserVARNAME, 0)
}

func (s *DefContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEQ, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterDef(s)
	}
}

func (s *DefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitDef(s)
	}
}

func (s *DefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Def() (localctx IDefContext) {
	localctx = NewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvalML4ParserRULE_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(EvalML4ParserVARNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(EvalML4ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.value(0)
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

// IEmptyPatternContext is an interface to support dynamic dispatch.
type IEmptyPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EMPTYLIST() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	Expr() IExprContext

	// IsEmptyPatternContext differentiates from other interfaces.
	IsEmptyPatternContext()
}

type EmptyPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyPatternContext() *EmptyPatternContext {
	var p = new(EmptyPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_emptyPattern
	return p
}

func InitEmptyEmptyPatternContext(p *EmptyPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_emptyPattern
}

func (*EmptyPatternContext) IsEmptyPatternContext() {}

func NewEmptyPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyPatternContext {
	var p = new(EmptyPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_emptyPattern

	return p
}

func (s *EmptyPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyPatternContext) EMPTYLIST() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEMPTYLIST, 0)
}

func (s *EmptyPatternContext) ARROW() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserARROW, 0)
}

func (s *EmptyPatternContext) Expr() IExprContext {
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

func (s *EmptyPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterEmptyPattern(s)
	}
}

func (s *EmptyPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitEmptyPattern(s)
	}
}

func (s *EmptyPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitEmptyPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) EmptyPattern() (localctx IEmptyPatternContext) {
	localctx = NewEmptyPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvalML4ParserRULE_emptyPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(EvalML4ParserEMPTYLIST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(EvalML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
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

// IConsPatternContext is an interface to support dynamic dispatch.
type IConsPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHeadVar returns the headVar token.
	GetHeadVar() antlr.Token

	// GetTailVar returns the tailVar token.
	GetTailVar() antlr.Token

	// SetHeadVar sets the headVar token.
	SetHeadVar(antlr.Token)

	// SetTailVar sets the tailVar token.
	SetTailVar(antlr.Token)

	// Getter signatures
	CONS() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	Expr() IExprContext
	AllVARNAME() []antlr.TerminalNode
	VARNAME(i int) antlr.TerminalNode

	// IsConsPatternContext differentiates from other interfaces.
	IsConsPatternContext()
}

type ConsPatternContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	headVar antlr.Token
	tailVar antlr.Token
}

func NewEmptyConsPatternContext() *ConsPatternContext {
	var p = new(ConsPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_consPattern
	return p
}

func InitEmptyConsPatternContext(p *ConsPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_consPattern
}

func (*ConsPatternContext) IsConsPatternContext() {}

func NewConsPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsPatternContext {
	var p = new(ConsPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_consPattern

	return p
}

func (s *ConsPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsPatternContext) GetHeadVar() antlr.Token { return s.headVar }

func (s *ConsPatternContext) GetTailVar() antlr.Token { return s.tailVar }

func (s *ConsPatternContext) SetHeadVar(v antlr.Token) { s.headVar = v }

func (s *ConsPatternContext) SetTailVar(v antlr.Token) { s.tailVar = v }

func (s *ConsPatternContext) CONS() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserCONS, 0)
}

func (s *ConsPatternContext) ARROW() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserARROW, 0)
}

func (s *ConsPatternContext) Expr() IExprContext {
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

func (s *ConsPatternContext) AllVARNAME() []antlr.TerminalNode {
	return s.GetTokens(EvalML4ParserVARNAME)
}

func (s *ConsPatternContext) VARNAME(i int) antlr.TerminalNode {
	return s.GetToken(EvalML4ParserVARNAME, i)
}

func (s *ConsPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterConsPattern(s)
	}
}

func (s *ConsPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitConsPattern(s)
	}
}

func (s *ConsPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitConsPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) ConsPattern() (localctx IConsPatternContext) {
	localctx = NewConsPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EvalML4ParserRULE_consPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)

		var _m = p.Match(EvalML4ParserVARNAME)

		localctx.(*ConsPatternContext).headVar = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(EvalML4ParserCONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _m = p.Match(EvalML4ParserVARNAME)

		localctx.(*ConsPatternContext).tailVar = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(EvalML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
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
	p.RuleIndex = EvalML4ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() antlr.Token { return s.param }

func (s *FunContext) SetParam(v antlr.Token) { s.param = v }

func (s *FunContext) GetBody() IExprContext { return s.body }

func (s *FunContext) SetBody(v IExprContext) { s.body = v }

func (s *FunContext) FUN() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserFUN, 0)
}

func (s *FunContext) ARROW() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserARROW, 0)
}

func (s *FunContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserVARNAME, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvalML4ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(EvalML4ParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)

		var _m = p.Match(EvalML4ParserVARNAME)

		localctx.(*FunContext).param = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(EvalML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)

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
	p.RuleIndex = EvalML4ParserRULE_recFun
	return p
}

func InitEmptyRecFunContext(p *RecFunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_recFun
}

func (*RecFunContext) IsRecFunContext() {}

func NewRecFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecFunContext {
	var p = new(RecFunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_recFun

	return p
}

func (s *RecFunContext) GetParser() antlr.Parser { return s.parser }

func (s *RecFunContext) GetFunName() antlr.Token { return s.funName }

func (s *RecFunContext) SetFunName(v antlr.Token) { s.funName = v }

func (s *RecFunContext) REC() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserREC, 0)
}

func (s *RecFunContext) EQ() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEQ, 0)
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
	return s.GetToken(EvalML4ParserVARNAME, 0)
}

func (s *RecFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterRecFun(s)
	}
}

func (s *RecFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitRecFun(s)
	}
}

func (s *RecFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitRecFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) RecFun() (localctx IRecFunContext) {
	localctx = NewRecFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EvalML4ParserRULE_recFun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(EvalML4ParserREC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)

		var _m = p.Match(EvalML4ParserVARNAME)

		localctx.(*RecFunContext).funName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(EvalML4ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
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
	p.RuleIndex = EvalML4ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalML4ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalML4ParserRULE_value

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

type EmptyListValueContext struct {
	ValueContext
}

func NewEmptyListValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyListValueContext {
	var p = new(EmptyListValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *EmptyListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyListValueContext) EMPTYLIST() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserEMPTYLIST, 0)
}

func (s *EmptyListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterEmptyListValue(s)
	}
}

func (s *EmptyListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitEmptyListValue(s)
	}
}

func (s *EmptyListValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitEmptyListValue(s)

	default:
		return t.VisitChildren(s)
	}
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
	return s.GetToken(EvalML4ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserLPAREN, 0)
}

func (s *FunValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserRPAREN, 0)
}

func (s *FunValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserLBRACKET, 0)
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
	return s.GetToken(EvalML4ParserRBRACKET, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterFunValue(s)
	}
}

func (s *FunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitFunValue(s)
	}
}

func (s *FunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserLPAREN, 0)
}

func (s *RecFunValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserRPAREN, 0)
}

func (s *RecFunValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserLBRACKET, 0)
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
	return s.GetToken(EvalML4ParserRBRACKET, 0)
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
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterRecFunValue(s)
	}
}

func (s *RecFunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitRecFunValue(s)
	}
}

func (s *RecFunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
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
	return s.GetToken(EvalML4ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConsValueContext struct {
	ValueContext
	head IValueContext
	tail IValueContext
}

func NewConsValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsValueContext {
	var p = new(ConsValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ConsValueContext) GetHead() IValueContext { return s.head }

func (s *ConsValueContext) GetTail() IValueContext { return s.tail }

func (s *ConsValueContext) SetHead(v IValueContext) { s.head = v }

func (s *ConsValueContext) SetTail(v IValueContext) { s.tail = v }

func (s *ConsValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsValueContext) CONS() antlr.TerminalNode {
	return s.GetToken(EvalML4ParserCONS, 0)
}

func (s *ConsValueContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ConsValueContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ConsValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.EnterConsValue(s)
	}
}

func (s *ConsValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalML4Listener); ok {
		listenerT.ExitConsValue(s)
	}
}

func (s *ConsValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalML4Visitor:
		return t.VisitConsValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalML4Parser) Value() (localctx IValueContext) {
	return p.value(0)
}

func (p *EvalML4Parser) value(_p int) (localctx IValueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, EvalML4ParserRULE_value, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(121)
			p.Match(EvalML4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewBoolValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(122)
			p.Match(EvalML4ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(EvalML4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalML4ParserVARNAME {
			{
				p.SetState(124)
				p.DefList()
			}

		}
		{
			p.SetState(127)
			p.Match(EvalML4ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(EvalML4ParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Fun()
		}
		{
			p.SetState(130)
			p.Match(EvalML4ParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewRecFunValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(EvalML4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalML4ParserVARNAME {
			{
				p.SetState(133)
				p.DefList()
			}

		}
		{
			p.SetState(136)
			p.Match(EvalML4ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(EvalML4ParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.RecFun()
		}
		{
			p.SetState(139)
			p.Match(EvalML4ParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewEmptyListValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(EvalML4ParserEMPTYLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConsValueContext(p, NewValueContext(p, _parentctx, _parentState))
			localctx.(*ConsValueContext).head = _prevctx

			p.PushNewRecursionContext(localctx, _startState, EvalML4ParserRULE_value)
			p.SetState(144)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(145)
				p.Match(EvalML4ParserCONS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(146)

				var _x = p.value(1)

				localctx.(*ConsValueContext).tail = _x
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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

func (p *EvalML4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 9:
		var t *ValueContext = nil
		if localctx != nil {
			t = localctx.(*ValueContext)
		}
		return p.Value_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvalML4Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *EvalML4Parser) Value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
