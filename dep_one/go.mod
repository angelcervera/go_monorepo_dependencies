module github.com/angelcervera/go_monorepo_dependencies/dep_one

go 1.21.0

replace github.com/angelcervera/go_monorepo_dependencies/dep_transitive => ../dep_transitive

require github.com/angelcervera/go_monorepo_dependencies/dep_transitive v0.0.0-00010101000000-000000000000

require github.com/ozgio/strutil v0.4.0 // indirect
