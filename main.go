package main

import "fmt"

type Player struct {
	basicHealth     int
	maxHealth       int
	currentHealth   int
	basicAttack     int
	currentAttack   int
	currentLocation int
	currentWeapon   Weapon
	currentChest    Chest
	currentBoots    Boots
	killedMonsters  map[string]int
	coins           int
}

func main() {
	fmt.Print("\033[H\033[2J")
	player := Player{
		basicHealth:     5,
		maxHealth:       5,
		currentHealth:   5,
		basicAttack:     1,
		currentAttack:   1,
		currentLocation: 0,
		currentWeapon:   Weapon{Name: "Палка", Attack: 0},
		currentChest:    Chest{Name: "Рубаха", Health: 0},
		currentBoots:    Boots{Name: "Лапти", Health: 0},
		killedMonsters: map[string]int{
			"Кентавр":      0,
			"Медведь":      0,
			"Краб":         0,
			"Зомби":        0,
			"Летучая мышь": 0,
			"Волк":         0,
			"Тролль":       0,
			"Кабан":        0,
		},
		coins: 0,
	}

	Quests := make(map[int]Quest)
	makeQuests(&Quests)

	World := make(map[int]Location)
	MakeWorld(&World)

	People := make(map[int]Npc)
	MakeNpc(&People)

	goods := map[string]int{
		"Дубина":  3,
		"Куртка":  5,
		"Тапочки": 2,
	}

	for {
		Action(&player, &World, &People, &Quests, &goods)
		if player.currentHealth <= 0 {
			fmt.Println("Вы погибли... Попробуйте заного!")
			break
		}
		if player.killedMonsters["Тролль"] > 0 {
			fmt.Println("Поздравляем! Вы победили тролля и спасли жителей деревни!")
			break
		}
	}
}

func showStats(player *Player) {
	fmt.Printf("Атака: %d\n", player.currentAttack)
	fmt.Printf("Здоровье: %d/%d\n", player.currentHealth, player.maxHealth)
	fmt.Printf("Оружие: %s\n", player.currentWeapon.Name)
	fmt.Printf("Нагрудник: %s\n", player.currentChest.Name)
	fmt.Printf("Обувь: %s\n", player.currentBoots.Name)
	fmt.Printf("Золото: %d\n", player.coins)
}

func showCurrentLocationInfo(player *Player, world *map[int]Location) {
	fmt.Printf("Вы находитесь в локации %s\n", (*world)[player.currentLocation].Name)

	if player.currentLocation > 0 {
		fmt.Printf("Вы видите, что на локации присутствуют:\n")
		for _, value := range (*world)[player.currentLocation].Monsters {
			fmt.Printf("%s (Атака: %d. Здоровье: %d)", value.Name, value.Attack, value.MaxHealth)
			fmt.Println()
		}
	}
	fmt.Println("-----------------------------------------------------")
}
