package azuread

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureAdDirectoryAuditReport(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_directory_audit_report",
		Description: "Represents the list of audit logs generated by Azure Active Directory.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdDirectoryAuditReport,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdDirectoryAuditReport,
		},

		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates the unique ID for the activity.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "activity_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Indicates the date and time the activity was performed.",
				Transform:   transform.FromField("Description.ActivityDateTime")},
			{
				Name:        "activity_display_name",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates the activity name or the operation name.",
				Transform:   transform.FromField("Description.ActivityDisplayName")},
			{
				Name:        "category",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates which resource category that's targeted by the activity.",
				Transform:   transform.FromField("Description.Category")},
			{
				Name:        "correlation_id",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.",
				Transform:   transform.FromField("Description.CorrelationId")},
			{
				Name:        "logged_by_service",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.",
				Transform:   transform.FromField("Description.LoggedByService")},
			{
				Name:        "operation_type",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates the type of operation that was performed. The possible values include but are not limited to the following: Add, Assign, Update, Unassign, and Delete.",
				Transform:   transform.FromField("Description.OperationType")},
			{
				Name:        "result",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.",
				Transform:   transform.FromField("Description.Result")},
			{
				Name:        "result_reason",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates the reason for failure if the result is failure or timeout.",
				Transform:   transform.FromField("Description.ResultReason")},

			// JSON fields
			{
				Name:        "additional_details",
				Type:        proto.ColumnType_JSON,
				Description: "Indicates additional details on the activity.",
				Transform:   transform.FromField("Description.AdditionalDetails")},
			{
				Name:        "initiated_by",
				Type:        proto.ColumnType_JSON,
				Description: "Indicates information about the user or app initiated the activity.",
				Transform:   transform.FromField("Description.InitiatedBy")},
			{
				Name:        "target_resources",
				Type:        proto.ColumnType_JSON,
				Description: "Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.",
				Transform:   transform.FromField("Description.TargetResources")},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
