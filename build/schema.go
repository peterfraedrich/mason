package build

var schemaMap = map[string]schemaField{
	"Name": schemaField{
		Name:         "Name",
		Value:        nil,
		DefaultValue: "",
		ValueType:    "string",
		Required:     true,
		NullValue:    "",
	},
	"DefaultShell": schemaField{
		Name:         "DefaultShell",
		Value:        nil,
		DefaultValue: "/bin/sh",
		ValueType:    "string",
		Required:     false,
		NullValue:    "",
	},
	"Parallel": schemaField{
		Name:         "Parallel",
		Value:        nil,
		DefaultValue: false,
		ValueType:    "bool",
		Required:     false,
		NullValue:    false,
	},
	"Env": schemaField{
		Name:         "Env",
		Value:        nil,
		DefaultValue: []map[string]string{},
		ValueType:    "[]map[string]string",
		Required:     false,
		NullValue:    nil,
	},
	"Vars": schemaField{
		Name:         "Vars",
		Value:        nil,
		DefaultValue: []map[string]string{},
		ValueType:    "[]map[string]string",
		Required:     false,
		NullValue:    nil,
	},
	"Steps": schemaField{
		Name:         "Steps",
		Value:        nil,
		DefaultValue: []Step{},
		ValueType:    "[]build.Step",
		Required:     true,
		NullValue:    nil,
	},
}
