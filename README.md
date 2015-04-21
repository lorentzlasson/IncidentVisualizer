# About
Visualizes incidents sent via IBM Internet of Things Foundation

Currently running on http://incident-visualizer.mybluemix.net/

# Getting started
Install go.

Add iotfcreds.json to in the root dir. It should have the following format:

```json
{
	"Org": "<org_id>",
	"Username": "<api-key>",
	"Password": "<auth-token>"
}
```

Run `go run *.go`

# Deploy to Bluemix
Install [cf cli](https://github.com/cloudfoundry/cli#downloads).

Modify the `host` attribute in `manifest.yml (needs to be unique).

Run `cf push` from the root directory.



