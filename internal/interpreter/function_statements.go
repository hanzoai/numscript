package interpreter

import (
	"github.com/formancehq/numscript/internal/parser"
	"github.com/formancehq/numscript/internal/utils"
)

func setTxMeta(st *programState, r parser.Range, args []Value) InterpreterError {
	p := NewArgsParser(args)
	key := p.parseArg(r, expectString)
	meta := p.parseArg(r, expectAnything)
	err := p.parse()
	if err != nil {
		return err
	}

	st.TxMeta[*key] = *meta
	return nil
}

func setAccountMeta(st *programState, r parser.Range, args []Value) InterpreterError {
	p := NewArgsParser(args)
	account := p.parseArg(r, expectAccount)
	key := p.parseArg(r, expectString)
	meta := p.parseArg(r, expectAnything)
	err := p.parse()
	if err != nil {
		return err
	}

	accountMeta := utils.MapGetOrPutDefault(st.SetAccountsMeta, *account, func() AccountMetadata {
		return AccountMetadata{}
	})

	accountMeta[*key] = (*meta).String()

	return nil
}
