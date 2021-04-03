package glog

// see https://rosettacode.org/wiki/Write_to_Windows_event_log#AutoHotkey

import (
	"errors"
	"fmt"
	"io"
	"syscall"
	"unsafe"
)

var _ io.Writer = &sysLogger{}

type sysLogger struct {
	procedure               uintptr
	eventhandle             uintptr
	eventlogInformationType uintptr // type: word
	category                uintptr // word
	eventID                 uintptr
	userSecurityID          uintptr
}

func SysLogger(application string) (sysLogger, error) {
	var logger sysLogger
	// system := syscall.NewLazyDLL("system.dll")
	// mscorlib := syscall.NewLazyDLL("mscorlib.dll")
	advapi, _ := syscall.LoadLibrary("Advapi32.dll")

	regEventSource, _ := syscall.GetProcAddress(advapi, "RegisterEventSource")
	reportEvent, _ := syscall.GetProcAddress(advapi, "ReportEvent")

	ret, _, err := syscall.Syscall(uintptr(regEventSource), 2, uintptr(0), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(application))))

	fmt.Println(ret, err)

	logger.eventhandle = ret
	logger.procedure = reportEvent
}

func (l *sysLogger) Write(data []byte) (int, error) {
	if l.procedure == 0 {
		return 0, ErrSysloggerNotValid
	}

	ret, _, callErr := syscall.Syscall9(uintptr(l.procedure), 9,
		l.eventhandle, // handle
		l.eventlogInformationType,
		l.category,
		l.eventID,
		l.userSecurityID,
		uintptr(1),     // WORD number of strings
		uintptr(99999), // dword datasize
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(string(data)))),
		uintptr(0),
	)
	fmt.Println(ret, callErr)
}

var ErrSysloggerNotValid = errors.New("the syslogger is not valid")

const (
	EVENTLOG_SUCCESS = 0x0004
	EVENTLOG_WARNING = 0x0002
	EVENTLOG_ERROR   = 0x0001
)
