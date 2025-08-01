package constant

import (
	app_constant "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/app"
	kafka_constant "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	postgre_constant "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
)

func init() {
	app_constant.InitAppEnv()
	postgre_constant.InitPostgre(app_constant.APP_ENV)
	kafka_constant.InitKafkaBridge()
}
