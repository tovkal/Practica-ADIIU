package main

import "errors"

func getFarmaciaById(id string) (farmacia Farmacias, err error) {
	farmacia = Farmacias{}
	if api.DB.First(&farmacia, id).Error != nil {
		err = errors.New("Farmacia not found for the id: " + id)
	}

	return
}
