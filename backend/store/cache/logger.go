package cache

import "github.com/sirupsen/logrus"

var logger = logrus.WithFields(logrus.Fields{
	"component":     "cache",
	"cache_version": "v1",
})
