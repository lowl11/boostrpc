package boostrpc

import "github.com/lowl11/lazylog/log"

func (app *App) Run(port string) {
	log.Fatal(app.server.Start(port))
}

func (app *App) RunAsync(port string) {
	go app.Run(port)
}

func (app *App) Close() {
	log.Error(app.server.Close(), "Shutdown gRPC server error")
}
