```go
package main

import (

        "fmt"
        "sync/atomic"
)

func main() {
        var counter int64

        for i := 0; i < 1000; i++ {
                atomic.AddInt64(&counter, int64(i))
        }

        fmt.Println("Counter:", counter)
}
```