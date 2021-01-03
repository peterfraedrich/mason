package build

type step struct {
	results StepResults
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
	URL    string `yaml:"url"`
	Status int    `yaml:"status"`
	Dest   string `yaml:"dest" required:"false"`
}

func (s *HTTPGet) do() error {
	return nil
}
