//go:build generate
// +build generate

// gen_soundid.go generates the enumeration of sound IDs.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"text/template"
)

const (
	version     = "1.17.1"
	protocolURL = "https://pokechu22.github.io/Burger/" + version + ".json"
	// language=gohtml
	soundTmpl = `// Code generated by gen_soundid.go. DO NOT EDIT.

package soundid

// SoundID represents a sound ID used in the minecraft protocol.
type SoundID int32

// SoundNames - map of ids to names for sounds.
var SoundNames = map[SoundID]string{ {{range .}}
	{{.ID}}: "{{.Name}}",{{end}}
}

// GetSoundNameByID helper method
func GetSoundNameByID(id SoundID) (string, bool) {
	name, ok := SoundNames[id]
	return name, ok
}`
)

type Sound struct {
	ID   int64
	Name string
}

//go:generate go run $GOFILE
//go:generate go fmt soundid.go
func main() {
	fmt.Println("generating soundid.go")
	sounds, err := downloadSoundInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	f, err := os.Create("soundid.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := template.Must(template.New("").Parse(soundTmpl)).Execute(f, sounds); err != nil {
		panic(err)
	}
}

func downloadSoundInfo() ([]*Sound, error) {
	resp, err := http.Get(protocolURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// I'm not sure why the response returns a list, it appears to ever only have a single object...
	var data []struct {
		Sounds map[string]Sound `json:"sounds"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	out := make([]*Sound, 0)
	for _, d := range data {
		if has := len(d.Sounds); has > 0 {
			for _, val := range d.Sounds {
				out = append(out, &Sound{ID: val.ID, Name: val.Name})
			}
		} else {
			return nil, fmt.Errorf("no sounds found in data from %s", protocolURL)
		}
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})

	return out, nil
}
