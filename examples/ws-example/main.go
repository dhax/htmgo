package main

import (
	"github.com/dhax/htmgo/extensions/websocket"
	ws2 "github.com/dhax/htmgo/extensions/websocket/opts"
	"github.com/dhax/htmgo/extensions/websocket/session"
	"github.com/dhax/htmgo/framework/h"
	"github.com/dhax/htmgo/framework/service"
	"io/fs"
	"net/http"
	"ws-example/__htmgo"
)

func main() {
	locator := service.NewLocator()

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {

			app.Use(func(ctx *h.RequestContext) {
				session.CreateSession(ctx)
			})

			websocket.EnableExtension(app, ws2.ExtensionOpts{
				WsPath: "/ws",
				RoomName: func(ctx *h.RequestContext) string {
					return "all"
				},
				SessionId: func(ctx *h.RequestContext) string {
					return ctx.QueryParam("sessionId")
				},
			})

			sub, err := fs.Sub(GetStaticAssets(), "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			__htmgo.Register(app.Router)
		},
	})
}
