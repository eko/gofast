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

type Templating struct {
    viewsDirectory  string
    assetsDirectory string
}

// Creates a new templating component instance
func NewTemplating() Templating {
    return Templating{}
}

// Sets templating views directory
func (t *Templating) SetViewsDirectory(name string) {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            log.Printf("Directory '%s' does not exists", name)
            os.Exit(1)
        }
    }

    t.viewsDirectory = name
}

// Returns templating views directory
func (t *Templating) GetViewsDirectory() string {
    return t.viewsDirectory
}

// Sets templating assets directory
func (t *Templating) SetAssetsDirectory(name string) {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            log.Printf("Directory '%s' does not exists", name)
            os.Exit(1)
        }
    }

    t.assetsDirectory = name
}

// Returns templating assets directory
func (t *Templating) GetAssetsDirectory() string {
    return t.assetsDirectory
}

// Renders a template
func (t *Templating) Render(context Context, name string) {
    var filename = fmt.Sprintf("%s/%s", t.GetViewsDirectory(), name)

    if _, err := os.Stat(filename); err != nil {
        if os.IsNotExist(err) {
            log.Printf("View '%s' does not exists", filename)
            os.Exit(1)
        }
    }

    var template = pongo2.Must(pongo2.FromFile(filename))
    template.ExecuteWriter(pongo2.Context{
        "request": context.GetRequest(),
        "response": context.GetResponse(),
    }, context.GetResponse())
}