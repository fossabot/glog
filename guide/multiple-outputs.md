---
toc: false
---

# Logging to multiple files at once

glog supports writing to as many outputs as you want. By default,
stdout and stderr are the only files used as outputs. You may want to
log your messages to other channels as well, which can be achieved by
just adding the output.

The output has to implement the `io.Writer` interface

```go
func main() {
	dbglogfh, err := glog.AddLogFile("debug.log", TRACE, DEBUG)
	if err != nil {
		glog.Errorf("cannot open file for debug logging")
	}
	defer dbglogfh.Close()
	// … any calls to glog.Debug() and glog.Trace() will be printed to
	// screen and written to the file `debug.log`, give that they are
	// enabled.
}
```

