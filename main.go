package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stealthpass/stealthpass/cmd"
	"github.com/stealthpass/stealthpass/internal/app"
	"github.com/stealthpass/stealthpass/internal/localproxy"
	"github.com/stealthpass/stealthpass/internal/manage"
	"github.com/stealthpass/stealthpass/internal/menu"
	"github.com/stealthpass/stealthpass/internal/monitor"
	"github.com/stealthpass/stealthpass/internal/telegram"
	"github.com/stealthpass/stealthpass/internal/utils"
	"github.com/stealthpass/stealthpass/internal/webui"
)

var logger = utils.NewLogger("info")

// main has two modes:
//
//   - Engine mode:  `stealthpass -c /etc/stealthpass/<name>.toml`
//     Runs a single tunnel (server or client). This is what the systemd
//     units execute. Behaviour is identical to the original engine.
//
//   - Menu mode:    `stealthpass`  (no arguments)
//     Opens the interactive management CLI on the VPS.
func main() {
	// Handled before the flags, because it is a subcommand with flags of its
	// own: `stealthpass node setup --panel ... --key ...`. The flag package would
	// stop at "node" and report the rest as unknown.
	if len(os.Args) > 1 && os.Args[1] == "node" {
		runNode(os.Args[2:])
		return
	}

	configPath := flag.String("c", "", "path to a tunnel configuration file (TOML) — runs in engine mode")
	showVersion := flag.Bool("v", false, "print the version and exit")
	restartAll := flag.Bool("restart-all", false, "restart every configured tunnel and exit (used by the auto-refresh job)")
	tgReport := flag.Bool("telegram-report", false, "send a Telegram status report and exit (used by the scheduled job)")
	webPanel := flag.Bool("webui", false, "run the web panel (used by the stealthpass-webui service)")
	monitorMode := flag.Bool("monitor", false, "run the watchdog, Telegram bot and alerts (used by the stealthpass-monitor service)")
	proxyMode := flag.Bool("proxy", false, "run the built-in SOCKS5/HTTP proxy (used by the stealthpass-proxy service)")
	flag.Parse()

	switch {
	case *showVersion:
		fmt.Println(app.Version)
		return
	case *restartAll:
		ok, failed := manage.RestartAll()
		fmt.Printf("restarted %d tunnels, %d failed\n", ok, failed)
		return
	case *tgReport:
		if err := telegram.SendStatusNow(); err != nil {
			logger.Errorf("telegram report failed: %v", err)
			os.Exit(1)
		}
		return
	case *webPanel:
		if err := webui.Serve(); err != nil {
			logger.Fatalf("web panel failed: %v", err)
		}
		return
	case *monitorMode:
		monitor.Run()
		return
	case *proxyMode:
		runProxy()
		return
	}

	// No config file -> interactive menu.
	if *configPath == "" {
		menu.Run()
		return
	}

	runEngine(*configPath)
}

// runProxy runs the built-in proxy until a termination signal arrives. The
// proxy is a plain loopback service; the tunnel forwards to it like any backend.
func runProxy() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go localproxy.Run(ctx)
	<-sigChan
	cancel()
	logger.Info("stealthpass proxy stopped")
}

// runEngine starts a single tunnel from a TOML config and blocks until a
// termination signal arrives, then shuts down gracefully.
func runEngine(configPath string) {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go cmd.Run(configPath, ctx)

	<-sigChan
	cancel()
	time.Sleep(1 * time.Second)
	logger.Info("stealthpass engine stopped")
}
