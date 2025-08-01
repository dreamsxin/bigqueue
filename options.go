package bigqueue

import "errors"

// Options the options struct
type Options struct {

	// size in bytes of a data page
	DataPageSize int

	// size in bytes of a index page
	indexPageSize int

	// the item count is  1 << IndexItemsPerPage
	IndexItemsPerPage int

	itemsPerPage int

	// if value > 0 then enable auto gc features and repeat process gc by the specified interval time in seconds.
	AutoGCBySeconds int
}

// Validate checks if the options are valid
func (o *Options) Validate() error {
	// 新增：验证IndexItemsPerPage必须是2的幂
	if o.IndexItemsPerPage <= 0 || (o.IndexItemsPerPage&(o.IndexItemsPerPage-1)) != 0 {
		return errors.New("IndexItemsPerPage must be a positive power of 2")
	}

	// 计算派生值
	o.itemsPerPage = 1 << o.IndexItemsPerPage
	o.indexPageSize = defaultIndexItemLen * o.itemsPerPage

	return nil
}
