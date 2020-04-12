# config

[![GoDoc](https://godoc.org/github.com/abemedia/config?status.png)](https://godoc.org/github.com/abemedia/config)

Simple config files in go. Supports reading from JSON & YAML files.
If you require a decent config package I'd look into [viper](https://github.com/spf13/viper) instead.

## Usage

```go
package main

import (
    "fmt"

    "github.com/abemedia/config"
)

type Config struct {
    Foo struct {
        Bar string
    }
}

func main() {
    c, err := config.New(&Config{})
    if err != nil {
        panic(err)
    }

    err = c.ReadFile("./config.yaml")
    if err != nil {
        panic(err)
    }

    val, err := c.Get("Foo.Bar")
    if err != nil {
        panic(err)
    }

    fmt.Println(val)

    err := c.Set("Foo.Bar", "baz")
    if err != nil {
        panic(err)
    }

    fmt.Println(val)
}

```
