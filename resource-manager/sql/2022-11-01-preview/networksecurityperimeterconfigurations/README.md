
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/networksecurityperimeterconfigurations` Documentation

The `networksecurityperimeterconfigurations` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/networksecurityperimeterconfigurations"
```


### Client Initialization

```go
client := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.Get`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "networkSecurityPerimeterConfigurationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.ListByServer`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.Reconcile`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "networkSecurityPerimeterConfigurationValue")

if err := client.ReconcileThenPoll(ctx, id); err != nil {
	// handle the error
}
```
