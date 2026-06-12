package entity

type Item int

const (
	ItemUnknown Item = iota

	// Ore
	ItemIronOre
	ItemCopperOre
	ItemLimestone
	ItemCoal

	ItemIronPlate
	ItemIronRod
)
