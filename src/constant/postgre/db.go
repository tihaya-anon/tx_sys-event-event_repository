package postgre_constant

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/util"
)

const (
	TEST_DB_URL_NAME = "TEST_DB_URL"
	DEV_DB_URL_NAME  = "DEV_DB_URL"
	PROD_DB_URL_NAME = "PROD_DB_URL"
)

var DB_URL string

func Init(appEnv string) {
	switch appEnv {
	case "test":
		DB_URL = util.MustGetEnv(TEST_DB_URL_NAME)
	case "dev":
		DB_URL = util.MustGetEnv(DEV_DB_URL_NAME)
	case "prod":
		DB_URL = util.MustGetEnv(PROD_DB_URL_NAME)
	default:
		DB_URL = util.MustGetEnv(TEST_DB_URL_NAME)
	}
}
