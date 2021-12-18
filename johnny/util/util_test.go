package util

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {

	checkComparisons := func(t testing.TB, got, want bool) {
        t.Helper()
        if got != want {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("proper http-s3 audio url", func(t *testing.T) {
		audio_url := "https://lol.yes-much-fake.amazonaws.com/9b015b18-443c-42f6-adad-8309d9659813.flac"
		want := true
		got := isValidURL(audio_url)
        checkComparisons(t, got, want)
    })


    t.Run("s3 bucket uri", func(t *testing.T) {
		audio_url := "s3://skit-fake-bucket/9b015b18-443c-42f6-adad-8309d9659813.flac"
		want := false
		got := isValidURL(audio_url)
        checkComparisons(t, got, want)
    })

    t.Run("s3 relative path uri", func(t *testing.T) {
		audio_url := "/skit-fake-bucket/9b015b18-443c-42f6-adad-8309d9659813.flac"
		want := false
		got := isValidURL(audio_url)
        checkComparisons(t, got, want)
    })

    t.Run("empty quote", func(t *testing.T) {
		audio_url := ""
		want := false
		got := isValidURL(audio_url)
        checkComparisons(t, got, want)
    })


}
