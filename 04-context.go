// See https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("Doing something!: myKey=[%s]\n", ctx.Value("myKey"))
}

func Ctxt() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}
