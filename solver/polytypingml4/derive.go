package main

import (
	"fmt"
)

type Judgement struct {
	Env      Env
	Exp      Exp
	Type     Type
	Sub      Substitution
	By       Rule
	Premises []Judgement
}

func (j Judgement) String() string {
	env := j.Env.String()
	if env != "" {
		env += " "
	}
	return fmt.Sprintf("%s|- %s : %s by %s", env, j.Exp, j.Type, j.By)
}

type Rule string

const (
	TInt    Rule = "T-Int"
	TBool   Rule = "T-Bool"
	TIf     Rule = "T-If"
	TPlus   Rule = "T-Plus"
	TMinus  Rule = "T-Minus"
	TMult   Rule = "T-Mult"
	TLT     Rule = "T-Lt"
	TVar    Rule = "T-Var"
	TLet    Rule = "T-Let"
	TAbs    Rule = "T-Abs"
	TApp    Rule = "T-App"
	TLetRec Rule = "T-LetRec"
	TNil    Rule = "T-Nil"
	TCons   Rule = "T-Cons"
	TMatch  Rule = "T-Match"
)
