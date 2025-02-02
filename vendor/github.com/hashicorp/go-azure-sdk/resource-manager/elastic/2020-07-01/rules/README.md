
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/rules` Documentation

The `rules` SDK allows for interaction with the Azure Resource Manager Service `elastic` (API Version `2020-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/rules"
```


### Client Initialization

```go
client := rules.NewRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `RulesClient.TagRulesCreateOrUpdate`

```go
ctx := context.TODO()
id := rules.NewTagRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue", "ruleSetValue")

payload := rules.MonitoringTagRules{
	// ...
}

read, err := client.TagRulesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RulesClient.TagRulesDelete`

```go
ctx := context.TODO()
id := rules.NewTagRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue", "ruleSetValue")
future, err := client.TagRulesDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `RulesClient.TagRulesGet`

```go
ctx := context.TODO()
id := rules.NewTagRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue", "ruleSetValue")
read, err := client.TagRulesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RulesClient.TagRulesList`

```go
ctx := context.TODO()
id := rules.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")
// alternatively `client.TagRulesList(ctx, id)` can be used to do batched pagination
items, err := client.TagRulesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
