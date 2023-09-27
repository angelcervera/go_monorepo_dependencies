## Structure

```text
dep-one
dep-two
dep-transitive
function
```

## References

- [Go Workspaces](https://go.dev/doc/tutorial/workspaces)
- [GoogleCloudPlatform/buildpacks](https://github.com/GoogleCloudPlatform/buildpacks)
- [Google Cloud's buildpacks](https://cloud.google.com/docs/buildpacks/overview)
- [Google Cloud's buildpacks builders](https://cloud.google.com/docs/buildpacks/builders)

## Commands

| Command                                                                                    | Description                                                                                                                   |
|--------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| `go mod init github.com/angelcervera/go_monorepo_dependencies/cli-app`                     | init a new module                                                                                                             |
| `go mod edit -replace github.com/angelcervera/go_monorepo_dependencies/dep_one=../dep_one` | add replace to allow access to the dependency                                                                                 |
| `go mod tidy`                                                                              | After adding an import in the `.go` file, this command will add required dependencies in the `go.mod`, with the right version |
|                                                                                            |                                                                                                                               |
|                                                                                            |                                                                                                                               |

## Notes
- **pseudo-version number** : It is a [temporal version](https://go.dev/doc/modules/version-numbers#pseudo-version-number) that `go mod tidy` will assign if the shared modules are not versioned.  
