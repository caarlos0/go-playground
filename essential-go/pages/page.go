package pages

import (
	"github.com/russross/blackfriday"
	"io"
	"io/ioutil"
	"text/template"
)

type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

func (p *Page) Render(layout string, w io.Writer) error {
	t, err := template.ParseFiles(layout)
	if err != nil {
		return err
	}
	return t.Execute(w, p)
}

func (p *Page) load(config string) error {
	c, err := LoadConfig(config)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}
	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()
	return nil
}

func NewPage(filename string, config string) (*Page, error) {
	p := &Page{File: filename}
	return p, p.load(config)
}
