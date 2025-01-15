grammar EvalML1;

question: eval EOF;

eval: expr 'evalto' value;

expr:
	BOOL														# BoolExpr
	| INT														# IntExpr
	| MINUS INT													# NegIntExpr
	| '(' expr ')'												# ParenExpr
	| left = expr op = TIMES right = expr						# BinOpExpr
	| left = expr op = (PLUS | MINUS) right = expr				# BinOpExpr
	| left = expr op = LT right = expr							# BinOpExpr
	| 'if' cond = expr 'then' then = expr 'else' else = expr	# IfExpr;

value: INT # IntValue | BOOL # BoolValue;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
INT: [0-9]+;
BOOL: TRUE | FALSE;
TRUE: 'true';
FALSE: 'false';
WS: [ \n\t\r]+ -> channel(HIDDEN);
