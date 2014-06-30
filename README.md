Gatekeeper
==========

Gatekeeper was born out of a need for a tiny service to authenticate web based deployment requests from services like circle-ci and tavis-ci.

I built it for myself, but I hope you find it useful as well.

## Security
Currently uses simple querystring auth_key over ssl (if configured)...
This is not meant to be Fort Knox, if you need a secure build server consider not using a SaaS based CI.

## Usage
Gatekeeper is intended to be used behind another web server like nginx which can proxy pass to 127.0.0.1:8000

## Example
This is an example configuration file for gatekeeper.. pretty simple.

```toml
# Example configuration for gatekeeper
[server]
bind_address = "127.0.0.1"
port = 8000
build_uri = "/build"
protocol = "https"

[applications]
    [applications.one]
        name = "Test Application 1"
        root = "/opt/gatekeeper/workspace/app1"
        api_key = "woooapikeeeeeey"
        build_steps = [
            "fab deploy"
        ]
```

## Roadmap (todo?)
I would like to clean this up a bit, it was a weekend project to scratch an itch.  If more people like this I will make great efforts in improving it.

## License
Gatekeeper is licensed under the new bsd license, see LICENSE.md.
