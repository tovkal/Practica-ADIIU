package main

import "errors"

func getEntradaById(id int64) (entrada EntradasJoin, err error) {
	entrada = EntradasJoin{}
	if api.DB.Table("entradas").Select("entradas.id, entradas.idmedicamento, medicamentos.nombre as nombremedicamento, entradas.cantidad, entradas.fechahora").Joins("INNER JOIN medicamentos ON entradas.idmedicamento = medicamentos.id").Where("entradas.id = ?", id).Scan(&entrada).Error != nil {
		err = errors.New("Salida not found")
	}

	return
}
