package examples

import "github.com/dhax/htmgo/framework/h"

func InputComponentExample(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &InputComponentSnippet)
	return Index(ctx)
}
