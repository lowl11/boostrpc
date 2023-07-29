package boostrpc

import "github.com/lowl11/lazylog/log"

func (app *App) Run(port string) error {
	return app.server.Start(port)
}

func (app *App) Close() {
	log.Error(app.server.Close(), "Shutdown gRPC server error")
}
