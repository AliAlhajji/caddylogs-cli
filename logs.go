package main

import (
	"github.com/AliAlhajji/caddylogs"
	logfilters "github.com/AliAlhajji/caddylogs/log_filters"
	"github.com/AliAlhajji/caddylogs/models"
)

type LogsManager struct {
	logs *caddylogs.AccessLogs
}

func NewLogsManager(logsPath string) (*LogsManager, error) {
	accessLogs, err := caddylogs.New(logsPath)
	if err != nil {
		return nil, err
	}

	logsManager := &LogsManager{
		logs: accessLogs,
	}

	return logsManager, nil
}

func (m *LogsManager) UrlContains(str string) {
	m.logs.StringFilter(logfilters.UrlContains, str)
}

func (m *LogsManager) RefererContains(str string) {
	m.logs.StringFilter(logfilters.RefererContains, str)
}

func (m *LogsManager) LoggerIs(str string) {
	m.logs.StringFilter(logfilters.LoggerIs, str)
}

func (m *LogsManager) StatusCode(statusCode int) {
	m.logs.IntFilter(logfilters.StatusCodeIs, statusCode)
}

func (m *LogsManager) RequestHeaderIs(headerKey string, headerValue string) {
	m.logs.KeyValueFilter(logfilters.RequestHeaderIs, headerKey, headerValue)
}

func (m *LogsManager) InfoLogs() {
	m.logs.Filter(logfilters.InfoLogs)
}

func (m *LogsManager) ErrorLogs() {
	m.logs.Filter(logfilters.ErrorLogs)
}

func (m *LogsManager) First(n int) {
	m.logs.First(n)
}

func (m *LogsManager) Last(n int) {
	m.logs.Last(n)
}

func (m *LogsManager) Reverse() {
	m.logs.Reverse()
}

func (m *LogsManager) GetLogs() []*models.Log {
	return m.logs.GetLogs()
}

func (m *LogsManager) GetLogsCount() int {
	return len(m.GetLogs())
}
