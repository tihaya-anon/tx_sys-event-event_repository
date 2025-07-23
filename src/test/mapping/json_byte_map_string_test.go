package mapping_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)

func TestMap2Bytes(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		wantErr  bool
		validate func(t *testing.T, data []byte)
	}{
		{
			name:    "nil map",
			input:   nil,
			wantErr: false,
			validate: func(t *testing.T, data []byte) {
				t.Helper()
				assert.Nil(t, data, "Expected nil for nil input map")
			},
		},
		{
			name:    "empty map",
			input:   map[string]string{},
			wantErr: false,
			validate: func(t *testing.T, data []byte) {
				t.Helper()
				assert.JSONEq(t, `{}`, string(data), "Expected empty JSON object for empty map")
			},
		},
		{
			name: "single key-value pair",
			input: map[string]string{
				"key1": "value1",
			},
			wantErr: false,
			validate: func(t *testing.T, data []byte) {
				t.Helper()
				assert.JSONEq(t, `{"key1":"value1"}`, string(data))
			},
		},
		{
			name: "multiple key-value pairs",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
			validate: func(t *testing.T, data []byte) {
				t.Helper()
				assert.JSONEq(t, `{"key1":"value1","key2":"value2"}`, string(data))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Map2Bytes(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.validate(t, result)
			}
		})
	}
}

func TestBytes2Map(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expected    map[string]string
		expectError bool
	}{
		{
			name:        "nil input",
			input:       nil,
			expectError: true,
		},
		{
			name:        "empty JSON object",
			input:       []byte(`{}`),
			expected:    map[string]string{},
			expectError: false,
		},
		{
			name:        "single key-value pair",
			input:       []byte(`{"key1":"value1"}`),
			expected:    map[string]string{"key1": "value1"},
			expectError: false,
		},
		{
			name:        "multiple key-value pairs",
			input:       []byte(`{"key1":"value1","key2":"value2"}`),
			expected:    map[string]string{"key1": "value1", "key2": "value2"},
			expectError: false,
		},
		{
			name:        "invalid JSON",
			input:       []byte(`{"key1":`),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Bytes2Map(tt.input)

			if tt.expectError {
				assert.Error(t, err, "Expected an error")
			} else {
				require.NoError(t, err, "Did not expect an error")
				assert.Equal(t, tt.expected, result, "Maps should match")
			}
		})
	}
}

func TestMap2Bytes_RoundTrip(t *testing.T) {
	// Test round-trip conversion
	original := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// Map to bytes
	bytes, err := Map2Bytes(original)
	require.NoError(t, err, "Map2Bytes should not return an error")

	// Bytes back to map
	result, err := Bytes2Map(bytes)
	require.NoError(t, err, "Bytes2Map should not return an error")

	// Should match original
	assert.Equal(t, original, result, "Round-trip conversion should preserve the original map")
}

func TestMap2Bytes_ErrorHandling(t *testing.T) {
	// This test ensures that non-string values in the map are handled correctly
	// by the JSON marshaling process
	invalidMap := map[string]interface{}{
		"key1": make(chan int), // Channels cannot be JSON marshaled
	}

	// We need to use json.Marshal directly to test the error case
	_, err := json.Marshal(invalidMap)
	assert.Error(t, err, "Should get an error when trying to marshal invalid data")

	// The actual Map2Bytes function only accepts map[string]string,
	// so we don't need to test the error case there as it's not possible with the type system
}
