package build

//Blueprint type
type Blueprint struct {
	Name string `yaml:"name"`

	// optional keys
	DefaultShell string `yaml:"default_shell"`
	Parallel     bool   `yaml:"parallel"`

	Env  []map[string]string `yaml:"env"`
	Vars []map[string]string `yaml:"vars"`

	Steps []Step `yaml:"steps"`
}

//Step type
type Step struct {
	GenericStep
	Name string `yaml:"name"`
}

type schemaField struct {
	Name         string
	Value        interface{}
	DefaultValue interface{}
	ValueType    string
	Required     bool
	NullValue    interface{}
}

//StepResults type
type StepResults struct {
}
