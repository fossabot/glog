# Reasoning

## Attach stderr levels to stdout if stdout is redirected to a file

When piping stdout into a file, chances are that you want to work with
a logfile and check something. Since there is no indication for the
end user that warning and above are sent to stderr, they are also sent
to stdout to allow easier debugging of issues for sysadmins working
with glog-software.

## The default format is strange

If that's an issue to you, feel free to change it.

I chose this format because it is simple to read (especially when
compared to JSON-logging) and allows working with the logfile using
cut/sed/awk without having to worry too much about what if `|` or `-`
is in a logmessage.  I consider `–` to be unlikely, while not being
too strange (`酒` would also be a solution but is way harder to enter)

## Caller is not printed for INFO and WARNING by default

This is because getting the caller is a relatively expensive operation
(19 µs vs 34 µs). Having written my fair share of software, I found
that most messages are either Info or Warning (Informational or
handled/expected errors). In both cases the caller is also usually not
relevant, as debugging happens primarily through the debug and trace
levels.

## The logrotation file is enormous

I consider 32 MiB a medium sized logfile, allowing for an ample amount
of logmessages (about 400.000 in my tests) while being fast enough to
compress to not unnecessarily slow down the program.

## Caller shows only the function, not the line

Lines are a very volatile mean of indicating a caller. Line 30 may
belong to function `A()` in version 1 and function `B()` in version 2.
However, if the functions do not change often or the potential
revision is known, lines may very well be better for debugging as they
allow pinpointing a message to their exact origin. This however, is
not always the case. The function name does not change as much and is
helpful in that it does not nearly change as often.
