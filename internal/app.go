package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	stdhttp "net/http"

	"github.com/alewkinr/example-app-design-review/internal/booking"
	"github.com/alewkinr/example-app-design-review/internal/config"
	"github.com/alewkinr/example-app-design-review/internal/http"
	"github.com/alewkinr/example-app-design-review/internal/orders"
	"github.com/alewkinr/example-app-design-review/pkg/logger"
	"github.com/alewkinr/example-app-design-review/pkg/store/inmemory"
)

// Application — main application container
type Application struct {
	cfg *config.Config
	log *slog.Logger
	srv *stdhttp.Server

	// HTTP API listed below
	ordersAPI http.Router
}

// NewApplication — constructor for application
func NewApplication() (*Application, error) {
	app := &Application{}
	app.cfg = config.MustNewConfig()

	log, err := logger.New(app.cfg.Log.Level)
	if err != nil {
		return nil, fmt.Errorf("creating logger: %w", err)
	}

	bookingMngr := booking.NewManager(log, inmemory.NewBookingRepository())
	ordMngr := orders.NewManager(log, bookingMngr, inmemory.NewOrdersRepository())

	app.ordersAPI = http.NewOrdersAPI(ordMngr)

	app.log = log
	app.srv = &stdhttp.Server{
		Addr:    net.JoinHostPort(app.cfg.Host, app.cfg.Port),
		Handler: http.NewRouter(app.ordersAPI),
	}

	return app, nil
}

// Run — run application
func (a *Application) Run() error {
	a.log.Info("✅ Server is running...", "host", a.cfg.Host, "port", a.cfg.Port)
	if runErr := a.srv.ListenAndServe(); !errors.Is(runErr, stdhttp.ErrServerClosed) {
		return fmt.Errorf("http server start: %w", runErr)
	}

	return nil
}

func (a *Application) Stop(ctx context.Context) {
	a.log.Info("🛑 Server is shutting down...")
	if shutdownErr := a.srv.Shutdown(ctx); shutdownErr != nil {
		a.log.Error("http server shutdown", "error", shutdownErr)
	}
}
