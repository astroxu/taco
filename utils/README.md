# utils
> 工具类

## Usage

you need to import the corresponding package name when use it. For example, if you use string-related functions,import the algorithm package like below:

```go
import "github.com/taco/utils/algorithm"
```

## Example

Here takes the algorithm function NewLRUCache as an example, and the algorithm package needs to be imported.

```go
package main

import (
    "github.com/taco/utils/algorithm"
)

func main() {
	cache := algorithm.NewLRUCache[int, int](2)

	cache.Put(1, 1)
	_, _ = cache.Get(0)
}
```