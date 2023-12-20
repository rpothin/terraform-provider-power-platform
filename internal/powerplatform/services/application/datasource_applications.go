package powerplatform

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/clients"
)

var (
	_ datasource.DataSource              = &ApplicationsDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationsDataSource{}
)

func NewApplicationsDataSource() datasource.DataSource {
	return &ApplicationsDataSource{
		ProviderTypeName: "powerplatform",
		TypeName:         "_applications",
	}
}

type ApplicationsDataSource struct {
	ApplicationClient ApplicationClient
	ProviderTypeName  string
	TypeName          string
}

type ApplicationsListDataSourceModel struct {
	EnvironmentId types.String                 `tfsdk:"environment_id"`
	Name          types.String                 `tfsdk:"name"`
	PublisherName types.String                 `tfsdk:"publisher_name"`
	Id            types.String                 `tfsdk:"id"`
	Applications  []ApplicationDataSourceModel `tfsdk:"applications"`
}

type ApplicationDataSourceModel struct {
	ApplicationId         types.String `tfsdk:"application_id"`
	Name                  types.String `tfsdk:"application_name"`
	UniqueName            types.String `tfsdk:"unique_name"`
	Version               types.String `tfsdk:"version"`
	Description           types.String `tfsdk:"description"`
	PublisherId           types.String `tfsdk:"publisher_id"`
	PublisherName         types.String `tfsdk:"publisher_name"`
	LearnMoreUrl          types.String `tfsdk:"learn_more_url"`
	State                 types.String `tfsdk:"state"`
	ApplicationVisibility types.String `tfsdk:"application_visibility"`
}

func (d *ApplicationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + d.TypeName
}

func (d *ApplicationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Fetches the list of Dynamics 365 applications in a tenant",
		MarkdownDescription: "Fetches the list of Dynamics 365 applications in a tenant",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Id of the read operation",
				Optional:    true,
			},
			"environment_id": schema.StringAttribute{
				Description: "Id of the Dynamics 365 environment",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Dynamics 365 application",
				Optional:    true,
			},
			"publisher_name": schema.StringAttribute{
				Description: "Publisher Name of the Dynamics 365 application",
				Optional:    true,
			},
			"applications": schema.ListNestedAttribute{
				Description:         "List of Applications",
				MarkdownDescription: "List of Applications",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"application_id": schema.StringAttribute{
							MarkdownDescription: "ApplicaitonId",
							Description:         "ApplicaitonId",
							Computed:            true,
						},
						"application_name": schema.StringAttribute{
							MarkdownDescription: "Name",
							Description:         "Name",
							Computed:            true,
						},
						"unique_name": schema.StringAttribute{
							MarkdownDescription: "Unique Name",
							Description:         "Unique Name",
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "Version",
							Description:         "Version",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Localized Description",
							Description:         "Localized Description",
							Computed:            true,
						},
						"publisher_id": schema.StringAttribute{
							MarkdownDescription: "Publisher Id",
							Description:         "Publisher Id",
							Computed:            true,
						},
						"publisher_name": schema.StringAttribute{
							MarkdownDescription: "Publisher Name",
							Description:         "Publisher Name",
							Computed:            true,
						},
						"learn_more_url": schema.StringAttribute{
							MarkdownDescription: "Learn More Url",
							Description:         "Learn More Url",
							Computed:            true,
						},
						"state": schema.StringAttribute{
							MarkdownDescription: "State",
							Description:         "State",
							Computed:            true,
						},
						"application_visibility": schema.StringAttribute{
							MarkdownDescription: "Application Visibility",
							Description:         "Application Visibility",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplicationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	clientBapi := req.ProviderData.(*clients.ProviderClient).PowerPlatformApi.Client
	if clientBapi == nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	d.ApplicationClient = NewApplicationClient(clientBapi)
}

func (d *ApplicationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var plan ApplicationsListDataSourceModel
	resp.State.Get(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("READ DATASOURCE APPLICATIONS START: %s", d.ProviderTypeName))

	plan.Id = types.StringValue(strconv.FormatInt(time.Now().Unix(), 10))
	plan.EnvironmentId = types.StringValue(plan.EnvironmentId.ValueString())
	plan.Name = types.StringValue(plan.Name.ValueString())
	plan.PublisherName = types.StringValue(plan.PublisherName.ValueString())

	applications, err := d.ApplicationClient.GetApplicationsByEnvironmentId(ctx, plan.EnvironmentId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Client error when reading %s", d.ProviderTypeName), err.Error())
		return
	}

	for _, application := range applications {
		if (plan.Name.ValueString() != "" && plan.Name.ValueString() != application.Name) ||
			(plan.PublisherName.ValueString() != "" && plan.PublisherName.ValueString() != application.PublisherName) {
			continue
		}
		plan.Applications = append(plan.Applications, ApplicationDataSourceModel{
			ApplicationId:         types.StringValue(application.ApplicationId),
			Name:                  types.StringValue(application.Name),
			UniqueName:            types.StringValue(application.UniqueName),
			Version:               types.StringValue(application.Version),
			Description:           types.StringValue(application.Description),
			PublisherId:           types.StringValue(application.PublisherId),
			PublisherName:         types.StringValue(application.PublisherName),
			LearnMoreUrl:          types.StringValue(application.LearnMoreUrl),
			State:                 types.StringValue(application.State),
			ApplicationVisibility: types.StringValue(application.ApplicationVisibility),
		})
	}

	diags := resp.State.Set(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("READ DATASOURCE APPLICATIONS END: %s", d.ProviderTypeName))

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}