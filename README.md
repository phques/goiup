# goiup
Go wrapper library for the IUP GUI library in C  
**This is a fork of [github.com/gonutz/goiup](https://github.com/gonutz/goiup)**

ps: since this uses `#include <iup.h>`, (not `#include <iup/iup.h>`),  
we must set `CGO_CFLAGS=-I/usr/include/iup` to build
