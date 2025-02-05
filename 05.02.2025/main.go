//ссылка на задачу: https://habr.com/ru/companies/yandex/articles/493966/
//A. День рождения Васи

//ЕЩЕ НЕ ДОРЕШАЛ!!!

package main

import "fmt"

type Dish struct {
	nameDish          string       //название блюдa
	quantity          int          //количество друзей, желающих отведать данное блюдо
	numberIngredients int          //количество ингредиентов необходимых для приготовления
	ingredients       []Ingredient //ингредиенты

}

type Ingredient struct {
	name            string  //название ингридиента
	uantity         float32 // требуемое количество для блюда
	unitMeasurement string  //единица измерения количества (l, ml, g, kg, cnt или tens)
}

// перевод из килограммов в граммы и обратно
func (i *Ingredient) TransfersKgG() {
	if i.unitMeasurement == "kg" {
		i.unitMeasurement = "g"
		i.uantity *= 1000
	} else {
		i.unitMeasurement = "kg"
		i.uantity /= 1000
	}
}

// перевод из миллилитров в литры и обратно
func (i *Ingredient) TransfersMlL() {
	if i.unitMeasurement == "l" {
		i.unitMeasurement = "ml"
		i.uantity *= 1000
	} else {
		i.unitMeasurement = "l"
		i.uantity /= 1000
	}
}

// справочник, где для каждого ингредиента указано его количество в упаковке и цена за упаковку
type PriceСatalog struct {
	numberIngredients int //количество ингредиентов

	//карта отображения <название ингридиента> - <количество, цена>
	packagingsIngredients map[string]struct {
		pricePackage    float32 //цена за упаковку
		quantityPackage float32 //количество в упаковке
		unitMeasurement string  //единица измерения количества (l, ml, g, kg, cnt или tens)
	}
}

// справочник, где для каждого ингредиента указано содержание белков, жиров,
// углеводов и энергетическая ценность некоторого количества данного ингредиента.
type FoodСatalog struct {

	//карта отображения <название ингридиента> - <и все что ниже =)>
	ingredients map[string]struct {
		ingredientQuantity float32 //количество ингредиента, для которого указаны характеристики ингредиента
		unitMeasurement    string  //единица измерения (l, ml, g, kg, cnt или tens)
		proteins           float32 //количесво белков
		fats               float32 //количесво жиров
		carbohydrates      float32 //количесво углеводов
		energyValue        float32 //энергетическая ценность

	}
}

func main() {
	var numberDishes int //количество блюд

	fmt.Print("Введите количество блюд: ")
	fmt.Scan(&numberDishes)

	dishes := make([]Dish, numberDishes)
	fmt.Println("Введите название блюда, количество друзей и количество ингредиентов: ")

	//итерируюсь по блюдам
	for i := 0; i < numberDishes; i++ {

		fmt.Scan(&dishes[i].nameDish, &dishes[i].quantity, &dishes[i].numberIngredients)
		dishes[i].ingredients = make([]Ingredient, dishes[i].numberIngredients)

		fmt.Println("Введите название ингредиента, количество и ед. измерения: ")
		//итерируюсь по ингредиентам, входящие в i блюдо
		for j := 0; j < dishes[i].numberIngredients; j++ {
			fmt.Scan(&dishes[i].ingredients[j].name, &dishes[i].ingredients[j].uantity, &dishes[i].ingredients[j].unitMeasurement)
		}

	}

	var priceСatalog PriceСatalog

	fmt.Print("Введите количество ингредиентов в каталоге цен: ")
	fmt.Scan(&priceСatalog.numberIngredients)
	fmt.Print("Введите навзание ингредиентов, стоимость, количество,единица измерения в каталоге цен: ")
	//итерируюсь по каталогу цен для ввода нужной информации
	for i := 0; i < priceСatalog.numberIngredients; i++ {
		var nameIngredient string
		fmt.Scan(&nameIngredient, &priceСatalog.packagingsIngredients[nameIngredient].pricePackage)

	}

	fmt.Println(dishes)

}
