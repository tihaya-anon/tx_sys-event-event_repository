package mapping

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"google.golang.org/protobuf/proto"
)

func PB2Bytes(e *pb.Event) ([]byte, error) {
	if e == nil {
		return nil, ErrNilInput
	}
	return proto.Marshal(e)	
}

func Bytes2PB(b []byte) (*pb.Event, error) {
	if b == nil {
		return nil, ErrNilInput
	}
	var obj pb.Event
	err := proto.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
