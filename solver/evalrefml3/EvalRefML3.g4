grammar EvalRefML3;

eval: (storeIn = store '/')? env? '|-' exp 'evalto' value (
		'/' storeOut = store
	)? EOF;

value:
	INT								# IntValue
	| BOOL							# BoolValue
	| LOC							# LocValue
	| '(' env? ')' '[' fun ']'		# FunValue
	| '(' env? ')' '[' recFun ']'	# RecFunValue;

env: bind (',' bind)*;
bind: IDENTIFIER '=' value;

store: assign (',' assign)*;
assign: LOC '=' value;

exp:
	'(' exp ')'													# ParenExp
	| 'ref' exp													# RefExp
	| '!' exp													# DerefExp
	| fun														# FunExp
	| fn = exp arg = exp										# AppExp
	| left = exp op = TIMES right = exp							# BinOpExp
	| left = exp op = (PLUS | MINUS) right = exp				# BinOpExp
	| left = exp op = LT right = exp							# BinOpExp
	| <assoc = right> left = exp ':=' right = exp				# AssignExp
	| 'if' cond = exp 'then' then = exp 'else' else = exp		# IfExp
	| 'let' var = IDENTIFIER '=' bindExp = exp 'in' body = exp	# LetExp
	| 'let' recFun 'in' body = exp								# LetRecExp
	| INT														# IntExp
	| BOOL														# BoolExp
	| IDENTIFIER												# VarExp;

fun: 'fun' param = IDENTIFIER '->' body = exp;
recFun: 'rec' funName = IDENTIFIER '=' fun;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
BOOL: 'true' | 'false';

// 優先順位低くするために後に定義する
LOC: '@' IDENTIFIER;
INT: MINUS? [0-9]+;
IDENTIFIER: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
