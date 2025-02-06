package main

type Dish struct {
	name       string       //название блюдa
	numFriends int          //количество друзей, желающих отведать данное блюдо
	numIngr    int          //количество ингредиентов необходимых для приготовления
	ingr       []Ingredient //ингредиенты

}

// Функция инициализации полей
func (d *Dish) InitDish(name string, numFr int, numIngr int) {
	d.name = name
	d.numFriends = numFr
	d.numIngr = numIngr
	d.ingr = make([]Ingredient, numIngr)
}

// фунуция инициализации массива ингрединетов, используемых в блюде
func (d *Dish) InitIngrDish(name string, num float32, unitMeasur string, count int) {
	d.ingr[count].InitIngr(name, num, unitMeasur)
}

type Ingredient struct {
	name            string  //название ингридиента
	number          float32 // требуемое количество для блюда
	unitMeasurement string  //единица измерения количества (l, ml, g, kg, cnt или tens)
}

// функция инициализации ингредиента
func (i *Ingredient) InitIngr(name string, num float32, unitMeasur string) {
	i.name = name
	i.number = num
	i.unitMeasurement = unitMeasur
}

// перевод из килограммов в граммы и обратно
func (i *Ingredient) TransfersKgG() {
	if i.unitMeasurement == "kg" {
		i.unitMeasurement = "g"
		i.number *= 1000
	} else {
		i.unitMeasurement = "kg"
		i.number /= 1000
	}
}

// перевод из миллилитров в литры и обратно
func (i *Ingredient) TransfersMlL() {
	if i.unitMeasurement == "l" {
		i.unitMeasurement = "ml"
		i.number *= 1000
	} else {
		i.unitMeasurement = "l"
		i.number /= 1000
	}
}

// справочник, где для каждого ингредиента указано его количество в упаковке и цена за упаковку
type PriceСatalog struct {
	numIngr int //количество ингредиентов в каталоге

	//карта отображения <название ингридиента> - <количество, цена>
	pkgIngr map[string]PriceIngredient
}

func (p *PriceСatalog) InitPriceCatalog(name string, price float32, num float32, unitMeasur string) {
	packIngr, ok := p.pkgIngr[name]
	if ok {
		packIngr.InitPackIngr(price, num, unitMeasur)
	}
	p.pkgIngr[name] = packIngr
}

type PriceIngredient struct {
	pricePkg        float32 //цена за упаковку
	numPkg          float32 //количество в упаковке
	unitMeasurement string  //единица измерения количества (l, ml, g, kg, cnt или tens)
}

func (p *PriceIngredient) InitPackIngr(price float32, num float32, unitMeasur string) {
	p.pricePkg = price
	p.numPkg = num
	p.unitMeasurement = unitMeasur
}

// справочник, где для каждого ингредиента указано содержание белков, жиров,
// углеводов и энергетическая ценность некоторого количества данного ингредиента.
type FoodСatalog struct {
	numIngr int //количество ингредиентов в каталоге

	//карта отображения <название ингридиента> - <и все что ниже =)>
	ingr map[string]FoodIngredient
}

func (f *FoodСatalog) InitFoodCataloge(name string, num float32, unitMeasur string, prot, fats, carb, ener float32) {
	ingr, ok := f.ingr[name]
	if ok {
		ingr.FoodIngr(num, unitMeasur, prot, fats, carb, ener)
	}
	f.ingr[name] = ingr
}

type FoodIngredient struct {
	numIngr         float32 //количество ингредиента, для которого указаны характеристики ингредиента
	unitMeasurement string  //единица измерения (l, ml, g, kg, cnt или tens)
	proteins        float32 //количесво белков
	fats            float32 //количесво жиров
	carbohydrates   float32 //количесво углеводов
	energyValue     float32 //энергетическая ценность
}

func (f *FoodIngredient) FoodIngr(num float32, unitMeasur string, prot, fats, carb, ener float32) {
	f.numIngr = num
	f.unitMeasurement = unitMeasur
	f.proteins = prot
	f.fats = fats
	f.carbohydrates = carb
	f.energyValue = ener
}
