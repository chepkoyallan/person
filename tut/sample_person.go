package tut

type Person struct {
	ID int
	FirstName string
	SecondName string
	LastName string
}

var(
	people []*Person
	nextID = 1
)

func GetPeople() []*Person {
	return people
}

func AddPerson(person Person) (Person, error){
	person.ID = nextID
	nextID++
	people = append(people, &person)
	return person, nil
}