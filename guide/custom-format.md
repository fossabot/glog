---
toc: false
---

# Using a custom log format

Changing the log format is as easy as setting a custom formatting.

```go
func main() {
	glog.LogFormatter = func(lvl Level, t time.Time, caller, message string) string {
		return fmt.Sprintf("%s logged %s on level %s at %s", caller, message, level, t.Format(glog.TimeFormat))
	}
	glog.Debug("Behold, my new format!")
}
```

[Open in playground](https://play.golang.org/p/6unREVq5DSv) (logging
as JSON)
