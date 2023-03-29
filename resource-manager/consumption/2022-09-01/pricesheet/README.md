
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-09-01/pricesheet` Documentation

The `pricesheet` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-09-01/pricesheet"
```


### Client Initialization

```go
client := pricesheet.NewPriceSheetClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheetClient.Get`

```go
ctx := context.TODO()
id := pricesheet.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.Get(ctx, id, pricesheet.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PriceSheetClient.GetByBillingPeriod`

```go
ctx := context.TODO()
id := pricesheet.NewBillingPeriodID("12345678-1234-9876-4563-123456789012", "billingPeriodValue")

read, err := client.GetByBillingPeriod(ctx, id, pricesheet.DefaultGetByBillingPeriodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```