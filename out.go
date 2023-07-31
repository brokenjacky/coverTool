package main

import (
    "bytes"
    "os"
    "fmt"
)

func outHTML(buff bytes.Buffer) {
    finfo, err := os.Stat(out)
    if err == nil && finfo.IsDir() == false {
        fmt.Println("out is a file already")
        return
    } else if err != nil {
        err = os.MkdirAll(out, os.ModePerm)
        if err != nil {
            fmt.Println("mkdir err", err)
            return

        }
    }

    outPath := fmt.Sprintf("%s/cover.html", out)
    os.WriteFile(outPath, buff.Bytes(), 0644)
    fmt.Println("write file to ", outPath)

    files := []string{"jqtree.css", "jquery.min.js", "tree.jquery.debug.js"}
    for _, file := range files {
        name := fmt.Sprintf("%s/%s", out, file)
        content, _ := rcFS.ReadFile(fmt.Sprintf("rc/%s", file))
        os.WriteFile(name, content, 0644)
        fmt.Println("write file to ", name)
    }

    fmt.Println("write file done")
}
