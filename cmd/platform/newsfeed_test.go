package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	feed.Add(Item{})
	result := feed.GetAll()

	if len(result) < 1 {
		t.Errorf("not get all")
	}

}
