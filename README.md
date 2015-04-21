# About
Visualizes incidents sent via IBM Internet of Things Foundation

# Getting started
Install go

Add iotfcreds.json to in the root dir. It should have the following format:

```json
{
	"Org": "<org_id>",
	"Username": "<api-key>",
	"Password": "<auth-token>"
}
```

Run `go run *.go`
