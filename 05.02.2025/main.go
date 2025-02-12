//ссылка на задачу: https://habr.com/ru/companies/yandex/articles/493966/
//A. День рождения Васи

//ЕЩЕ НЕ ДОРЕШАЛ!!!

package main

import (
	"fmt"
	//"log"
)

func main() {

	//Заполняем информацию про блюда
	//=================================================
	var numDishes int //количество блюд

	fmt.Scan(&numDishes)

	dishes := make([]Dish, numDishes)


	//итерируюсь по блюдам
	for i := 0; i < numDishes; i++ {
		var name string
		var numFr, numIngr int

		fmt.Scan(&name, &numFr, &numIngr)
		dishes[i].InitDish(name, numFr, numIngr)

		//итерируюсь по ингредиентам, входящие в i блюдо
		for j := 0; j < numIngr; j++ {
			var unitMeasur string
			var num float32

			fmt.Scan(&name, &num, &unitMeasur)
			dishes[i].InitIngrDish(name, num, unitMeasur, j)
		}

	}
	//=================================================

	//заполняю информация в каталоге цен
	//=================================================
	var priceСatalog PriceСatalog
	var numIngr int

	fmt.Scan(&numIngr)
	priceСatalog.NumIngr(numIngr)

	//итерируюсь по каталогу цен для ввода нужной информации
	for i := 0; i < numIngr; i++ {
		var name, unitMeasur string
		var price, num float32

		fmt.Scan(&name, &price, &num, &unitMeasur)
		priceСatalog.InitPriceCatalog(name, price, num, unitMeasur)
	}
	//=================================================


	//заполняю информацию в каталоге еды
	//=================================================
	var foodCatalog FoodСatalog

	fmt.Scan(&numIngr)
	foodCatalog.NumIngr(numIngr)

	//итерируюсь по каталогу еды для ввода нужной информации
	for i := 0; i < numIngr; i++ {
		var name, unitMeasur string
		var num, prot, fats, carb, ener float32

		fmt.Scan(&name, &num, &unitMeasur)
		fmt.Scan(&prot, &fats, &carb, &ener)
		foodCatalog.InitFoodCataloge(name, num, unitMeasur, prot, fats, carb, ener)
	}
	//=================================================

	//вычисляю стоимость всех блюд
	//=================================================
	var costAllDishes = CostAllDishes(dishes,&priceСatalog)
	fmt.Println("Стоимость всех блюд:",costAllDishes)
	//=================================================

	//Вывод необходимого количества ингредиентов
	//=================================================
	fmt.Println("Вывод необходимого количества ингредиентов:")
	listNumIngr:=priceСatalog.RequiredNumberIngredients(dishes)

	for key,value:=range listNumIngr{
		fmt.Println(key,value)
	}
	//=================================================

	//Вычисляю характеристики блюда
	//=================================================
	fmt.Println("Xарактеристики блюда:")
	for _,dish:=range dishes{
		prot, fats, carb, ener := dish.CharacteristicsDish(&foodCatalog)
		fmt.Println(dish.name, prot, fats, carb, ener)
	}
}
