package cmn

import "log"

func HandleError(err error, fn ...func(err error) bool) {
	if err != nil {
		for _, f := range fn {
			if f(err) {
				return
			}
		}
		log.Panic(err)
	}
}
