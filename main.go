package main

import (
    "context"
    "fmt"
    "time"
    "os"

	"github.com/kokizzu/gotro/L"
    "dagger.io/dagger"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Second)
	defer cancel()
    if err := build(ctx); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Building with Dagger")

    L.Print(time.Now())
    // initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithWorkdir("."), dagger.WithLogOutput(os.Stdout))
    if L.IsError(err,"dagger.Connect") {
        return err
    }
    defer client.Close()
    L.Print(time.Now())

	entries, err := client.Host().Directory(".").Entries(ctx)
    if L.IsError(err,"Entries") {
        return err
    }
    L.Print(time.Now())

    fmt.Println(entries)
    return nil
}
