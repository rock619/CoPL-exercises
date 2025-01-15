grammar EvalML2;

question: eval EOF;

eval: defList? '|-' expr 'evalto' value;

expr:
	BOOL															# BoolExpr
	| INT															# IntExpr
	| MINUS INT														# NegIntExpr
	| VARNAME														# VarExpr
	| '(' expr ')'													# ParenExpr
	| left = expr op = TIMES right = expr							# BinOpExpr
	| left = expr op = (PLUS | MINUS) right = expr					# BinOpExpr
	| left = expr op = LT right = expr								# BinOpExpr
	| 'if' cond = expr 'then' then = expr 'else' else = expr		# IfExpr
	| 'let' bindName = VARNAME '=' bindExpr = expr 'in' body = expr	# LetExpr;

defList: def (COMMA def)*;

def: VARNAME '=' value;

value: INT # IntValue | BOOL # BoolValue;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
BOOL: TRUE | FALSE;
TRUE: 'true';
FALSE: 'false';
COMMA: ',';
TURNSTILE: '|-';

INT: [0-9]+;
VARNAME: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
