// Code generated from TypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TypingML4
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

type TypingML4Parser struct {
	*antlr.BaseParser
}

var TypingML4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func typingml4ParserInit() {
	staticData := &TypingML4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'|'", "", "'::'", "'match'", "'with'", "'+'", "'-'", "'*'", "'<'",
		"'if'", "'then'", "'else'", "", "'true'", "'false'", "','", "'('", "')'",
		"'let'", "'in'", "'fun'", "'='", "'rec'", "'->'", "'['", "']'", "'|-'",
		"':'", "'bool'", "'int'", "'list'",
	}
	staticData.SymbolicNames = []string{
		"", "OR", "EMPTYLIST", "CONS", "MATCH", "WITH", "PLUS", "MINUS", "TIMES",
		"LT", "IF", "THEN", "ELSE", "BOOL", "TRUE", "FALSE", "COMMA", "LPAREN",
		"RPAREN", "LET", "IN", "FUN", "EQ", "REC", "ARROW", "LBRACKET", "RBRACKET",
		"TURNSTILE", "COLON", "BOOLTYPE", "INTTYPE", "LISTTYPE", "INT", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"judgement", "type", "expr", "env", "bind", "emptyPattern", "consPattern",
		"fun", "recFun",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 136, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 3, 0, 20, 8, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 42, 8, 1, 10,
		1, 12, 1, 45, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 83, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 99, 8, 2, 10, 2,
		12, 2, 102, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 107, 8, 3, 10, 3, 12, 3, 110,
		9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 0, 2, 2, 4, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 6,
		7, 146, 0, 19, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 82, 1, 0, 0, 0, 6, 103,
		1, 0, 0, 0, 8, 111, 1, 0, 0, 0, 10, 115, 1, 0, 0, 0, 12, 119, 1, 0, 0,
		0, 14, 125, 1, 0, 0, 0, 16, 130, 1, 0, 0, 0, 18, 20, 3, 6, 3, 0, 19, 18,
		1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 5, 27, 0, 0,
		22, 23, 3, 4, 2, 0, 23, 24, 5, 28, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5,
		0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0, 28, 29, 5, 17, 0, 0, 29,
		30, 3, 2, 1, 0, 30, 31, 5, 18, 0, 0, 31, 35, 1, 0, 0, 0, 32, 35, 5, 29,
		0, 0, 33, 35, 5, 30, 0, 0, 34, 27, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34,
		33, 1, 0, 0, 0, 35, 43, 1, 0, 0, 0, 36, 37, 10, 1, 0, 0, 37, 38, 5, 24,
		0, 0, 38, 42, 3, 2, 1, 1, 39, 40, 10, 2, 0, 0, 40, 42, 5, 31, 0, 0, 41,
		36, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0,
		0, 43, 44, 1, 0, 0, 0, 44, 3, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 6,
		2, -1, 0, 47, 48, 5, 17, 0, 0, 48, 49, 3, 4, 2, 0, 49, 50, 5, 18, 0, 0,
		50, 83, 1, 0, 0, 0, 51, 83, 3, 14, 7, 0, 52, 53, 5, 4, 0, 0, 53, 54, 3,
		4, 2, 0, 54, 55, 5, 5, 0, 0, 55, 56, 3, 10, 5, 0, 56, 57, 5, 1, 0, 0, 57,
		58, 3, 12, 6, 0, 58, 83, 1, 0, 0, 0, 59, 83, 5, 2, 0, 0, 60, 61, 5, 10,
		0, 0, 61, 62, 3, 4, 2, 0, 62, 63, 5, 11, 0, 0, 63, 64, 3, 4, 2, 0, 64,
		65, 5, 12, 0, 0, 65, 66, 3, 4, 2, 6, 66, 83, 1, 0, 0, 0, 67, 68, 5, 19,
		0, 0, 68, 69, 5, 33, 0, 0, 69, 70, 5, 22, 0, 0, 70, 71, 3, 4, 2, 0, 71,
		72, 5, 20, 0, 0, 72, 73, 3, 4, 2, 5, 73, 83, 1, 0, 0, 0, 74, 75, 5, 19,
		0, 0, 75, 76, 3, 16, 8, 0, 76, 77, 5, 20, 0, 0, 77, 78, 3, 4, 2, 4, 78,
		83, 1, 0, 0, 0, 79, 83, 5, 32, 0, 0, 80, 83, 5, 13, 0, 0, 81, 83, 5, 33,
		0, 0, 82, 46, 1, 0, 0, 0, 82, 51, 1, 0, 0, 0, 82, 52, 1, 0, 0, 0, 82, 59,
		1, 0, 0, 0, 82, 60, 1, 0, 0, 0, 82, 67, 1, 0, 0, 0, 82, 74, 1, 0, 0, 0,
		82, 79, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 100, 1,
		0, 0, 0, 84, 85, 10, 13, 0, 0, 85, 99, 3, 4, 2, 14, 86, 87, 10, 10, 0,
		0, 87, 88, 5, 3, 0, 0, 88, 99, 3, 4, 2, 10, 89, 90, 10, 9, 0, 0, 90, 91,
		5, 8, 0, 0, 91, 99, 3, 4, 2, 10, 92, 93, 10, 8, 0, 0, 93, 94, 7, 0, 0,
		0, 94, 99, 3, 4, 2, 9, 95, 96, 10, 7, 0, 0, 96, 97, 5, 9, 0, 0, 97, 99,
		3, 4, 2, 8, 98, 84, 1, 0, 0, 0, 98, 86, 1, 0, 0, 0, 98, 89, 1, 0, 0, 0,
		98, 92, 1, 0, 0, 0, 98, 95, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1,
		0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 5, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0,
		103, 108, 3, 8, 4, 0, 104, 105, 5, 16, 0, 0, 105, 107, 3, 8, 4, 0, 106,
		104, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 7, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 5, 33,
		0, 0, 112, 113, 5, 28, 0, 0, 113, 114, 3, 2, 1, 0, 114, 9, 1, 0, 0, 0,
		115, 116, 5, 2, 0, 0, 116, 117, 5, 24, 0, 0, 117, 118, 3, 4, 2, 0, 118,
		11, 1, 0, 0, 0, 119, 120, 5, 33, 0, 0, 120, 121, 5, 3, 0, 0, 121, 122,
		5, 33, 0, 0, 122, 123, 5, 24, 0, 0, 123, 124, 3, 4, 2, 0, 124, 13, 1, 0,
		0, 0, 125, 126, 5, 21, 0, 0, 126, 127, 5, 33, 0, 0, 127, 128, 5, 24, 0,
		0, 128, 129, 3, 4, 2, 0, 129, 15, 1, 0, 0, 0, 130, 131, 5, 23, 0, 0, 131,
		132, 5, 33, 0, 0, 132, 133, 5, 22, 0, 0, 133, 134, 3, 14, 7, 0, 134, 17,
		1, 0, 0, 0, 8, 19, 34, 41, 43, 82, 98, 100, 108,
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

// TypingML4ParserInit initializes any static state used to implement TypingML4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTypingML4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TypingML4ParserInit() {
	staticData := &TypingML4ParserStaticData
	staticData.once.Do(typingml4ParserInit)
}

// NewTypingML4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewTypingML4Parser(input antlr.TokenStream) *TypingML4Parser {
	TypingML4ParserInit()
	this := new(TypingML4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TypingML4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TypingML4.g4"

	return this
}

// TypingML4Parser tokens.
const (
	TypingML4ParserEOF        = antlr.TokenEOF
	TypingML4ParserOR         = 1
	TypingML4ParserEMPTYLIST  = 2
	TypingML4ParserCONS       = 3
	TypingML4ParserMATCH      = 4
	TypingML4ParserWITH       = 5
	TypingML4ParserPLUS       = 6
	TypingML4ParserMINUS      = 7
	TypingML4ParserTIMES      = 8
	TypingML4ParserLT         = 9
	TypingML4ParserIF         = 10
	TypingML4ParserTHEN       = 11
	TypingML4ParserELSE       = 12
	TypingML4ParserBOOL       = 13
	TypingML4ParserTRUE       = 14
	TypingML4ParserFALSE      = 15
	TypingML4ParserCOMMA      = 16
	TypingML4ParserLPAREN     = 17
	TypingML4ParserRPAREN     = 18
	TypingML4ParserLET        = 19
	TypingML4ParserIN         = 20
	TypingML4ParserFUN        = 21
	TypingML4ParserEQ         = 22
	TypingML4ParserREC        = 23
	TypingML4ParserARROW      = 24
	TypingML4ParserLBRACKET   = 25
	TypingML4ParserRBRACKET   = 26
	TypingML4ParserTURNSTILE  = 27
	TypingML4ParserCOLON      = 28
	TypingML4ParserBOOLTYPE   = 29
	TypingML4ParserINTTYPE    = 30
	TypingML4ParserLISTTYPE   = 31
	TypingML4ParserINT        = 32
	TypingML4ParserIDENTIFIER = 33
	TypingML4ParserWS         = 34
)

// TypingML4Parser rules.
const (
	TypingML4ParserRULE_judgement    = 0
	TypingML4ParserRULE_type         = 1
	TypingML4ParserRULE_expr         = 2
	TypingML4ParserRULE_env          = 3
	TypingML4ParserRULE_bind         = 4
	TypingML4ParserRULE_emptyPattern = 5
	TypingML4ParserRULE_consPattern  = 6
	TypingML4ParserRULE_fun          = 7
	TypingML4ParserRULE_recFun       = 8
)

// IJudgementContext is an interface to support dynamic dispatch.
type IJudgementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TURNSTILE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EOF() antlr.TerminalNode
	Env() IEnvContext

	// IsJudgementContext differentiates from other interfaces.
	IsJudgementContext()
}

type JudgementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJudgementContext() *JudgementContext {
	var p = new(JudgementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_judgement
	return p
}

func InitEmptyJudgementContext(p *JudgementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_judgement
}

func (*JudgementContext) IsJudgementContext() {}

func NewJudgementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JudgementContext {
	var p = new(JudgementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_judgement

	return p
}

func (s *JudgementContext) GetParser() antlr.Parser { return s.parser }

func (s *JudgementContext) TURNSTILE() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserTURNSTILE, 0)
}

func (s *JudgementContext) Expr() IExprContext {
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

func (s *JudgementContext) COLON() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserCOLON, 0)
}

func (s *JudgementContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *JudgementContext) EOF() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserEOF, 0)
}

func (s *JudgementContext) Env() IEnvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *JudgementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JudgementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JudgementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterJudgement(s)
	}
}

func (s *JudgementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitJudgement(s)
	}
}

func (s *JudgementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitJudgement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Judgement() (localctx IJudgementContext) {
	localctx = NewJudgementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TypingML4ParserRULE_judgement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TypingML4ParserIDENTIFIER {
		{
			p.SetState(18)
			p.Env()
		}

	}
	{
		p.SetState(21)
		p.Match(TypingML4ParserTURNSTILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.expr(0)
	}
	{
		p.SetState(23)
		p.Match(TypingML4ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.type_(0)
	}
	{
		p.SetState(25)
		p.Match(TypingML4ParserEOF)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolTypeContext struct {
	TypeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) BOOLTYPE() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserBOOLTYPE, 0)
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListTypeContext struct {
	TypeContext
	elementType ITypeContext
}

func NewListTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListTypeContext {
	var p = new(ListTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ListTypeContext) GetElementType() ITypeContext { return s.elementType }

func (s *ListTypeContext) SetElementType(v ITypeContext) { s.elementType = v }

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) LISTTYPE() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserLISTTYPE, 0)
}

func (s *ListTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitListType(s)
	}
}

func (s *ListTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitListType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunTypeContext struct {
	TypeContext
	paramType  ITypeContext
	returnType ITypeContext
}

func NewFunTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunTypeContext {
	var p = new(FunTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *FunTypeContext) GetParamType() ITypeContext { return s.paramType }

func (s *FunTypeContext) GetReturnType() ITypeContext { return s.returnType }

func (s *FunTypeContext) SetParamType(v ITypeContext) { s.paramType = v }

func (s *FunTypeContext) SetReturnType(v ITypeContext) { s.returnType = v }

func (s *FunTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunTypeContext) ARROW() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserARROW, 0)
}

func (s *FunTypeContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *FunTypeContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *FunTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterFunType(s)
	}
}

func (s *FunTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitFunType(s)
	}
}

func (s *FunTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitFunType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenTypeContext struct {
	TypeContext
}

func NewParenTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenTypeContext {
	var p = new(ParenTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ParenTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserLPAREN, 0)
}

func (s *ParenTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParenTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserRPAREN, 0)
}

func (s *ParenTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterParenType(s)
	}
}

func (s *ParenTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitParenType(s)
	}
}

func (s *ParenTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitParenType(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntTypeContext struct {
	TypeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) INTTYPE() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserINTTYPE, 0)
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Type_() (localctx ITypeContext) {
	return p.type_(0)
}

func (p *TypingML4Parser) type_(_p int) (localctx ITypeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TypingML4ParserRULE_type, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TypingML4ParserLPAREN:
		localctx = NewParenTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(28)
			p.Match(TypingML4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.type_(0)
		}
		{
			p.SetState(30)
			p.Match(TypingML4ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TypingML4ParserBOOLTYPE:
		localctx = NewBoolTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(TypingML4ParserBOOLTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TypingML4ParserINTTYPE:
		localctx = NewIntTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(TypingML4ParserINTTYPE)
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
	p.SetState(43)
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
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFunTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
				localctx.(*FunTypeContext).paramType = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_type)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(37)
					p.Match(TypingML4ParserARROW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(38)

					var _x = p.type_(1)

					localctx.(*FunTypeContext).returnType = _x
				}

			case 2:
				localctx = NewListTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
				localctx.(*ListTypeContext).elementType = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_type)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(40)
					p.Match(TypingML4ParserLISTTYPE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(45)
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
	p.RuleIndex = TypingML4ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_expr

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
	return s.GetToken(TypingML4ParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserIF, 0)
}

func (s *IfExprContext) THEN() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserTHEN, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserELSE, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExprContext struct {
	ExprContext
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

func (s *LetExprContext) GetBindExpr() IExprContext { return s.bindExpr }

func (s *LetExprContext) GetBody() IExprContext { return s.body }

func (s *LetExprContext) SetBindExpr(v IExprContext) { s.bindExpr = v }

func (s *LetExprContext) SetBody(v IExprContext) { s.body = v }

func (s *LetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExprContext) LET() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserLET, 0)
}

func (s *LetExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, 0)
}

func (s *LetExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserEQ, 0)
}

func (s *LetExprContext) IN() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIN, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterLetExpr(s)
	}
}

func (s *LetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitLetExpr(s)
	}
}

func (s *LetExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitLetExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetRecExprContext struct {
	ExprContext
	body IExprContext
}

func NewLetRecExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetRecExprContext {
	var p = new(LetRecExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LetRecExprContext) GetBody() IExprContext { return s.body }

func (s *LetRecExprContext) SetBody(v IExprContext) { s.body = v }

func (s *LetRecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetRecExprContext) LET() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserLET, 0)
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
	return s.GetToken(TypingML4ParserIN, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterLetRecExpr(s)
	}
}

func (s *LetRecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitLetRecExpr(s)
	}
}

func (s *LetRecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterAppExpr(s)
	}
}

func (s *AppExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitAppExpr(s)
	}
}

func (s *AppExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserEMPTYLIST, 0)
}

func (s *EmptyListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterEmptyListExpr(s)
	}
}

func (s *EmptyListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitEmptyListExpr(s)
	}
}

func (s *EmptyListExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserCONS, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterConsExpr(s)
	}
}

func (s *ConsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitConsExpr(s)
	}
}

func (s *ConsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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

func (s *VarExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, 0)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (s *VarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserTIMES, 0)
}

func (s *BinOpExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserPLUS, 0)
}

func (s *BinOpExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserMINUS, 0)
}

func (s *BinOpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserLT, 0)
}

func (s *BinOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterBinOpExpr(s)
	}
}

func (s *BinOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitBinOpExpr(s)
	}
}

func (s *BinOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	return s.GetToken(TypingML4ParserLPAREN, 0)
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
	return s.GetToken(TypingML4ParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterFunExpr(s)
	}
}

func (s *FunExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitFunExpr(s)
	}
}

func (s *FunExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitFunExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExprContext struct {
	ExprContext
}

func NewMatchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExprContext {
	var p = new(MatchExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MatchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExprContext) MATCH() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserMATCH, 0)
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

func (s *MatchExprContext) WITH() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserWITH, 0)
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
	return s.GetToken(TypingML4ParserOR, 0)
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

func (s *MatchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterMatchExpr(s)
	}
}

func (s *MatchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitMatchExpr(s)
	}
}

func (s *MatchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitMatchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TypingML4Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, TypingML4ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(47)
			p.Match(TypingML4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.expr(0)
		}
		{
			p.SetState(49)
			p.Match(TypingML4ParserRPAREN)
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
			p.SetState(51)
			p.Fun()
		}

	case 3:
		localctx = NewMatchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(TypingML4ParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.expr(0)
		}
		{
			p.SetState(54)
			p.Match(TypingML4ParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.EmptyPattern()
		}
		{
			p.SetState(56)
			p.Match(TypingML4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.ConsPattern()
		}

	case 4:
		localctx = NewEmptyListExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(TypingML4ParserEMPTYLIST)
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
			p.SetState(60)
			p.Match(TypingML4ParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _x = p.expr(0)

			localctx.(*IfExprContext).cond = _x
		}
		{
			p.SetState(62)
			p.Match(TypingML4ParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)

			var _x = p.expr(0)

			localctx.(*IfExprContext).then = _x
		}
		{
			p.SetState(64)
			p.Match(TypingML4ParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)

			var _x = p.expr(6)

			localctx.(*IfExprContext).else_ = _x
		}

	case 6:
		localctx = NewLetExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(TypingML4ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(TypingML4ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(TypingML4ParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)

			var _x = p.expr(0)

			localctx.(*LetExprContext).bindExpr = _x
		}
		{
			p.SetState(71)
			p.Match(TypingML4ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)

			var _x = p.expr(5)

			localctx.(*LetExprContext).body = _x
		}

	case 7:
		localctx = NewLetRecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(TypingML4ParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.RecFun()
		}
		{
			p.SetState(76)
			p.Match(TypingML4ParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)

			var _x = p.expr(4)

			localctx.(*LetRecExprContext).body = _x
		}

	case 8:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)
			p.Match(TypingML4ParserINT)
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
			p.SetState(80)
			p.Match(TypingML4ParserBOOL)
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
			p.SetState(81)
			p.Match(TypingML4ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AppExprContext).fn = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(85)

					var _x = p.expr(14)

					localctx.(*AppExprContext).arg = _x
				}

			case 2:
				localctx = NewConsExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ConsExprContext).head = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_expr)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(87)
					p.Match(TypingML4ParserCONS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(88)

					var _x = p.expr(10)

					localctx.(*ConsExprContext).tail = _x
				}

			case 3:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_expr)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(90)

					var _m = p.Match(TypingML4ParserTIMES)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(91)

					var _x = p.expr(10)

					localctx.(*BinOpExprContext).right = _x
				}

			case 4:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_expr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(93)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TypingML4ParserPLUS || _la == TypingML4ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(94)

					var _x = p.expr(9)

					localctx.(*BinOpExprContext).right = _x
				}

			case 5:
				localctx = NewBinOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinOpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TypingML4ParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(96)

					var _m = p.Match(TypingML4ParserLT)

					localctx.(*BinOpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)

					var _x = p.expr(8)

					localctx.(*BinOpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

// IEnvContext is an interface to support dynamic dispatch.
type IEnvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBind() []IBindContext
	Bind(i int) IBindContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEnvContext differentiates from other interfaces.
	IsEnvContext()
}

type EnvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvContext() *EnvContext {
	var p = new(EnvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_env
	return p
}

func InitEmptyEnvContext(p *EnvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_env
}

func (*EnvContext) IsEnvContext() {}

func NewEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvContext {
	var p = new(EnvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_env

	return p
}

func (s *EnvContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvContext) AllBind() []IBindContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindContext); ok {
			len++
		}
	}

	tst := make([]IBindContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindContext); ok {
			tst[i] = t.(IBindContext)
			i++
		}
	}

	return tst
}

func (s *EnvContext) Bind(i int) IBindContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindContext); ok {
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

	return t.(IBindContext)
}

func (s *EnvContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TypingML4ParserCOMMA)
}

func (s *EnvContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TypingML4ParserCOMMA, i)
}

func (s *EnvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterEnv(s)
	}
}

func (s *EnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitEnv(s)
	}
}

func (s *EnvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitEnv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Env() (localctx IEnvContext) {
	localctx = NewEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TypingML4ParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Bind()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TypingML4ParserCOMMA {
		{
			p.SetState(104)
			p.Match(TypingML4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Bind()
		}

		p.SetState(110)
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

// IBindContext is an interface to support dynamic dispatch.
type IBindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsBindContext differentiates from other interfaces.
	IsBindContext()
}

type BindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindContext() *BindContext {
	var p = new(BindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_bind
	return p
}

func InitEmptyBindContext(p *BindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_bind
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, 0)
}

func (s *BindContext) COLON() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserCOLON, 0)
}

func (s *BindContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TypingML4ParserRULE_bind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(TypingML4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(TypingML4ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.type_(0)
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
	p.RuleIndex = TypingML4ParserRULE_emptyPattern
	return p
}

func InitEmptyEmptyPatternContext(p *EmptyPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_emptyPattern
}

func (*EmptyPatternContext) IsEmptyPatternContext() {}

func NewEmptyPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyPatternContext {
	var p = new(EmptyPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_emptyPattern

	return p
}

func (s *EmptyPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyPatternContext) EMPTYLIST() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserEMPTYLIST, 0)
}

func (s *EmptyPatternContext) ARROW() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserARROW, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterEmptyPattern(s)
	}
}

func (s *EmptyPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitEmptyPattern(s)
	}
}

func (s *EmptyPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitEmptyPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) EmptyPattern() (localctx IEmptyPatternContext) {
	localctx = NewEmptyPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TypingML4ParserRULE_emptyPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(TypingML4ParserEMPTYLIST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(TypingML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
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
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

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
	p.RuleIndex = TypingML4ParserRULE_consPattern
	return p
}

func InitEmptyConsPatternContext(p *ConsPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_consPattern
}

func (*ConsPatternContext) IsConsPatternContext() {}

func NewConsPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsPatternContext {
	var p = new(ConsPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_consPattern

	return p
}

func (s *ConsPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsPatternContext) GetHeadVar() antlr.Token { return s.headVar }

func (s *ConsPatternContext) GetTailVar() antlr.Token { return s.tailVar }

func (s *ConsPatternContext) SetHeadVar(v antlr.Token) { s.headVar = v }

func (s *ConsPatternContext) SetTailVar(v antlr.Token) { s.tailVar = v }

func (s *ConsPatternContext) CONS() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserCONS, 0)
}

func (s *ConsPatternContext) ARROW() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserARROW, 0)
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

func (s *ConsPatternContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TypingML4ParserIDENTIFIER)
}

func (s *ConsPatternContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, i)
}

func (s *ConsPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterConsPattern(s)
	}
}

func (s *ConsPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitConsPattern(s)
	}
}

func (s *ConsPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitConsPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) ConsPattern() (localctx IConsPatternContext) {
	localctx = NewConsPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TypingML4ParserRULE_consPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)

		var _m = p.Match(TypingML4ParserIDENTIFIER)

		localctx.(*ConsPatternContext).headVar = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(TypingML4ParserCONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)

		var _m = p.Match(TypingML4ParserIDENTIFIER)

		localctx.(*ConsPatternContext).tailVar = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(TypingML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
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
	IDENTIFIER() antlr.TerminalNode
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
	p.RuleIndex = TypingML4ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() antlr.Token { return s.param }

func (s *FunContext) SetParam(v antlr.Token) { s.param = v }

func (s *FunContext) GetBody() IExprContext { return s.body }

func (s *FunContext) SetBody(v IExprContext) { s.body = v }

func (s *FunContext) FUN() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserFUN, 0)
}

func (s *FunContext) ARROW() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserARROW, 0)
}

func (s *FunContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TypingML4ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(TypingML4ParserFUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)

		var _m = p.Match(TypingML4ParserIDENTIFIER)

		localctx.(*FunContext).param = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(TypingML4ParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)

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
	IDENTIFIER() antlr.TerminalNode

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
	p.RuleIndex = TypingML4ParserRULE_recFun
	return p
}

func InitEmptyRecFunContext(p *RecFunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TypingML4ParserRULE_recFun
}

func (*RecFunContext) IsRecFunContext() {}

func NewRecFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecFunContext {
	var p = new(RecFunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypingML4ParserRULE_recFun

	return p
}

func (s *RecFunContext) GetParser() antlr.Parser { return s.parser }

func (s *RecFunContext) GetFunName() antlr.Token { return s.funName }

func (s *RecFunContext) SetFunName(v antlr.Token) { s.funName = v }

func (s *RecFunContext) REC() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserREC, 0)
}

func (s *RecFunContext) EQ() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserEQ, 0)
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

func (s *RecFunContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TypingML4ParserIDENTIFIER, 0)
}

func (s *RecFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.EnterRecFun(s)
	}
}

func (s *RecFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypingML4Listener); ok {
		listenerT.ExitRecFun(s)
	}
}

func (s *RecFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypingML4Visitor:
		return t.VisitRecFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypingML4Parser) RecFun() (localctx IRecFunContext) {
	localctx = NewRecFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TypingML4ParserRULE_recFun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TypingML4ParserREC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)

		var _m = p.Match(TypingML4ParserIDENTIFIER)

		localctx.(*RecFunContext).funName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(TypingML4ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
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

func (p *TypingML4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *TypeContext = nil
		if localctx != nil {
			t = localctx.(*TypeContext)
		}
		return p.Type__Sempred(t, predIndex)

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

func (p *TypingML4Parser) Type__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TypingML4Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
