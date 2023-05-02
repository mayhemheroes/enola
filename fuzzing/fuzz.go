package fuzz

import (
	"context"
	"github.com/theyahya/enola"
)

type responseMsg enola.Result

func Fuzz(usr []byte) int{
	ctx := context.Background()
	sh, err := enola.New(ctx)

	if err != nil {
		panic(err)
		return 1
	}

	_, err = sh.SetSite("Crevado").Check(string(usr))

	if err != nil {
		panic(err)
		return 1
	}
	return 0
}