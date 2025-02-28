package main

import (
	"fmt"
	"github.com/dhax/htmgo/framework/h"
	"github.com/dhax/htmgo/framework/service"
	"hackernews/__htmgo"
	"io/fs"
	"net/http"
)

func main() {
	locator := service.NewLocator()

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(GetStaticAssets(), "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/item", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				id := r.URL.Query().Get("id")
				w.Header().Set("Location", fmt.Sprintf("/?item=%s", id))
				w.WriteHeader(302)
			}))
			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			__htmgo.Register(app.Router)
		},
	})
}
