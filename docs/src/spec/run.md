# Run

Specs containing a `run` field can be executed using [Shims](shim.md).

## Alias

Alias bridges a command while retaining the argument completion.

```yaml
name: log-by-author
run: "[git, log, --author]"
```

![](./run-alias.cast)

## Script

Script macro is executed with `sh` on unix systems and `pwsh` on windows.
Flags are used for [environment substitution](https://github.com/drone/envsubst) and positional arguments are passed to the script.

```yaml
name: ls-remote
run: "$(git ls-remote --sort='${C_FLAG_SORT:-HEAD}' $@)"
flags:
  --sort=: field name to sort on
completion:
  flag:
    sort: [version:refname, authordate]
  positional:
    - ["$_tools.git.RepositorySearch"]
  positionalany: ["$_tools.git.LsRemoteRefs({url: '${C_ARG0}', branches: true, tags: true})"]
```

![](./run-script.cast)