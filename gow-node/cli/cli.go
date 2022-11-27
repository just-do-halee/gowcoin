package cli

import (
	"log"

	"github.com/just-do-halee/opt"
)

func Execute() {
	err := opt.Args().
		Version("v0.1.0").
		Author("just-do-halee").
		About("A simple example").
		Build(new(root))

	if err != nil {
		log.Print(err)
	}
}
