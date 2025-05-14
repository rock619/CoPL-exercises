// Code generated from EvalContML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML4
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

type EvalContML4Parser struct {
	*antlr.BaseParser
}

var EvalContML4ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalcontml4ParserInit() {
	staticData := &EvalContML4ParserStaticData
	staticData.LiteralNames = []string{
		"", "'|-'", "'evalto'", "'('", "')'", "'['", "']'", "'::'", "','", "'='",
		"'match'", "'with'", "'->'", "'|'", "'if'", "'then'", "'else'", "'let'",
		"'in'", "'letcc'", "'fun'", "'rec'", "'_'", "'{'", "'}'", "'>>'", "'+'",
		"'-'", "'*'", "'<'", "", "'true'", "'false'", "'[]'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "PLUS", "MINUS", "TIMES", "LT",
		"BOOL", "TRUE", "FALSE", "NIL_LIST", "INT", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"eval", "value", "env", "bind", "exp", "fun", "recFun", "cont",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 275, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 3, 0, 18, 8, 0, 1, 0, 1, 0,
		1, 0, 3, 0, 23, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 34, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 43,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		55, 8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63, 9, 1, 1,
		2, 1, 2, 1, 2, 5, 2, 68, 8, 2, 10, 2, 12, 2, 71, 9, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		123, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 139, 8, 4, 10, 4, 12, 4, 142, 9, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		3, 7, 157, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 166, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 175, 8, 7, 1, 7, 1,
		7, 3, 7, 179, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 191, 8, 7, 1, 7, 1, 7, 3, 7, 195, 8, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 207, 8, 7, 1, 7, 1, 7,
		3, 7, 211, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 219, 8, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 227, 8, 7, 1, 7, 1, 7, 3, 7, 231,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 240, 8, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 249, 8, 7, 1, 7, 1, 7, 3, 7,
		253, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 271, 8, 7, 3, 7, 273, 8, 7, 1,
		7, 0, 2, 2, 8, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 2, 1, 0, 26, 27, 1, 0,
		26, 29, 316, 0, 17, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 4, 64, 1, 0, 0, 0, 6,
		72, 1, 0, 0, 0, 8, 122, 1, 0, 0, 0, 10, 143, 1, 0, 0, 0, 12, 148, 1, 0,
		0, 0, 14, 272, 1, 0, 0, 0, 16, 18, 3, 4, 2, 0, 17, 16, 1, 0, 0, 0, 17,
		18, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 20, 5, 1, 0, 0, 20, 22, 3, 8, 4,
		0, 21, 23, 3, 14, 7, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24,
		1, 0, 0, 0, 24, 25, 5, 2, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27, 5, 0, 0, 1,
		27, 1, 1, 0, 0, 0, 28, 29, 6, 1, -1, 0, 29, 55, 5, 34, 0, 0, 30, 55, 5,
		30, 0, 0, 31, 33, 5, 3, 0, 0, 32, 34, 3, 4, 2, 0, 33, 32, 1, 0, 0, 0, 33,
		34, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 5, 4, 0, 0, 36, 37, 5, 5, 0,
		0, 37, 38, 3, 10, 5, 0, 38, 39, 5, 6, 0, 0, 39, 55, 1, 0, 0, 0, 40, 42,
		5, 3, 0, 0, 41, 43, 3, 4, 2, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 45, 5, 4, 0, 0, 45, 46, 5, 5, 0, 0, 46, 47, 3,
		12, 6, 0, 47, 48, 5, 6, 0, 0, 48, 55, 1, 0, 0, 0, 49, 55, 5, 33, 0, 0,
		50, 51, 5, 5, 0, 0, 51, 52, 3, 14, 7, 0, 52, 53, 5, 6, 0, 0, 53, 55, 1,
		0, 0, 0, 54, 28, 1, 0, 0, 0, 54, 30, 1, 0, 0, 0, 54, 31, 1, 0, 0, 0, 54,
		40, 1, 0, 0, 0, 54, 49, 1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 55, 61, 1, 0, 0,
		0, 56, 57, 10, 2, 0, 0, 57, 58, 5, 7, 0, 0, 58, 60, 3, 2, 1, 2, 59, 56,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 3, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 69, 3, 6, 3, 0, 65, 66, 5, 8,
		0, 0, 66, 68, 3, 6, 3, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 5, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		72, 73, 5, 35, 0, 0, 73, 74, 5, 9, 0, 0, 74, 75, 3, 2, 1, 0, 75, 7, 1,
		0, 0, 0, 76, 77, 6, 4, -1, 0, 77, 78, 5, 3, 0, 0, 78, 79, 3, 8, 4, 0, 79,
		80, 5, 4, 0, 0, 80, 123, 1, 0, 0, 0, 81, 123, 3, 10, 5, 0, 82, 83, 5, 10,
		0, 0, 83, 84, 3, 8, 4, 0, 84, 85, 5, 11, 0, 0, 85, 86, 5, 33, 0, 0, 86,
		87, 5, 12, 0, 0, 87, 88, 3, 8, 4, 0, 88, 89, 5, 13, 0, 0, 89, 90, 5, 35,
		0, 0, 90, 91, 5, 7, 0, 0, 91, 92, 5, 35, 0, 0, 92, 93, 5, 12, 0, 0, 93,
		94, 3, 8, 4, 13, 94, 123, 1, 0, 0, 0, 95, 123, 5, 33, 0, 0, 96, 97, 5,
		14, 0, 0, 97, 98, 3, 8, 4, 0, 98, 99, 5, 15, 0, 0, 99, 100, 3, 8, 4, 0,
		100, 101, 5, 16, 0, 0, 101, 102, 3, 8, 4, 7, 102, 123, 1, 0, 0, 0, 103,
		104, 5, 17, 0, 0, 104, 105, 5, 35, 0, 0, 105, 106, 5, 9, 0, 0, 106, 107,
		3, 8, 4, 0, 107, 108, 5, 18, 0, 0, 108, 109, 3, 8, 4, 6, 109, 123, 1, 0,
		0, 0, 110, 111, 5, 17, 0, 0, 111, 112, 3, 12, 6, 0, 112, 113, 5, 18, 0,
		0, 113, 114, 3, 8, 4, 5, 114, 123, 1, 0, 0, 0, 115, 116, 5, 19, 0, 0, 116,
		117, 5, 35, 0, 0, 117, 118, 5, 18, 0, 0, 118, 123, 3, 8, 4, 4, 119, 123,
		5, 34, 0, 0, 120, 123, 5, 30, 0, 0, 121, 123, 5, 35, 0, 0, 122, 76, 1,
		0, 0, 0, 122, 81, 1, 0, 0, 0, 122, 82, 1, 0, 0, 0, 122, 95, 1, 0, 0, 0,
		122, 96, 1, 0, 0, 0, 122, 103, 1, 0, 0, 0, 122, 110, 1, 0, 0, 0, 122, 115,
		1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 121, 1, 0,
		0, 0, 123, 140, 1, 0, 0, 0, 124, 125, 10, 14, 0, 0, 125, 139, 3, 8, 4,
		15, 126, 127, 10, 11, 0, 0, 127, 128, 5, 7, 0, 0, 128, 139, 3, 8, 4, 11,
		129, 130, 10, 10, 0, 0, 130, 131, 5, 28, 0, 0, 131, 139, 3, 8, 4, 11, 132,
		133, 10, 9, 0, 0, 133, 134, 7, 0, 0, 0, 134, 139, 3, 8, 4, 10, 135, 136,
		10, 8, 0, 0, 136, 137, 5, 29, 0, 0, 137, 139, 3, 8, 4, 9, 138, 124, 1,
		0, 0, 0, 138, 126, 1, 0, 0, 0, 138, 129, 1, 0, 0, 0, 138, 132, 1, 0, 0,
		0, 138, 135, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 9, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5,
		20, 0, 0, 144, 145, 5, 35, 0, 0, 145, 146, 5, 12, 0, 0, 146, 147, 3, 8,
		4, 0, 147, 11, 1, 0, 0, 0, 148, 149, 5, 21, 0, 0, 149, 150, 5, 35, 0, 0,
		150, 151, 5, 9, 0, 0, 151, 152, 3, 10, 5, 0, 152, 13, 1, 0, 0, 0, 153,
		273, 5, 22, 0, 0, 154, 156, 5, 23, 0, 0, 155, 157, 3, 4, 2, 0, 156, 155,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 5, 1,
		0, 0, 159, 160, 5, 22, 0, 0, 160, 161, 7, 1, 0, 0, 161, 162, 3, 8, 4, 0,
		162, 165, 5, 24, 0, 0, 163, 164, 5, 25, 0, 0, 164, 166, 3, 14, 7, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 273, 1, 0, 0, 0, 167, 168,
		5, 23, 0, 0, 168, 169, 3, 2, 1, 0, 169, 170, 7, 1, 0, 0, 170, 171, 5, 22,
		0, 0, 171, 174, 5, 24, 0, 0, 172, 173, 5, 25, 0, 0, 173, 175, 3, 14, 7,
		0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 273, 1, 0, 0, 0, 176,
		178, 5, 23, 0, 0, 177, 179, 3, 4, 2, 0, 178, 177, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 5, 1, 0, 0, 181, 182, 5, 14,
		0, 0, 182, 183, 5, 22, 0, 0, 183, 184, 5, 15, 0, 0, 184, 185, 3, 8, 4,
		0, 185, 186, 5, 16, 0, 0, 186, 187, 3, 8, 4, 0, 187, 190, 5, 24, 0, 0,
		188, 189, 5, 25, 0, 0, 189, 191, 3, 14, 7, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 273, 1, 0, 0, 0, 192, 194, 5, 23, 0, 0, 193, 195,
		3, 4, 2, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 197, 5, 1, 0, 0, 197, 198, 5, 17, 0, 0, 198, 199, 5, 35, 0,
		0, 199, 200, 5, 9, 0, 0, 200, 201, 5, 22, 0, 0, 201, 202, 5, 18, 0, 0,
		202, 203, 3, 8, 4, 0, 203, 206, 5, 24, 0, 0, 204, 205, 5, 25, 0, 0, 205,
		207, 3, 14, 7, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 273,
		1, 0, 0, 0, 208, 210, 5, 23, 0, 0, 209, 211, 3, 4, 2, 0, 210, 209, 1, 0,
		0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 5, 1, 0, 0,
		213, 214, 5, 22, 0, 0, 214, 215, 3, 8, 4, 0, 215, 218, 5, 24, 0, 0, 216,
		217, 5, 25, 0, 0, 217, 219, 3, 14, 7, 0, 218, 216, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 273, 1, 0, 0, 0, 220, 221, 5, 23, 0, 0, 221, 222, 3, 2,
		1, 0, 222, 223, 5, 22, 0, 0, 223, 226, 5, 24, 0, 0, 224, 225, 5, 25, 0,
		0, 225, 227, 3, 14, 7, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		273, 1, 0, 0, 0, 228, 230, 5, 23, 0, 0, 229, 231, 3, 4, 2, 0, 230, 229,
		1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 5, 1,
		0, 0, 233, 234, 5, 22, 0, 0, 234, 235, 5, 7, 0, 0, 235, 236, 3, 8, 4, 0,
		236, 239, 5, 24, 0, 0, 237, 238, 5, 25, 0, 0, 238, 240, 3, 14, 7, 0, 239,
		237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 273, 1, 0, 0, 0, 241, 242,
		5, 23, 0, 0, 242, 243, 3, 2, 1, 0, 243, 244, 5, 7, 0, 0, 244, 245, 5, 22,
		0, 0, 245, 248, 5, 24, 0, 0, 246, 247, 5, 25, 0, 0, 247, 249, 3, 14, 7,
		0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 273, 1, 0, 0, 0, 250,
		252, 5, 23, 0, 0, 251, 253, 3, 4, 2, 0, 252, 251, 1, 0, 0, 0, 252, 253,
		1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 5, 1, 0, 0, 255, 256, 5, 10,
		0, 0, 256, 257, 5, 22, 0, 0, 257, 258, 5, 11, 0, 0, 258, 259, 5, 33, 0,
		0, 259, 260, 5, 12, 0, 0, 260, 261, 3, 8, 4, 0, 261, 262, 5, 13, 0, 0,
		262, 263, 5, 35, 0, 0, 263, 264, 5, 7, 0, 0, 264, 265, 5, 35, 0, 0, 265,
		266, 5, 12, 0, 0, 266, 267, 3, 8, 4, 0, 267, 270, 5, 24, 0, 0, 268, 269,
		5, 25, 0, 0, 269, 271, 3, 14, 7, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1,
		0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 153, 1, 0, 0, 0, 272, 154, 1, 0, 0,
		0, 272, 167, 1, 0, 0, 0, 272, 176, 1, 0, 0, 0, 272, 192, 1, 0, 0, 0, 272,
		208, 1, 0, 0, 0, 272, 220, 1, 0, 0, 0, 272, 228, 1, 0, 0, 0, 272, 241,
		1, 0, 0, 0, 272, 250, 1, 0, 0, 0, 273, 15, 1, 0, 0, 0, 26, 17, 22, 33,
		42, 54, 61, 69, 122, 138, 140, 156, 165, 174, 178, 190, 194, 206, 210,
		218, 226, 230, 239, 248, 252, 270, 272,
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

// EvalContML4ParserInit initializes any static state used to implement EvalContML4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEvalContML4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalContML4ParserInit() {
	staticData := &EvalContML4ParserStaticData
	staticData.once.Do(evalcontml4ParserInit)
}

// NewEvalContML4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewEvalContML4Parser(input antlr.TokenStream) *EvalContML4Parser {
	EvalContML4ParserInit()
	this := new(EvalContML4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalContML4ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "EvalContML4.g4"

	return this
}

// EvalContML4Parser tokens.
const (
	EvalContML4ParserEOF        = antlr.TokenEOF
	EvalContML4ParserT__0       = 1
	EvalContML4ParserT__1       = 2
	EvalContML4ParserT__2       = 3
	EvalContML4ParserT__3       = 4
	EvalContML4ParserT__4       = 5
	EvalContML4ParserT__5       = 6
	EvalContML4ParserT__6       = 7
	EvalContML4ParserT__7       = 8
	EvalContML4ParserT__8       = 9
	EvalContML4ParserT__9       = 10
	EvalContML4ParserT__10      = 11
	EvalContML4ParserT__11      = 12
	EvalContML4ParserT__12      = 13
	EvalContML4ParserT__13      = 14
	EvalContML4ParserT__14      = 15
	EvalContML4ParserT__15      = 16
	EvalContML4ParserT__16      = 17
	EvalContML4ParserT__17      = 18
	EvalContML4ParserT__18      = 19
	EvalContML4ParserT__19      = 20
	EvalContML4ParserT__20      = 21
	EvalContML4ParserT__21      = 22
	EvalContML4ParserT__22      = 23
	EvalContML4ParserT__23      = 24
	EvalContML4ParserT__24      = 25
	EvalContML4ParserPLUS       = 26
	EvalContML4ParserMINUS      = 27
	EvalContML4ParserTIMES      = 28
	EvalContML4ParserLT         = 29
	EvalContML4ParserBOOL       = 30
	EvalContML4ParserTRUE       = 31
	EvalContML4ParserFALSE      = 32
	EvalContML4ParserNIL_LIST   = 33
	EvalContML4ParserINT        = 34
	EvalContML4ParserIDENTIFIER = 35
	EvalContML4ParserWS         = 36
)

// EvalContML4Parser rules.
const (
	EvalContML4ParserRULE_eval   = 0
	EvalContML4ParserRULE_value  = 1
	EvalContML4ParserRULE_env    = 2
	EvalContML4ParserRULE_bind   = 3
	EvalContML4ParserRULE_exp    = 4
	EvalContML4ParserRULE_fun    = 5
	EvalContML4ParserRULE_recFun = 6
	EvalContML4ParserRULE_cont   = 7
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
	Env() IEnvContext
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
	p.RuleIndex = EvalContML4ParserRULE_eval
	return p
}

func InitEmptyEvalContext(p *EvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_eval
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_eval

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
	return s.GetToken(EvalContML4ParserEOF, 0)
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitEval(s)
	}
}

func (s *EvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitEval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EvalContML4ParserRULE_eval)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalContML4ParserIDENTIFIER {
		{
			p.SetState(16)
			p.Env()
		}

	}
	{
		p.SetState(19)
		p.Match(EvalContML4ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.exp(0)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == EvalContML4ParserT__21 || _la == EvalContML4ParserT__22 {
		{
			p.SetState(21)
			p.Cont()
		}

	}
	{
		p.SetState(24)
		p.Match(EvalContML4ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.value(0)
	}
	{
		p.SetState(26)
		p.Match(EvalContML4ParserEOF)
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
	p.RuleIndex = EvalContML4ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_value

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
	return s.GetToken(EvalContML4ParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterFunValue(s)
	}
}

func (s *FunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitFunValue(s)
	}
}

func (s *FunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterRecFunValue(s)
	}
}

func (s *RecFunValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitRecFunValue(s)
	}
}

func (s *RecFunValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitRecFunValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContValueContext struct {
	ValueContext
}

func NewContValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContValueContext {
	var p = new(ContValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ContValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContValueContext) Cont() IContContext {
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

func (s *ContValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterContValue(s)
	}
}

func (s *ContValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitContValue(s)
	}
}

func (s *ContValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitContValue(s)

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
	return s.GetToken(EvalContML4ParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilValueContext struct {
	ValueContext
}

func NewNilValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilValueContext {
	var p = new(NilValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NilValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilValueContext) NIL_LIST() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserNIL_LIST, 0)
}

func (s *NilValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterNilValue(s)
	}
}

func (s *NilValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitNilValue(s)
	}
}

func (s *NilValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitNilValue(s)

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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterConsValue(s)
	}
}

func (s *ConsValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitConsValue(s)
	}
}

func (s *ConsValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitConsValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Value() (localctx IValueContext) {
	return p.value(0)
}

func (p *EvalContML4Parser) value(_p int) (localctx IValueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, EvalContML4ParserRULE_value, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(29)
			p.Match(EvalContML4ParserINT)
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
			p.SetState(30)
			p.Match(EvalContML4ParserBOOL)
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
			p.SetState(31)
			p.Match(EvalContML4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(32)
				p.Env()
			}

		}
		{
			p.SetState(35)
			p.Match(EvalContML4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(EvalContML4ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Fun()
		}
		{
			p.SetState(38)
			p.Match(EvalContML4ParserT__5)
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
			p.SetState(40)
			p.Match(EvalContML4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(41)
				p.Env()
			}

		}
		{
			p.SetState(44)
			p.Match(EvalContML4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(EvalContML4ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.RecFun()
		}
		{
			p.SetState(47)
			p.Match(EvalContML4ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewNilValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(EvalContML4ParserNIL_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewContValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(EvalContML4ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Cont()
		}
		{
			p.SetState(52)
			p.Match(EvalContML4ParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

			p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_value)
			p.SetState(56)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(57)
				p.Match(EvalContML4ParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(58)

				var _x = p.value(2)

				localctx.(*ConsValueContext).tail = _x
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.RuleIndex = EvalContML4ParserRULE_env
	return p
}

func InitEmptyEnvContext(p *EnvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_env
}

func (*EnvContext) IsEnvContext() {}

func NewEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvContext {
	var p = new(EnvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_env

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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterEnv(s)
	}
}

func (s *EnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitEnv(s)
	}
}

func (s *EnvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitEnv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Env() (localctx IEnvContext) {
	localctx = NewEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EvalContML4ParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Bind()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == EvalContML4ParserT__7 {
		{
			p.SetState(65)
			p.Match(EvalContML4ParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Bind()
		}

		p.SetState(71)
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
	p.RuleIndex = EvalContML4ParserRULE_bind
	return p
}

func InitEmptyBindContext(p *BindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_bind
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EvalContML4ParserRULE_bind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(EvalContML4ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(EvalContML4ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
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
	p.RuleIndex = EvalContML4ParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_exp

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
	return s.GetToken(EvalContML4ParserBOOL, 0)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (s *BoolExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterConsExp(s)
	}
}

func (s *ConsExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitConsExp(s)
	}
}

func (s *ConsExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterFunExp(s)
	}
}

func (s *FunExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitFunExp(s)
	}
}

func (s *FunExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	return s.GetToken(EvalContML4ParserTIMES, 0)
}

func (s *BinOpExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserPLUS, 0)
}

func (s *BinOpExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserMINUS, 0)
}

func (s *BinOpExpContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserLT, 0)
}

func (s *BinOpExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterBinOpExp(s)
	}
}

func (s *BinOpExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitBinOpExp(s)
	}
}

func (s *BinOpExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterLetRecExp(s)
	}
}

func (s *LetRecExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitLetRecExp(s)
	}
}

func (s *LetRecExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitLetRecExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetCCExpContext struct {
	ExpContext
	var_ antlr.Token
	body IExpContext
}

func NewLetCCExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetCCExpContext {
	var p = new(LetCCExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *LetCCExpContext) GetVar_() antlr.Token { return s.var_ }

func (s *LetCCExpContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *LetCCExpContext) GetBody() IExpContext { return s.body }

func (s *LetCCExpContext) SetBody(v IExpContext) { s.body = v }

func (s *LetCCExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetCCExpContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
}

func (s *LetCCExpContext) Exp() IExpContext {
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

func (s *LetCCExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterLetCCExp(s)
	}
}

func (s *LetCCExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitLetCCExp(s)
	}
}

func (s *LetCCExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitLetCCExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExpContext struct {
	ExpContext
	matchedExp IExpContext
	nilExp     IExpContext
	headVar    antlr.Token
	tailVar    antlr.Token
	consExp    IExpContext
}

func NewMatchExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExpContext {
	var p = new(MatchExpContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *MatchExpContext) GetHeadVar() antlr.Token { return s.headVar }

func (s *MatchExpContext) GetTailVar() antlr.Token { return s.tailVar }

func (s *MatchExpContext) SetHeadVar(v antlr.Token) { s.headVar = v }

func (s *MatchExpContext) SetTailVar(v antlr.Token) { s.tailVar = v }

func (s *MatchExpContext) GetMatchedExp() IExpContext { return s.matchedExp }

func (s *MatchExpContext) GetNilExp() IExpContext { return s.nilExp }

func (s *MatchExpContext) GetConsExp() IExpContext { return s.consExp }

func (s *MatchExpContext) SetMatchedExp(v IExpContext) { s.matchedExp = v }

func (s *MatchExpContext) SetNilExp(v IExpContext) { s.nilExp = v }

func (s *MatchExpContext) SetConsExp(v IExpContext) { s.consExp = v }

func (s *MatchExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExpContext) NIL_LIST() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserNIL_LIST, 0)
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

func (s *MatchExpContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EvalContML4ParserIDENTIFIER)
}

func (s *MatchExpContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, i)
}

func (s *MatchExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterMatchExp(s)
	}
}

func (s *MatchExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitMatchExp(s)
	}
}

func (s *MatchExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitMatchExp(s)

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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterIfExp(s)
	}
}

func (s *IfExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitIfExp(s)
	}
}

func (s *IfExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterAppExp(s)
	}
}

func (s *AppExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitAppExp(s)
	}
}

func (s *AppExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterLetExp(s)
	}
}

func (s *LetExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitLetExp(s)
	}
}

func (s *LetExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
}

func (s *VarExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterVarExp(s)
	}
}

func (s *VarExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitVarExp(s)
	}
}

func (s *VarExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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
	return s.GetToken(EvalContML4ParserINT, 0)
}

func (s *IntExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterIntExp(s)
	}
}

func (s *IntExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitIntExp(s)
	}
}

func (s *IntExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
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

func (s *NilExpContext) NIL_LIST() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserNIL_LIST, 0)
}

func (s *NilExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterNilExp(s)
	}
}

func (s *NilExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitNilExp(s)
	}
}

func (s *NilExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitNilExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *EvalContML4Parser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, EvalContML4ParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(77)
			p.Match(EvalContML4ParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.exp(0)
		}
		{
			p.SetState(79)
			p.Match(EvalContML4ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Fun()
		}

	case 3:
		localctx = NewMatchExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Match(EvalContML4ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)

			var _x = p.exp(0)

			localctx.(*MatchExpContext).matchedExp = _x
		}
		{
			p.SetState(84)
			p.Match(EvalContML4ParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(EvalContML4ParserNIL_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(EvalContML4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)

			var _x = p.exp(0)

			localctx.(*MatchExpContext).nilExp = _x
		}
		{
			p.SetState(88)
			p.Match(EvalContML4ParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*MatchExpContext).headVar = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(EvalContML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*MatchExpContext).tailVar = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(EvalContML4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)

			var _x = p.exp(13)

			localctx.(*MatchExpContext).consExp = _x
		}

	case 4:
		localctx = NewNilExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(EvalContML4ParserNIL_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIfExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(EvalContML4ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _x = p.exp(0)

			localctx.(*IfExpContext).cond = _x
		}
		{
			p.SetState(98)
			p.Match(EvalContML4ParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)

			var _x = p.exp(0)

			localctx.(*IfExpContext).then = _x
		}
		{
			p.SetState(100)
			p.Match(EvalContML4ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)

			var _x = p.exp(7)

			localctx.(*IfExpContext).else_ = _x
		}

	case 6:
		localctx = NewLetExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(103)
			p.Match(EvalContML4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*LetExpContext).var_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(EvalContML4ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _x = p.exp(0)

			localctx.(*LetExpContext).bindExp = _x
		}
		{
			p.SetState(107)
			p.Match(EvalContML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)

			var _x = p.exp(6)

			localctx.(*LetExpContext).body = _x
		}

	case 7:
		localctx = NewLetRecExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.Match(EvalContML4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.RecFun()
		}
		{
			p.SetState(112)
			p.Match(EvalContML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _x = p.exp(5)

			localctx.(*LetRecExpContext).body = _x
		}

	case 8:
		localctx = NewLetCCExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(115)
			p.Match(EvalContML4ParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*LetCCExpContext).var_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(EvalContML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.exp(4)

			localctx.(*LetCCExpContext).body = _x
		}

	case 9:
		localctx = NewIntExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(119)
			p.Match(EvalContML4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewBoolExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(120)
			p.Match(EvalContML4ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewVarExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(121)
			p.Match(EvalContML4ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(140)
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
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAppExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*AppExpContext).fn = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_exp)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(125)

					var _x = p.exp(15)

					localctx.(*AppExpContext).arg = _x
				}

			case 2:
				localctx = NewConsExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*ConsExpContext).head = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_exp)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(127)
					p.Match(EvalContML4ParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)

					var _x = p.exp(11)

					localctx.(*ConsExpContext).tail = _x
				}

			case 3:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_exp)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(130)

					var _m = p.Match(EvalContML4ParserTIMES)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)

					var _x = p.exp(11)

					localctx.(*BinOpExpContext).right = _x
				}

			case 4:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_exp)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(133)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinOpExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == EvalContML4ParserPLUS || _la == EvalContML4ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinOpExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(134)

					var _x = p.exp(10)

					localctx.(*BinOpExpContext).right = _x
				}

			case 5:
				localctx = NewBinOpExpContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*BinOpExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, EvalContML4ParserRULE_exp)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(136)

					var _m = p.Match(EvalContML4ParserLT)

					localctx.(*BinOpExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(137)

					var _x = p.exp(9)

					localctx.(*BinOpExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(142)
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
	p.RuleIndex = EvalContML4ParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetParam() antlr.Token { return s.param }

func (s *FunContext) SetParam(v antlr.Token) { s.param = v }

func (s *FunContext) GetBody() IExpContext { return s.body }

func (s *FunContext) SetBody(v IExpContext) { s.body = v }

func (s *FunContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitFun(s)
	}
}

func (s *FunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EvalContML4ParserRULE_fun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(EvalContML4ParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)

		var _m = p.Match(EvalContML4ParserIDENTIFIER)

		localctx.(*FunContext).param = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(EvalContML4ParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)

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
	p.RuleIndex = EvalContML4ParserRULE_recFun
	return p
}

func InitEmptyRecFunContext(p *RecFunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_recFun
}

func (*RecFunContext) IsRecFunContext() {}

func NewRecFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecFunContext {
	var p = new(RecFunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_recFun

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
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
}

func (s *RecFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterRecFun(s)
	}
}

func (s *RecFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitRecFun(s)
	}
}

func (s *RecFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitRecFun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) RecFun() (localctx IRecFunContext) {
	localctx = NewRecFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EvalContML4ParserRULE_recFun)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(EvalContML4ParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)

		var _m = p.Match(EvalContML4ParserIDENTIFIER)

		localctx.(*RecFunContext).funName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(EvalContML4ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
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
	p.RuleIndex = EvalContML4ParserRULE_cont
	return p
}

func InitEmptyContContext(p *ContContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = EvalContML4ParserRULE_cont
}

func (*ContContext) IsContContext() {}

func NewContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContContext {
	var p = new(ContContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = EvalContML4ParserRULE_cont

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
	next  IContContext
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

func (s *IfContContext) GetNext() IContContext { return s.next }

func (s *IfContContext) SetThen(v IExpContext) { s.then = v }

func (s *IfContContext) SetElse_(v IExpContext) { s.else_ = v }

func (s *IfContContext) SetNext(v IContContext) { s.next = v }

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

func (s *IfContContext) Env() IEnvContext {
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterIfCont(s)
	}
}

func (s *IfContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitIfCont(s)
	}
}

func (s *IfContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitIfCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConsExpContContext struct {
	ContContext
	next IContContext
}

func NewConsExpContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsExpContContext {
	var p = new(ConsExpContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *ConsExpContContext) GetNext() IContContext { return s.next }

func (s *ConsExpContContext) SetNext(v IContContext) { s.next = v }

func (s *ConsExpContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsExpContContext) Exp() IExpContext {
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

func (s *ConsExpContContext) Env() IEnvContext {
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

func (s *ConsExpContContext) Cont() IContContext {
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

func (s *ConsExpContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterConsExpCont(s)
	}
}

func (s *ConsExpContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitConsExpCont(s)
	}
}

func (s *ConsExpContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitConsExpCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpContContext struct {
	ContContext
	op   antlr.Token
	next IContContext
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

func (s *ExpContContext) GetNext() IContContext { return s.next }

func (s *ExpContContext) SetNext(v IContContext) { s.next = v }

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
	return s.GetToken(EvalContML4ParserPLUS, 0)
}

func (s *ExpContContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserMINUS, 0)
}

func (s *ExpContContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserTIMES, 0)
}

func (s *ExpContContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserLT, 0)
}

func (s *ExpContContext) Env() IEnvContext {
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterExpCont(s)
	}
}

func (s *ExpContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitExpCont(s)
	}
}

func (s *ExpContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitExpCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type TerminalContContext struct {
	ContContext
}

func NewTerminalContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TerminalContContext {
	var p = new(TerminalContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *TerminalContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminalContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterTerminalCont(s)
	}
}

func (s *TerminalContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitTerminalCont(s)
	}
}

func (s *TerminalContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitTerminalCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppValueContContext struct {
	ContContext
	next IContContext
}

func NewAppValueContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppValueContContext {
	var p = new(AppValueContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *AppValueContContext) GetNext() IContContext { return s.next }

func (s *AppValueContContext) SetNext(v IContContext) { s.next = v }

func (s *AppValueContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppValueContContext) Value() IValueContext {
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

func (s *AppValueContContext) Cont() IContContext {
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

func (s *AppValueContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterAppValueCont(s)
	}
}

func (s *AppValueContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitAppValueCont(s)
	}
}

func (s *AppValueContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitAppValueCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchContContext struct {
	ContContext
	nilExp  IExpContext
	headVar antlr.Token
	tailVar antlr.Token
	consExp IExpContext
	next    IContContext
}

func NewMatchContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchContContext {
	var p = new(MatchContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *MatchContContext) GetHeadVar() antlr.Token { return s.headVar }

func (s *MatchContContext) GetTailVar() antlr.Token { return s.tailVar }

func (s *MatchContContext) SetHeadVar(v antlr.Token) { s.headVar = v }

func (s *MatchContContext) SetTailVar(v antlr.Token) { s.tailVar = v }

func (s *MatchContContext) GetNilExp() IExpContext { return s.nilExp }

func (s *MatchContContext) GetConsExp() IExpContext { return s.consExp }

func (s *MatchContContext) GetNext() IContContext { return s.next }

func (s *MatchContContext) SetNilExp(v IExpContext) { s.nilExp = v }

func (s *MatchContContext) SetConsExp(v IExpContext) { s.consExp = v }

func (s *MatchContContext) SetNext(v IContContext) { s.next = v }

func (s *MatchContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchContContext) NIL_LIST() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserNIL_LIST, 0)
}

func (s *MatchContContext) AllExp() []IExpContext {
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

func (s *MatchContContext) Exp(i int) IExpContext {
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

func (s *MatchContContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(EvalContML4ParserIDENTIFIER)
}

func (s *MatchContContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, i)
}

func (s *MatchContContext) Env() IEnvContext {
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

func (s *MatchContContext) Cont() IContContext {
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

func (s *MatchContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterMatchCont(s)
	}
}

func (s *MatchContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitMatchCont(s)
	}
}

func (s *MatchContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitMatchCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContContext struct {
	ContContext
	op   antlr.Token
	next IContContext
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

func (s *ValueContContext) GetNext() IContContext { return s.next }

func (s *ValueContContext) SetNext(v IContContext) { s.next = v }

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
	return s.GetToken(EvalContML4ParserPLUS, 0)
}

func (s *ValueContContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserMINUS, 0)
}

func (s *ValueContContext) TIMES() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserTIMES, 0)
}

func (s *ValueContContext) LT() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserLT, 0)
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
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterValueCont(s)
	}
}

func (s *ValueContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitValueCont(s)
	}
}

func (s *ValueContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitValueCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetContContext struct {
	ContContext
	var_ antlr.Token
	next IContContext
}

func NewLetContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetContContext {
	var p = new(LetContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *LetContContext) GetVar_() antlr.Token { return s.var_ }

func (s *LetContContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *LetContContext) GetNext() IContContext { return s.next }

func (s *LetContContext) SetNext(v IContContext) { s.next = v }

func (s *LetContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetContContext) Exp() IExpContext {
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

func (s *LetContContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(EvalContML4ParserIDENTIFIER, 0)
}

func (s *LetContContext) Env() IEnvContext {
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

func (s *LetContContext) Cont() IContContext {
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

func (s *LetContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterLetCont(s)
	}
}

func (s *LetContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitLetCont(s)
	}
}

func (s *LetContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitLetCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConsValueContContext struct {
	ContContext
	next IContContext
}

func NewConsValueContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsValueContContext {
	var p = new(ConsValueContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *ConsValueContContext) GetNext() IContContext { return s.next }

func (s *ConsValueContContext) SetNext(v IContContext) { s.next = v }

func (s *ConsValueContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsValueContContext) Value() IValueContext {
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

func (s *ConsValueContContext) Cont() IContContext {
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

func (s *ConsValueContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterConsValueCont(s)
	}
}

func (s *ConsValueContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitConsValueCont(s)
	}
}

func (s *ConsValueContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitConsValueCont(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppExpContContext struct {
	ContContext
	next IContContext
}

func NewAppExpContContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppExpContContext {
	var p = new(AppExpContContext)

	InitEmptyContContext(&p.ContContext)
	p.parser = parser
	p.CopyAll(ctx.(*ContContext))

	return p
}

func (s *AppExpContContext) GetNext() IContContext { return s.next }

func (s *AppExpContContext) SetNext(v IContContext) { s.next = v }

func (s *AppExpContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppExpContContext) Exp() IExpContext {
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

func (s *AppExpContContext) Env() IEnvContext {
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

func (s *AppExpContContext) Cont() IContContext {
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

func (s *AppExpContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.EnterAppExpCont(s)
	}
}

func (s *AppExpContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EvalContML4Listener); ok {
		listenerT.ExitAppExpCont(s)
	}
}

func (s *AppExpContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case EvalContML4Visitor:
		return t.VisitAppExpCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *EvalContML4Parser) Cont() (localctx IContContext) {
	localctx = NewContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EvalContML4ParserRULE_cont)
	var _la int

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTerminalContContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpContContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(155)
				p.Env()
			}

		}
		{
			p.SetState(158)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpContContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpContContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(161)
			p.exp(0)
		}
		{
			p.SetState(162)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(163)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(164)

				var _x = p.Cont()

				localctx.(*ExpContContext).next = _x
			}

		}

	case 3:
		localctx = NewValueContContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.value(0)
		}
		{
			p.SetState(169)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueContContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueContContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(170)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(172)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(173)

				var _x = p.Cont()

				localctx.(*ValueContContext).next = _x
			}

		}

	case 4:
		localctx = NewIfContContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(177)
				p.Env()
			}

		}
		{
			p.SetState(180)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(EvalContML4ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(EvalContML4ParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)

			var _x = p.exp(0)

			localctx.(*IfContContext).then = _x
		}
		{
			p.SetState(185)
			p.Match(EvalContML4ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _x = p.exp(0)

			localctx.(*IfContContext).else_ = _x
		}
		{
			p.SetState(187)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(188)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(189)

				var _x = p.Cont()

				localctx.(*IfContContext).next = _x
			}

		}

	case 5:
		localctx = NewLetContContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(192)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(193)
				p.Env()
			}

		}
		{
			p.SetState(196)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(EvalContML4ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*LetContContext).var_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(EvalContML4ParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(EvalContML4ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.exp(0)
		}
		{
			p.SetState(203)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(204)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(205)

				var _x = p.Cont()

				localctx.(*LetContContext).next = _x
			}

		}

	case 6:
		localctx = NewAppExpContContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(208)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(209)
				p.Env()
			}

		}
		{
			p.SetState(212)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.exp(0)
		}
		{
			p.SetState(215)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(216)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(217)

				var _x = p.Cont()

				localctx.(*AppExpContContext).next = _x
			}

		}

	case 7:
		localctx = NewAppValueContContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(220)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.value(0)
		}
		{
			p.SetState(222)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(224)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(225)

				var _x = p.Cont()

				localctx.(*AppValueContContext).next = _x
			}

		}

	case 8:
		localctx = NewConsExpContContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(228)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(229)
				p.Env()
			}

		}
		{
			p.SetState(232)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(EvalContML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.exp(0)
		}
		{
			p.SetState(236)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(237)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)

				var _x = p.Cont()

				localctx.(*ConsExpContContext).next = _x
			}

		}

	case 9:
		localctx = NewConsValueContContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(241)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.value(0)
		}
		{
			p.SetState(243)
			p.Match(EvalContML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(246)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(247)

				var _x = p.Cont()

				localctx.(*ConsValueContContext).next = _x
			}

		}

	case 10:
		localctx = NewMatchContContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(250)
			p.Match(EvalContML4ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserIDENTIFIER {
			{
				p.SetState(251)
				p.Env()
			}

		}
		{
			p.SetState(254)
			p.Match(EvalContML4ParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(EvalContML4ParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(EvalContML4ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Match(EvalContML4ParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(EvalContML4ParserNIL_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(EvalContML4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.exp(0)

			localctx.(*MatchContContext).nilExp = _x
		}
		{
			p.SetState(261)
			p.Match(EvalContML4ParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*MatchContContext).headVar = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(EvalContML4ParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)

			var _m = p.Match(EvalContML4ParserIDENTIFIER)

			localctx.(*MatchContContext).tailVar = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(EvalContML4ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)

			var _x = p.exp(0)

			localctx.(*MatchContContext).consExp = _x
		}
		{
			p.SetState(267)
			p.Match(EvalContML4ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == EvalContML4ParserT__24 {
			{
				p.SetState(268)
				p.Match(EvalContML4ParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(269)

				var _x = p.Cont()

				localctx.(*MatchContContext).next = _x
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

func (p *EvalContML4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ValueContext = nil
		if localctx != nil {
			t = localctx.(*ValueContext)
		}
		return p.Value_Sempred(t, predIndex)

	case 4:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EvalContML4Parser) Value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *EvalContML4Parser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
