gatekeeper
==========

Golang based server for serving servers with served services.

## Example
This is an example configuration file for gatekeeper.. pretty simple.

```toml
# Example configuration for gatekeeper
[server]
bind_address = "0.0.0.0"
port = 8000
build_uri = "/build"

[applications]
    [applications.one]
        name = "Test Application 1"
        root = "/opt/gatekeeper/workspace/app1"
        build_steps = [
            "fab deploy"
        ]
```