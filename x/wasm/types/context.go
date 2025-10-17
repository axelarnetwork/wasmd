package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: remove during next coordinated upgrade
const contextKeyCallDepthBuggy contextKey = iota

// TODO: remove during next coordinated upgrade
func WithCallDepthBuggy(ctx sdk.Context, counter uint32) sdk.Context {
	return ctx.WithValue(contextKeyCallDepthBuggy, counter)
}

// TODO: remove during next coordinated upgrade
func CallDepthBuggy(ctx sdk.Context) (uint32, bool) {
	val, ok := ctx.Value(contextKeyCallDepthBuggy).(uint32)
	return val, ok
}

func WithCallDepth(ctx sdk.Context, counter uint32) sdk.Context {
	return ctx.WithValue(contextKeyCallDepth, counter)
}

func CallDepth(ctx sdk.Context) (uint32, bool) {
	val, ok := ctx.Value(contextKeyCallDepth).(uint32)
	return val, ok
}
