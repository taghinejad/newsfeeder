package newsfeed

import (
	"errors"
)

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(Item)
	GetAll() []Item
}
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{Items: []Item{}}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}

func (r *Repo) GetByTitle(title string) (Item, error) {
	for i, j := range r.Items {
		if j.Title == title {
			return r.Items[i], nil
		}
	}
	var t Item
	return t, errors.New("not found")
}
