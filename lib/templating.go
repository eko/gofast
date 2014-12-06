// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "github.com/flosch/pongo2"
    "fmt"
    "log"
    "os"
)

type templating struct {
    directory string
}

// Creates a new templating component instance
func NewTemplating() templating {
    return templating{}
}

// Sets templating base directory
func (t *templating) SetDirectory(name string) {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            log.Printf("Directory '%s' does not exists", name)
            os.Exit(1)
        }
    }

    t.directory = name
}

// Returns templating base directory
func (t *templating) GetDirectory() string {
    return t.directory
}

// Renders a template
func (t *templating) Render(context *context, name string) {
    var filename = fmt.Sprintf("%s/%s", t.GetDirectory(), name)

    if _, err := os.Stat(filename); err != nil {
        if os.IsNotExist(err) {
            log.Printf("View '%s' does not exists", filename)
            os.Exit(1)
        }
    }

    var template = pongo2.Must(pongo2.FromFile(filename))
    template.ExecuteWriter(pongo2.Context{"context": context}, context.GetResponse())
}