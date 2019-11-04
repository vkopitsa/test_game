package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	gameServer "server/pkg/server"
	"syscall"
)

// var countCluters = 5
// var countWorldsInCluster = 100

var (
	listenAddress string
	// listenAddressSocket string
	// listenAddressSSH    string
	// netrisBinary        string
	// debugAddress        string

	logDebug   bool
	logVerbose bool

	done = make(chan bool)
)

func init() {
	log.SetFlags(0)

	flag.StringVar(&listenAddress, "listen-tcp", ":8080", "host server on network address")
	// flag.StringVar(&listenAddressSocket, "listen-socket", "", "host server on socket path")
	// flag.StringVar(&listenAddressSSH, "listen-ssh", "", "host SSH server on network address")
	// flag.StringVar(&netrisBinary, "netris", "", "path to netris client")
	// flag.StringVar(&debugAddress, "debug-address", "", "address to serve debug info")
	flag.BoolVar(&logDebug, "debug", false, "enable debug logging")
	flag.BoolVar(&logVerbose, "verbose", false, "enable verbose logging")
}

func main() {
	flag.Parse()

	logLevel := gameServer.LogStandard
	if logVerbose {
		logLevel = gameServer.LogVerbose
	} else if logDebug {
		logLevel = gameServer.LogDebug
	}

	server := gameServer.New(logLevel)

	go server.Listen(listenAddress)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc

		done <- true
	}()

	<-done

	server.StopListening()

	// var frame_rate int64 = 30

	// var clusters []cluster.Cluster = make([]cluster.Cluster, countCluters, countCluters)

	// 	go func() {
	// 		// for i := 0; i < len(clusters); i++ {
	// 		// 	clusters[i] = cluster.New()
	// 		// }
	// 		// for {
	// 		// 	fmt.Println("Update world")

	// 		// 	time.Sleep(time.Millisecond / time.Millisecond * time.Duration(frame_rate))
	// 		// }
	// 	}()

	// 	// Echo instance
	// 	e := echo.New()

	// 	// Middleware
	// 	e.Use(middleware.Logger())
	// 	e.Use(middleware.Recover())

	// 	// Routes
	// 	e.GET("/", hello)

	// 	// Start server
	// 	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
