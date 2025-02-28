package ws

import (
	"github.com/dhax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/dhax/htmgo/framework/h"
)

func ManagerFromCtx(ctx *h.RequestContext) *wsutil.SocketManager {
	return wsutil.SocketManagerFromCtx(ctx)
}
