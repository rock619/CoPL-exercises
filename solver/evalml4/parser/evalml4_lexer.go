// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type EvalML4Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EvalML4LexerLexerStaticData struct {
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

func evalml4lexerLexerInit() {
	staticData := &EvalML4LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"OR", "EMPTYLIST", "CONS", "MATCH", "WITH", "PLUS", "MINUS", "TIMES",
		"LT", "IF", "THEN", "ELSE", "BOOL", "TRUE", "FALSE", "COMMA", "LPAREN",
		"RPAREN", "LET", "IN", "FUN", "EQ", "REC", "ARROW", "LBRACKET", "RBRACKET",
		"TURNSTILE", "EVALTO", "INT", "VARNAME", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 106, 8, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 3, 28, 160,
		8, 28, 1, 28, 4, 28, 163, 8, 28, 11, 28, 12, 28, 164, 1, 29, 1, 29, 5,
		29, 169, 8, 29, 10, 29, 12, 29, 172, 9, 29, 1, 30, 4, 30, 175, 8, 30, 11,
		30, 12, 30, 176, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 1, 0, 4, 1,
		0, 48, 57, 2, 0, 95, 95, 97, 122, 5, 0, 39, 39, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 184, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 65, 1,
		0, 0, 0, 5, 68, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 82,
		1, 0, 0, 0, 13, 84, 1, 0, 0, 0, 15, 86, 1, 0, 0, 0, 17, 88, 1, 0, 0, 0,
		19, 90, 1, 0, 0, 0, 21, 93, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 105, 1,
		0, 0, 0, 27, 107, 1, 0, 0, 0, 29, 112, 1, 0, 0, 0, 31, 118, 1, 0, 0, 0,
		33, 120, 1, 0, 0, 0, 35, 122, 1, 0, 0, 0, 37, 124, 1, 0, 0, 0, 39, 128,
		1, 0, 0, 0, 41, 131, 1, 0, 0, 0, 43, 135, 1, 0, 0, 0, 45, 137, 1, 0, 0,
		0, 47, 141, 1, 0, 0, 0, 49, 144, 1, 0, 0, 0, 51, 146, 1, 0, 0, 0, 53, 148,
		1, 0, 0, 0, 55, 151, 1, 0, 0, 0, 57, 159, 1, 0, 0, 0, 59, 166, 1, 0, 0,
		0, 61, 174, 1, 0, 0, 0, 63, 64, 5, 124, 0, 0, 64, 2, 1, 0, 0, 0, 65, 66,
		3, 49, 24, 0, 66, 67, 3, 51, 25, 0, 67, 4, 1, 0, 0, 0, 68, 69, 5, 58, 0,
		0, 69, 70, 5, 58, 0, 0, 70, 6, 1, 0, 0, 0, 71, 72, 5, 109, 0, 0, 72, 73,
		5, 97, 0, 0, 73, 74, 5, 116, 0, 0, 74, 75, 5, 99, 0, 0, 75, 76, 5, 104,
		0, 0, 76, 8, 1, 0, 0, 0, 77, 78, 5, 119, 0, 0, 78, 79, 5, 105, 0, 0, 79,
		80, 5, 116, 0, 0, 80, 81, 5, 104, 0, 0, 81, 10, 1, 0, 0, 0, 82, 83, 5,
		43, 0, 0, 83, 12, 1, 0, 0, 0, 84, 85, 5, 45, 0, 0, 85, 14, 1, 0, 0, 0,
		86, 87, 5, 42, 0, 0, 87, 16, 1, 0, 0, 0, 88, 89, 5, 60, 0, 0, 89, 18, 1,
		0, 0, 0, 90, 91, 5, 105, 0, 0, 91, 92, 5, 102, 0, 0, 92, 20, 1, 0, 0, 0,
		93, 94, 5, 116, 0, 0, 94, 95, 5, 104, 0, 0, 95, 96, 5, 101, 0, 0, 96, 97,
		5, 110, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 108,
		0, 0, 100, 101, 5, 115, 0, 0, 101, 102, 5, 101, 0, 0, 102, 24, 1, 0, 0,
		0, 103, 106, 3, 27, 13, 0, 104, 106, 3, 29, 14, 0, 105, 103, 1, 0, 0, 0,
		105, 104, 1, 0, 0, 0, 106, 26, 1, 0, 0, 0, 107, 108, 5, 116, 0, 0, 108,
		109, 5, 114, 0, 0, 109, 110, 5, 117, 0, 0, 110, 111, 5, 101, 0, 0, 111,
		28, 1, 0, 0, 0, 112, 113, 5, 102, 0, 0, 113, 114, 5, 97, 0, 0, 114, 115,
		5, 108, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 101, 0, 0, 117, 30,
		1, 0, 0, 0, 118, 119, 5, 44, 0, 0, 119, 32, 1, 0, 0, 0, 120, 121, 5, 40,
		0, 0, 121, 34, 1, 0, 0, 0, 122, 123, 5, 41, 0, 0, 123, 36, 1, 0, 0, 0,
		124, 125, 5, 108, 0, 0, 125, 126, 5, 101, 0, 0, 126, 127, 5, 116, 0, 0,
		127, 38, 1, 0, 0, 0, 128, 129, 5, 105, 0, 0, 129, 130, 5, 110, 0, 0, 130,
		40, 1, 0, 0, 0, 131, 132, 5, 102, 0, 0, 132, 133, 5, 117, 0, 0, 133, 134,
		5, 110, 0, 0, 134, 42, 1, 0, 0, 0, 135, 136, 5, 61, 0, 0, 136, 44, 1, 0,
		0, 0, 137, 138, 5, 114, 0, 0, 138, 139, 5, 101, 0, 0, 139, 140, 5, 99,
		0, 0, 140, 46, 1, 0, 0, 0, 141, 142, 5, 45, 0, 0, 142, 143, 5, 62, 0, 0,
		143, 48, 1, 0, 0, 0, 144, 145, 5, 91, 0, 0, 145, 50, 1, 0, 0, 0, 146, 147,
		5, 93, 0, 0, 147, 52, 1, 0, 0, 0, 148, 149, 5, 124, 0, 0, 149, 150, 5,
		45, 0, 0, 150, 54, 1, 0, 0, 0, 151, 152, 5, 101, 0, 0, 152, 153, 5, 118,
		0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 108, 0, 0, 155, 156, 5, 116,
		0, 0, 156, 157, 5, 111, 0, 0, 157, 56, 1, 0, 0, 0, 158, 160, 3, 13, 6,
		0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161,
		163, 7, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 58, 1, 0, 0, 0, 166, 170, 7, 1,
		0, 0, 167, 169, 7, 2, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0,
		170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 60, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 173, 175, 7, 3, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0,
		178, 179, 6, 30, 0, 0, 179, 62, 1, 0, 0, 0, 6, 0, 105, 159, 164, 170, 176,
		1, 0, 1, 0,
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

// EvalML4LexerInit initializes any static state used to implement EvalML4Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEvalML4Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalML4LexerInit() {
	staticData := &EvalML4LexerLexerStaticData
	staticData.once.Do(evalml4lexerLexerInit)
}

// NewEvalML4Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEvalML4Lexer(input antlr.CharStream) *EvalML4Lexer {
	EvalML4LexerInit()
	l := new(EvalML4Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EvalML4LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "EvalML4.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EvalML4Lexer tokens.
const (
	EvalML4LexerOR        = 1
	EvalML4LexerEMPTYLIST = 2
	EvalML4LexerCONS      = 3
	EvalML4LexerMATCH     = 4
	EvalML4LexerWITH      = 5
	EvalML4LexerPLUS      = 6
	EvalML4LexerMINUS     = 7
	EvalML4LexerTIMES     = 8
	EvalML4LexerLT        = 9
	EvalML4LexerIF        = 10
	EvalML4LexerTHEN      = 11
	EvalML4LexerELSE      = 12
	EvalML4LexerBOOL      = 13
	EvalML4LexerTRUE      = 14
	EvalML4LexerFALSE     = 15
	EvalML4LexerCOMMA     = 16
	EvalML4LexerLPAREN    = 17
	EvalML4LexerRPAREN    = 18
	EvalML4LexerLET       = 19
	EvalML4LexerIN        = 20
	EvalML4LexerFUN       = 21
	EvalML4LexerEQ        = 22
	EvalML4LexerREC       = 23
	EvalML4LexerARROW     = 24
	EvalML4LexerLBRACKET  = 25
	EvalML4LexerRBRACKET  = 26
	EvalML4LexerTURNSTILE = 27
	EvalML4LexerEVALTO    = 28
	EvalML4LexerINT       = 29
	EvalML4LexerVARNAME   = 30
	EvalML4LexerWS        = 31
)
