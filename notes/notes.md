### go - lang ( notes )

"Pardon me for the horrible code that i have written, although suggestions
are always welcomed in any form " - `A Go Lang Beginner`

### Static typed vs Dynamic typed :

Before we dive deep into understanding the core concepts of `go` let me 
shout out lound in order to remind the fact that code written in`go` 
can be both :

+ Staticly typed  ( Example: `C`, `C++`, `Java` ) 
+ Dynamicly typed ( Example: `Python`, `Ruby` )

In order to understand what i am trying to say let me give an example

```go
    // staticly typed
    var x, y int            // integer type
    var x, y int = 10,5     // integer type
    var str string          // string type
    var str string = " "    // string type

    // dynamicly typed
    var x, y            // integer type
    x, y := 10,5        // integer type
    var str = " "       // string type
    str := " "          // string type
```

`NOTE` : As we see above the advantage of using `go` is we can both write
code which is statictly typed as well as dynamicly typed.

### Using package aliases :

Since we might face a situation where in we can have multiple packages that
end with the same name, then we can use package aliases.


```go
    // this is error prone
    import (
        "lib/hector/server"
        "lib/akhil/server"
    )

    // this solves the problem
    import (
        hsrv "lib/hector/server"
        asrv "lib/akhil/server"
    )

    // we can also use thing for the inbuilt packages
    import (
        f "fmt"
        o "os"
        h "net/http"
    )

```
Okay, So if we use the package alias in `go` we can essentially reduce the
scope for multiple packages with the same name.

### 

