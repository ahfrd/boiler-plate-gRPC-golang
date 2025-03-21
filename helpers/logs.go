package helpers

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func LogRequest(ctx context.Context, requestString string, uuid string, path string) string {
	logger := logrus.New()
	currentTime := time.Now()
	logger.Info("[Start][RequestId]= " + uuid + ", [Path]= " + path + ", [IP]= " + GetPeerIP(ctx) + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString)
	return "[Start][RequestId]= " + uuid + ", [Path]= " + path + ", [IP]= " + GetPeerIP(ctx) + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString
}

func LogResponse(ctx context.Context, responseString string, uuid string, path string) string {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Info("[Stop][RequestId]= " + uuid + ", [Path]= " + path + ", [IP]= " + GetPeerIP(ctx) + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString)
	return "[Stop][RequestId]= " + uuid + ", [Path]= " + path + ", [IP]= " + GetPeerIP(ctx) + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString
}

func LogError(ctx context.Context, responseString string, uuid string, path string) {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Error("[Stop][RequestId]= " + uuid + ", [Path]= " + path + ", [IP]= " + GetPeerIP(ctx) + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Error]= " + responseString)
}

func LogScrapStart(requestString string, uuid string, path string) string {
	logger := logrus.New()
	currentTime := time.Now()
	logger.Info("[Start][RequestId]= " + uuid + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString)
	return "[Start][RequestId]= " + uuid + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString
}

func LogScrapEnd(responseString string, uuid string, path string) string {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Info("[Stop][RequestId]= " + uuid + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString)
	return "[Stop][RequestId]= " + uuid + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString
}

func LogScrapError(responseString string, uuid string, path string) {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Error("[Stop][RequestId]= " + uuid + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Error]= " + responseString)
}
