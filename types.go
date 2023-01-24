package sink

import (
	"context"

	"github.com/streamingfast/bstream"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
)

type BlockScopeDataHandler = func(ctx context.Context, cursor *Cursor, data *pbsubstreams.BlockScopedData) error

type Cursor struct {
	Cursor string
	Block  bstream.BlockRef
}

func NewCursor(cursor string, block bstream.BlockRef) *Cursor {
	return &Cursor{cursor, block}
}

func NewBlankCursor() *Cursor {
	return NewCursor("", bstream.BlockRefEmpty)
}

func (c *Cursor) IsBlank() bool {
	return c.Cursor == ""
}

func (c *Cursor) IsEqualTo(other *Cursor) bool {
	if c == nil && other == nil {
		return true
	}

	// We know both are not equal, so if either side is `nil`, we are sure the other is not, so not equal
	if c == nil || other == nil {
		return false
	}

	// Both side are non-nil here
	return c.Cursor == other.Cursor
}

func (c *Cursor) String() string {
	if c == nil {
		return "<Unset>"
	}

	return c.Cursor
}

//go:generate go-enum -f=$GOFILE --marshal --names

// ENUM(
//
//	Development
//	Production
//
// )
type SubstreamsMode uint
