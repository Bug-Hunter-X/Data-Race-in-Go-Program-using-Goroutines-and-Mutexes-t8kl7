```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var mutex sync.Mutex
        var counter int

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mutex.Lock()
                        counter += i
                        mutex.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println("Counter:", counter)
}
```