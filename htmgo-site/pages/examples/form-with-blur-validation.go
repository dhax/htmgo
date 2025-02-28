package examples

import (
	"github.com/dhax/htmgo/framework/h"
)

func FormWithBlurValidation(ctx *h.RequestContext) *h.Page {
	SetSnippet(ctx, &FormWithBlurValidationSnippet)
	return Index(ctx)
}
