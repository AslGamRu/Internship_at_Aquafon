//ссылка на задачу: https://habr.com/ru/companies/yandex/articles/493966/
//A. День рождения Васи

//ЕЩЕ НЕ ДОРЕШАЛ!!!

package main

import (
	"fmt"
)

func main() {

	//=================================================
	var numDishes int //количество блюд

	fmt.Print("Введите количество блюд: ")
	fmt.Scan(&numDishes)

	dishes := make([]Dish, numDishes)

	fmt.Println("Введите название блюда, количество друзей и количество ингредиентов: ")

	//итерируюсь по блюдам
	for i := 0; i < numDishes; i++ {
		var name string
		var numFr, numIngr int

		fmt.Scan(&name, &numFr, &numIngr)
		dishes[i].InitDish(name, numFr, numIngr)
		fmt.Println("Введите название ингредиента, количество и ед. измерения: ")

		//итерируюсь по ингредиентам, входящие в i блюдо
		for j := 0; j < numIngr; j++ {
			var unitMeasur string
			var num float32

			fmt.Scan(&name, &num, &unitMeasur)
			dishes[i].InitIngrDish(name, num, unitMeasur, j)
		}

	}
	//=================================================

	//=================================================
	var priceСatalog PriceСatalog

	fmt.Print("Введите количество ингредиентов в каталоге цен: ")
	fmt.Scan(&priceСatalog.numIngr)
	fmt.Println("Введите навзание ингредиентов, стоимость, количество,единица измерения в каталоге цен: ")

	//итерируюсь по каталогу цен для ввода нужной информации
	for i := 0; i < priceСatalog.numIngr; i++ {
		var name, unitMeasur string
		var price, num float32

		fmt.Scan(&name, &price, &num, &unitMeasur)
		priceСatalog.InitPriceCatalog(name, price, num, unitMeasur)
	}
	//=================================================

	//=================================================
	var foodCatalog FoodСatalog

	fmt.Print("Введите количество ингредиентов в каталоге еды: ")
	fmt.Scan(&foodCatalog.numIngr)
	fmt.Println("Введите название, количество,ед.измерения и содержание белков, жиров, углеводов и энергетическая ценность ингредиентов: ")

	//итерируюсь по каталогу еды для ввода нужной информации
	for i := 0; i < foodCatalog.numIngr; i++ {
		var name, unitMeasur string
		var num, prot, fats, carb, ener float32

		fmt.Scan(&name, &num, &unitMeasur)
		fmt.Scan(&prot, &fats, &carb, &ener)
		foodCatalog.InitFoodCataloge(name, num, unitMeasur, prot, fats, carb, ener)
	}

	//=================================================
	fmt.Println(dishes)
	fmt.Println(priceСatalog)
	fmt.Println(foodCatalog)

}
