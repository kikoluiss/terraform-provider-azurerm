
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2021-09-03-preview/desktop` Documentation

The `desktop` SDK allows for interaction with the Azure Resource Manager Service `desktopvirtualization` (API Version `2021-09-03-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2021-09-03-preview/desktop"
```


### Client Initialization

```go
client := desktop.NewDesktopClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DesktopClient.Get`

```go
ctx := context.TODO()
id := desktop.NewDesktopID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGroupValue", "desktopValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DesktopClient.List`

```go
ctx := context.TODO()
id := desktop.NewApplicationGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGroupValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DesktopClient.Update`

```go
ctx := context.TODO()
id := desktop.NewDesktopID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGroupValue", "desktopValue")

payload := desktop.DesktopPatch{
	// ...
}

read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
