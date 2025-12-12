package actions

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure github_actions_secret resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_actions_runner_group", func(r *config.Resource) {

		r.ShortGroup = "actions"
		// Runner group name/selected_repository_ids should remain plain values (IDs),
		// not Crossplane references. Drop auto-injected refs to avoid type mismatch.
		delete(r.References, "name")
		delete(r.References, "selected_repository_ids")
	})
}
