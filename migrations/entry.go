package main

import (
	"marketplacego/migrations/list"

	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
)

func main() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateUserTable{},
	}
}
