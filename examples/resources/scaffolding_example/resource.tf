resource "spotify_playlist" "example" {
  name = "some-name"
  track {
    title = "some-title"
    album = "some-album" 
  }
  track {
    title = "another-title"
    album = "another-album" 
  }
}
