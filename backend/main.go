package main

import (
	"fmt"
	"log"
	"net/http"

	"calculator/api/apiconnect"
	"calculator/config"
	"calculator/internal/service"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置CORS头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Connect-Protocol-Version")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.NewDefaultConfig()
	calculatorService := service.NewCalculatorService()
	path, handler := apiconnect.NewCalculatorServiceHandler(calculatorService)

	mux := http.NewServeMux()
	mux.Handle(path, corsMiddleware(handler))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	log.Printf("服务器启动在端口 %d", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
