package main

import (
    "fmt"
    "os"

    "golang.org/x/net/html"
    "github.com/PuerkitoBio/goquery"
    "encoding/json"
    "text/template"
    "bytes"
    "embed"
    "flag"
)

//go:embed rc
var rcFS embed.FS

type option struct {
    value   string
    content string
}

type OutPut struct {
    Content string
    Data    string
}

var (
    out    string
    server string

    cover string
)

func init() {
    flag.StringVar(&out, "out", "", "--out=dirname")
    flag.StringVar(&cover, "cover", "cover.html", "--cover=cover.html")
    flag.StringVar(&server, "server", ":8080", "--server=:8080")
}

func main() {
    flag.Parse()
    filePath := cover

    fmt.Println("filePath:", filePath)
    content := getContent(filePath)

    options := extractOptions(filePath)
    root := &Tree{}
    for _, opt := range options {
        root.AddNode(opt.content, opt.value)
    }

    if len(root.Child) == 0 {
        fmt.Println("no data")
        return
    }
    root = root.Child[0]
    js := getJsTree(root)
    j, _ := json.Marshal([]*JsTree{js})

    outPut := OutPut{
        Content: content,
        Data:    string(j),
    }

    t, err := template.ParseFS(rcFS, "rc/cover.template.html")
    if err != nil {
        fmt.Println("error parsing template:", err)
        os.Exit(-1)
    }

    var buff bytes.Buffer
    err = t.Execute(&buff, outPut)
    if err != nil {
        fmt.Println("error executing template:", err)
        os.Exit(-1)
    }

    if len(out) > 0 {
        outHTML(buff)
    }
    startServer(buff)
}

func extractOptions(filePath string) []option {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil
    }
    defer file.Close()

    doc, err := html.Parse(file)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return nil
    }

    var options []option
    var traverse func(*html.Node)
    traverse = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "option" {
            var opt option
            for _, attr := range n.Attr {
                if attr.Key == "value" {
                    opt.value = attr.Val
                }
            }
            opt.content = n.FirstChild.Data
            options = append(options, opt)
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            traverse(c)
        }
    }
    traverse(doc)

    return options
}

func getContent(filePath string) string {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return ""
    }
    defer file.Close()

    node, err := html.Parse(file)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return ""
    }

    doc := goquery.NewDocumentFromNode(node)

    divContent, _ := doc.Find("#content").Html()
    return divContent

}

func printOptions(options []option) {
    for _, opt := range options {
        fmt.Println("Value:", opt.value)
        fmt.Println("Content:", opt.content)
    }
}
