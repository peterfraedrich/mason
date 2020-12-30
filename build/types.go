package build

//Blueprint type
type Blueprint struct {
	Name string `yaml:"name"`

	// optional keys
	DefaultShell string `yaml:"default_shell"`
	Parallel     bool   `yaml:"parallel"`

	Env  map[string]string `yaml:"env"`
	Vars map[string]string `yaml:"vars"`

	Steps []BlueprintStep `yaml:"steps"`
}

//BlueprintStep type
type BlueprintStep struct {
	Name string `yaml:"name"`
}
