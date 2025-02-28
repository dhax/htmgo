package examples

import (
	"github.com/dhax/htmgo/framework/h"
)

func JsSetTextOnClickPage(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &JsSetTextOnClick)
	return Index(ctx)
}
