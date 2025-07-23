package mapping_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)

func TestErrNilInput(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr bool
	}{
		{
			name:    "compare with ErrNilInput",
			err:     ErrNilInput,
			wantErr: true,
		},
		{
			name:    "compare with wrapped error",
			err:     errors.New("wrapped: " + ErrNilInput.Error()),
			wantErr: false,
		},
		{
			name:    "compare with different error",
			err:     errors.New("some other error"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.ErrorIs(t, tt.err, ErrNilInput, "Error should be ErrNilInput")
			} else {
				err := errors.Unwrap(tt.err)
				if err == nil {
					err = tt.err
				}
				assert.NotErrorIs(t, err, ErrNilInput, "Error should not be ErrNilInput")
			}
		})
	}
}

func TestErrNilInput_ErrorMethod(t *testing.T) {
	// Test that the error message is as expected
	err := ErrNilInput
	expectedMsg := "input is nil"
	assert.Equal(t, expectedMsg, err.Error(), "Error message should match")
}
