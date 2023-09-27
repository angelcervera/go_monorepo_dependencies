module github.com/angelcervera/go_monorepo_dependencies/cli_app

go 1.21.0

replace github.com/angelcervera/go_monorepo_dependencies/dep_one => ../dep_one

replace github.com/angelcervera/go_monorepo_dependencies/dep_two => ../dep_two

require (
	github.com/angelcervera/go_monorepo_dependencies/dep_one v0.0.0-00010101000000-000000000000
	github.com/angelcervera/go_monorepo_dependencies/dep_two v0.0.0-00010101000000-000000000000
)
