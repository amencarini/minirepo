package differ

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
	// ArtifactPaths string `yaml:"artifact_paths"`
	// Env           struct {
	// 	BUILDKITEDOCKERCOMPOSECONTAINER string `yaml:"BUILDKITE_DOCKER_COMPOSE_CONTAINER"`
	// } `yaml:"env"`
}

func CalculateDiffingFolders() {
	steps := []Step{}
	steps = append(steps, Step{
		Label:   ":hammer: Tests",
		Command: "echo 'Teeests'",
	})
	pipeline := Pipeline{
		Steps: steps,
	}
	d, _ := yaml.Marshal(&pipeline)
	fmt.Println(string(d))
}
