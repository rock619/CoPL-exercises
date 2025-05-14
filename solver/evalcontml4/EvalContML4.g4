grammar EvalContML4;

eval: env? '|-' exp cont? 'evalto' value EOF;

value:
	INT													# IntValue
	| BOOL												# BoolValue
	| '(' env? ')' '[' fun ']'							# FunValue
	| '(' env? ')' '[' recFun ']'						# RecFunValue
	| NIL_LIST											# NilValue
	| <assoc = right> head = value '::' tail = value	# ConsValue
	| '[' cont ']'										# ContValue;

env: bind (',' bind)*;

bind: IDENTIFIER '=' value;

exp:
	'(' exp ')'				# ParenExp
	| fun					# FunExp
	| fn = exp arg = exp	# AppExp
	| 'match' matchedExp = exp 'with' NIL_LIST '->' nilExp = exp '|' headVar = IDENTIFIER '::'
		tailVar = IDENTIFIER '->' consExp = exp					# MatchExp
	| NIL_LIST													# NilExp
	| <assoc = right> head = exp '::' tail = exp				# ConsExp
	| left = exp op = TIMES right = exp							# BinOpExp
	| left = exp op = (PLUS | MINUS) right = exp				# BinOpExp
	| left = exp op = LT right = exp							# BinOpExp
	| 'if' cond = exp 'then' then = exp 'else' else = exp		# IfExp
	| 'let' var = IDENTIFIER '=' bindExp = exp 'in' body = exp	# LetExp
	| 'let' recFun 'in' body = exp								# LetRecExp
	| 'letcc' var = IDENTIFIER 'in' body = exp					# LetCCExp
	| INT														# IntExp
	| BOOL														# BoolExp
	| IDENTIFIER												# VarExp;

fun: 'fun' param = IDENTIFIER '->' body = exp;
recFun: 'rec' funName = IDENTIFIER '=' fun;

cont:
	'_' # TerminalCont
	| '{' env? '|-' '_' op = (PLUS | MINUS | TIMES | LT) exp '}' (
		'>>' next = cont
	)? # ExpCont
	| '{' value op = (PLUS | MINUS | TIMES | LT) '_' '}' (
		'>>' next = cont
	)? # ValueCont
	| '{' env? '|-' 'if' '_' 'then' then = exp 'else' else = exp '}' (
		'>>' next = cont
	)? # IfCont
	| '{' env? '|-' 'let' var = IDENTIFIER '=' '_' 'in' exp '}' (
		'>>' next = cont
	)?														# LetCont
	| '{' env? '|-' '_' exp '}' ('>>' next = cont)?			# AppExpCont
	| '{' value '_' '}' ('>>' next = cont)?					# AppValueCont
	| '{' env? '|-' '_' '::' exp '}' ('>>' next = cont)?	# ConsExpCont
	| '{' value '::' '_' '}' ('>>' next = cont)?			# ConsValueCont
	| '{' env? '|-' 'match' '_' 'with' NIL_LIST '->' nilExp = exp '|' headVar = IDENTIFIER '::'
		tailVar = IDENTIFIER '->' consExp = exp '}' (
		'>>' next = cont
	)? # MatchCont;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
BOOL: TRUE | FALSE;
TRUE: 'true';
FALSE: 'false';
NIL_LIST: '[]';

// 優先順位低くするために後に定義する
INT: MINUS? [0-9]+;
IDENTIFIER: [a-z_][a-zA-Z0-9_']*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
