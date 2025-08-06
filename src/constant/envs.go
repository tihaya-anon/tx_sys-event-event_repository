package constant

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/app"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
)

var DB_URL string
var APP_ENV string

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	app.Init()
	postgre.Init(app.APP_ENV)
	DB_URL = postgre.DB_URL
	APP_ENV = app.APP_ENV
}
