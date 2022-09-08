terraform {
    required_providers {
        spotify = {
            source = "local/plattenschieber/spotify"
            version = "0.1.0"
        }
    }
}

resource "spotify_example" "example" {
  configurable_attribute = "some-value"
}
