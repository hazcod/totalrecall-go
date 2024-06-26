package recall

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func IsRecallEnabled(l *logrus.Logger, username string) (bool, error) {
	if l == nil {
		l = logrus.New()
	}
	_, _, err := getRecallPaths(l, username)
	return !errors.Is(err, NotEnabledError), nil
}
