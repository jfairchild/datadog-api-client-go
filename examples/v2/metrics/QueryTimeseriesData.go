// Query timeseries data across multiple products returns "OK" response

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
	body := datadogV2.TimeseriesFormulaQueryRequest{
		Data: datadogV2.TimeseriesFormulaRequest{
			Attributes: datadogV2.TimeseriesFormulaRequestAttributes{
				Formulas: []datadogV2.QueryFormula{
					{
						Formula: "a+b",
						Limit: &datadogV2.FormulaLimit{
							Count: datadog.PtrInt32(10),
							Order: datadogV2.QUERYSORTORDER_DESC.Ptr(),
						},
					},
				},
				From:     1568899800000,
				Interval: datadog.PtrInt64(5000),
				Queries: []datadogV2.TimeseriesQuery{
					datadogV2.TimeseriesQuery{
						MetricsTimeseriesQuery: &datadogV2.MetricsTimeseriesQuery{
							DataSource: datadogV2.METRICSDATASOURCE_METRICS,
							Query:      "avg:system.cpu.user{*} by {env}",
						}},
				},
				To: 1568923200000,
			},
			Type: datadogV2.TIMESERIESFORMULAREQUESTTYPE_TIMESERIES_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.QueryTimeseriesData(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryTimeseriesData`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.QueryTimeseriesData`:\n%s\n", responseContent)
}
