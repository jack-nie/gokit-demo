package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUpperCaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
