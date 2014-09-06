package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/parser"
	"github.com/creack/goray/rt"
	"github.com/creack/goray/utils"
)

func init() {
	parser.RegisterParser("json", &JSONParser{})
}

type JSONParser struct {
}

func (jp *JSONParser) Extensions() []string {
	return []string{
		"json",
	}
}

func toObjectConfig(in objectConfig) objects.ObjectConfig {
	out := objects.ObjectConfig{
		Type: in.Type,
		Position: objects.Point{
			X: in.Position.X,
			Y: in.Position.Y,
			Z: in.Position.Z,
		},
		Rotation: objects.Vector{
			X: in.Rotation.X,
			Y: in.Rotation.Y,
			Z: in.Rotation.Z,
		},
		R: in.R,
	}
	c, err := utils.DecodeColor(in.Color)
	if err != nil {
		log.Printf("Error decoding color: %s", err)
	}
	out.Color = c
	return out
}

func (jp *JSONParser) Parse(filename string) (*rt.SceneConfig, error) {
	var conf config

	if filename == "-" {
		if err := json.NewDecoder(os.Stdin).Decode(&conf); err != nil {
			return nil, err
		}

	} else {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(content, &conf); err != nil {
			return nil, err
		}
	}

	eye := &rt.Eye{
		Position: objects.Point{
			X: conf.Eye.Position.X,
			Y: conf.Eye.Position.Y,
			Z: conf.Eye.Position.Z,
		},
		Rotation: objects.Vector{
			X: conf.Eye.Rotation.X,
			Y: conf.Eye.Rotation.Y,
			Z: conf.Eye.Rotation.Z,
		},
	}
	objs := []objects.Object{}
	for _, obj := range conf.Objects {
		newObjFct, ok := objects.ObjectList[obj.Type]
		if !ok {
			log.Printf("Unkown section: %s, skipping", obj.Type)
			continue
		}
		obj, err := newObjFct(toObjectConfig(obj))
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}

	if conf.Window.Width == 0 {
		conf.Window.Width = 800
	}
	if conf.Window.Height == 0 {
		conf.Window.Height = 600
	}

	return &rt.SceneConfig{
		Height:  conf.Window.Height,
		Width:   conf.Window.Width,
		Eye:     eye,
		Objects: objs,
	}, nil
}

type config struct {
	Window struct {
		Height int `json:"h"`
		Width  int `json:"w"`
	} `json:"window"`
	Eye     objectConfig   `json:"eye"`
	Objects []objectConfig `json:"objects"`
}

type objectConfig struct {
	Type     string `json:"type"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	} `json:"position"`
	Rotation struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"rotation"`
	Color string `json:"color"`
	R     int    `json:"R"`
}