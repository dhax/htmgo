package ws

import (
	"github.com/dhax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/dhax/htmgo/framework/h"
)

type Metrics struct {
	Manager wsutil.ManagerMetrics
	Handler HandlerMetrics
}

func MetricsFromCtx(ctx *h.RequestContext) Metrics {
	manager := ManagerFromCtx(ctx)
	return Metrics{
		Manager: manager.Metrics(),
		Handler: GetHandlerMetics(),
	}
}
