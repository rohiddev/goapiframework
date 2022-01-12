package main

import (
	"customerapp/services/customer"
	"database/sql"
	"flag"
	"fmt"
	"github.com/shijuvar/go-distsys/gokitdemo/pkg/oc"
	"net/http"
	"os"
	"os/signal"
	"syscall"


	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"

	_ "customerapp/services/customer"
	"customerapp/services/customer/cockroachdb"
	customersvc "customerapp/services/customer/implementation"
	"customerapp/services/customer/middleware"
	"customerapp/services/customer/transport"
	httptransport "customerapp/services/customer/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":9002", "HTTP listen address")
	)
	flag.Parse()
	// initialize our OpenCensus configuration and defer a clean-up
	//defer oc.Setup("customer").Close()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "customer",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		// Connect to the "ordersdb" database
		db, err = sql.Open("postgres",
			"postgresql://127.0.0.1:5432?sslmode=disable")

		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	// Create Order Service
	var svc customer.Service
	{
		repository, err := cockroachdb.New(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = customersvc.NewService(repository, logger)
		// Add service middleware here
		// Logging middleware
		svc = middleware.LoggingMiddleware(logger)(svc)
	}
	// Create Go kit endpoints for the Order Service
	// Then decorates with endpoint middlewares
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(svc)
		// Add endpoint level middlewares here
		// Trace server side endpoints with open census
		endpoints = transport.Endpoints{
			CreateCustomer:       oc.ServerEndpoint("CreateCustomer")(endpoints.CreateCustomer),

		}

	}
	var h http.Handler
	{
		//ocTracing := kitoc.HTTPServerTrace()
		//serverOptions := []kithttp.ServerOption{ocTracing}
		//h = httptransport.NewService(endpoints, serverOptions, logger)

		h = httptransport.NewService(endpoints, nil,  logger)


	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()
	level.Error(logger).Log("exit", <-errs)

}

