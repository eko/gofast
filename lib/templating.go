// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "errors"
    "fmt"
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
            errors.New(fmt.Sprintf("Directory '%s' does not exists", name))
        }
    }

    t.directory = name
}

// Returns templating base directory
func (t *templating) GetDirectory() string {
    return t.directory
}