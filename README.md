# optinator
Go packages are generally start with a main struct and the package initiates and fills that struct in the beginning. There are so many ways to fill that struct.  

In this repo I wanted to show the idiomatic way to fill a struct. This is generally used with "options" parameters. So I called it optinator. Hope you find it helpful.

Example usage:
```
func main() {
	req := NewReq(
		WithAddress("https://yakuter.com"),
		WithTimeout(30*time.Second),
		WithContentType("application/json"),
	)

	fmt.Printf("%+v", req)
}
```

Sources:  
- [Different Ways to Initialize Go structs](https://asankov.dev/blog/2022/01/29/different-ways-to-initialize-go-structs/)  
- [github.com/sethvargo/go-githubactions](https://github.com/sethvargo/go-githubactions/blob/main/options.go)
- [github.com/binalyze/httpreq](https://github.com/binalyze/httpreq)