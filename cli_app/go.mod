module github.com/angelcervera/go_monorepo_dependencies/cli_app

go 1.21.0

replace (
	github.com/angelcervera/go_monorepo_dependencies/dep_one => ../dep_one
	github.com/angelcervera/go_monorepo_dependencies/dep_two => ../dep_two
	github.com/angelcervera/go_monorepo_dependencies/dep_transitive => ../dep_transitive
)

require (
	github.com/angelcervera/go_monorepo_dependencies/dep_one v0.0.0-00010101000000-000000000000
	github.com/angelcervera/go_monorepo_dependencies/dep_two v0.0.0-00010101000000-000000000000
)

require (
	github.com/angelcervera/go_monorepo_dependencies/dep_transitive v0.0.0-00010101000000-000000000000 // indirect
	github.com/ozgio/strutil v0.4.0 // indirect
)
