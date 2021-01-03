package build

//GenericStep interface
type GenericStep interface {
	Results()
	do()
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

func (s *Template) do() error {
	return nil
}

//HTTPGet type
type HTTPGet struct {
	step
	URL       string              `yaml:"url"`
	Status    int                 `yaml:"status"`
	Dest      string              `yaml:"dest" required:"false"`
	Headers   []map[string]string `yaml:"headers"`
	IgnoreSSL bool                `yaml:"ignore_ssl"`
}

func (s *HTTPGet) do() error {
	return nil
}

//HTTPPost type
type HTTPPost struct {
	step
	URL       string              `yaml:"url"`
	Payload   string              `yaml:"payload"`
	Status    int                 `yaml:"status"`
	Headers   []map[string]string `yaml:"headers"`
	IgnoreSSL bool                `yaml:"ignore_ssl"`
}

func (s *HTTPPost) do() error {
	return nil
}

//Run runs a shell command
type Run struct {
	step
	CMD    string `yaml:"cmd"`
	Shell  string `yaml:"shell"`
	AsUser string `yaml:"as_user"`
}

func (s *Run) do() error {
	return nil
}
