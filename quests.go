package main

type Quest struct {
	title       string
	discription string
	status      int
	need        int
	ready       bool
	done        bool
	coinReward  int
	otherReward string
}

func makeQuests(quests *map[int]Quest) {
	(*quests)[0] = Quest{
		title:       "Тролль-людоед",
		discription: "Нашей деревне угрожает страшный тролль. Если кто-то его не победит, мы все умрем! Спаси нас!",
		status:      0, //1
		need:        1,
		ready:       false,
		done:        false,
		coinReward:  0,
		otherReward: "Голова тролля",
	}
	(*quests)[1] = Quest{
		title:       "Пропажа овец",
		discription: "Волки повадились вороввать наших овец! Убей трех волков, иначе в деревне скоро нечего будет есть...",
		status:      0, //3
		need:        3,
		ready:       false,
		done:        false,
		coinReward:  3,
		otherReward: "",
	}
	(*quests)[2] = Quest{
		title:       "Телега на кладбище",
		discription: "Во время очередной поставки товара недалеко от кладбища на обоз напали зомби. Пожалуйста, уничтожь нечисть, чтобы я мог забрать товар.",
		status:      0, //3
		need:        3,
		ready:       false,
		done:        false,
		coinReward:  0,
		otherReward: "Сапоги",
	}
	(*quests)[3] = Quest{
		title:       "Потерянная сумка",
		discription: "Во время патрулирования побережья я был ранен и вынужден бежать в деревню, оставив все свои пожитки. Найди мою сумку, она дорога мне, как память.",
		status:      0, //4
		need:        4,
		ready:       false,
		done:        false,
		coinReward:  2,
		otherReward: "",
	}
	(*quests)[4] = Quest{
		title:       "Новая кольчуга",
		discription: "Если хочешь, чтобы я выковал тебе новую кольчугу, тебе придется найти несколько кусков качественного металла. Я слышал у кентавров нагорья есть такой.",
		status:      0, //5
		need:        5,
		ready:       false,
		done:        false,
		coinReward:  0,
		otherReward: "Кольчуга",
	}
	(*quests)[5] = Quest{
		title:       "Шкура медведя",
		discription: "Моя жена выгнала меня из дома и сказа не возвращаться без медвежьей шубы... Если достанешь мне шкуру медведя, я готов даже поменять ее на свой фамильный меч!",
		status:      0, //1
		need:        1,
		ready:       false,
		done:        false,
		coinReward:  0,
		otherReward: "Меч",
	}
	(*quests)[6] = Quest{
		title:       "Недостающие ингридиенты",
		discription: "Я уже стара и мне сложно добывать некоторые ингридиенты для моих зелий. Если ты принесешь мне несколько крыльев летучей мыши, я щедро вознагражу тебя",
		status:      0, //5
		need:        5,
		ready:       false,
		done:        false,
		coinReward:  3,
		otherReward: "",
	}
	(*quests)[7] = Quest{
		title:       "Пропавший ребенок",
		discription: "Помоги мне, добрый молодец! Мой сын пошел играть на болото и не вернулся. Найди его, прошу тебя!",
		status:      0, //3
		need:        3,
		ready:       false,
		done:        false,
		coinReward:  2,
		otherReward: "",
	}
}

//0 Тролль-людоед
//1 Пропажа овец
//2 Телега на кладбище
//3 Потерянная сумка
//4 Новая кольчуга
//5 Шкура медведя
//6 Недостающие ингридиенты
//7 Пропавший ребенок
