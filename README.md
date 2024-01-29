# isdebug

Go library to check if your code is currently running in a debugger.

Usage:

```go
import "github.com/ReallyLiri/go-isdebug"

// ...

if isdebug.IsDebugging()() {
    // do something
}
```
