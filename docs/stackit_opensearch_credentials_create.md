## stackit opensearch credentials create

Creates credentials for an OpenSearch instance

### Synopsis

Creates credentials (username and password) for an OpenSearch instance.

```
stackit opensearch credentials create [flags]
```

### Examples

```
  Create credentials for an OpenSearch instance
  $ stackit opensearch credentials create --instance-id xxx

  Create credentials for an OpenSearch instance and show the password in the output
  $ stackit opensearch credentials create --instance-id xxx --show-password
```

### Options

```
  -h, --help                 Help for "stackit opensearch credentials create"
      --instance-id string   Instance ID
  -s, --show-password        Show password in output
```

### Options inherited from parent commands

```
  -y, --assume-yes             If set, skips all confirmation prompts
      --async                  If set, runs the command asynchronously
  -o, --output-format string   Output format, one of ["json" "pretty" "none" "yaml"]
  -p, --project-id string      Project ID
      --region string          Target region for region-specific requests
      --verbosity string       Verbosity of the CLI, one of ["debug" "info" "warning" "error"] (default "info")
```

### SEE ALSO

* [stackit opensearch credentials](./stackit_opensearch_credentials.md)	 - Provides functionality for OpenSearch credentials

