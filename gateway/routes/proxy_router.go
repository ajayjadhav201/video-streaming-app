package routes

import (
	"encoding/json"
	"gateway/utils"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyServer struct {
	uploadVideosProxy *httputil.ReverseProxy
	watchVideosProxy  *httputil.ReverseProxy
}

func NewProxyServer(uploadVideosUrl string, watchVideosUrl string) *ProxyServer {
	//
	uploadUrl, err := url.Parse(uploadVideosUrl)
	utils.FatalIfError(err, "Unable to parse url: %s", uploadVideosUrl)
	watchUrl, err := url.Parse(watchVideosUrl)
	utils.FatalIfError(err, "Unable to parse url: %s", watchVideosUrl)
	//
	//
	uploadProxy := httputil.NewSingleHostReverseProxy(uploadUrl)
	watchProxy := httputil.NewSingleHostReverseProxy(watchUrl)
	//
	return &ProxyServer{
		uploadVideosProxy: uploadProxy,
		watchVideosProxy:  watchProxy,
	}
}

func (ps *ProxyServer) UploadVideosProxyHandler(w http.ResponseWriter, r *http.Request) {
	//
	ps.uploadVideosProxy.ServeHTTP(w, r)
}

func (ps *ProxyServer) WatchVideosProxy(w http.ResponseWriter, r *http.Request) {
	//
	//
	path := strings.TrimPrefix(r.URL.Path, "/videos/")
	if path == "" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"ERROR": "videoID not provided"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//
	segments := strings.Split(path, "/")
	// Check if more than one path segment is provided
	if len(segments) > 1 && segments[1] != "" {
		http.Error(w, "Too many path segments", http.StatusBadRequest)
		return
	}

	//
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"videoID": "videoID"}
	json.NewEncoder(w).Encode(response)
	//
	// ps.watchVideosProxy.ServeHTTP(w, r)
}
