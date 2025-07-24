package constant

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/app"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
)

func init() {
	app.InitAppEnv()
	postgre.InitPostgre(app.APP_ENV)
	kafka.InitKafkaBridge()
}
