package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/spotify"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.Provider = &spotifyProvider{}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type spotifyProvider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	// TODO: If appropriate, implement upstream provider SDK or HTTP client.
	// client vendorsdk.ExampleClient
	client *spotify.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Example types.String `tfsdk:"example"`
}

func (p *spotifyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	spotifyToken := os.Getenv("SPOTIFY_TOKEN")
	if spotifyToken == "" {
		resp.Diagnostics.AddError("Spotify token is missing", "The spotify token is blank, please ensure the env `SPOTIFY_TOKEN` is set correctly.")
	}
	userId := os.Getenv("SPOTIFY_USER_ID")
	if userId == "" {
		resp.Diagnostics.AddError("Spotify user id missing", "The spotify user id is blank, please ensure the env `SPOTIFY_USER_ID` is set correctly.")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	p.client = spotify.NewClient(spotifyToken, userId)
	p.configured = true
}

func (p *spotifyProvider) GetResources(ctx context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"spotify_playlist": spotifyPlaylistResourceType{},
	}, nil
}

func (p *spotifyProvider) GetDataSources(ctx context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{
		"spotify_example": spotifyTrackDataSourceType{},
	}, nil
}

func (p *spotifyProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"example": {
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &spotifyProvider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*scaffoldingProvider)), however using this can prevent
// potential panics.
func convertProviderType(in provider.Provider) (spotifyProvider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*spotifyProvider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return spotifyProvider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return spotifyProvider{}, diags
	}

	return *p, diags
}
