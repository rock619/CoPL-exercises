// Code generated from PolyTypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PolyTypingML4
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

type PolyTypingML4Parser struct {
	*antlr.BaseParser
}

var PolyTypingML4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func polytypingml4ParserInit() {
	staticData := &PolyTypingML4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'|-'", "':'", "'''", "'bool'", "'int'", "'list'", "'->'", "'('",
		"')'", "'.'", "','", "'::'", "'if'", "'then'", "'else'", "'let'", "'='",
		"'in'", "'rec'", "'match'", "'with'", "'[]'", "'|'", "'fun'", "'+'",
		"'-'", "'*'", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES", "LT", "BOOL",
		"INT", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"eval", "var", "tVar", "type", "tyScheme", "env", "bind", "exp", "fun",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 147, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 39, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 4, 4, 4, 52, 8, 4, 11, 4, 12,
		4, 53, 1, 4, 1, 4, 3, 4, 58, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5,
		65, 8, 5, 10, 5, 12, 5, 68, 9, 5, 3, 5, 70, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 121, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 137, 8, 7, 10, 7, 12, 7, 140, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 0, 2, 6, 14, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0,
		1, 1, 0, 25, 26, 160, 0, 18, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 27, 1, 0,
		0, 0, 6, 38, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 71,
		1, 0, 0, 0, 14, 120, 1, 0, 0, 0, 16, 141, 1, 0, 0, 0, 18, 19, 3, 10, 5,
		0, 19, 20, 5, 1, 0, 0, 20, 21, 3, 14, 7, 0, 21, 22, 5, 2, 0, 0, 22, 23,
		3, 6, 3, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 5, 31, 0, 0,
		26, 3, 1, 0, 0, 0, 27, 28, 5, 3, 0, 0, 28, 29, 5, 31, 0, 0, 29, 5, 1, 0,
		0, 0, 30, 31, 6, 3, -1, 0, 31, 39, 3, 4, 2, 0, 32, 39, 5, 4, 0, 0, 33,
		39, 5, 5, 0, 0, 34, 35, 5, 8, 0, 0, 35, 36, 3, 6, 3, 0, 36, 37, 5, 9, 0,
		0, 37, 39, 1, 0, 0, 0, 38, 30, 1, 0, 0, 0, 38, 32, 1, 0, 0, 0, 38, 33,
		1, 0, 0, 0, 38, 34, 1, 0, 0, 0, 39, 47, 1, 0, 0, 0, 40, 41, 10, 2, 0, 0,
		41, 42, 5, 7, 0, 0, 42, 46, 3, 6, 3, 2, 43, 44, 10, 3, 0, 0, 44, 46, 5,
		6, 0, 0, 45, 40, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47,
		45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 7, 1, 0, 0, 0, 49, 47, 1, 0, 0,
		0, 50, 52, 3, 4, 2, 0, 51, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 5, 10, 0, 0,
		56, 58, 1, 0, 0, 0, 57, 51, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 3, 6, 3, 0, 60, 9, 1, 0, 0, 0, 61, 66, 3, 12, 6, 0, 62,
		63, 5, 11, 0, 0, 63, 65, 3, 12, 6, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0,
		0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66,
		1, 0, 0, 0, 69, 61, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 11, 1, 0, 0, 0,
		71, 72, 3, 2, 1, 0, 72, 73, 5, 2, 0, 0, 73, 74, 3, 8, 4, 0, 74, 13, 1,
		0, 0, 0, 75, 76, 6, 7, -1, 0, 76, 121, 3, 16, 8, 0, 77, 78, 5, 13, 0, 0,
		78, 79, 3, 14, 7, 0, 79, 80, 5, 14, 0, 0, 80, 81, 3, 14, 7, 0, 81, 82,
		5, 15, 0, 0, 82, 83, 3, 14, 7, 9, 83, 121, 1, 0, 0, 0, 84, 85, 5, 16, 0,
		0, 85, 86, 3, 2, 1, 0, 86, 87, 5, 17, 0, 0, 87, 88, 3, 14, 7, 0, 88, 89,
		5, 18, 0, 0, 89, 90, 3, 14, 7, 8, 90, 121, 1, 0, 0, 0, 91, 92, 5, 16, 0,
		0, 92, 93, 5, 19, 0, 0, 93, 94, 3, 2, 1, 0, 94, 95, 5, 17, 0, 0, 95, 96,
		3, 16, 8, 0, 96, 97, 5, 18, 0, 0, 97, 98, 3, 14, 7, 7, 98, 121, 1, 0, 0,
		0, 99, 100, 5, 20, 0, 0, 100, 101, 3, 14, 7, 0, 101, 102, 5, 21, 0, 0,
		102, 103, 5, 22, 0, 0, 103, 104, 5, 7, 0, 0, 104, 105, 3, 14, 7, 0, 105,
		106, 5, 23, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108, 5, 12, 0, 0, 108, 109,
		3, 2, 1, 0, 109, 110, 5, 7, 0, 0, 110, 111, 3, 14, 7, 6, 111, 121, 1, 0,
		0, 0, 112, 121, 5, 22, 0, 0, 113, 121, 5, 30, 0, 0, 114, 121, 5, 29, 0,
		0, 115, 121, 3, 2, 1, 0, 116, 117, 5, 8, 0, 0, 117, 118, 3, 14, 7, 0, 118,
		119, 5, 9, 0, 0, 119, 121, 1, 0, 0, 0, 120, 75, 1, 0, 0, 0, 120, 77, 1,
		0, 0, 0, 120, 84, 1, 0, 0, 0, 120, 91, 1, 0, 0, 0, 120, 99, 1, 0, 0, 0,
		120, 112, 1, 0, 0, 0, 120, 113, 1, 0, 0, 0, 120, 114, 1, 0, 0, 0, 120,
		115, 1, 0, 0, 0, 120, 116, 1, 0, 0, 0, 121, 138, 1, 0, 0, 0, 122, 123,
		10, 14, 0, 0, 123, 137, 3, 14, 7, 15, 124, 125, 10, 13, 0, 0, 125, 126,
		5, 12, 0, 0, 126, 137, 3, 14, 7, 13, 127, 128, 10, 12, 0, 0, 128, 129,
		5, 27, 0, 0, 129, 137, 3, 14, 7, 13, 130, 131, 10, 11, 0, 0, 131, 132,
		7, 0, 0, 0, 132, 137, 3, 14, 7, 12, 133, 134, 10, 10, 0, 0, 134, 135, 5,
		28, 0, 0, 135, 137, 3, 14, 7, 11, 136, 122, 1, 0, 0, 0, 136, 124, 1, 0,
		0, 0, 136, 127, 1, 0, 0, 0, 136, 130, 1, 0, 0, 0, 136, 133, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		15, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 24, 0, 0, 142, 143,
		3, 2, 1, 0, 143, 144, 5, 7, 0, 0, 144, 145, 3, 14, 7, 0, 145, 17, 1, 0,
		0, 0, 10, 38, 45, 47, 53, 57, 66, 69, 120, 136, 138,
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

// PolyTypingML4ParserInit initializes any static state used to implement PolyTypingML4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPolyTypingML4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PolyTypingML4ParserInit() {
	staticData := &PolyTypingML4ParserStaticData
	staticData.once.Do(polytypingml4ParserInit)
}

// NewPolyTypingML4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewPolyTypingML4Parser(input antlr.TokenStream) *PolyTypingML4Parser {
	PolyTypingML4ParserInit()
	this := new(PolyTypingML4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PolyTypingML4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PolyTypingML4.g4"

	return this
}

// PolyTypingML4Parser tokens.
const (
	PolyTypingML4ParserEOF   = antlr.TokenEOF
	PolyTypingML4ParserT__0  = 1
	PolyTypingML4ParserT__1  = 2
	PolyTypingML4ParserT__2  = 3
	PolyTypingML4ParserT__3  = 4
	PolyTypingML4ParserT__4  = 5
	PolyTypingML4ParserT__5  = 6
	PolyTypingML4ParserT__6  = 7
	PolyTypingML4ParserT__7  = 8
	PolyTypingML4ParserT__8  = 9
	PolyTypingML4ParserT__9  = 10
	PolyTypingML4ParserT__10 = 11
	PolyTypingML4ParserT__11 = 12
	PolyTypingML4ParserT__12 = 13
	PolyTypingML4ParserT__13 = 14
	PolyTypingML4ParserT__14 = 15
	PolyTypingML4ParserT__15 = 16
	PolyTypingML4ParserT__16 = 17
	PolyTypingML4ParserT__17 = 18
	PolyTypingML4ParserT__18 = 19
	PolyTypingML4ParserT__19 = 20
	PolyTypingML4ParserT__20 = 21
	PolyTypingML4ParserT__21 = 22
	PolyTypingML4ParserT__22 = 23
	PolyTypingML4ParserT__23 = 24
	PolyTypingML4ParserPLUS  = 25
	PolyTypingML4ParserMINUS = 26
	PolyTypingML4ParserTIMES = 27
	PolyTypingML4ParserLT    = 28
	PolyTypingML4ParserBOOL  = 29
	PolyTypingML4ParserINT   = 30
	PolyTypingML4ParserID    = 31
	PolyTypingML4ParserWS    = 32
)

// PolyTypingML4Parser rules.
const (
	PolyTypingML4ParserRULE_eval     = 0
	PolyTypingML4ParserRULE_var      = 1
	PolyTypingML4ParserRULE_tVar     = 2
	PolyTypingML4ParserRULE_type     = 3
	PolyTypingML4ParserRULE_tyScheme = 4
	PolyTypingML4ParserRULE_env      = 5
	PolyTypingML4ParserRULE_bind     = 6
	PolyTypingML4ParserRULE_exp      = 7
	PolyTypingML4ParserRULE_fun      = 8
)

// IEvalContext is an interface to support dynamic dispatch.
type IEvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Env() IEnvContext
	Exp() IExpContext
	Type_() ITypeContext
	EOF() antlr.TerminalNode

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
	p.RuleIndex = PolyTypingML4ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) Env() IEnvContext {
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

func (s *EvalContext) Type_() ITypeContext {
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

func (s *EvalContext) EOF() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserEOF, 0)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PolyTypingML4ParserRULE_eval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Env()
	}
	{
		p.SetState(19)
		p.Match(PolyTypingML4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.exp(0)
	}
	{
		p.SetState(21)
		p.Match(PolyTypingML4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.type_(0)
	}
	{
		p.SetState(23)
		p.Match(PolyTypingML4ParserEOF)
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

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) ID() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserID, 0)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PolyTypingML4ParserRULE_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(PolyTypingML4ParserID)
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

// ITVarContext is an interface to support dynamic dispatch.
type ITVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTVarContext differentiates from other interfaces.
	IsTVarContext()
}

type TVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTVarContext() *TVarContext {
	var p = new(TVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_tVar
	return p
}

func InitEmptyTVarContext(p *TVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_tVar
}

func (*TVarContext) IsTVarContext() {}

func NewTVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TVarContext {
	var p = new(TVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_tVar

	return p
}

func (s *TVarContext) GetParser() antlr.Parser { return s.parser }

func (s *TVarContext) ID() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserID, 0)
}

func (s *TVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterTVar(s)
	}
}

func (s *TVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitTVar(s)
	}
}

func (s *TVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitTVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) TVar() (localctx ITVarContext) {
	localctx = NewTVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PolyTypingML4ParserRULE_tVar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(PolyTypingML4ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(PolyTypingML4ParserID)
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
	p.RuleIndex = PolyTypingML4ParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_type

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

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeVarContext struct {
	TypeContext
}

func NewTypeVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVarContext {
	var p = new(TypeVarContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *TypeVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVarContext) TVar() ITVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITVarContext)
}

func (s *TypeVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterTypeVar(s)
	}
}

func (s *TypeVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitTypeVar(s)
	}
}

func (s *TypeVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitTypeVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListTypeContext struct {
	TypeContext
	elem ITypeContext
}

func NewListTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListTypeContext {
	var p = new(ListTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ListTypeContext) GetElem() ITypeContext { return s.elem }

func (s *ListTypeContext) SetElem(v ITypeContext) { s.elem = v }

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
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
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitListType(s)
	}
}

func (s *ListTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitListType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunTypeContext struct {
	TypeContext
	param   ITypeContext
	return_ ITypeContext
}

func NewFunTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunTypeContext {
	var p = new(FunTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *FunTypeContext) GetParam() ITypeContext { return s.param }

func (s *FunTypeContext) GetReturn_() ITypeContext { return s.return_ }

func (s *FunTypeContext) SetParam(v ITypeContext) { s.param = v }

func (s *FunTypeContext) SetReturn_(v ITypeContext) { s.return_ = v }

func (s *FunTypeContext) GetRuleContext() antlr.RuleContext {
	return s
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
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterFunType(s)
	}
}

func (s *FunTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitFunType(s)
	}
}

func (s *FunTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
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

func (s *ParenTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterParenType(s)
	}
}

func (s *ParenTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitParenType(s)
	}
}

func (s *ParenTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
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

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Type_() (localctx ITypeContext) {
	return p.type_(0)
}

func (p *PolyTypingML4Parser) type_(_p int) (localctx ITypeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, PolyTypingML4ParserRULE_type, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PolyTypingML4ParserT__2:
		localctx = NewTypeVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(31)
			p.TVar()
		}

	case PolyTypingML4ParserT__3:
		localctx = NewBoolTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(PolyTypingML4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PolyTypingML4ParserT__4:
		localctx = NewIntTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(PolyTypingML4ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PolyTypingML4ParserT__7:
		localctx = NewParenTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(PolyTypingML4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.type_(0)
		}
		{
			p.SetState(36)
			p.Match(PolyTypingML4ParserT__8)
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
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFunTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
				localctx.(*FunTypeContext).param = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_type)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(41)
					p.Match(PolyTypingML4ParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(42)

					var _x = p.type_(2)

					localctx.(*FunTypeContext).return_ = _x
				}

			case 2:
				localctx = NewListTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
				localctx.(*ListTypeContext).elem = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_type)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(44)
					p.Match(PolyTypingML4ParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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

// ITySchemeContext is an interface to support dynamic dispatch.
type ITySchemeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	AllTVar() []ITVarContext
	TVar(i int) ITVarContext

	// IsTySchemeContext differentiates from other interfaces.
	IsTySchemeContext()
}

type TySchemeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTySchemeContext() *TySchemeContext {
	var p = new(TySchemeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_tyScheme
	return p
}

func InitEmptyTySchemeContext(p *TySchemeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_tyScheme
}

func (*TySchemeContext) IsTySchemeContext() {}

func NewTySchemeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TySchemeContext {
	var p = new(TySchemeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_tyScheme

	return p
}

func (s *TySchemeContext) GetParser() antlr.Parser { return s.parser }

func (s *TySchemeContext) Type_() ITypeContext {
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

func (s *TySchemeContext) AllTVar() []ITVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITVarContext); ok {
			len++
		}
	}

	tst := make([]ITVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITVarContext); ok {
			tst[i] = t.(ITVarContext)
			i++
		}
	}

	return tst
}

func (s *TySchemeContext) TVar(i int) ITVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITVarContext); ok {
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

	return t.(ITVarContext)
}

func (s *TySchemeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TySchemeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TySchemeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterTyScheme(s)
	}
}

func (s *TySchemeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitTyScheme(s)
	}
}

func (s *TySchemeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitTyScheme(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) TyScheme() (localctx ITySchemeContext) {
	localctx = NewTySchemeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PolyTypingML4ParserRULE_tyScheme)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PolyTypingML4ParserT__2 {
			{
				p.SetState(50)
				p.TVar()
			}

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(55)
			p.Match(PolyTypingML4ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(59)
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

// IEnvContext is an interface to support dynamic dispatch.
type IEnvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBind() []IBindContext
	Bind(i int) IBindContext

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
	p.RuleIndex = PolyTypingML4ParserRULE_env
	return p
}

func InitEmptyEnvContext(p *EnvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_env
}

func (*EnvContext) IsEnvContext() {}

func NewEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvContext {
	var p = new(EnvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_env

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

func (s *EnvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterEnv(s)
	}
}

func (s *EnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitEnv(s)
	}
}

func (s *EnvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitEnv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Env() (localctx IEnvContext) {
	localctx = NewEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PolyTypingML4ParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PolyTypingML4ParserID {
		{
			p.SetState(61)
			p.Bind()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PolyTypingML4ParserT__10 {
			{
				p.SetState(62)
				p.Match(PolyTypingML4ParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)
				p.Bind()
			}

			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IBindContext is an interface to support dynamic dispatch.
type IBindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	TyScheme() ITySchemeContext

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
	p.RuleIndex = PolyTypingML4ParserRULE_bind
	return p
}

func InitEmptyBindContext(p *BindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_bind
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *BindContext) TyScheme() ITySchemeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITySchemeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITySchemeContext)
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PolyTypingML4ParserRULE_bind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Var_()
	}
	{
		p.SetState(72)
		p.Match(PolyTypingML4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.TyScheme()
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
	p.RuleIndex = PolyTypingML4ParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_exp

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
	return s.GetToken(PolyTypingML4ParserBOOL, 0)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (s *BoolExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitBoolExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConsExpContext struct {
	ExpContext
	head IExpContext
	tail IExpContext
}

func NewConsExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsExpContext {
	var p = new(ConsExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ConsExpContext) GetHead() IExpContext { return s.head }

func (s *ConsExpContext) GetTail() IExpContext { return s.tail }

func (s *ConsExpContext) SetHead(v IExpContext) { s.head = v }

func (s *ConsExpContext) SetTail(v IExpContext) { s.tail = v }

func (s *ConsExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsExpContext) AllExp() []IExpContext {
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

func (s *ConsExpContext) Exp(i int) IExpContext {
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

func (s *ConsExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterConsExp(s)
	}
}

func (s *ConsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitConsExp(s)
	}
}

func (s *ConsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitConsExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunExpContext struct {
	ExpContext
}

func NewFunExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunExpContext {
	var p = new(FunExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *FunExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunExpContext) Fun() IFunContext {
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

func (s *FunExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterFunExp(s)
	}
}

func (s *FunExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitFunExp(s)
	}
}

func (s *FunExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitFunExp(s)

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
	return s.GetToken(PolyTypingML4ParserTIMES, 0)
}

func (s *BinOpExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserPLUS, 0)
}

func (s *BinOpExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserMINUS, 0)
}

func (s *BinOpExpContext) LT() antlr.TerminalNode {
	return s.GetToken(PolyTypingML4ParserLT, 0)
}

func (s *BinOpExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterBinOpExp(s)
	}
}

func (s *BinOpExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitBinOpExp(s)
	}
}

func (s *BinOpExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitBinOpExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetRecExpContext struct {
	ExpContext
	body IExpContext
}

func NewLetRecExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetRecExpContext {
	var p = new(LetRecExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *LetRecExpContext) GetBody() IExpContext { return s.body }

func (s *LetRecExpContext) SetBody(v IExpContext) { s.body = v }

func (s *LetRecExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetRecExpContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *LetRecExpContext) Fun() IFunContext {
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

func (s *LetRecExpContext) Exp() IExpContext {
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

func (s *LetRecExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterLetRecExp(s)
	}
}

func (s *LetRecExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitLetRecExp(s)
	}
}

func (s *LetRecExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitLetRecExp(s)

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
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterIfExp(s)
	}
}

func (s *IfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitIfExp(s)
	}
}

func (s *IfExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitIfExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExpContext struct {
	ExpContext
	arg      IExpContext
	nilCase  IExpContext
	head     IVarContext
	tail     IVarContext
	consCase IExpContext
}

func NewMatchExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExpContext {
	var p = new(MatchExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *MatchExpContext) GetArg() IExpContext { return s.arg }

func (s *MatchExpContext) GetNilCase() IExpContext { return s.nilCase }

func (s *MatchExpContext) GetHead() IVarContext { return s.head }

func (s *MatchExpContext) GetTail() IVarContext { return s.tail }

func (s *MatchExpContext) GetConsCase() IExpContext { return s.consCase }

func (s *MatchExpContext) SetArg(v IExpContext) { s.arg = v }

func (s *MatchExpContext) SetNilCase(v IExpContext) { s.nilCase = v }

func (s *MatchExpContext) SetHead(v IVarContext) { s.head = v }

func (s *MatchExpContext) SetTail(v IVarContext) { s.tail = v }

func (s *MatchExpContext) SetConsCase(v IExpContext) { s.consCase = v }

func (s *MatchExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExpContext) AllExp() []IExpContext {
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

func (s *MatchExpContext) Exp(i int) IExpContext {
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

func (s *MatchExpContext) AllVar_() []IVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarContext); ok {
			len++
		}
	}

	tst := make([]IVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarContext); ok {
			tst[i] = t.(IVarContext)
			i++
		}
	}

	return tst
}

func (s *MatchExpContext) Var_(i int) IVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
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

	return t.(IVarContext)
}

func (s *MatchExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterMatchExp(s)
	}
}

func (s *MatchExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitMatchExp(s)
	}
}

func (s *MatchExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitMatchExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppExpContext struct {
	ExpContext
	fn  IExpContext
	arg IExpContext
}

func NewAppExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppExpContext {
	var p = new(AppExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *AppExpContext) GetFn() IExpContext { return s.fn }

func (s *AppExpContext) GetArg() IExpContext { return s.arg }

func (s *AppExpContext) SetFn(v IExpContext) { s.fn = v }

func (s *AppExpContext) SetArg(v IExpContext) { s.arg = v }

func (s *AppExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppExpContext) AllExp() []IExpContext {
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

func (s *AppExpContext) Exp(i int) IExpContext {
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

func (s *AppExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterAppExp(s)
	}
}

func (s *AppExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitAppExp(s)
	}
}

func (s *AppExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitAppExp(s)

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
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExpContext struct {
	ExpContext
	value IExpContext
	body  IExpContext
}

func NewLetExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetExpContext {
	var p = new(LetExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *LetExpContext) GetValue() IExpContext { return s.value }

func (s *LetExpContext) GetBody() IExpContext { return s.body }

func (s *LetExpContext) SetValue(v IExpContext) { s.value = v }

func (s *LetExpContext) SetBody(v IExpContext) { s.body = v }

func (s *LetExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExpContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *LetExpContext) AllExp() []IExpContext {
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

func (s *LetExpContext) Exp(i int) IExpContext {
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

func (s *LetExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterLetExp(s)
	}
}

func (s *LetExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitLetExp(s)
	}
}

func (s *LetExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitLetExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExpContext struct {
	ExpContext
}

func NewVarExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExpContext {
	var p = new(VarExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *VarExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExpContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *VarExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterVarExp(s)
	}
}

func (s *VarExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitVarExp(s)
	}
}

func (s *VarExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitVarExp(s)

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
	return s.GetToken(PolyTypingML4ParserINT, 0)
}

func (s *IntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterIntExp(s)
	}
}

func (s *IntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitIntExp(s)
	}
}

func (s *IntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitIntExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExpContext struct {
	ExpContext
}

func NewNilExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExpContext {
	var p = new(NilExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *NilExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterNilExp(s)
	}
}

func (s *NilExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitNilExp(s)
	}
}

func (s *NilExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitNilExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *PolyTypingML4Parser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, PolyTypingML4ParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(76)
			p.Fun()
		}

	case 2:
		localctx = NewIfExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(77)
			p.Match(PolyTypingML4ParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)

			var _x = p.exp(0)

			localctx.(*IfExpContext).cond = _x
		}
		{
			p.SetState(79)
			p.Match(PolyTypingML4ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

			var _x = p.exp(0)

			localctx.(*IfExpContext).then = _x
		}
		{
			p.SetState(81)
			p.Match(PolyTypingML4ParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)

			var _x = p.exp(9)

			localctx.(*IfExpContext).else_ = _x
		}

	case 3:
		localctx = NewLetExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(PolyTypingML4ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Var_()
		}
		{
			p.SetState(86)
			p.Match(PolyTypingML4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)

			var _x = p.exp(0)

			localctx.(*LetExpContext).value = _x
		}
		{
			p.SetState(88)
			p.Match(PolyTypingML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)

			var _x = p.exp(8)

			localctx.(*LetExpContext).body = _x
		}

	case 4:
		localctx = NewLetRecExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(PolyTypingML4ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(PolyTypingML4ParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Var_()
		}
		{
			p.SetState(94)
			p.Match(PolyTypingML4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Fun()
		}
		{
			p.SetState(96)
			p.Match(PolyTypingML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _x = p.exp(7)

			localctx.(*LetRecExpContext).body = _x
		}

	case 5:
		localctx = NewMatchExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(PolyTypingML4ParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)

			var _x = p.exp(0)

			localctx.(*MatchExpContext).arg = _x
		}
		{
			p.SetState(101)
			p.Match(PolyTypingML4ParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(PolyTypingML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(PolyTypingML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)

			var _x = p.exp(0)

			localctx.(*MatchExpContext).nilCase = _x
		}
		{
			p.SetState(105)
			p.Match(PolyTypingML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _x = p.Var_()

			localctx.(*MatchExpContext).head = _x
		}
		{
			p.SetState(107)
			p.Match(PolyTypingML4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)

			var _x = p.Var_()

			localctx.(*MatchExpContext).tail = _x
		}
		{
			p.SetState(109)
			p.Match(PolyTypingML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)

			var _x = p.exp(6)

			localctx.(*MatchExpContext).consCase = _x
		}

	case 6:
		localctx = NewNilExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.Match(PolyTypingML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)
			p.Match(PolyTypingML4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewBoolExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(PolyTypingML4ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewVarExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(115)
			p.Var_()
		}

	case 10:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.Match(PolyTypingML4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.exp(0)
		}
		{
			p.SetState(118)
			p.Match(PolyTypingML4ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*AppExpContext).fn = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_exp)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(123)

					var _x = p.exp(15)

					localctx.(*AppExpContext).arg = _x
				}

			case 2:
				localctx = NewConsExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*ConsExpContext).head = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_exp)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(125)
					p.Match(PolyTypingML4ParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(126)

					var _x = p.exp(13)

					localctx.(*ConsExpContext).tail = _x
				}

			case 3:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_exp)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(128)

					var _m = p.Match(PolyTypingML4ParserTIMES)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(129)

					var _x = p.exp(13)

					localctx.(*BinOpExpContext).right = _x
				}

			case 4:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_exp)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(131)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PolyTypingML4ParserPLUS || _la == PolyTypingML4ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(132)

					var _x = p.exp(12)

					localctx.(*BinOpExpContext).right = _x
				}

			case 5:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PolyTypingML4ParserRULE_exp)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(134)

					var _m = p.Match(PolyTypingML4ParserLT)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(135)

					var _x = p.exp(11)

					localctx.(*BinOpExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParam returns the param rule contexts.
	GetParam() IVarContext

	// GetBody returns the body rule contexts.
	GetBody() IExpContext

	// SetParam sets the param rule contexts.
	SetParam(IVarContext)

	// SetBody sets the body rule contexts.
	SetBody(IExpContext)

	// Getter signatures
	Var_() IVarContext
	Exp() IExpContext

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	param  IVarContext
	body   IExpContext
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PolyTypingML4ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolyTypingML4ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() IVarContext { return s.param }

func (s *FunContext) GetBody() IExpContext { return s.body }

func (s *FunContext) SetParam(v IVarContext) { s.param = v }

func (s *FunContext) SetBody(v IExpContext) { s.body = v }

func (s *FunContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *FunContext) Exp() IExpContext {
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

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolyTypingML4Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PolyTypingML4Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PolyTypingML4Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PolyTypingML4ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(PolyTypingML4ParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)

		var _x = p.Var_()

		localctx.(*FunContext).param = _x
	}
	{
		p.SetState(143)
		p.Match(PolyTypingML4ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)

		var _x = p.exp(0)

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

func (p *PolyTypingML4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *TypeContext = nil
		if localctx != nil {
			t = localctx.(*TypeContext)
		}
		return p.Type__Sempred(t, predIndex)

	case 7:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PolyTypingML4Parser) Type__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PolyTypingML4Parser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
