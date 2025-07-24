package redis

import (
	"strconv"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/util"
)

const (
	REDIS_ADDR_NAME    = "REDIS_ADDR"
	REDIS_ADDR_DEFAULT = "http://localhost:6379"
)

const (
	REDIS_PWD_NAME    = "REDIS_PWD"
	REDIS_PWD_DEFAULT = ""
)

const (
	REDIS_DB_NAME    = "REDIS_DB"
	REDIS_DB_DEFAULT = "0"
)

// NO TAILING SLASH
var REDIS_ADDR string
var REDIS_PWD string
var REDIS_DB int

func Init() {
	REDIS_ADDR = util.GetEnv(REDIS_ADDR_NAME, REDIS_ADDR_DEFAULT)
	REDIS_PWD = util.GetEnv(REDIS_PWD_NAME, REDIS_PWD_DEFAULT)
	REDIS_DB, _ = strconv.Atoi(util.GetEnv(REDIS_DB_NAME, REDIS_DB_DEFAULT))
}
