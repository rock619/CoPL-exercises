// Code generated from EvalRefML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalRefML3
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

type EvalRefML3Parser struct {
	*antlr.BaseParser
}

var EvalRefML3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalrefml3ParserInit() {
	staticData := &EvalRefML3ParserStaticData
	staticData.LiteralNames = []string{
		"", "'/'", "'|-'", "'evalto'", "'('", "')'", "'['", "']'", "','", "'='",
		"'ref'", "'!'", "':='", "'if'", "'then'", "'else'", "'let'", "'in'",
		"'fun'", "'->'", "'rec'", "'+'", "'-'", "'*'", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "PLUS", "MINUS", "TIMES", "LT", "BOOL", "LOC", "INT",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"eval", "value", "env", "bind", "store", "assign", "exp", "fun", "recFun",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 147, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 3,
		0, 22, 8, 0, 1, 0, 3, 0, 25, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		3, 0, 33, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 42, 8,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 58, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 63, 8, 2, 10,
		2, 12, 2, 66, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 75,
		8, 4, 10, 4, 12, 4, 78, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 116, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 132,
		8, 6, 10, 6, 12, 6, 135, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 0, 1, 12, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0,
		1, 1, 0, 21, 22, 162, 0, 21, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 59, 1, 0,
		0, 0, 6, 67, 1, 0, 0, 0, 8, 71, 1, 0, 0, 0, 10, 79, 1, 0, 0, 0, 12, 115,
		1, 0, 0, 0, 14, 136, 1, 0, 0, 0, 16, 141, 1, 0, 0, 0, 18, 19, 3, 8, 4,
		0, 19, 20, 5, 1, 0, 0, 20, 22, 1, 0, 0, 0, 21, 18, 1, 0, 0, 0, 21, 22,
		1, 0, 0, 0, 22, 24, 1, 0, 0, 0, 23, 25, 3, 4, 2, 0, 24, 23, 1, 0, 0, 0,
		24, 25, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 27, 5, 2, 0, 0, 27, 28, 3,
		12, 6, 0, 28, 29, 5, 3, 0, 0, 29, 32, 3, 2, 1, 0, 30, 31, 5, 1, 0, 0, 31,
		33, 3, 8, 4, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 35, 5, 0, 0, 1, 35, 1, 1, 0, 0, 0, 36, 58, 5, 27, 0, 0, 37, 58,
		5, 25, 0, 0, 38, 58, 5, 26, 0, 0, 39, 41, 5, 4, 0, 0, 40, 42, 3, 4, 2,
		0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44,
		5, 5, 0, 0, 44, 45, 5, 6, 0, 0, 45, 46, 3, 14, 7, 0, 46, 47, 5, 7, 0, 0,
		47, 58, 1, 0, 0, 0, 48, 50, 5, 4, 0, 0, 49, 51, 3, 4, 2, 0, 50, 49, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 5, 5, 0, 0, 53,
		54, 5, 6, 0, 0, 54, 55, 3, 16, 8, 0, 55, 56, 5, 7, 0, 0, 56, 58, 1, 0,
		0, 0, 57, 36, 1, 0, 0, 0, 57, 37, 1, 0, 0, 0, 57, 38, 1, 0, 0, 0, 57, 39,
		1, 0, 0, 0, 57, 48, 1, 0, 0, 0, 58, 3, 1, 0, 0, 0, 59, 64, 3, 6, 3, 0,
		60, 61, 5, 8, 0, 0, 61, 63, 3, 6, 3, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1,
		0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 67, 68, 5, 28, 0, 0, 68, 69, 5, 9, 0, 0, 69, 70, 3, 2,
		1, 0, 70, 7, 1, 0, 0, 0, 71, 76, 3, 10, 5, 0, 72, 73, 5, 8, 0, 0, 73, 75,
		3, 10, 5, 0, 74, 72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		76, 77, 1, 0, 0, 0, 77, 9, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5, 26,
		0, 0, 80, 81, 5, 9, 0, 0, 81, 82, 3, 2, 1, 0, 82, 11, 1, 0, 0, 0, 83, 84,
		6, 6, -1, 0, 84, 85, 5, 4, 0, 0, 85, 86, 3, 12, 6, 0, 86, 87, 5, 5, 0,
		0, 87, 116, 1, 0, 0, 0, 88, 89, 5, 10, 0, 0, 89, 116, 3, 12, 6, 14, 90,
		91, 5, 11, 0, 0, 91, 116, 3, 12, 6, 13, 92, 116, 3, 14, 7, 0, 93, 94, 5,
		13, 0, 0, 94, 95, 3, 12, 6, 0, 95, 96, 5, 14, 0, 0, 96, 97, 3, 12, 6, 0,
		97, 98, 5, 15, 0, 0, 98, 99, 3, 12, 6, 6, 99, 116, 1, 0, 0, 0, 100, 101,
		5, 16, 0, 0, 101, 102, 5, 28, 0, 0, 102, 103, 5, 9, 0, 0, 103, 104, 3,
		12, 6, 0, 104, 105, 5, 17, 0, 0, 105, 106, 3, 12, 6, 5, 106, 116, 1, 0,
		0, 0, 107, 108, 5, 16, 0, 0, 108, 109, 3, 16, 8, 0, 109, 110, 5, 17, 0,
		0, 110, 111, 3, 12, 6, 4, 111, 116, 1, 0, 0, 0, 112, 116, 5, 27, 0, 0,
		113, 116, 5, 25, 0, 0, 114, 116, 5, 28, 0, 0, 115, 83, 1, 0, 0, 0, 115,
		88, 1, 0, 0, 0, 115, 90, 1, 0, 0, 0, 115, 92, 1, 0, 0, 0, 115, 93, 1, 0,
		0, 0, 115, 100, 1, 0, 0, 0, 115, 107, 1, 0, 0, 0, 115, 112, 1, 0, 0, 0,
		115, 113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 133, 1, 0, 0, 0, 117,
		118, 10, 11, 0, 0, 118, 132, 3, 12, 6, 12, 119, 120, 10, 10, 0, 0, 120,
		121, 5, 23, 0, 0, 121, 132, 3, 12, 6, 11, 122, 123, 10, 9, 0, 0, 123, 124,
		7, 0, 0, 0, 124, 132, 3, 12, 6, 10, 125, 126, 10, 8, 0, 0, 126, 127, 5,
		24, 0, 0, 127, 132, 3, 12, 6, 9, 128, 129, 10, 7, 0, 0, 129, 130, 5, 12,
		0, 0, 130, 132, 3, 12, 6, 7, 131, 117, 1, 0, 0, 0, 131, 119, 1, 0, 0, 0,
		131, 122, 1, 0, 0, 0, 131, 125, 1, 0, 0, 0, 131, 128, 1, 0, 0, 0, 132,
		135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 13, 1,
		0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 18, 0, 0, 137, 138, 5, 28,
		0, 0, 138, 139, 5, 19, 0, 0, 139, 140, 3, 12, 6, 0, 140, 15, 1, 0, 0, 0,
		141, 142, 5, 20, 0, 0, 142, 143, 5, 28, 0, 0, 143, 144, 5, 9, 0, 0, 144,
		145, 3, 14, 7, 0, 145, 17, 1, 0, 0, 0, 11, 21, 24, 32, 41, 50, 57, 64,
		76, 115, 131, 133,
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

// EvalRefML3ParserInit initializes any static state used to implement EvalRefML3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalRefML3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalRefML3ParserInit() {
	staticData := &EvalRefML3ParserStaticData
	staticData.once.Do(evalrefml3ParserInit)
}

// NewEvalRefML3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalRefML3Parser(input antlr.TokenStream) *EvalRefML3Parser {
	EvalRefML3ParserInit()
	this := new(EvalRefML3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalRefML3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalRefML3.g4"

	return this
}

// EvalRefML3Parser tokens.
const (
	EvalRefML3ParserEOF        = antlr.TokenEOF
	EvalRefML3ParserT__0       = 1
	EvalRefML3ParserT__1       = 2
	EvalRefML3ParserT__2       = 3
	EvalRefML3ParserT__3       = 4
	EvalRefML3ParserT__4       = 5
	EvalRefML3ParserT__5       = 6
	EvalRefML3ParserT__6       = 7
	EvalRefML3ParserT__7       = 8
	EvalRefML3ParserT__8       = 9
	EvalRefML3ParserT__9       = 10
	EvalRefML3ParserT__10      = 11
	EvalRefML3ParserT__11      = 12
	EvalRefML3ParserT__12      = 13
	EvalRefML3ParserT__13      = 14
	EvalRefML3ParserT__14      = 15
	EvalRefML3ParserT__15      = 16
	EvalRefML3ParserT__16      = 17
	EvalRefML3ParserT__17      = 18
	EvalRefML3ParserT__18      = 19
	EvalRefML3ParserT__19      = 20
	EvalRefML3ParserPLUS       = 21
	EvalRefML3ParserMINUS      = 22
	EvalRefML3ParserTIMES      = 23
	EvalRefML3ParserLT         = 24
	EvalRefML3ParserBOOL       = 25
	EvalRefML3ParserLOC        = 26
	EvalRefML3ParserINT        = 27
	EvalRefML3ParserIDENTIFIER = 28
	EvalRefML3ParserWS         = 29
)

// EvalRefML3Parser rules.
const (
	EvalRefML3ParserRULE_eval   = 0
	EvalRefML3ParserRULE_value  = 1
	EvalRefML3ParserRULE_env    = 2
	EvalRefML3ParserRULE_bind   = 3
	EvalRefML3ParserRULE_store  = 4
	EvalRefML3ParserRULE_assign = 5
	EvalRefML3ParserRULE_exp    = 6
	EvalRefML3ParserRULE_fun    = 7
	EvalRefML3ParserRULE_recFun = 8
)

// IEvalContext is an interface to support dynamic dispatch.
type IEvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStoreIn returns the storeIn rule contexts.
	GetStoreIn() IStoreContext

	// GetStoreOut returns the storeOut rule contexts.
	GetStoreOut() IStoreContext

	// SetStoreIn sets the storeIn rule contexts.
	SetStoreIn(IStoreContext)

	// SetStoreOut sets the storeOut rule contexts.
	SetStoreOut(IStoreContext)

	// Getter signatures
	Exp() IExpContext
	Value() IValueContext
	EOF() antlr.TerminalNode
	Env() IEnvContext
	AllStore() []IStoreContext
	Store(i int) IStoreContext

	// IsEvalContext differentiates from other interfaces.
	IsEvalContext()
}

type EvalContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	storeIn  IStoreContext
	storeOut IStoreContext
}

func NewEmptyEvalContext() *EvalContext {
	var p = new(EvalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) GetStoreIn() IStoreContext { return s.storeIn }

func (s *EvalContext) GetStoreOut() IStoreContext { return s.storeOut }

func (s *EvalContext) SetStoreIn(v IStoreContext) { s.storeIn = v }

func (s *EvalContext) SetStoreOut(v IStoreContext) { s.storeOut = v }

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
	return s.GetToken(EvalRefML3ParserEOF, 0)
}

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

func (s *EvalContext) AllStore() []IStoreContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStoreContext); ok {
			len++
		}
	}

	tst := make([]IStoreContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStoreContext); ok {
			tst[i] = t.(IStoreContext)
			i++
		}
	}

	return tst
}

func (s *EvalContext) Store(i int) IStoreContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStoreContext); ok {
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

	return t.(IStoreContext)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalRefML3ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalRefML3ParserLOC {
		{
			p.SetState(18)

			var _x = p.Store()

			localctx.(*EvalContext).storeIn = _x
		}
		{
			p.SetState(19)
			p.Match(EvalRefML3ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalRefML3ParserIDENTIFIER {
		{
			p.SetState(23)
			p.Env()
		}

	}
	{
		p.SetState(26)
		p.Match(EvalRefML3ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.exp(0)
	}
	{
		p.SetState(28)
		p.Match(EvalRefML3ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Value()
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalRefML3ParserT__0 {
		{
			p.SetState(30)
			p.Match(EvalRefML3ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)

			var _x = p.Store()

			localctx.(*EvalContext).storeOut = _x
		}

	}
	{
		p.SetState(34)
		p.Match(EvalRefML3ParserEOF)
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
	p.RuleIndex = EvalRefML3ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_value

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

type LocValueContext struct {
	ValueContext
}

func NewLocValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocValueContext {
	var p = new(LocValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *LocValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocValueContext) LOC() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserLOC, 0)
}

func (s *LocValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterLocValue(s)
	}
}

func (s *LocValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitLocValue(s)
	}
}

func (s *LocValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitLocValue(s)

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
	return s.GetToken(EvalRefML3ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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

func (s *FunValueContext) Env() IEnvContext {
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

func (s *FunValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterFunValue(s)
	}
}

func (s *FunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitFunValue(s)
	}
}

func (s *FunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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

func (s *RecFunValueContext) Env() IEnvContext {
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

func (s *RecFunValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterRecFunValue(s)
	}
}

func (s *RecFunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitRecFunValue(s)
	}
}

func (s *RecFunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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
	return s.GetToken(EvalRefML3ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EvalRefML3ParserRULE_value)
	var _la int

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Match(EvalRefML3ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(EvalRefML3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewLocValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Match(EvalRefML3ParserLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFunValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Match(EvalRefML3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalRefML3ParserIDENTIFIER {
			{
				p.SetState(40)
				p.Env()
			}

		}
		{
			p.SetState(43)
			p.Match(EvalRefML3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(EvalRefML3ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Fun()
		}
		{
			p.SetState(46)
			p.Match(EvalRefML3ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewRecFunValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(EvalRefML3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalRefML3ParserIDENTIFIER {
			{
				p.SetState(49)
				p.Env()
			}

		}
		{
			p.SetState(52)
			p.Match(EvalRefML3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(EvalRefML3ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.RecFun()
		}
		{
			p.SetState(55)
			p.Match(EvalRefML3ParserT__6)
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
	p.RuleIndex = EvalRefML3ParserRULE_env
	return p
}

func InitEmptyEnvContext(p *EnvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_env
}

func (*EnvContext) IsEnvContext() {}

func NewEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvContext {
	var p = new(EnvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_env

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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterEnv(s)
	}
}

func (s *EnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitEnv(s)
	}
}

func (s *EnvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitEnv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Env() (localctx IEnvContext) {
	localctx = NewEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EvalRefML3ParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Bind()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalRefML3ParserT__7 {
		{
			p.SetState(60)
			p.Match(EvalRefML3ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
	Value() IValueContext

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
	p.RuleIndex = EvalRefML3ParserRULE_bind
	return p
}

func InitEmptyBindContext(p *BindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_bind
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserIDENTIFIER, 0)
}

func (s *BindContext) Value() IValueContext {
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

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalRefML3ParserRULE_bind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(EvalRefML3ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(EvalRefML3ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
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

// IStoreContext is an interface to support dynamic dispatch.
type IStoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssign() []IAssignContext
	Assign(i int) IAssignContext

	// IsStoreContext differentiates from other interfaces.
	IsStoreContext()
}

type StoreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStoreContext() *StoreContext {
	var p = new(StoreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_store
	return p
}

func InitEmptyStoreContext(p *StoreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_store
}

func (*StoreContext) IsStoreContext() {}

func NewStoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreContext {
	var p = new(StoreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_store

	return p
}

func (s *StoreContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreContext) AllAssign() []IAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignContext); ok {
			len++
		}
	}

	tst := make([]IAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignContext); ok {
			tst[i] = t.(IAssignContext)
			i++
		}
	}

	return tst
}

func (s *StoreContext) Assign(i int) IAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
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

	return t.(IAssignContext)
}

func (s *StoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterStore(s)
	}
}

func (s *StoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitStore(s)
	}
}

func (s *StoreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitStore(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Store() (localctx IStoreContext) {
	localctx = NewStoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EvalRefML3ParserRULE_store)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Assign()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalRefML3ParserT__7 {
		{
			p.SetState(72)
			p.Match(EvalRefML3ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Assign()
		}

		p.SetState(78)
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

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOC() antlr.TerminalNode
	Value() IValueContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) LOC() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserLOC, 0)
}

func (s *AssignContext) Value() IValueContext {
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

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvalRefML3ParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(EvalRefML3ParserLOC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(EvalRefML3ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
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
	p.RuleIndex = EvalRefML3ParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_exp

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
	return s.GetToken(EvalRefML3ParserBOOL, 0)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (s *BoolExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitBoolExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RefExpContext struct {
	ExpContext
}

func NewRefExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefExpContext {
	var p = new(RefExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *RefExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefExpContext) Exp() IExpContext {
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

func (s *RefExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterRefExp(s)
	}
}

func (s *RefExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitRefExp(s)
	}
}

func (s *RefExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitRefExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DerefExpContext struct {
	ExpContext
}

func NewDerefExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerefExpContext {
	var p = new(DerefExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *DerefExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerefExpContext) Exp() IExpContext {
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

func (s *DerefExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterDerefExp(s)
	}
}

func (s *DerefExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitDerefExp(s)
	}
}

func (s *DerefExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitDerefExp(s)

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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterFunExp(s)
	}
}

func (s *FunExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitFunExp(s)
	}
}

func (s *FunExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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
	return s.GetToken(EvalRefML3ParserTIMES, 0)
}

func (s *BinOpExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserPLUS, 0)
}

func (s *BinOpExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserMINUS, 0)
}

func (s *BinOpExpContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserLT, 0)
}

func (s *BinOpExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterBinOpExp(s)
	}
}

func (s *BinOpExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitBinOpExp(s)
	}
}

func (s *BinOpExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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

func (s *LetRecExpContext) RecFun() IRecFunContext {
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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterLetRecExp(s)
	}
}

func (s *LetRecExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitLetRecExp(s)
	}
}

func (s *LetRecExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitLetRecExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExpContext struct {
	ExpContext
	left  IExpContext
	right IExpContext
}

func NewAssignExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExpContext {
	var p = new(AssignExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *AssignExpContext) GetLeft() IExpContext { return s.left }

func (s *AssignExpContext) GetRight() IExpContext { return s.right }

func (s *AssignExpContext) SetLeft(v IExpContext) { s.left = v }

func (s *AssignExpContext) SetRight(v IExpContext) { s.right = v }

func (s *AssignExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpContext) AllExp() []IExpContext {
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

func (s *AssignExpContext) Exp(i int) IExpContext {
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

func (s *AssignExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterAssignExp(s)
	}
}

func (s *AssignExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitAssignExp(s)
	}
}

func (s *AssignExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitAssignExp(s)

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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterIfExp(s)
	}
}

func (s *IfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitIfExp(s)
	}
}

func (s *IfExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitIfExp(s)

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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterAppExp(s)
	}
}

func (s *AppExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitAppExp(s)
	}
}

func (s *AppExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetExpContext struct {
	ExpContext
	var_    antlr.Token
	bindExp IExpContext
	body    IExpContext
}

func NewLetExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetExpContext {
	var p = new(LetExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *LetExpContext) GetVar_() antlr.Token { return s.var_ }

func (s *LetExpContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *LetExpContext) GetBindExp() IExpContext { return s.bindExp }

func (s *LetExpContext) GetBody() IExpContext { return s.body }

func (s *LetExpContext) SetBindExp(v IExpContext) { s.bindExp = v }

func (s *LetExpContext) SetBody(v IExpContext) { s.body = v }

func (s *LetExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExpContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterLetExp(s)
	}
}

func (s *LetExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitLetExp(s)
	}
}

func (s *LetExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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

func (s *VarExpContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserIDENTIFIER, 0)
}

func (s *VarExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterVarExp(s)
	}
}

func (s *VarExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitVarExp(s)
	}
}

func (s *VarExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
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
	return s.GetToken(EvalRefML3ParserINT, 0)
}

func (s *IntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterIntExp(s)
	}
}

func (s *IntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitIntExp(s)
	}
}

func (s *IntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitIntExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *EvalRefML3Parser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, EvalRefML3ParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(84)
			p.Match(EvalRefML3ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.exp(0)
		}
		{
			p.SetState(86)
			p.Match(EvalRefML3ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewRefExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(EvalRefML3ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.exp(14)
		}

	case 3:
		localctx = NewDerefExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(EvalRefML3ParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.exp(13)
		}

	case 4:
		localctx = NewFunExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Fun()
		}

	case 5:
		localctx = NewIfExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Match(EvalRefML3ParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)

			var _x = p.exp(0)

			localctx.(*IfExpContext).cond = _x
		}
		{
			p.SetState(95)
			p.Match(EvalRefML3ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)

			var _x = p.exp(0)

			localctx.(*IfExpContext).then = _x
		}
		{
			p.SetState(97)
			p.Match(EvalRefML3ParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)

			var _x = p.exp(6)

			localctx.(*IfExpContext).else_ = _x
		}

	case 6:
		localctx = NewLetExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(EvalRefML3ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)

			var _m = p.Match(EvalRefML3ParserIDENTIFIER)

			localctx.(*LetExpContext).var_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(EvalRefML3ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _x = p.exp(0)

			localctx.(*LetExpContext).bindExp = _x
		}
		{
			p.SetState(104)
			p.Match(EvalRefML3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _x = p.exp(5)

			localctx.(*LetExpContext).body = _x
		}

	case 7:
		localctx = NewLetRecExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(107)
			p.Match(EvalRefML3ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.RecFun()
		}
		{
			p.SetState(109)
			p.Match(EvalRefML3ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)

			var _x = p.exp(4)

			localctx.(*LetRecExpContext).body = _x
		}

	case 8:
		localctx = NewIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.Match(EvalRefML3ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewBoolExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)
			p.Match(EvalRefML3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewVarExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(EvalRefML3ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*AppExpContext).fn = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalRefML3ParserRULE_exp)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(118)

					var _x = p.exp(12)

					localctx.(*AppExpContext).arg = _x
				}

			case 2:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalRefML3ParserRULE_exp)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

					var _m = p.Match(EvalRefML3ParserTIMES)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(121)

					var _x = p.exp(11)

					localctx.(*BinOpExpContext).right = _x
				}

			case 3:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalRefML3ParserRULE_exp)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(123)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalRefML3ParserPLUS || _la == EvalRefML3ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(124)

					var _x = p.exp(10)

					localctx.(*BinOpExpContext).right = _x
				}

			case 4:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalRefML3ParserRULE_exp)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(126)

					var _m = p.Match(EvalRefML3ParserLT)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(127)

					var _x = p.exp(9)

					localctx.(*BinOpExpContext).right = _x
				}

			case 5:
				localctx = NewAssignExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*AssignExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalRefML3ParserRULE_exp)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(129)
					p.Match(EvalRefML3ParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)

					var _x = p.exp(7)

					localctx.(*AssignExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

	// GetParam returns the param token.
	GetParam() antlr.Token

	// SetParam sets the param token.
	SetParam(antlr.Token)

	// GetBody returns the body rule contexts.
	GetBody() IExpContext

	// SetBody sets the body rule contexts.
	SetBody(IExpContext)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Exp() IExpContext

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	param  antlr.Token
	body   IExpContext
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() antlr.Token { return s.param }

func (s *FunContext) SetParam(v antlr.Token) { s.param = v }

func (s *FunContext) GetBody() IExpContext { return s.body }

func (s *FunContext) SetBody(v IExpContext) { s.body = v }

func (s *FunContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalRefML3ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvalRefML3ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(EvalRefML3ParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)

		var _m = p.Match(EvalRefML3ParserIDENTIFIER)

		localctx.(*FunContext).param = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(EvalRefML3ParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)

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
	p.RuleIndex = EvalRefML3ParserRULE_recFun
	return p
}

func InitEmptyRecFunContext(p *RecFunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalRefML3ParserRULE_recFun
}

func (*RecFunContext) IsRecFunContext() {}

func NewRecFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecFunContext {
	var p = new(RecFunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalRefML3ParserRULE_recFun

	return p
}

func (s *RecFunContext) GetParser() antlr.Parser { return s.parser }

func (s *RecFunContext) GetFunName() antlr.Token { return s.funName }

func (s *RecFunContext) SetFunName(v antlr.Token) { s.funName = v }

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
	return s.GetToken(EvalRefML3ParserIDENTIFIER, 0)
}

func (s *RecFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.EnterRecFun(s)
	}
}

func (s *RecFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalRefML3Listener); ok {
		listenerT.ExitRecFun(s)
	}
}

func (s *RecFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalRefML3Visitor:
		return t.VisitRecFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalRefML3Parser) RecFun() (localctx IRecFunContext) {
	localctx = NewRecFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EvalRefML3ParserRULE_recFun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(EvalRefML3ParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)

		var _m = p.Match(EvalRefML3ParserIDENTIFIER)

		localctx.(*RecFunContext).funName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(EvalRefML3ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
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

func (p *EvalRefML3Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvalRefML3Parser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

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
