# Defaults

## glog

- *Time Format*:      ISO-8601
- *Show caller*:      TRACE, DEBUG, ERROR, FATAL
- *Short caller*:     yes
- *Show caller line*: no

## logrotation

- *Permissions*:               600
- *Maximum Filesize*:          32 MiB[^1]
- *Retention (kept old logs)*: 2
- *KeptPercent*:               5
- *Compression*:               GZip
- *Synchronous write*:         yes

[^1]: see also: [reasoning](docs/reasoning.md)
