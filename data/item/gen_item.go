//go:build generate
// +build generate

// gen_item.go generates item information.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	version = "1.20"
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/" + version + "/items.json"
	// language=gohtml
	itemTmpl = `// Code generated by gen_item.go DO NOT EDIT.
// Package item stores information about items in Minecraft.
package item

// ID describes the numeric ID of an item.
type ID uint32

// Item describes information about a type of item.
type Item struct {
	ID          ID
	DisplayName string
	Name        string
	StackSize   uint
}

var (
	{{- range .}}
	{{.CamelName}} = Item{
		ID: {{.ID}},
		DisplayName: "{{.DisplayName}}",
		Name: "{{.Name}}",
		StackSize: {{.StackSize}},
	}{{end}}
)

// ByID is an index of minecraft items by their ID.
var ByID = map[ID]*Item{ {{range .}}
	{{.ID}}: &{{.CamelName}},{{end}}
}`
)

type Item struct {
	ID          uint32 `json:"id"`
	CamelName   string `json:"-"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	StackSize   uint   `json:"stackSize"`
}

func downloadInfo() ([]*Item, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []*Item
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for _, d := range data {
		d.CamelName = strcase.ToCamel(d.Name)
	}
	return data, nil
}

//go:generate go run $GOFILE
//go:generate go fmt item.go
func main() {
	fmt.Println("generating item.go")
	items, err := downloadInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	f, err := os.Create("item.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := template.Must(template.New("").Parse(itemTmpl)).Execute(f, items); err != nil {
		panic(err)
	}
}
