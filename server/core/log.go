package core

import (
	"blog/config"
	"blog/global"
	"blog/utils"
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	oplogging "github.com/op/go-logging"
	"io"
	"os"
	"strings"
	"time"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "blog"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func init() {
	c := global.GVA_CONFIG.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := oplogging.MustGetLogger(module)
	var backends []oplogging.Backend
	backends = registerStdout(c, backends)
	backends = registerFile(c, backends)

	oplogging.SetBackend(backends...)
	global.GVA_LOG = logger

}

func registerStdout(c config.Log, backends []oplogging.Backend) []oplogging.Backend {
	if c.Stdout != "" {
		level, err := oplogging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(os.Stdout, c, level))
	}
	return backends
}

func registerFile(c config.Log, backends []oplogging.Backend) []oplogging.Backend {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotatelogs.New(
			logDir+string(os.PathSeparator)+"%y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(logSoftLink),
			rotatelogs.WithMaxAge(7*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
			return backends
		}
		level, err := oplogging.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(fileWriter, c, level))
	}
	return backends
}

func createBackend(w io.Writer, c config.Log, level oplogging.Level) oplogging.Backend {
	backend := oplogging.NewLogBackend(w, c.Prefix, 0)
	stdoutWrite := false
	if w == os.Stdout {
		stdoutWrite = true
	}
	format := getLogFormatter(c, stdoutWrite)
	backendleveled := oplogging.AddModuleLevel(oplogging.NewBackendFormatter(backend, format))
	backendleveled.SetLevel(level, module)
	return backendleveled
}
func getLogFormatter(c config.Log, stdouWriter bool) oplogging.Formatter {
	pattern := defaultFormatter
	if !stdouWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return oplogging.MustStringFormatter(pattern)
}
