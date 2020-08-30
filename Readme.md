```go

package main

import (
    "fmt"

    im "github.com/frozen/immutable_map"
)

func main() {
    m := im.New()
    m2 := m.Insert([]byte("value"), 5)
    
    fmt.Println(m.Contains([]byte("value"))) // false

    fmt.Println(m2.Contains([]byte("value"))) // true

    fmt.Println(m2.Get([]byte("value"))) // 5, true
}





```