package main

import (
	"os"
	"strings"
	"text/template"
)

var tpl = `package {{.name}}

import (
	"image/color"

	"github.com/creack/goray/objects"
)

func init() {
	objects.RegisterObject("{{.name}}", New{{.Name}})
}

// {{.Name}} is the object's implemetation for a {{.Name}}.
type {{.Name}} struct {
	position objects.Point
	color    color.Color
}

// New{{.Name}} instanciate the {{.Name}} object.
func New{{.Name}}(obj objects.ObjectConfig) (objects.Object, error) {
	return (&{{.Name}}{}).Parse(obj)
}

// Color returns the Object's color
func (p *{{.Name}}) Color() color.Color {
	return p.color
}

// Parse populates the {{.Name}}'s values from the given configuration object.
// If the {{.name}} is nil, instanciate it.
func (p *{{.Name}}) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if p == nil {
		p = &{{.Name}}{}
	}
	p.position = obj.Position
	p.color = obj.Color
	return p, nil
}

// Intersect calculates the distance between the eye and the Object.
func (p *{{.Name}}) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(p.position)
	defer eye.Add(p.position)

	return 0.
}
`

func main() {
	objName := strings.ToLower("cube")
	objNameCapitalized := string([]byte{objName[0] - ' '}) + objName[1:]
	template.Must(template.New("tpl").Parse(tpl)).Execute(os.Stdout, map[string]string{
		"Name": objNameCapitalized,
		"name": objName,
	})
}
