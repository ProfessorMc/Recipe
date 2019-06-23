package recipe

import (
	"sync"
)

type SyncCookBook struct {
	*CookBook
	mux sync.Mutex
}

func NewSyncCookBook(name string) *SyncCookBook {
	return &SyncCookBook{CookBook: NewCookBook(name)}
}

func (c *SyncCookBook) AddRecipe(recipe *Recipe) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.CookBook.AddRecipe(recipe)
}

func (c *SyncCookBook) DeleteRecipe(name string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.CookBook.DeleteRecipe(name)
}

func (c *SyncCookBook) GetRecipe(name string) (*Recipe, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.CookBook.GetRecipe(name)
}

func (c *SyncCookBook) PrintAll() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.CookBook.PrintAll()
}
