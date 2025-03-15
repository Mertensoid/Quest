package main

type Weapon struct {
	Name   string
	Attack int
}

type Chest struct {
	Name   string
	Health int
}

type Boots struct {
	Name   string
	Health int
}

func EquipWeapon(weapon Weapon, player *Player) {
	player.currentAttack = player.basicAttack + weapon.Attack
	player.currentWeapon = weapon
}

func EquipChest(chest Chest, player *Player) {
	player.maxHealth = player.basicHealth + player.currentBoots.Health + chest.Health
	player.currentChest = chest
}

func EquipBoots(boots Boots, player *Player) {
	player.maxHealth = player.basicHealth + player.currentChest.Health + boots.Health
	player.currentBoots = boots
}
