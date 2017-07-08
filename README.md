m1
==

`m1` stands for "Module 1" or "First Module" and it is designed to be used with
[Hello World Go WebApp][1].

[Hello World Go WebApp][1] modules should export a [Map literal][2] whose keys
are module relative routes and values [http.ServeMux][3] handler functions.

```go
var Routes = map[string]func(w http.ResponseWriter, req *http.Request){
    "/": func(w httpResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Module home")
    }
}
```

[1]: https://github.com/PauloASilva/hello-world
[2]:
[3]:
