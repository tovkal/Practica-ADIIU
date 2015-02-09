package main

import "errors"

func getSalidaById(id int64) (salida SalidasJoin, err error) {
	salida = SalidasJoin{}
	if api.DB.Table("salidas").Select("salidas.id, salidas.idmedicamento, medicamentos.nombre as nombremedicamento, salidas.fechahora, salidas.cantidad, salidas.idfarmacia, farmacias.nik as nombrefarmacia").Joins("INNER JOIN medicamentos ON salidas.idmedicamento = medicamentos.id INNER JOIN farmacias ON salidas.idfarmacia = farmacias.id").Where("salidas.id = ?", id).Scan(&salida).Error != nil {
		err = errors.New("Salida not found")
	}

	return
}
