package seed

import "gorm.io/gorm"

var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(db *gorm.DB)

type Seeder struct {
	Func SeederFunc
	name string
}

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}