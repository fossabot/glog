# Defaults

## glog

- **Time Format**:      ISO-8601 (2006-01-02T15:04:05.000000-0700)
- **Show caller**:      TRACE, DEBUG, ERROR, FATAL
- **Short caller**:     yes
- **Show caller line**: no
- **Colors**:

| Level | Color    |
|-------|----------|
| Trace | Blue     |
| Debug | Cyan     |
| Info  | Green    |
| Warn  | Yellow   |
| Error | Red      |
| Fatal | Red+Bold |

## logrotation

- **Permissions**:               600
- **Maximum Filesize**:          4 MiB
- **Retention (kept old logs)**: 2
- **KeptPercent**:               5
- **Compression**:               GZip
- **Synchronous write**:         yes
