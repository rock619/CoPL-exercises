grammar EvalContML1;

eval: exp cont? 'evalto' value EOF;

value: INT # IntValue | BOOL # BoolValue;

exp:
	INT														# IntExp
	| BOOL													# BoolExp
	| left = exp op = TIMES right = exp						# BinOpExp
	| left = exp op = (PLUS | MINUS) right = exp			# BinOpExp
	| left = exp op = LT right = exp						# BinOpExp
	| 'if' cond = exp 'then' then = exp 'else' else = exp	# IfExp
	| '(' exp ')'											# ParenExp;

cont:
	'>>' '_'															# UnaryCont
	| '>>' '{' '_' op = (PLUS | MINUS | TIMES | LT) exp '}' cont?		# ExpCont
	| '>>' '{' value op = (PLUS | MINUS | TIMES | LT) '_' '}' cont?		# ValueCont
	| '>>' '{' 'if' '_' 'then' then = exp 'else' else = exp '}' cont?	# IfCont;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
BOOL: TRUE | FALSE;
TRUE: 'true';
FALSE: 'false';

// 優先順位低くするために後に定義する
INT: MINUS? [0-9]+;
WS: [ \n\t\r]+ -> channel(HIDDEN);
