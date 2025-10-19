package main

import (
	"context"
	"fmt"
)

const (
	dataKey     = "data"
	userDataKey = "data"
)

type key struct {
}

type userKey struct {
}

func keyConflict() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, dataKey, "some data")

	fmt.Println("data: %s\n", ctx.Value(dataKey))

	ctx = context.WithValue(ctx, userDataKey, "user data")

	fmt.Println("user data: %s\n", ctx.Value(userDataKey))
	fmt.Println("data: %s\n", ctx.Value(dataKey))
}

func keyConflictFix() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, key{}, "some data")

	fmt.Println("data: %s\n", ctx.Value(key{}))

	ctx = context.WithValue(ctx, userKey{}, "user data")

	fmt.Println("user data: %s\n", ctx.Value(key{}))
	fmt.Println("data: %s\n", ctx.Value(userKey{}))
}
