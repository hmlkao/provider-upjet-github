package actionsorganizationpermissions

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure github_actions_organization_permissions resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_actions_organization_permissions", func(r *config.Resource) {

		r.ShortGroup = "actions"
		// repository_ids are numeric IDs, keep them as plain ints
		delete(r.References, "enabled_repositories_config.repository_ids")
	})
}
