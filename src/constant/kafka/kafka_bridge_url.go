package constant_kafka

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant"
)

const KAFKA_BRIDGE_HOST_NAME = "KAFKA_BRIDGE_HOST"

// NO TAILING SLASH
var KAFKA_BRIDGE_HOST string

func init() {
	KAFKA_BRIDGE_HOST = constant.GetEnv(KAFKA_BRIDGE_HOST_NAME, "http://kafka-bridge")
}
