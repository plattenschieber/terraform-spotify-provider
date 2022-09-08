package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.DataSourceType = spotifyTrackDataSourceType{}
var _ datasource.DataSource = spotifyTrackDataSource{}

type spotifyTrackDataSourceType struct{}

func (t spotifyTrackDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Spotify track data source",

		Attributes: map[string]tfsdk.Attribute{
			"title": {
				MarkdownDescription: "Track title",
				Required:            true,
				Type:                types.StringType,
			},
			"artist": {
				MarkdownDescription: "Track artist",
				Required:            true,
				Type:                types.StringType,
			},
			"uri": {
				MarkdownDescription: "Track URI",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t spotifyTrackDataSourceType) NewDataSource(ctx context.Context, in provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return spotifyTrackDataSource{
		provider: provider,
	}, diags
}

type spotifyTrackDataSourceData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
}

type spotifyTrackDataSource struct {
	provider spotifyProvider
}

func (d spotifyTrackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data spotifyTrackDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.ReadExample(...)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.String{Value: "example-id"}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
