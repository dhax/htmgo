package examples

import (
	"github.com/dhax/htmgo/framework/h"
)

func FormWithLoadingState(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &FormWithLoadingStateSnippet)
	return Index(ctx)
}
