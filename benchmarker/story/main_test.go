package story

import (
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	story, err := NewStory("localhost:9292")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	err = story.Main(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(story.stdout.String())
	t.Log(story.errors.GetMessages())
}
