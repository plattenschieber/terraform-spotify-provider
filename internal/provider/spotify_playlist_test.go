package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPlaylistResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExampleResourceConfig("one"),
				Check:  resource.ComposeAggregateTestCheckFunc(
				// resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "one"),
				// resource.TestCheckResourceAttr("scaffolding_example.test", "id", "example-id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccExampleResourceConfig(configurableAttribute string) string {
	return `
	resource "spotify_playlist" "test" {
	name = "test name"
	description = "some description"
	public = false
}`
}
