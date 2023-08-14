# Simple Golang in-memory cache

This package let you save get and delete values from in-memory cache

Installation:

`go get -u github.com/NIRumiantsev/golang-memory-cache`

Usage:
```
func main() {
    cache := cache.New()

    cache.Set("userId", 42)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
}
```