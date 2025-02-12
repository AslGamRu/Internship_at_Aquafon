package main

import (
	"errors"
	//"fmt"
)
//-----------------------DISH-----------------------
//==================================================
//структура блюда
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


//функция считает стоимость блюда
func CostAllDishes(dishes []Dish,prcCat *PriceСatalog) float32{
	var costAllDishes float32
	listNumIngr:=prcCat.RequiredNumberIngredients(dishes)
	for nameIngr,ingr:=range prcCat.pkgIngr{
		costAllDishes+=ingr.pricePkg * float32(listNumIngr[nameIngr])
	}
	return costAllDishes
}



//функция преобразования ед.измерения
func ConvOneUnitMeasur (dishIngr *Ingredient)error{

	unitMeasur:=dishIngr.unitMeasurement

	switch {
	case unitMeasur=="l" || unitMeasur=="ml" :
		dishIngr.TransfersMlL()
	case unitMeasur=="g" || unitMeasur=="kg" :
		dishIngr.TransfersKgG()
	case unitMeasur=="cnt" || unitMeasur=="tens" :
		dishIngr.TransferCntTens()
	default:
		return errors.New("Невозможно преобразовать ед. измерения.")
	}
	return nil
}

func (d *Dish)CharacteristicsDish(foodCat *FoodСatalog)(float32,float32,float32,float32){
	var prot,fats,carb,ener float32
	for _,ingr:=range d.ingr{
		foodIngr:=foodCat.ingr[ingr.name]
		if foodIngr.unitMeasurement!=ingr.unitMeasurement{
			ConvOneUnitMeasur(&ingr)
		}
		prot+=(foodIngr.proteins*ingr.number)/foodIngr.numIngr  
		fats+=(foodIngr.fats*ingr.number)/foodIngr.numIngr  
		carb+=(foodIngr.carbohydrates*ingr.number)/foodIngr.numIngr 
		ener+=(foodIngr.energyValue*ingr.number)/foodIngr.numIngr   
	}
	return prot,fats,carb,ener
}

//==================================================

//--------------------INGREDIENT--------------------
//==================================================
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

// перевод из килограммов в граммы (если нужно)
func (i *Ingredient) TransfersKgG() {
    if i.unitMeasurement == "kg" {
        i.unitMeasurement = "g"
        i.number *= 1000
    } else if i.unitMeasurement == "g" { // Добавлено условие
        i.unitMeasurement = "kg"
        i.number /= 1000
    }
}

// перевод из миллилитров в литры (если нужно)
func (i *Ingredient) TransfersMlL() {
    if i.unitMeasurement == "ml" {
        i.unitMeasurement = "l"
        i.number /= 1000
    } else if i.unitMeasurement == "l"{
        i.unitMeasurement = "ml"
        i.number *= 1000
    }
}

// перевод из единиц в десятки (и обратно)
func (i *Ingredient) TransferCntTens() {
    if i.unitMeasurement == "cnt" {
        i.unitMeasurement = "tens"
        i.number /= 10
    } else if i.unitMeasurement == "tens" {
        i.unitMeasurement = "cnt"
        i.number *= 10
    }
}
//==================================================

//-------------------PRICECATALOG-------------------
//==================================================
// справочник, где для каждого ингредиента указано его количество в упаковке и цена за упаковку
type PriceСatalog struct {
	numIngr int //количество ингредиентов в каталоге

	//карта отображения <название ингридиента> - <количество, цена>
	pkgIngr map[string]PriceIngredient
}

//Функция инициализации карты
func(p *PriceСatalog) NumIngr(numIngr int){
	p.numIngr=numIngr
	p.pkgIngr=make(map[string]PriceIngredient)
}

//функция заполнения карты каталога цен
func (p *PriceСatalog) InitPriceCatalog(name string, price float32, num float32, unitMeasur string) {
	packIngr, ok := p.pkgIngr[name]
	if !ok {
		packIngr.InitPackIngr(price, num, unitMeasur)
		p.pkgIngr[name] = packIngr
	}
}

//эта структура создана специально под PriceСatalog
type PriceIngredient struct {
	pricePkg        float32 //цена за упаковку
	numPkg          float32 //количество в упаковке
	unitMeasurement string  //единица измерения количества (l, ml, g, kg, cnt или tens)
}
 
//функция инициализации ингредиентов из PriceСatalog
func (p *PriceIngredient) InitPackIngr(price float32, num float32, unitMeasur string) {
	p.pricePkg = price
	p.numPkg = num
	p.unitMeasurement = unitMeasur
}

func (p *PriceСatalog) RequiredNumberIngredients(dishes []Dish)map[string]int{
	var reqNumIngr  =make(map[string]float32)
	var listNumIngr=make( map[string]int)

	//итерируюсь по всем блюдам
	for _,dish:=range dishes{
		//итерируюсь по всем ингредиентам в блюде
		for _,ingr:=range dish.ingr{

			_,ok:=reqNumIngr[ingr.name]
			priceIngr:=p.pkgIngr[ingr.name]

			if priceIngr.unitMeasurement!=ingr.unitMeasurement{
				ConvOneUnitMeasur(&ingr)
			}

			if ok {
				reqNumIngr[ingr.name]+=ingr.number * float32(dish.numFriends)
			}else{
				reqNumIngr[ingr.name]=ingr.number * float32(dish.numFriends)
			}
		}
	}

	for key,value:=range p.pkgIngr{
		numIngr,ok:=reqNumIngr[key]
		if ok{
			listNumIngr[key]=1
			for i:=1;value.numPkg*float32(i)<=numIngr;i++{
				listNumIngr[key]+=1
			}
		}else{
			listNumIngr[key]=0
		}

	}
return listNumIngr
}
//=================================================

//-------------------FOODCDTDLOG-------------------
//=================================================
// справочник, где для каждого ингредиента указано содержание белков, жиров,
// углеводов и энергетическая ценность некоторого количества данного ингредиента.
type FoodСatalog struct {
	numIngr int //количество ингредиентов в каталоге

	//карта отображения <название ингридиента> - <и все что ниже =)>
	ingr map[string]FoodIngredient
}

//Функция инициализации карты
func(f *FoodСatalog) NumIngr(numIngr int){
	f.numIngr=numIngr
	f.ingr=make(map[string]FoodIngredient)
}

func (f *FoodСatalog) InitFoodCataloge(name string, num float32, unitMeasur string, prot, fats, carb, ener float32) {
	ingr, ok := f.ingr[name]
	if !ok {
		ingr.FoodIngr(num, unitMeasur, prot, fats, carb, ener)
		f.ingr[name] = ingr
	}
}
//эта структура создана специально под FoodСatalog
type FoodIngredient struct {
	numIngr         float32 //количество ингредиента, для которого указаны характеристики ингредиента
	unitMeasurement string  //единица измерения (l, ml, g, kg, cnt или tens)
	proteins        float32 //количесво белков
	fats            float32 //количесво жиров
	carbohydrates   float32 //количесво углеводов
	energyValue     float32 //энергетическая ценность
}

//функция инициализации ингредиентов из FoodСatalog
func (f *FoodIngredient) FoodIngr(num float32, unitMeasur string, prot, fats, carb, ener float32) {
	f.numIngr = num
	f.unitMeasurement = unitMeasur
	f.proteins = prot
	f.fats = fats
	f.carbohydrates = carb
	f.energyValue = ener
}
//=================================================





