---
toc: false
---

# Log panics with go

Reporting a panic is not that easy, if you don't know what lead to it.
That's why glog comes with a dead-simple panic handler. Just defer the
call to it at the start of main and forget about it.

The file it the panic is logged to can be changed using
`glog.PanicLogFile`.  Remember that all errors that occur while
handling the panic are not logged and discarded instead.

```go
func main() {
	defer glog.PanicHandler()
	// … do you program stuff …
}
```

