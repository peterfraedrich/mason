package build

//GenericStep interface
type GenericStep interface {
	Results() (StepResults, error)
	Do() error
}

type step struct {
	Count      int    `yaml:"count"`
	Retry      int    `yaml:"retry"`
	Register   string `yaml:"register"`
	When       string `yaml:"when"`
	IgnoreFail bool   `yaml:"ignore_fail"`
	Creates    string `yaml:"creates"`
	results    StepResults
}

func (t *step) Results() (StepResults, error) {
	return t.results, nil
}

//Template type
type Template struct {
	step
	Src   string `yaml:"src"`
	Dest  string `yaml:"dest"`
	Force bool   `yaml:"force" required:"false"`
}

//Do the thing
func (s *Template) Do() error {
	return nil
}

//HTTPGet type
type HTTPGet struct {
	step
	URL       string              `yaml:"url" required:"true"`
	Status    int                 `yaml:"status"`
	Dest      string              `yaml:"dest"`
	Headers   []map[string]string `yaml:"headers"`
	IgnoreSSL bool                `yaml:"ignore_ssl"`
}

func (s *HTTPGet) Do() error {
	return nil
}

//HTTPPost type
type HTTPPost struct {
	step
	URL       string              `yaml:"url" required:"true"`
	Payload   string              `yaml:"payload" required:"true"`
	Status    int                 `yaml:"status"`
	Headers   []map[string]string `yaml:"headers"`
	IgnoreSSL bool                `yaml:"ignore_ssl"`
}

//Do the thing
func (s *HTTPPost) Do() error {
	return nil
}

//Run runs a shell command
type Run struct {
	step
	CMD    string `yaml:"cmd" required:"true"`
	Shell  string `yaml:"shell"`
	AsUser string `yaml:"as_user"`
}

//Do the thing
func (s *Run) Do() error {
	return nil
}

//Package func
type Package struct {
	step
	Name    string `yaml:"name"`
	Version string `yaml:"version" required:"false"`
}

//Do the thing
func (s *Package) Do() error {
	return nil
}
