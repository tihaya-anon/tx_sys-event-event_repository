package mapping

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"google.golang.org/protobuf/proto"
)

func PB2Bytes(e *pb.Event) ([]byte, error) {
	return proto.Marshal(e)
}

func Bytes2PB(b []byte) (*pb.Event, error) {
	var obj pb.Event
	err := proto.Unmarshal(b, &obj)
	return &obj, err
}
