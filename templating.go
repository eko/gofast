// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"fmt"
	"os"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
)

type Templating struct {
	viewsDirectory  string
	assetsDirectory string
}

// NewTemplating creates a new templating component instance
func NewTemplating() Templating {
	return Templating{}
}

// SetViewsDirectory sets templating views directory
func (t *Templating) SetViewsDirectory(name string) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			logrus.Warn(fmt.Sprintf("Directory '%s' does not exists", name))
			os.Exit(1)
		}
	}

	t.viewsDirectory = name
}

// GetViewsDirectory returns templating views directory
func (t *Templating) GetViewsDirectory() string {
	return t.viewsDirectory
}

// SetAssetsDirectory sets templating assets directory
func (t *Templating) SetAssetsDirectory(name string) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			logrus.Warn(fmt.Sprintf("Directory '%s' does not exists", name))
			os.Exit(1)
		}
	}

	t.assetsDirectory = name
}

// GetAssetsDirectory returns templating assets directory
func (t *Templating) GetAssetsDirectory() string {
	return t.assetsDirectory
}

// Render renders a template
func (t *Templating) Render(context Context, name string) {
	var filename = fmt.Sprintf("%s/%s", t.GetViewsDirectory(), name)

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			logrus.Warn(fmt.Sprintf("View '%s' does not exists", filename))
			os.Exit(1)
		}
	}

	var template = pongo2.Must(pongo2.FromFile(filename))
	template.ExecuteWriter(pongo2.Context{
		"request":  context.GetRequest(),
		"response": context.GetResponse(),
	}, context.GetResponse())
}
