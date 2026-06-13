package building

import (
	"github.com/biohuns/factory/pkg/entity"
)

type StorageInterface interface {
	Accept(item entity.Item) bool
	Item() entity.Item
	Amount() int
	Clear()
	Tick(tick int)
}

type Storage struct {
	// item 保管しているアイテム
	item entity.Item

	// amount 保管量
	amount int
}

func (s *Storage) Accept(item entity.Item) bool {
	// 投入済みのアイテムがある場合、そのアイテム以外は投入できない
	if s.item != entity.ItemUnknown && s.item != item {
		return false
	}

	s.item = item
	s.amount += 1

	return true
}

func (s *Storage) Item() entity.Item {
	return s.item
}

func (s *Storage) Amount() int {
	return s.amount
}

func (s *Storage) Clear() {
	s.item = entity.ItemUnknown
	s.amount = 0
}

func (s *Storage) Tick(tick int) {
	// not implement
}

func NewStorage() *Storage {
	return &Storage{}
}
