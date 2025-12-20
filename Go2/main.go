package lesson

import (
	"fmt"
)

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

func (a animal) MakeSound() string { return a.sound }
func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

type animal struct {
	name    string
	species string
	age     int
	sound   string
}

func NewAnimal(name, species string, age int, sound string) animal {
	return animal{name, species, age, sound}
}

func ZooShow(animals []Animal) {

	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println(animal.MakeSound())
	}
}

type ZooKeeper struct {
}

func (zk ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", animal.GetName(), animal.MakeSound())
}

func main() {
	fmt.Println("Началось")
}
