package main

import (
	"fmt"
)

type Npc struct {
	Name   string
	Quests []int
}

func MakeNpc(people *map[int]Npc) {
	(*people)[0] = Npc{Name: "Староста", Quests: []int{0, 1, 5}}
	(*people)[1] = Npc{Name: "Торговец", Quests: []int{2}}
	(*people)[2] = Npc{Name: "Знахарка", Quests: []int{6}}
	(*people)[3] = Npc{Name: "Кузнец", Quests: []int{4}}
	(*people)[4] = Npc{Name: "Жительница деревни", Quests: []int{7}}
	(*people)[5] = Npc{Name: "Раненый солдат", Quests: []int{3}}
}

func chooseNpc(people *map[int]Npc, quests *map[int]Quest, player *Player, goods *map[string]int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("С кем Вы хотите поговорить?")
	for index, value := range *people {
		fmt.Print(index+1, ". ", value.Name)
		fmt.Println()
	}
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	if tempAnswer <= len(*people) && tempAnswer > 0 {
		Talk(tempAnswer-1, people, quests, player, goods)
	}
}

func Talk(npc int, people *map[int]Npc, quests *map[int]Quest, player *Player, goods *map[string]int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("О чем Вы хотите поговорить?")
	for index, value := range (*people)[npc].Quests {
		fmt.Print(index+1, ". ", (*quests)[value].title)
		fmt.Println()
	}
	switch npc {
	case 1:
		fmt.Print(len((*people)[npc].Quests)+1, ". Торговать")
		fmt.Println()
		fmt.Print(len((*people)[npc].Quests)+2, ". Уйти")
		fmt.Println()
	case 2:
		fmt.Print(len((*people)[npc].Quests)+1, ". Восстановить здоровье")
		fmt.Println()
		fmt.Print(len((*people)[npc].Quests)+2, ". Уйти")
		fmt.Println()
	default:
		fmt.Print(len((*people)[npc].Quests)+1, ". Уйти")
		fmt.Println()
	}
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	if tempAnswer > 0 && tempAnswer < len((*people)[npc].Quests)+1 {
		QuestDescriptor(npc, tempAnswer, people, quests, player, goods)
	} else if npc == 1 && tempAnswer == len((*people)[npc].Quests)+1 {
		TradeDescriptor(goods, player)
	} else if npc == 2 && tempAnswer == len((*people)[npc].Quests)+1 {
		player.currentHealth = player.maxHealth
	}
}

func QuestDescriptor(npc int, answer int, people *map[int]Npc, quests *map[int]Quest, player *Player, goods *map[string]int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	currentQuest := (*quests)[(*people)[npc].Quests[answer-1]]
	if currentQuest.done {
		fmt.Println("Спасибо! Ты мне очень помог!")
		fmt.Println("1. Обращайся")
	} else if currentQuest.ready {
		fmt.Println(currentQuest.discription)
		fmt.Println("1. Я выполнил задание")
	} else {
		fmt.Println(currentQuest.discription)
		fmt.Println("1. Хорошо")
	}
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	if tempAnswer == 1 {
		if currentQuest.ready && !currentQuest.done {
			currentQuest.done = true
			(*quests)[(*people)[npc].Quests[answer-1]] = currentQuest
			if currentQuest.coinReward > 0 {
				player.coins += currentQuest.coinReward
			} else {
				getReward(currentQuest.otherReward, player)
			}
		}
		Talk(npc, people, quests, player, goods)
	} else {
		fmt.Println("Выберите из предложенных вариантов ответа...")
		fmt.Scanf("%d\n", &tempAnswer)
	}
}

func TradeDescriptor(goods *map[string]int, player *Player) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------------------------------------------")
	showStats(player)
	fmt.Println("-----------------------------------------------------")
	i := 1
	var tempSlice []string
	for key, value := range *goods {
		fmt.Printf("%d. %s (Цена: %d)\n", i, key, value)
		i++
		tempSlice = append(tempSlice, key)
	}
	fmt.Printf("%d. Уйти\n", i)
	var tempAnswer int
	fmt.Scanf("%d\n", &tempAnswer)
	switch {
	case tempAnswer == i:
		break
	case tempAnswer < i && tempAnswer > 0:
		switch tempSlice[tempAnswer-1] {
		case "Дубина":
			if player.currentWeapon.Name == "Меч" {
				fmt.Println("Ваше оружие лучше")
			} else if player.coins >= (*goods)["Дубина"] {
				EquipWeapon(Weapon{"Дубина", 2}, player)
				player.coins -= (*goods)["Дубина"]
				delete(*goods, "Дубина")
			} else {
				fmt.Println("Недостаточно денег")
			}
		case "Куртка":
			if player.currentChest.Name == "Кольчуга" {
				fmt.Println("Ваш нагрудник лучше")
			} else if player.coins >= (*goods)["Куртка"] {
				EquipChest(Chest{"Куртка", 5}, player)
				player.coins -= (*goods)["Куртка"]
				delete(*goods, "Куртка")
			} else {
				fmt.Println("Недостаточно денег")
			}
		case "Тапочки":
			if player.currentBoots.Name == "Сапоги" {
				fmt.Println("Ваша обувь лучше")
			} else if player.coins >= (*goods)["Тапочки"] {
				EquipBoots(Boots{"Тапочки", 2}, player)
				player.coins -= (*goods)["Тапочки"]
				delete(*goods, "Тапочки")
			} else {
				fmt.Println("Недостаточно денег")
			}
		default:
		}
	default:
		fmt.Println("Выберите из предложенных вариантов ответа...")
		fmt.Scanf("%d\n", &tempAnswer)
	}
}

func getReward(reward string, player *Player) {
	switch reward {
	case "Меч":
		rewardWeapon := Weapon{Name: "Меч", Attack: 4}
		EquipWeapon(rewardWeapon, player)
	case "Кольчуга":
		rewardChest := Chest{Name: "Кольчуга", Health: 10}
		EquipChest(rewardChest, player)
	case "Сапоги":
		rewardBoots := Boots{Name: "Сапоги", Health: 5}
		EquipBoots(rewardBoots, player)
	default:
	}
}

//0 Староста
//1 Торговец
//2 Знахарка
//3 Кузнец
//4 Жительница
//5 Раненый солдат

//0 Тролль-людоед
//1 Пропажа овец
//2 Телега на кладбище
//3 Потерянная сумка
//4 Новая кольчуга
//5 Шкура медведя
//6 Недостающие ингридиенты
//7 Пропавший ребенок
