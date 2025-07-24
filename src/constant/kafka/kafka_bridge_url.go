package kafka

import "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/util"

const KAFKA_BRIDGE_HOST_NAME = "KAFKA_BRIDGE_HOST"
const KAFKA_BRIDGE_HOST_DEFAULT = "http://kafka-bridge"

// NO TAILING SLASH
var KAFKA_BRIDGE_HOST string

func Init() {
	KAFKA_BRIDGE_HOST = util.GetEnv(KAFKA_BRIDGE_HOST_NAME, KAFKA_BRIDGE_HOST_DEFAULT)
}
