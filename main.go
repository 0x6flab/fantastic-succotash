package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/0x6flab/fantastic-succotash/internal"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
)

var (
	cKey      = os.Getenv("MPESA_CONSUMER_KEY")
	cSecret   = os.Getenv("MPESA_CONSUMER_SECRET")
	passKey   = os.Getenv("MPESA_PASS_KEY")
	levelText = os.Getenv("LOG_LEVEL")
)

const (
	defLogLevel     = "info"
	defBaseURL      = "https://sandbox.safaricom.co.ke"
	defShortCode    = 174379
	defSendInterval = 5 * time.Second
	defDOSInterval  = 30 * time.Minute
)

func main() {
	if levelText == "" {
		levelText = defLogLevel
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(levelText)); err != nil {
		fmt.Println("Failed to unmarshal log level", err)

		return
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	slog.SetDefault(slog.New(logHandler))

	conf := mpesa.Config{
		BaseURL:   defBaseURL,
		AppKey:    cKey,
		AppSecret: cSecret,
	}

	mp, err := mpesa.NewSDK(conf)
	if err != nil {
		slog.Error("Failed to create new mpesa sdk", slog.Any("error", err))

		return
	}

	slog.Info("Starting mpesa_dos")

	for {
		mpesa_dos(mp)
		time.Sleep(defDOSInterval)
	}
}

func mpesa_dos(sdk mpesa.SDK) {
	for _, contact := range internal.Contacts {
		qrReq := mpesa.ExpressSimulateReq{
			PassKey:           passKey,
			BusinessShortCode: defShortCode,
			TransactionType:   "CustomerPayBillOnline",
			PhoneNumber:       contact,
			Amount:            uint64(rand.Int31n(1000)),
			PartyA:            contact,
			PartyB:            defShortCode,
			CallBackURL:       "https://da828a91-bbdd-4f80-97e7-7a13e6caa42f.ngrok.io",
			AccountReference:  "CompanyXLTD",
			TransactionDesc:   "Payment of X",
		}

		if _, err := sdk.ExpressSimulate(qrReq); err != nil {
			slog.Error("Failed to simulate mpesa express", slog.Any("error", err))

			continue
		}
		slog.Info("Simulated mpesa express", slog.Any("contact", contact))
		time.Sleep(defSendInterval)
	}
}
