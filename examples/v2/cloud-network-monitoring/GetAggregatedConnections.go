// Get all aggregated connections returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetAggregatedConnections", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudNetworkMonitoringApi(apiClient)
	resp, r, err := api.GetAggregatedConnections(ctx, *datadogV2.NewGetAggregatedConnectionsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNetworkMonitoringApi.GetAggregatedConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudNetworkMonitoringApi.GetAggregatedConnections`:\n%s\n", responseContent)
}
