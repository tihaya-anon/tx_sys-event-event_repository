package constant

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/app"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
)

func init() {
	go kafka.Init()
	go redis.Init()
	app.Init()
	postgre.Init(app.APP_ENV)
}
