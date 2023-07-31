package main

import (
    "fmt"
    "strings"
)

type Tree struct {
    Name  string
    Child []*Tree
    Value string
    Full  string
    Depth int
}

func (root *Tree) AddNode(key, v string) bool {
    keys := strings.Split(key, "/")
    t := root
    for i, k := range keys {
        if t.Name == k && i == 0 {
            continue
        }

        find := false
        for _, child := range t.Child {
            if child.Name == k {
                t = child
                find = true
                continue
            }
        }
        if find {
            continue
        } else {
            child := &Tree{
                Name:  k,
                Full:  fmt.Sprintf("%s/%s", t.Full, k),
                Depth: i,
            }
            t.Child = append(t.Child, child)
            t = child
        }
    }
    t.Value = v
    return true
}

type JsTree struct {
    Name     string    `json:"name"`
    Id       string    `json:"id,omitempty"`
    Children []*JsTree `json:"children,omitempty"`
}

func getJsTree(tree *Tree) *JsTree {
    if tree == nil {
        return nil
    }
    jsTree := JsTree{
        Name:     tree.Name,
        Id:       tree.Value,
        Children: []*JsTree{},
    }
    for _, child := range tree.Child {
        if child == nil {
            continue
        }
        jsTree.Children = append(jsTree.Children, getJsTree(child))
    }
    return &jsTree
}
