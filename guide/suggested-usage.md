---
toc: false
---
# Suggested usage of glog

glog is designed to stay out of your way, by doing things as you
probably want them to work, while at the same time allowing complete
customization.  I therefore suggest the follwing usage. (This is
obviously not applicable to every use, but should be an okay base to
work of)

- defer the PanicHandler
- Set the loglevel using `glog.SetLevel`
	- I'd suggest using INFO or WARNING as to not slow down your
	  program to much
- Add a logfile for errors as output for ERROR and FATAL
- Log as much as you seem fit
