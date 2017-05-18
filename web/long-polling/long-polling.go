package main

import (
	"container/list"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	address = flag.String("address", ":3001", "Server port")
)

func ChannelsKeeper(c chan chan string) {
	channels := list.New()

	go func() {
		for {
			select {
			case nc := <-c:
				// Inserts a new client at the back of list clients and returns client
				channels.PushBack(nc)
				fmt.Printf("New client: %d\n", channels.Len())
			}
		}
	}()
}

func InitSignalHandlers(s chan os.Signal) {
	go func() {
		sig := <-s
		switch sig {
		case syscall.SIGINT:
			fmt.Printf("\nCtrl-C signalled\n")
			os.Exit(0)
		}
	}()
}

func LonPollingHandler(c chan chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set Headers
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// We are going to return json:
		w.Header().Set("Content-Type", "application/json")

		// Don't cache response:
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")

		//CloseNotify returns a channel that receives at most a
		// single value (true) when the client connection has gone
		// away.
		cn, _ := w.(http.CloseNotifier)

		// Create channel messages
		message := make(chan string, 1)

		// Create [for] and execute code for obtain the data of what your want return it

		// Notify clients new client
		c <- message

		select {
		case <-time.After(60e9):
			// Request Timeout
			fmt.Println("Timeout")
			io.WriteString(w, "Timeout!\n")
		case msg := <-message:
			// Receive o get message
			io.WriteString(w, msg)
		case <- cn.CloseNotify():
			// Client closed connection
			return

		}
	}
}

func CreateHttpServer(clients chan chan string) {
	http.HandleFunc("/", LonPollingHandler(clients))
	log.Println("ListenOn " + *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal("Listen And Serve: ", err.Error())
	}
}

func main() {
	// Set proc
	runtime.GOMAXPROCS(runtime.NumCPU() * 2 + 1)
	flag.Parse()

	// Create channels
	clients := make(chan chan string, 1)
	signals := make(chan os.Signal, 1)

	// Registers the given channel to receive notifications of the specified signals.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGUSR1)

	ChannelsKeeper(clients)
	InitSignalHandlers(signals)
	CreateHttpServer(clients)
}