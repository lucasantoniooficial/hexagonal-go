package handler_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "error message"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"error message"}`), result)
}
