package mapping

func Metadata2Headers(metadata map[string]string) []map[string]string {
	headers := make([]map[string]string, 0)
	for k, v := range metadata {
		headers = append(headers, map[string]string{"key": k, "value": v})
	}
	return headers
}

func Headers2Metadata(headers []map[string]string) map[string]string {
	metadata := make(map[string]string)
	for _, header := range headers {
		key, value := header["key"], header["value"]
		if key == "" || value == "" {
			continue
		}
		metadata[key] = value
	}
	return metadata
}
