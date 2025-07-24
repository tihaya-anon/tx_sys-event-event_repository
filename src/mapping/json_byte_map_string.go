package mapping

import "encoding/json"

func Map2Bytes(m map[string]string) ([]byte, error) {
	if m == nil {
		return nil, ErrNilInput
	}
	return json.Marshal(m)
}

func Bytes2Map(data []byte) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
