package names

import (
	"sbebe.ch/baby-name-guesser/pkg/postgres"
	"sbebe.ch/baby-name-guesser/pkg/utils"
)

type Names struct {
	Names []string `json:"names"`
}

// GetNames returns a list of names.
func GetNames() ([]string, error) {
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.GetAllBabyNames()
}

// GetNames returns a list of names.
func AddNames(names []string) ([]string, error) {
	utils.Logger.Sugar().Infof("Adding names: %v", names)
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	err = c.AddBabyNames(names)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// DeleteName deletes a name from the list.
func DeleteName(name string) error {
	utils.Logger.Sugar().Infof("Deleting name: %v\n", name)
	c, err := postgres.NewPostgresClient()
	if err != nil {
		return err
	}
	defer c.Close()
	err = c.DeleteBabyName(name)
	if err != nil {
		return err
	}
	return nil
}
