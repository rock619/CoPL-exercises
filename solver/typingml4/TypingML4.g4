grammar TypingML4;

judgement: env? TURNSTILE expr COLON type EOF;

type:
	LPAREN type RPAREN											# ParenType
	| BOOLTYPE													# BoolType
	| INTTYPE													# IntType
	| elementType = type LISTTYPE								# ListType
	| <assoc = right> paramType = type ARROW returnType = type	# FunType;

expr:
	LPAREN expr RPAREN									# ParenExpr
	| fun												# FunExpr
	| fn = expr arg = expr								# AppExpr
	| MATCH expr WITH emptyPattern OR consPattern		# MatchExpr
	| EMPTYLIST											# EmptyListExpr
	| <assoc = right> head = expr CONS tail = expr		# ConsExpr
	| left = expr op = TIMES right = expr				# BinOpExpr
	| left = expr op = (PLUS | MINUS) right = expr		# BinOpExpr
	| left = expr op = LT right = expr					# BinOpExpr
	| IF cond = expr THEN then = expr ELSE else = expr	# IfExpr
	| LET IDENTIFIER EQ bindExpr = expr IN body = expr	# LetExpr
	| LET recFun IN body = expr							# LetRecExpr
	| INT												# IntExpr
	| BOOL												# BoolExpr
	| IDENTIFIER										# VarExpr;

env: bind (COMMA bind)*;

bind: IDENTIFIER COLON type;

emptyPattern: EMPTYLIST ARROW expr;
consPattern:
	headVar = IDENTIFIER CONS tailVar = IDENTIFIER ARROW expr;

fun: FUN param = IDENTIFIER ARROW body = expr;
recFun: REC funName = IDENTIFIER EQ fun;

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
COLON: ':';
BOOLTYPE: 'bool';
INTTYPE: 'int';
LISTTYPE: 'list';

// 優先順位低くするために後に定義する
INT: MINUS? [0-9]+;
IDENTIFIER: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
