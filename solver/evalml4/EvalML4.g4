grammar EvalML4;

question: eval EVALTO value EOF;

eval: defList? TURNSTILE expr;

expr:
	LPAREN expr RPAREN											# ParenExpr
	| fun														# FunExpr
	| fn = expr arg = expr										# AppExpr
	| MATCH matchExpr = expr WITH emptyPattern OR consPattern	# MatchExpr
	| EMPTYLIST													# EmptyListExpr
	| <assoc = right> head = expr CONS tail = expr				# ConsExpr
	| left = expr op = TIMES right = expr						# BinOpExpr
	| left = expr op = (PLUS | MINUS) right = expr				# BinOpExpr
	| left = expr op = LT right = expr							# BinOpExpr
	| IF cond = expr THEN then = expr ELSE else = expr			# IfExpr
	| LET bindName = VARNAME EQ bindExpr = expr IN body = expr	# LetExpr
	| LET recFun IN expr										# LetRecExpr
	| INT														# IntExpr
	| BOOL														# BoolExpr
	| VARNAME													# VarExpr;

defList: def (COMMA def)*;

def: VARNAME EQ value;

emptyPattern: EMPTYLIST ARROW expr;
consPattern:
	headVar = VARNAME CONS tailVar = VARNAME ARROW expr;

fun: FUN param = VARNAME ARROW body = expr;
recFun: REC funName = VARNAME EQ fun;

value:
	INT													# IntValue
	| BOOL												# BoolValue
	| LPAREN defList? RPAREN LBRACKET fun RBRACKET		# FunValue
	| LPAREN defList? RPAREN LBRACKET recFun RBRACKET	# RecFunValue
	| EMPTYLIST											# EmptyListValue
	| <assoc = right> head = value CONS tail = value	# ConsValue;

OR: '|';
EMPTYLIST: LBRACKET RBRACKET;
CONS: '::';
MATCH: 'match';
WITH: 'with';
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
INT: MINUS? [0-9]+;
VARNAME: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
