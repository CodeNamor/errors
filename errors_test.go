package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constError_Error(t *testing.T) {
	s2 := testError.Error()

	var err = "error"

	assert.Equal(t, s2, err)
}

func Test_constError_Is(t *testing.T) {
	err := fmt.Errorf(string(testError))

	var e2 = err
	s2 := testError.Is(e2)

	assert.Equal(t, true, s2)
}

func Test_constError_Wrap(t *testing.T) {
	err := fmt.Errorf(string(testError))

	var e2 = err
	s2 := testError.Wrap(e2)

	assert.Error(t, s2)
}

func Test_WrapError_Error(t *testing.T) {
	t.Run("Error populated", func(t *testing.T) {
		err := wrapError{
			err: testError,
			msg: "error",
		}

		e2 := err.Error()
		e := "error: error"

		assert.Equal(t, e, e2)
	})
	t.Run("Error is nil", func(t *testing.T) {
		err := wrapError{
			err: nil,
			msg: "error",
		}

		e2 := err.Error()
		e := "error"

		assert.Equal(t, e, e2)
	})
}

func Test_WrapError_Unwrap(t *testing.T) {
	err := wrapError{
		err: testError,
		msg: "error",
	}

	e2 := err.Unwrap()
	e := err.err

	assert.Equal(t, e, e2)
}

func Test_WrapError_Is(t *testing.T) {
	err := wrapError{
		err: testError,
		msg: "error",
	}

	err2 := err.Is(err)

	assert.Equal(t, true, err2)
}
