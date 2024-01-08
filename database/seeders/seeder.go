package seeders

import (
	"golearning/pkg/seed"
)

func Initialize(){
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
