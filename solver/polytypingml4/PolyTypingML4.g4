grammar PolyTypingML4;

eval: env '|-' exp ':' type EOF;

var: ID;

tVar: '\'' ID;

type:
	tVar												# TypeVar
	| 'bool'											# BoolType
	| 'int'												# IntType
	| elem = type 'list'								# ListType
	| <assoc = right> param = type '->' return = type	# FunType
	| '(' type ')'										# ParenType;

tyScheme: (tVar+ '.')? type;

env: (bind (',' bind)*)?;
bind: var ':' tyScheme;

exp:
	fun														# FunExp
	| fn = exp arg = exp									# AppExp
	| <assoc = right> head = exp '::' tail = exp			# ConsExp
	| left = exp op = TIMES right = exp						# BinOpExp
	| left = exp op = (PLUS | MINUS) right = exp			# BinOpExp
	| left = exp op = LT right = exp						# BinOpExp
	| 'if' cond = exp 'then' then = exp 'else' else = exp	# IfExp
	| 'let' var '=' value = exp 'in' body = exp				# LetExp
	| 'let' 'rec' var '=' fun 'in' body = exp				# LetRecExp
	| 'match' arg = exp 'with' '[]' '->' nilCase = exp '|' head = var '::' tail = var '->' consCase
		= exp		# MatchExp
	| '[]'			# NilExp
	| INT			# IntExp
	| BOOL			# BoolExp
	| var			# VarExp
	| '(' exp ')'	# ParenExp;

fun: 'fun' param = var '->' body = exp;

PLUS: '+';
MINUS: '-';
TIMES: '*';
LT: '<';
BOOL: 'true' | 'false';

// 優先順位低くするために後に定義する
INT: MINUS? [0-9]+;
ID: [a-z_][a-zA-Z0-9_]*;
WS: [ \n\t\r]+ -> channel(HIDDEN);
