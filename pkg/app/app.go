package app

import (
	"context"
	"net"
	"time"

	"apisrv/pkg/newsportal"

	"apisrv/pkg/db"
	"apisrv/pkg/embedlog"
	"apisrv/pkg/vt"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/rpcgen/v2/typescript"
	"github.com/vmkteam/vfs"
	"github.com/vmkteam/zenrpc/v2"
)

type Config struct {
	Database *pg.Options
	Server   struct {
		Host      string
		Port      int
		IsDevel   bool
		EnableVFS bool
	}
	Sentry struct {
		Environment string
		DSN         string
	}
	VFS vfs.Config
}

type App struct {
	embedlog.Logger
	appName string
	cfg     Config
	dbo     db.DB
	dbc     *pg.DB
	nr      db.NewsRepo
	echo    *echo.Echo
	nm      *newsportal.Manager
	vtsrv   zenrpc.Server
}

func New(appName string, verbose bool, cfg Config, dbo db.DB, dbc *pg.DB) *App {
	a := &App{
		appName: appName,
		cfg:     cfg,
		dbo:     dbo,
		dbc:     dbc,
		echo:    echo.New(),
	}
	a.nr = db.NewNewsRepo(a.dbc)
	a.nm = newsportal.NewManager(a.nr)
	a.SetStdLoggers(verbose)
	a.echo.HideBanner = true
	a.echo.HidePort = true
	_, mask, _ := net.ParseCIDR("0.0.0.0/0")
	a.echo.IPExtractor = echo.ExtractIPFromRealIPHeader(echo.TrustIPRange(mask))
	a.nm = newsportal.NewManager(a.nr)
	a.vtsrv = vt.New(a.dbo, a.Logger, a.cfg.Server.IsDevel, a.nm)
	return a
}

// Run is a function that runs application.
func (a *App) Run() error {
	a.registerMetrics()
	a.registerHandlers()
	a.registerDebugHandlers()
	a.registerAPIHandlers()
	a.registerVTApiHandlers()
	return a.runHTTPServer(a.cfg.Server.Host, a.cfg.Server.Port)
}

// VTTypeScriptClient returns TypeScript client for VT.
func (a *App) VTTypeScriptClient() ([]byte, error) {
	gen := rpcgen.FromSMD(a.vtsrv.SMD())
	tsSettings := typescript.Settings{ExcludedNamespace: []string{NSVFS}, WithClasses: true}
	return gen.TSCustomClient(tsSettings).Generate()
}

// Shutdown is a function that gracefully stops HTTP server.
func (a *App) Shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := a.echo.Shutdown(ctx); err != nil {
		a.Errorf("shutting down server err=%q", err)
	}
}
