package project

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("gitlab_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
