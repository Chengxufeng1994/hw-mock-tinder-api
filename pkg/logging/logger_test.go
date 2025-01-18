package logging

import (
	"testing"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/pkg/errors"
)

func TestLogger(t *testing.T) {
	logger := InitializationLogger(&config.Logging{
		Format:       "json",
		Level:        "debug",
		Name:         "test",
		Outputs:      []string{"stdout"},
		ErrorOutputs: []string{"stderr"},
	})

	t.Run("default", func(t *testing.T) {})

	t.Run("with_fields", func(t *testing.T) {
		fields := make(Fields, 3)
		fields["key1"] = "value1"
		fields["key2"] = "value2"
		fields["key3"] = "value3"
		logger.WithFields(fields).Info("with_fields")
	})

	t.Run("with_error", func(t *testing.T) {
		err := errors.New("with_error")
		wrappedErr := errors.Wrap(err, "wrapped_error")
		logger.WithError(wrappedErr).Info("with_error")
	})
}
