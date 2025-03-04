package entraid

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-entraid/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableEntraIdGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_group",
		Description: "Represents an Azure AD group.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Invalid filter clause"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "display_name", Require: plugin.Optional},
				{Name: "mail", Require: plugin.Optional},
				{Name: "mail_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "on_premises_sync_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "security_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.", Transform: transform.FromField("Description.DisplayName")},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the group.", Transform: transform.FromField("Description.ID")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description for the group.", Transform: transform.FromField("Description.Description")},

			{Name: "classification", Type: proto.ColumnType_STRING, Description: "Describes a classification for the group (such as low, medium or high business impact).", Transform: transform.FromField("Description.Classification")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time at which the group was created.", Transform: transform.FromField("Description.CreatedDateTime")},
			{Name: "expiration_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group is set to expire.", Transform: transform.FromField("Description.ExpirationDateTime")},
			{Name: "is_assignable_to_role", Type: proto.ColumnType_BOOL, Description: "Indicates whether this group can be assigned to an Azure Active Directory role or not.", Transform: transform.FromField("Description.IsAssignableToRole")},
			{Name: "is_subscribed_by_mail", Type: proto.ColumnType_BOOL, Description: "Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true.", Transform: transform.FromField("Description.IsAssignableToRole")},
			{Name: "mail", Type: proto.ColumnType_STRING, Description: "The SMTP address for the group, for example, \"serviceadmins@contoso.onmicrosoft.com\".", Transform: transform.FromField("Description.Mail")},
			{Name: "mail_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is mail-enabled.", Transform: transform.FromField("Description.MailEnabled")},
			{Name: "mail_nickname", Type: proto.ColumnType_STRING, Description: "The mail alias for the user.", Transform: transform.FromField("Description.MailNickname")},
			{Name: "membership_rule", Type: proto.ColumnType_STRING, Description: "The mail alias for the group, unique in the organization.", Transform: transform.FromField("Description.MembershipRule")},
			{Name: "membership_rule_processing_state", Type: proto.ColumnType_STRING, Description: "Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused.", Transform: transform.FromField("Description.MembershipRuleProcessingState")},
			{Name: "on_premises_domain_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises Domain name synchronized from the on-premises directory.", Transform: transform.FromField("Description.OnPremisesDomainName")},
			{Name: "on_premises_last_sync_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Indicates the last time at which the group was synced with the on-premises directory.", Transform: transform.FromField("Description.OnPremisesLastSyncDateTime")},
			{Name: "on_premises_net_bios_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises NetBiosName synchronized from the on-premises directory.", Transform: transform.FromField("Description.OnPremisesNetBiosName")},
			{Name: "on_premises_sam_account_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises SAM account name synchronized from the on-premises directory.", Transform: transform.FromField("Description.OnPremisesSamAccountName")},
			{Name: "on_premises_security_identifier", Type: proto.ColumnType_STRING, Description: "Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud.", Transform: transform.FromField("Description.OnPremisesSecurityIdentifier")},
			{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "True if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default).", Transform: transform.FromField("Description.OnPremisesSyncEnabled")},
			{Name: "renewed_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action.", Transform: transform.FromField("Description.RenewedDateTime")},
			{Name: "security_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is a security group.", Transform: transform.FromField("Description.SecurityEnabled")},
			{Name: "security_identifier", Type: proto.ColumnType_STRING, Description: "Security identifier of the group, used in Windows scenarios.", Transform: transform.FromField("Description.SecurityIdentifier")},
			{Name: "visibility", Type: proto.ColumnType_STRING, Description: "Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership.", Transform: transform.FromField("Description.Visibility")},
			{Name: "assigned_labels", Type: proto.ColumnType_JSON, Description: "The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group.", Transform: transform.FromField("Description.AssignedLabels")},
			{Name: "group_types", Type: proto.ColumnType_JSON, Description: "Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see [groups overview](https://docs.microsoft.com/en-us/graph/api/resources/groups-overview?view=graph-rest-1.0).", Transform: transform.FromField("Description.GroupTypes")},
			{Name: "member_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.MemberIDs"), Description: "Id of Users and groups that are members of this group."},
			{Name: "owner_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.OwnerIDs"), Description: "Id od the owners of the group. The owners are a set of non-admin users who are allowed to modify this object."},
			{Name: "proxy_addresses", Type: proto.ColumnType_JSON, Description: "Email addresses for the group that direct to the same group mailbox. For example: [\"SMTP: bob@contoso.com\", \"smtp: bob@sales.contoso.com\"]. The any operator is required to filter expressions on multi-valued properties.", Transform: transform.FromField("Description.ProxyAddresses")},
			{Name: "resource_behavior_options", Type: proto.ColumnType_JSON, Description: "Specifies the group behaviors that can be set for a Microsoft 365 group during creation. Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled.", Transform: transform.FromField("Description.ResourceBehaviorOptions")},
			{Name: "resource_provisioning_options", Type: proto.ColumnType_JSON, Description: "Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team.", Transform: transform.FromField("Description.ResourceProvisioningOptions")},
			{Name: "nested_groups", Type: proto.ColumnType_JSON, Description: "Members which are group.", Transform: transform.FromField("Description.NestedGroups")},

			{Name: "tags", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTags, Transform: transform.From(adGroupTags)},
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.From(adGroupTitle)},
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

func adGroupTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	group := d.HydrateItem.(opengovernance.AdGroup).Description

	if group.AssignedLabels == nil {
		return nil, nil
	}

	assignedLabels := group.AssignedLabels
	if len(assignedLabels) == 0 {
		return nil, nil
	}

	var tags = map[*string]*string{}
	for _, i := range assignedLabels {
		tags[i.LabelId] = i.DisplayName
	}

	return tags, nil
}

func adGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AdGroup).Description

	title := data.DisplayName
	//if title == nil {
	//title = data.DirectoryObject.GetId()
	//}

	return title, nil
}
