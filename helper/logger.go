package helper

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func Logger(request *http.Request) {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	logger.Info().
		Str("Method", request.Method).
		Str("URI", request.RequestURI).
		Msg("Request received")
}
