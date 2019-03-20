RedisDB
======
## Example (How to use)

```go
package main

import (
    "github.com/liontail/redisdb"
)

redis, err := Initial(url, password)
...

// Set Key = "test" : value = "Hello"
redis.Set("test", "Hello")

...

// Get Key = "test"

result, err := redis.Get("test")

...
```