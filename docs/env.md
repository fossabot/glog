---
title: "Environment"
---
# Environment variables

- `GLOG_COLOR` – set whether color should be used 
- `GLOG_LEVEL` – set the loglevel at start. Can be overwritten by the
  program
- `GLOG_METALOGGER` – enable the built-in metalogger
- `NO_COLOR` – disables color if set to any value. Can be overwritten
  by `GLOG_COLOR`

## `GLOG_COLOR`

- `ALWAYS`, `ON`, `1`
	- automatically sets OverrideColor to 1, thereby enabling
	  colour even on file outputs

- `NEVER`, `OFF`, `-1`
	- automatically sets OverrideColor to -1, thereby disabling
	  colour even on terminals

## `GLOG_LEVEL`

Sets the level to the specified value (case insensitive). Special
cases:

- `VERBOSE`
	- an alias for `TRACE`

- `MUTE`, `SILENT`
	- disable every and all log-output

## `GLOG_METALOGGER`

- `1`
	- enable the built-in meta-logger (the logger, logging
	  activities of the logging library). *DO NOT USE THIS UNLESS
	  YOU HAVE GOOD REASON!* It looks horrible.

