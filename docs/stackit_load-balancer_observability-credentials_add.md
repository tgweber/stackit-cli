## stackit load-balancer observability-credentials add

Adds observability credentials to Load Balancer

### Synopsis

Adds existing observability credentials (username and password) to Load Balancer. The credentials can be for Observability or another monitoring tool.

```
stackit load-balancer observability-credentials add [flags]
```

### Examples

```
  Add observability credentials to a load balancer with username "xxx" and display name "yyy". The password is entered using the terminal
  $ stackit load-balancer observability-credentials add --username xxx --display-name yyy

  Add observability credentials to a load balancer with username "xxx" and display name "yyy", providing the path to a file with the password as flag
  $ stackit load-balancer observability-credentials add --username xxx --password @./password.txt --display-name yyy
```

### Options

```
      --display-name string   Credentials display name
  -h, --help                  Help for "stackit load-balancer observability-credentials add"
      --password string       Password. Can be a string or a file path, if prefixed with "@" (example: @./password.txt).
      --username string       Username
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

* [stackit load-balancer observability-credentials](./stackit_load-balancer_observability-credentials.md)	 - Provides functionality for Load Balancer observability credentials

