package pipeline

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type Pipeline struct {
	Steps []interface{} `yaml:"steps"`
}

const wait = "wait"

type CommandStep struct {
	Label    string   `yaml:"label"`
	Command  string   `yaml:"command"`
	Branches []string `yaml:"branches,omitempty"`
}

type BlockStep struct {
	Block    string   `yaml:"block"`
	Branches []string `yaml:"branches,omitempty"`
}

func Generate() {
	steps := []interface{}{}
	steps = append(steps, CommandStep{
		Label:   ":hammer: Tests",
		Command: "go test ./...",
	})
	steps = append(steps, wait)
	steps = append(steps, BlockStep{
		Block:    ":shipit: Deploy",
		Branches: []string{"master"},
	})
	steps = append(steps, CommandStep{
		Label:    ":kubernetes: Deploy",
		Command:  ".buildkite/deploy.sh",
		Branches: []string{"master"},
	})

	pipeline := Pipeline{
		Steps: steps,
	}
	d, _ := yaml.Marshal(&pipeline)
	fmt.Println(string(d))
}
