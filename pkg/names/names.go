package names

import (
	"fmt"

	"sbebe.ch/baby-name-guesser/pkg/postgres"
)

// GetNames returns a list of names.
func GetNames() ([]string, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	return c.GetAllBabyNames()
}

// GetNames returns a list of names.
func AddNames(names []string) ([]string, error) {
	fmt.Println("adding names")
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	err = c.AddBabyNames(names)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// DeleteName deletes a name from the list.
func DeleteName(name string) error {
	fmt.Printf("Deleting name: %v\n", name)
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return err
	}
	err = c.DeleteBabyName(name)
	if err != nil {
		return err
	}
	return nil
}
