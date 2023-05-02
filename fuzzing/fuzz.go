package fuzz

import (
	"context"
	"github.com/theyahya/enola"
)

func Fuzz(usr []byte) int{
	ctx := context.Background()
	sh, err := enola.New(ctx)
	if err != nil {
		panic(err)
	}

	_, err = sh.SetSite("").Check(string(usr))
	if err != nil {
		panic(err)
	}
	return 0
}