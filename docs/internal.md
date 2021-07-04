# Inner Workings

## Initialization

Initialization happens as soon as the program starts.

### `env.go`

- check if the metalogger is requested
- activate ANSI-Escape codes of on Windows
- check if a color override is set and apply it
- check if a specific loglevel is requested and set it

### `init.go`

- check if stdout is piped into a file
	- if it is, add this output to the stderr levels (warning+)

## Receiving a log message

### Preparation

- directly after receiving the request a check is performed whether
  the loglevel is enabled
	- if it is not, the message is discarded
- required values (level, timestamp, caller (optional) and message)
  are collected
- the collected values are then assembled to form a log entry.

### Printing

- the outputs of the specified level are iterated over
- writes happen concurrently
- if the output is not a terminal or color was turned off, all ANSI
  codes are stripped from the entry
- in case of errors, these errors are collected

### Finishing

- in the last step caught errors are written to all outputs, hoping
  that at least one will be in able to be written.
- if that is not the case, the errors are discarded

## Panic Handler

The `PanicHandler()` allows logging panics to a special file and are
then forwarded to be handled by either the runtime or the developer

- `recover()` and check if there was a `panic()`
	- if there was, a dedicated file is opened panic message and
	  stack trace are then written to this file.
