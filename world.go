package main

import (
	"fmt"
	"math"
)

type Location struct {
	Name        string
	Connections []int
	Monsters    map[int]Monster
	Actions     []string
}

type Monster struct {
	Name      string
	MaxHealth int
	Attack    int
}

func MakeWorld(world *map[int]Location) {
	(*world)[0] = Location{
		Name:        "Деревня",
		Connections: []int{1, 2, 4, 6},
		Actions:     []string{"Говорить", "Перейти в другую локацию"},
	}
	(*world)[1] = Location{
		Name:        "Побережье",
		Connections: []int{0},
		Monsters: map[int]Monster{
			0: {Name: "Краб", MaxHealth: 2, Attack: 1},
			1: {Name: "Чайка", MaxHealth: 1, Attack: 1},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[2] = Location{
		Name:        "Предгорья",
		Connections: []int{0, 3},
		Monsters: map[int]Monster{
			0: {Name: "Сатир", MaxHealth: 3, Attack: 1},
			1: {Name: "Кентавр", MaxHealth: 4, Attack: 2},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[3] = Location{
		Name:        "Логово тролля",
		Connections: []int{2},
		Monsters: map[int]Monster{
			0: {Name: "Тролль", MaxHealth: 27, Attack: 3},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[4] = Location{
		Name:        "Болото",
		Connections: []int{0, 5},
		Monsters: map[int]Monster{
			0: {Name: "Кабан", MaxHealth: 2, Attack: 2},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[5] = Location{
		Name:        "Кладбище",
		Connections: []int{4},
		Monsters: map[int]Monster{
			0: {Name: "Зомби", MaxHealth: 5, Attack: 2},
			1: {Name: "Скелет", MaxHealth: 3, Attack: 3},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[6] = Location{
		Name:        "Лес",
		Connections: []int{0, 7},
		Monsters: map[int]Monster{
			0: {Name: "Волк", MaxHealth: 4, Attack: 1},
			1: {Name: "Медведь", MaxHealth: 7, Attack: 2},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
	(*world)[7] = Location{
		Name:        "Пещера",
		Connections: []int{6},
		Monsters: map[int]Monster{
			0: {Name: "Летучая мышь", MaxHealth: 2, Attack: 4},
		},
		Actions: []string{"Атаковать", "Перейти в другую локацию"},
	}
}

func Action(player *Player, world *map[int]Location, people *map[int]Npc, quests *map[int]Quest, goods *map[string]int) {

	fmt.Println("Какое действие вы хотите совершить?")
	for index, value := range (*world)[player.currentLocation].Actions {
		fmt.Print(index+1, ". ", value)
		fmt.Println()
	}
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	switch tempAnswer {
	case 1:
		if player.currentLocation == 0 {
			chooseNpc(people, quests, player, goods)
		} else {
			chooseEnemy(player, world, quests)
		}
	case 2:
		Walk(player, world)
	default:
	}
}

func chooseEnemy(player *Player, world *map[int]Location, quests *map[int]Quest) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("На кого Вы хотите напасть?")
	for index, value := range (*world)[player.currentLocation].Monsters {
		fmt.Print(index+1, ". ", value.Name)
		fmt.Println()
	}
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	if tempAnswer <= len((*world)[player.currentLocation].Monsters) && tempAnswer > 0 {
		currentMonster := (*world)[player.currentLocation].Monsters[tempAnswer-1]
		fight(player, &currentMonster, quests)
	}
}

func fight(player *Player, monster *Monster, quests *map[int]Quest) {
	turnsToKillMonster := int(math.Ceil(float64(monster.MaxHealth / player.currentAttack)))
	turnsToKillPlayer := int(math.Ceil(float64(player.currentHealth / monster.Attack)))
	if turnsToKillMonster < turnsToKillPlayer {
		player.currentHealth -= turnsToKillMonster * monster.Attack
		fmt.Println("Вы победили! У Вас ", player.currentHealth, "очков здоровья")
		switch monster.Name {
		case "Кентавр":
			if player.killedMonsters["Кентавр"] < 4 {
				player.killedMonsters["Кентавр"]++
				fmt.Println("Вы находите в походном мешке кентавра кусок качественного металла")
			} else if player.killedMonsters["Кентавр"] == 4 {
				player.killedMonsters["Кентавр"]++
				fmt.Println("Вы находите в походном мешке кентавра кусок качественного металла")
				currentQuest := (*quests)[4]
				currentQuest.ready = true
				(*quests)[4] = currentQuest
			}
		case "Медведь":
			if player.killedMonsters["Медведь"] < 1 {
				player.killedMonsters["Медведь"]++
				fmt.Println("Вы аккуратно срезаете шкуру медведя")
				currentQuest := (*quests)[5]
				currentQuest.ready = true
				(*quests)[5] = currentQuest
			}
		case "Краб":
			if player.killedMonsters["Краб"] < 3 {
				player.killedMonsters["Краб"]++
				fmt.Println("К сожалению Вам пока не удалось найти сумку солдата")
			} else if player.killedMonsters["Краб"] == 3 {
				player.killedMonsters["Краб"]++
				fmt.Println("Рядом с трупом очередного краба Вы видите краешек сумки, торчащей из-под песка...")
				currentQuest := (*quests)[3]
				currentQuest.ready = true
				(*quests)[3] = currentQuest
			}
		case "Зомби":
			if player.killedMonsters["Зомби"] < 2 {
				player.killedMonsters["Зомби"]++
				fmt.Println("После уничтожения очередного зомбы Вы понимаете, что еще не все...")
			} else if player.killedMonsters["Зомби"] == 2 {
				player.killedMonsters["Зомби"]++
				fmt.Println("Думаю, теперь зомби не помешают торговцу забрать свои пожитки")
				currentQuest := (*quests)[2]
				currentQuest.ready = true
				(*quests)[2] = currentQuest
			}
		case "Летучая мышь":
			if player.killedMonsters["Летучая мышь"] < 4 {
				player.killedMonsters["Летучая мышь"]++
				fmt.Println("Вы отрезаете крыло летучей мыши, но этого пока мало")
			} else if player.killedMonsters["Летучая мышь"] == 4 {
				player.killedMonsters["Летучая мышь"]++
				fmt.Println("Кажется, знахарке этого должно быть достаточно")
				currentQuest := (*quests)[6]
				currentQuest.ready = true
				(*quests)[6] = currentQuest
			}
		case "Волк":
			if player.killedMonsters["Волк"] < 2 {
				player.killedMonsters["Волк"]++
				fmt.Println("Волки все еще угрожают безопасности деревенских овец...")
			} else if player.killedMonsters["Волк"] == 2 {
				player.killedMonsters["Волк"]++
				fmt.Println("Почти все волки истреблены, и овцы могут спать спокойно")
				currentQuest := (*quests)[1]
				currentQuest.ready = true
				(*quests)[1] = currentQuest
			}
		case "Тролль":
			if player.killedMonsters["Тролль"] < 1 {
				player.killedMonsters["Тролль"]++
			}
		case "Кабан":
			if player.killedMonsters["Кабан"] < 2 {
				player.killedMonsters["Кабан"]++
				fmt.Println("Ребенка нигде не видно...")
			} else if player.killedMonsters["Кабан"] == 2 {
				player.killedMonsters["Кабан"]++
				fmt.Println("Как только очередной кабан падает замертво, из кустов неподалеку выбегает мальчишка, живой и здоровый!")
				currentQuest := (*quests)[7]
				currentQuest.ready = true
				(*quests)[7] = currentQuest
			}
		default:
		}
	} else {
		player.currentHealth = 0
	}
}

func Walk(player *Player, world *map[int]Location) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Куда хотите пойти?")
	for index, value := range (*world)[player.currentLocation].Connections {
		fmt.Print(index+1, ". ", (*world)[value].Name)
		fmt.Println()
	}
	fmt.Print(len((*world)[player.currentLocation].Connections)+1, ". Назад")
	fmt.Println()

	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	if tempAnswer > 0 && tempAnswer < len((*world)[player.currentLocation].Connections)+1 {
		WalkDescriptor(tempAnswer, player, world)
	}
}

func WalkDescriptor(answer int, player *Player, world *map[int]Location) {
	player.currentLocation = (*world)[player.currentLocation].Connections[answer-1]
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	showCurrentLocationInfo(player, world)
	//fmt.Println("Вы перешли в локацию", (*world)[player.currentLocation].Name)
}

//0 Деревня
//1 Побережье
//2 Предгорья
//3 Логово тролля
//4 Болото
//5 Кладбище
//6 Лес
//7 Пещера
