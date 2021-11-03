package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JEDIAC/server/conf"
	"github.com/JEDIAC/server/internal/db"
	"github.com/JEDIAC/server/internal/log"
	"github.com/JEDIAC/server/internal/mw"
	"github.com/JEDIAC/server/internal/service"
	"github.com/JEDIAC/server/router"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var engine *gin.Engine

func main() {
	rootCmd := &cobra.Command{
		Use:     "server",
		Short:   "JediAutocare backend server",
		PreRunE: Initialize,
		Run:     Run,
	}
	conf.InitFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("root cmd execute error. %v", err)
	}
}

// Initialize ...
func Initialize(cmd *cobra.Command, _ []string) error {
	if err := conf.Init(cmd); err != nil {
		return err
	}
	if err := log.Init(conf.LogFile, conf.LogLevel); err != nil {
		return err
	}
	if err := db.Init(conf.DBConnStr, conf.DBName); err != nil {
		return err
	}
	if err := service.Init(); err != nil {
		return err
	}
	if err := service.MigrateModel(); err != nil {
		return err
	}
	initEngine()
	log.Info("app initialize complete")
	return nil
}

// Run ...
func Run(_ *cobra.Command, _ []string) {
	server := http.Server{
		Addr:    conf.Addr,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("fail to start http server. %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	log.Info("closing server in 3 seconds...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	<-ctx.Done()
	server.Close()
	log.Info("server shutdown")
}

func initEngine() {
	engine = gin.New()
	engine.Use(ginzap.Ginzap(log.L(), time.RFC3339, false))
	engine.Use(ginzap.RecoveryWithZap(log.L(), true))
	engine.Use(mw.Cors())
	// init routers
	r := engine.Group("/api")
	{
		router.InitStaffRouter(r)
		router.InitCustomerRouter(r)
		router.InitAppRouter(r)
	}
}
