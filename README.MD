# Wrapped Logger
The package supplies adapters to various logging libraries.
This allows the application not to be tied to any particular library, but, if necessary, to painlessly replace one with another.

This is implemented due to the fact that your application is tied to this wrapper, and not a specific library.
The wrapper, in turn, has a standardized logging interface that adapts to a specific library.

When there is a need to try some new logging library, you can use a ready-made adapter or write your own, which is much easier than rewriting the entire application.

# Documentation
- [Russian](README.RU.MD)
- [English](README.MD)

# Available adapters
- STD Log (https://go.dev/pkg/log)
- Zap (https://github.com/uber-go/zap)
- Nop (for use in tests)

# Features of the wdlogger
- zap adapter is able to output logs in the **GELF** format (others will have this opportunity in the future)

# Project status
It is currently under experimental development and is used in my own projects for convenience, but this does not prevent you from proposing your own changes if you wish.