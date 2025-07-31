package test

import (
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener"
)

// TestPollMessages tests the pollMessages function
func TestPollMessages(t *testing.T) {
	// Create test consumer info
	testConsumerInfo := &listener.KafkaConsumerInfo{
		GroupId:  "test-topic",
		Name:     "test-consumer",
		MaxBytes: 1024,
		BaseURI:  "http://test-uri",
		PodName:  "test-pod",
	}

	// Create a mock HTTP client for testing
	originalHTTPGet := listener.HTTPGet
	defer func() { listener.HTTPGet = originalHTTPGet }()

	// Test successful polling
	t.Run("Successful polling", func(t *testing.T) {
		// Mock HTTP response
		listener.HTTPGet = func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       newMockReadCloser(`[{"topic":"test-topic","key":"test-key","value":"test-value"}]`),
			}, nil
		}

		// Call the function
		messages, err := listener.PollMessages(testConsumerInfo)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, messages, 1)
		assert.Equal(t, "test-topic", messages[0]["topic"])
		assert.Equal(t, "test-key", messages[0]["key"])
		assert.Equal(t, "test-value", messages[0]["value"])
	})

	// Test HTTP error
	t.Run("HTTP error", func(t *testing.T) {
		// Mock HTTP error
		listener.HTTPGet = func(url string) (*http.Response, error) {
			return nil, errors.New("connection error")
		}

		// Call the function
		messages, err := listener.PollMessages(testConsumerInfo)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, messages)
		assert.Contains(t, err.Error(), "connection error")
	})

	// Test non-200 status code
	t.Run("Non-200 status code", func(t *testing.T) {
		// Mock HTTP response with non-200 status
		listener.HTTPGet = func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       newMockReadCloser(`{"error":"not found"}`),
			}, nil
		}

		// Call the function
		messages, err := listener.PollMessages(testConsumerInfo)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, messages)
		assert.Contains(t, err.Error(), "404")
	})

	// Test invalid JSON response
	t.Run("Invalid JSON response", func(t *testing.T) {
		// Mock HTTP response with invalid JSON
		listener.HTTPGet = func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       newMockReadCloser(`invalid json`),
			}, nil
		}

		// Call the function
		messages, err := listener.PollMessages(testConsumerInfo)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, messages)
		assert.Contains(t, err.Error(), "unmarshal")
	})
}

// Skip TestCreateListener for now as it requires more mocking setup
// We'll focus on the PollMessages test which is more straightforward

// Helper functions

// mockReadCloser is a mock io.ReadCloser for testing
type mockReadCloser struct {
	data   string
	offset int
}

func newMockReadCloser(data string) *mockReadCloser {
	return &mockReadCloser{data: data}
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	if m.offset >= len(m.data) {
		return 0, io.EOF
	}
	n = copy(p, m.data[m.offset:])
	m.offset += n
	return n, nil
}

func (m *mockReadCloser) Close() error {
	return nil
}
