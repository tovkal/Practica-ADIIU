package main

import "errors"

func getFarmaciaByNik(nik string) (farmacia Farmacias, err error) {
	farmacia = Farmacias{}
	if api.DB.Where("nik = ?", nik).Find(&farmacia).Error != nil {
		err = errors.New("Farmacia not found")
	}

	return
}
