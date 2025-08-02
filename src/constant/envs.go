package constant

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/app"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
)

func init() {
	app.Init()
	postgre.Init(app.APP_ENV)
}
