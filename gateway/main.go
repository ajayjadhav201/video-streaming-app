package main

import (
	"gateway/mqtt"
	"gateway/routes"
	"gateway/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	MqttHosturl = ":1883"
	HttpHosturl = ":8080"
	uploadUrl   = "localhost:8081"
	watchurl    = "localhost:8082"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	//
	// start mqtt server
	server := mqtt.NewMqttServer(MqttHosturl)
	defer server.Close()
	go func() {
		server.InitInlineClient()
		//
		utils.FatalIfError(server.Run(), "Failed to start mqtt server at: %s", MqttHosturl)
	}()
	//  create Mqtt inline Client
	// if services are not available this
	go func() {
		//
		time.Sleep(5 * time.Second)
		//
		_ = mqtt.NewInlineClient(server)
		//
	}()
	//
	// start http server
	proxyServer := routes.NewProxyServer(uploadUrl, watchurl)
	go func() {
		http.HandleFunc("/videos/", proxyServer.WatchVideosProxy)
		http.HandleFunc("/api/v1/upload", proxyServer.UploadVideosProxyHandler)
		utils.FatalIfError(http.ListenAndServe(HttpHosturl, nil), "Failed to start http server at: ", HttpHosturl)
	}()
	//
	<-done
	//
	utils.LogPrintln("Server shut down started...")
	server.Close()
	time.Sleep(1 * time.Second)
	utils.LogPrintln("Server stopped...")
	//
}
