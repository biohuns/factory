package building

import (
	"github.com/biohuns/factory/pkg/entity"
)

type DepositInterface interface {
	SetOutput(accepter Accepter)
	Amount() int
	Tick(tick int)
}

type Deposit struct {
	// interval 出力する間隔の tick 数
	interval int
	// item 出力するアイテム
	item entity.Item

	// output 出力先の Accepter
	output Accepter

	// amount 保管量
	amount int
	// elapsed 経過 tick 数
	elapsed int
}

func (d *Deposit) SetOutput(accepter Accepter) {
	d.output = accepter
}

func (d *Deposit) Amount() int {
	return d.amount
}

func (d *Deposit) Tick(tick int) {
	d.elapsed += tick

	if d.interval <= d.elapsed {
		d.amount += int(d.elapsed / d.interval)
		d.elapsed = d.elapsed - d.interval
		for ; d.amount > 0; d.amount-- {
			if !d.output.Accept(d.item) {
				break
			}
		}
	}
}

func NewIronOreDeposit() DepositInterface {
	return &Deposit{
		interval: 1000,
		item:     entity.ItemIronOre,
	}
}
