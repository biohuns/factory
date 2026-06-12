package building

import "github.com/biohuns/factory/pkg/entity"

// Accepter アイテムを受け取る構造物用の interface
type Accepter interface {
	Accept(item entity.Item) bool
}
