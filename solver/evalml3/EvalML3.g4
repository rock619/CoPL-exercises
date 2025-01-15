grammar EvalML3;

question: eval EVALTO value EOF;

eval: defList? TURNSTILE expr;

expr:
	LPAREN expr RPAREN											# ParenExpr
	| fun														# FunExpr
	| expr expr													# AppExpr
	| left = expr op = TIMES right = expr						# BinOpExpr
	| left = expr op = (PLUS | MINUS) right = expr				# BinOpExpr
	| left = expr op = LT right = expr							# BinOpExpr
	| IF cond = expr THEN then = expr ELSE else = expr			# IfExpr
	| LET bindName = VARNAME EQ bindExpr = expr IN body = expr	# LetExpr
	| LET recFun IN expr										# LetRecExpr
	| BOOL														# BoolExpr
	| INT														# IntExpr
	| NEGINT													# NegIntExpr
	| VARNAME													# VarExpr;

defList: def (COMMA def)*;

def: VARNAME EQ value;

fun: FUN param = VARNAME ARROW body = expr;
recFun: REC funName = VARNAME EQ fun;

value:
	INT													# IntValue
	| BOOL												# BoolValue
	| LPAREN defList? RPAREN LBRACKET fun RBRACKET		# FunValue
	| LPAREN defList? RPAREN LBRACKET recFun RBRACKET	# RecFunValue;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
IF: 'if';
THEN: 'then';
ELSE: 'else';
BOOL: TRUE | FALSE;
TRUE: 'true';
FALSE: 'false';
COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LET: 'let';
IN: 'in';
FUN: 'fun';
EQ: '=';
REC: 'rec';
ARROW: '->';
LBRACKET: '[';
RBRACKET: ']';
TURNSTILE: '|-';
EVALTO: 'evalto';
// 優先順位低くするために後に定義する
NEGINT: '-[0-9]+';
INT: [0-9]+;
VARNAME: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
