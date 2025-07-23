package mapping_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)

func TestMetadata2Headers(t *testing.T) {
	tests := []struct {
		name     string
		metadata map[string]string
		expected []map[string]string
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			expected: nil,
		},
		{
			name:     "empty metadata",
			metadata: map[string]string{},
			expected: []map[string]string{},
		},
		{
			name: "single key-value pair",
			metadata: map[string]string{
				"key1": "value1",
			},
			expected: []map[string]string{{"key": "key1", "value": "value1"}},
		},
		{
			name: "multiple key-value pairs",
			metadata: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: []map[string]string{
				{"key": "key1", "value": "value1"},
				{"key": "key2", "value": "value2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Metadata2Headers(tt.metadata)
			assert.ElementsMatch(t, tt.expected, result, "Metadata2Headers result should match expected")
		})
	}
}

func TestHeaders2Metadata(t *testing.T) {
	tests := []struct {
		name     string
		headers  []map[string]string
		expected map[string]string
	}{
		{
			name:     "nil headers",
			headers:  nil,
			expected: map[string]string{},
		},
		{
			name:     "empty headers",
			headers:  []map[string]string{},
			expected: map[string]string{},
		},
		{
			name: "single header",
			headers: []map[string]string{
				{"key": "key1", "value": "value1"},
			},
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			name: "multiple headers",
			headers: []map[string]string{
				{"key": "key1", "value": "value1"},
				{"key": "key2", "value": "value2"},
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "headers with missing key or value",
			headers: []map[string]string{
				{"key": "key1"},                    // missing value
				{"value": "value2"},                 // missing key
				{"key": "key3", "value": "value3"}, // valid
			},
			expected: map[string]string{
				"key3": "value3", // only valid pairs should be included
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Headers2Metadata(tt.headers)
			assert.Equal(t, tt.expected, result, "Headers2Metadata result should match expected")
		})
	}
}

func TestHeaders2Metadata_EdgeCases(t *testing.T) {
	// Test with nil slice
	t.Run("nil slice", func(t *testing.T) {
		result := Headers2Metadata(nil)
		assert.Empty(t, result, "Should return empty map for nil slice")
	})

	// Test with empty slice
	t.Run("empty slice", func(t *testing.T) {
		result := Headers2Metadata([]map[string]string{})
		assert.Empty(t, result, "Should return empty map for empty slice")
	})

	// Test with nil map in slice
	t.Run("nil map in slice", func(t *testing.T) {
		result := Headers2Metadata([]map[string]string{nil, {"key": "key1", "value": "value1"}})
		expected := map[string]string{"key1": "value1"}
		assert.Equal(t, expected, result, "Should skip nil maps in slice")
	})
}
