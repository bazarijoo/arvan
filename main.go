package main

import (
	"arvan/wallet/Entity"
	"arvan/wallet/Reository"
	newHTTP "arvan/wallet/pkg/http"
	"arvan/wallet/pkg/service"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	var db *gorm.DB
	{
		dbDriver := "sqlite3"
		dbName := "demo.db"

		dbLoaded, err := gorm.Open(dbDriver, dbName) // Rename the variable
		if err != nil {
			_ = level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		dbLoaded.AutoMigrate(&Entity.UserEntity{})

		db = dbLoaded // Set `db` of the outer scope to `dbLoaded` of this scope
	} // dbLoaded is lost here, but it can be accessed using `db`

	flag.Parse()
	ctx := context.Background()
	var srv service.WalletService
	{
		repository := Reository.NewRepository(db, logger)

		srv = service.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := newHTTP.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := newHTTP.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)

}
