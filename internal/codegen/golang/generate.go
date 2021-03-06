package golang

import (
	_ "embed"
	"fmt"
	"github.com/jschaf/pggen/internal/casing"
	"github.com/jschaf/pggen/internal/codegen"
	"path/filepath"
	"text/template"
)

// GenerateOptions are options to control generated Go output.
type GenerateOptions struct {
	GoPkg     string
	OutputDir string
	// A map of lowercase acronyms to the upper case equivalent, like:
	// "api" => "API".
	Acronyms map[string]string
	// A map from a Postgres type name to a fully qualified Go type.
	TypeOverrides map[string]string
}

// Generate emits generated Go files for each of the queryFiles.
func Generate(opts GenerateOptions, queryFiles []codegen.QueryFile) error {
	pkgName := opts.GoPkg
	if pkgName == "" {
		pkgName = filepath.Base(opts.OutputDir)
	}
	caser := casing.NewCaser()
	caser.AddAcronyms(opts.Acronyms)
	templater := NewTemplater(TemplaterOpts{
		Caser:    caser,
		Resolver: NewTypeResolver(caser, opts.TypeOverrides),
		Pkg:      pkgName,
	})
	templatedFiles, err := templater.TemplateAll(queryFiles)
	if err != nil {
		return fmt.Errorf("template all: %w", err)
	}

	tmpl, err := parseQueryTemplate()
	if err != nil {
		return fmt.Errorf("parse generated Go code template: %w", err)
	}
	emitter := NewEmitter(opts.OutputDir, tmpl)
	for _, qf := range templatedFiles {
		if err := emitter.EmitQueryFile(qf); err != nil {
			return fmt.Errorf("emit generated Go code: %w", err)
		}
	}
	return nil
}

//go:embed query.gotemplate
var queryTemplate string

func parseQueryTemplate() (*template.Template, error) {
	tmpl, err := template.New("gen_query").Parse(queryTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse query.gotemplate: %w", err)
	}
	return tmpl, nil
}
