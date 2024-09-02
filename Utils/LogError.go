package Utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

type ErrorDetail struct {
	FileName     string
	Line         int
	FunctionName string
	ErrorText    string
}

func FormatError(errorText string) ErrorDetail {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "UnknownFile"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "UnknownFunction()"
	} else {
		fnName = strings.TrimLeft(filepath.Ext(fn.Name()), ".") + "()"
	}

	return ErrorDetail{FileName: filepath.Base(file), Line: line, FunctionName: fnName, ErrorText: errorText}
}
