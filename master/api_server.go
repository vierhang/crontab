package master

import (
	"net"
	"net/http"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	//GAppServer 单例对象
	GAppServer *ApiServer
)

func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	// 启动tpc
	if listener, err = net.Listen("tcp", ":8070"); err != nil {
		return
	}
	// 创建http 服务对象
	httpServer = &http.Server{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}
	GAppServer = &ApiServer{httpServer: httpServer}

	go httpServer.Serve(listener)
	return
}
