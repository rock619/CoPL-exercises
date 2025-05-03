// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EvalContML1Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EvalContML1LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalcontml1lexerLexerInit() {
	staticData := &EvalContML1LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'evalto'", "'if'", "'then'", "'else'", "'('", "')'", "'>>'", "'_'",
		"'{'", "'}'", "'+'", "'-'", "'*'", "'<'", "", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES",
		"LT", "BOOL", "TRUE", "FALSE", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "PLUS", "MINUS", "TIMES", "LT", "BOOL", "TRUE", "FALSE", "INT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 110, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 83, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 3, 17, 97, 8, 17,
		1, 17, 4, 17, 100, 8, 17, 11, 17, 12, 17, 101, 1, 18, 4, 18, 105, 8, 18,
		11, 18, 12, 18, 106, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9,
		10, 13, 13, 32, 32, 113, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 49, 1, 0, 0,
		0, 7, 54, 1, 0, 0, 0, 9, 59, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13, 63, 1,
		0, 0, 0, 15, 66, 1, 0, 0, 0, 17, 68, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21,
		72, 1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 76, 1, 0, 0, 0, 27, 78, 1, 0, 0,
		0, 29, 82, 1, 0, 0, 0, 31, 84, 1, 0, 0, 0, 33, 89, 1, 0, 0, 0, 35, 96,
		1, 0, 0, 0, 37, 104, 1, 0, 0, 0, 39, 40, 5, 101, 0, 0, 40, 41, 5, 118,
		0, 0, 41, 42, 5, 97, 0, 0, 42, 43, 5, 108, 0, 0, 43, 44, 5, 116, 0, 0,
		44, 45, 5, 111, 0, 0, 45, 2, 1, 0, 0, 0, 46, 47, 5, 105, 0, 0, 47, 48,
		5, 102, 0, 0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 116, 0, 0, 50, 51, 5, 104,
		0, 0, 51, 52, 5, 101, 0, 0, 52, 53, 5, 110, 0, 0, 53, 6, 1, 0, 0, 0, 54,
		55, 5, 101, 0, 0, 55, 56, 5, 108, 0, 0, 56, 57, 5, 115, 0, 0, 57, 58, 5,
		101, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60, 5, 40, 0, 0, 60, 10, 1, 0, 0, 0,
		61, 62, 5, 41, 0, 0, 62, 12, 1, 0, 0, 0, 63, 64, 5, 62, 0, 0, 64, 65, 5,
		62, 0, 0, 65, 14, 1, 0, 0, 0, 66, 67, 5, 95, 0, 0, 67, 16, 1, 0, 0, 0,
		68, 69, 5, 123, 0, 0, 69, 18, 1, 0, 0, 0, 70, 71, 5, 125, 0, 0, 71, 20,
		1, 0, 0, 0, 72, 73, 5, 43, 0, 0, 73, 22, 1, 0, 0, 0, 74, 75, 5, 45, 0,
		0, 75, 24, 1, 0, 0, 0, 76, 77, 5, 42, 0, 0, 77, 26, 1, 0, 0, 0, 78, 79,
		5, 60, 0, 0, 79, 28, 1, 0, 0, 0, 80, 83, 3, 31, 15, 0, 81, 83, 3, 33, 16,
		0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 30, 1, 0, 0, 0, 84, 85,
		5, 116, 0, 0, 85, 86, 5, 114, 0, 0, 86, 87, 5, 117, 0, 0, 87, 88, 5, 101,
		0, 0, 88, 32, 1, 0, 0, 0, 89, 90, 5, 102, 0, 0, 90, 91, 5, 97, 0, 0, 91,
		92, 5, 108, 0, 0, 92, 93, 5, 115, 0, 0, 93, 94, 5, 101, 0, 0, 94, 34, 1,
		0, 0, 0, 95, 97, 3, 23, 11, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 99, 1, 0, 0, 0, 98, 100, 7, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 36, 1, 0, 0,
		0, 103, 105, 7, 1, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		6, 18, 0, 0, 109, 38, 1, 0, 0, 0, 5, 0, 82, 96, 101, 106, 1, 0, 1, 0,
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

// EvalContML1LexerInit initializes any static state used to implement EvalContML1Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEvalContML1Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalContML1LexerInit() {
	staticData := &EvalContML1LexerLexerStaticData
	staticData.once.Do(evalcontml1lexerLexerInit)
}

// NewEvalContML1Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEvalContML1Lexer(input antlr.CharStream) *EvalContML1Lexer {
	EvalContML1LexerInit()
	l := new(EvalContML1Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EvalContML1LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "EvalContML1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EvalContML1Lexer tokens.
const (
	EvalContML1LexerT__0  = 1
	EvalContML1LexerT__1  = 2
	EvalContML1LexerT__2  = 3
	EvalContML1LexerT__3  = 4
	EvalContML1LexerT__4  = 5
	EvalContML1LexerT__5  = 6
	EvalContML1LexerT__6  = 7
	EvalContML1LexerT__7  = 8
	EvalContML1LexerT__8  = 9
	EvalContML1LexerT__9  = 10
	EvalContML1LexerPLUS  = 11
	EvalContML1LexerMINUS = 12
	EvalContML1LexerTIMES = 13
	EvalContML1LexerLT    = 14
	EvalContML1LexerBOOL  = 15
	EvalContML1LexerTRUE  = 16
	EvalContML1LexerFALSE = 17
	EvalContML1LexerINT   = 18
	EvalContML1LexerWS    = 19
)
