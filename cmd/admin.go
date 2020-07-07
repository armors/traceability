package main

import (
	log "github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/sword"
	"traceability/internal/admin"
)

const (
	version   = "1.0.0"
)
func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfgPath := rootPath + "/config.json"

	admin := admin.New(cfgPath)
	go func() {
		_ = admin.Gin.Run(":9033")
	}()

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			admin.Close()
			log.Infof("admin [version: %s] exit", version)
			log.Flush()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}