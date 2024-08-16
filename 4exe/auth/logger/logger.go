package logger

import (

	"log"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	mLog := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	txt, err := os.OpenFile("auth.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed opening file: %v", err)
	}
	logger := slog.New(slog.NewJSONHandler(txt, &mLog))

	return logger	
}
