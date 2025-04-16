package entraid

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-entraid/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableEntraIdGroupMembership(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_group_membership",
		Description: "Represents an Azure AD group membership.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdGroupMembership,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdGroupMembership,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Invalid filter clause"}),
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.",
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the user.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The group id.",
				Transform:   transform.FromField("Description.GroupId")},

			{
				Name:        "account_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "is the account enabled or not",
				Transform:   transform.FromField("Description.AccountEnabled")},
			{
				Name:        "user_principal_name",
				Type:        proto.ColumnType_STRING,
				Description: "user principal name.",
				Transform:   transform.FromField("Description.UserPrincipalName")},
			{
				Name:        "member_type",
				Type:        proto.ColumnType_STRING,
				Description: "user type.",
				Transform:   transform.FromField("Description.UserType")},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "state.",
				Transform:   transform.FromField("Description.State")},
			{
				Name:        "security_identifier",
				Type:        proto.ColumnType_STRING,
				Description: "state.",
				Transform:   transform.FromField("Description.SecurityIdentifier")},
			{
				Name:        "proxy_addresses",
				Type:        proto.ColumnType_STRING,
				Description: "user proxy addresses.",
				Transform:   transform.FromField("Description.ProxyAddresses")},
			{
				Name:        "mail",
				Type:        proto.ColumnType_STRING,
				Description: "user email.",
				Transform:   transform.FromField("Description.Mail")},

			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.DisplayName")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromField("Description.TenantID")},
			{
				Name:        "metadata",
				Description: "Metadata of the Azure resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata").Transform(marshalJSON),
			},
			{
				Name:        "platform_account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Platform Account ID in which the resource is located.",
				Transform:   transform.FromField("Metadata.SourceID")},
			{
				Name:        "platform_resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the resource in opengovernance.",
				Transform:   transform.FromField("ID")},
		}),
	}
}
