package cli

import (
	"errors"

	"github.com/just-do-halee/gowcoin/tree/main/gow-node/explorer"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/rest"
	"github.com/just-do-halee/opt"
)

type root struct {
	Port opt.Option[int]    `msg:"Set port of the server" opt:"l,s"`
	Mode opt.Option[string] `msg:"Choose between 'html' and 'rest'" opt:"l,s"`

	Help        opt.Option[opt.Help]  `msg:"Show help"`
	HelpCommand opt.Command[opt.Help] `msg:"Show help" rename:"help"`
}

func (r *root) Before() error {
	opt.Set(&r.Port, 4000)
	opt.Set(&r.Mode, "rest")
	return nil
}

func (r *root) After() error {
	var err error

	err = opt.Validate(&r.Port, opt.IsMinMax(1024, 65535))
	if err != nil {
		return err
	}

	return opt.Validate(&r.Mode, func(s string) error {
		port := r.Port.Get()
		switch s {
		case "html":
			explorer.Start(port)
		case "rest":
			rest.Start(port)
		default:
			return errors.New("Mode is not supported")
		}
		return nil
	})
}

func (r *root) Run() error {
	return errors.New(r.Help.GetUsage())
}
