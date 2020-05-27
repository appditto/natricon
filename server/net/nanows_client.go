package net

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/recws-org/recws"
)

type wsSubscribe struct {
	Action  string              `json:"action"`
	Topic   string              `json:"topic"`
	Ack     bool                `json:"ack"`
	Options map[string][]string `json:"options"`
}

type subscribeResponse struct {
	Topic   string                 `json:"topic"`
	Message map[string]interface{} `json:"message"`
	Block   map[string]interface{} `json:"block"`
}

func StartNanoWSClient(wsUrl string, account string) {
	ctx, cancel := context.WithCancel(context.Background())
	sentSubscribe := false
	ws := recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	// Nano subscription request
	subRequest := wsSubscribe{
		Action: "subscribe",
		Topic:  "confirmation",
		Ack:    false,
		Options: map[string][]string{
			"accounts": {
				account,
			},
		},
	}
	ws.Dial(wsUrl, nil)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer func() {
		signal.Stop(sigc)
		cancel()
	}()

	for {
		select {
		case <-sigc:
			cancel()
		case <-ctx.Done():
			go ws.Close()
			glog.Infof("Websocket closed %s", ws.GetURL())
			return
		default:
			if !ws.IsConnected() {
				sentSubscribe = false
				glog.Infof("Websocket disconnected %s", ws.GetURL())
				continue
			}

			// Sent subscribe with ack
			if !sentSubscribe {
				if err := ws.WriteJSON(subRequest); err != nil {
					glog.Infof("Error sending subscribe request %s", ws.GetURL())
					return
				}
			}

			var confMessage subscribeResponse
			err := ws.ReadJSON(&confMessage)
			if err != nil {
				glog.Infof("Error: ReadJSON %s", ws.GetURL())
				return
			}

			glog.Infof("Received callback WS hash %s", confMessage.Message["hash"])
		}
	}
}
