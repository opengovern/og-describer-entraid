package describers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"

	"github.com/opengovern/og-util/pkg/describe/enums"

	hamiltonAuth "github.com/manicminer/hamilton/auth"
)

const SubscriptionBatchSize = 100

type GenericResourceGraph struct {
	Table string
	Type  string
}

func (d GenericResourceGraph) DescribeResources(ctx context.Context, cred *azidentity.ClientSecretCredential, _ hamiltonAuth.Authorizer, tempSubscriptions []string, tenantId string, triggerType enums.DescribeTriggerType, stream *models.StreamSender) ([]models.Resource, error) {
	ctx = WithTriggerType(ctx, triggerType)
	query := fmt.Sprintf("%s | where type == \"%s\"", d.Table, strings.ToLower(d.Type))

	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource

	var subscriptions []*string
	for _, subscription := range tempSubscriptions {
		subscriptions = append(subscriptions, &subscription)
	}

	// Group the subscriptions to batches with a max size
	for i := 0; i < len(subscriptions); i = i + SubscriptionBatchSize {
		j := i + SubscriptionBatchSize
		if j > len(subscriptions) {
			j = len(subscriptions)
		}

		resultFormat := armresourcegraph.ResultFormatObjectArray
		subs := subscriptions[i:j]
		request := armresourcegraph.QueryRequest{
			Subscriptions: subs,
			Query:         &query,
			Options: &armresourcegraph.QueryRequestOptions{
				ResultFormat: &resultFormat,
			},
		}

		// Fetch all resources by paging through all the results
		for first, skipToken := true, (*string)(nil); first || skipToken != nil; {
			request.Options.SkipToken = skipToken

			response, err := client.Resources(ctx, request, nil)
			if err != nil {
				return nil, err
			}

			// No Need to wait for quota
			//
			//quotaRemaining, untilResets, err := quota()
			//if err != nil {
			//	return nil, err
			//}
			//if quotaRemaining == 0 {
			//	time.Sleep(untilResets)
			//}

			for _, v := range response.Data.([]interface{}) {
				m := v.(map[string]interface{})
				loc := "global"
				if v, ok := m["location"]; ok {
					if vStr, ok := v.(string); ok {
						loc = vStr
					}
				}
				resource := models.Resource{
					ID:          m["id"].(string),
					Location:    loc,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

				values = append(values)
			}
			first, skipToken = false, response.SkipToken
		}
	}

	return values, nil
}

// quota parses the Azure throttling headers.
// See https://docs.microsoft.com/en-us/azure/governance/resource-graph/concepts/guidance-for-throttled-requests#understand-throttling-headers
func quota(header http.Header) (int, time.Duration, error) {
	remainingHeader := header[http.CanonicalHeaderKey("x-ms-user-quota-remaining")]
	if len(remainingHeader) == 0 {
		return 0, 0, errors.New("header 'x-ms-user-quota-remaining' missing")
	}

	remaining, err := strconv.Atoi(remainingHeader[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse 'x-ms-user-quota-remaining':  %w", err)
	}

	afterHeader := header[http.CanonicalHeaderKey("x-ms-user-quota-resets-after")]
	if len(afterHeader) == 0 {
		return 0, 0, errors.New("header 'x-ms-user-quota-resets-after' missing")
	}

	t, err := time.Parse("15:04:05", afterHeader[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse 'x-ms-user-quota-resets-after'")
	}

	t = t.UTC()
	after := time.Duration(t.Second())*time.Second +
		time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Hour())*time.Hour

	return remaining, after, nil
}
