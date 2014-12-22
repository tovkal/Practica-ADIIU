package main

import (
	"time"
)

type Categorias struct {
	Id     int64  `db:"id" json:"id"`
	Nombre string `db:"nombre" json:"nombre"`
	Texto  string `db:"texto" json:"texto"`
	Imagen string `db:"imagen" json:"imagen"`
}

func compareCategorias(a, b Categorias) bool {
	if &a == &b {
		return true
	}

	if a.Nombre != b.Nombre {
		return false
	}

	if a.Texto != b.Texto {
		return false
	}

	if a.Imagen != b.Imagen {
		return false
	}

	return true
}

type Entradas struct {
	Id            int64     `db:"id" json:"id"`
	Idmedicamento int64     `db:"idmedicamento" json:"idmedicamento"`
	Cantidad      int64     `db:"cantidad" json:"cantidad"`
	Fechahora     time.Time `db:"fechahora" json:"fechahora"`
}

func compareEntradas(a, b Entradas) bool {
	if &a == &b {
		return true
	}

	if a.Idmedicamento != b.Idmedicamento {
		return false
	}

	if a.Cantidad != b.Cantidad {
		return false
	}

	if a.Fechahora != b.Fechahora {
		return false
	}

	return true
}

type Farmacias struct {
	Id    int64  `db:"id" json:"id"`
	Nik   string `db:"nik" json:"nik"`
	Pass  string `db:"pass" json:"pass"`
	Nivel int64  `db:"nivel" json:"nivel"`
}

type Medicamentos struct {
	Id          int64  `db:"id" json:"id"`
	Idcategoria int64  `db:"idcategoria" json:"idcategoria"`
	Texto       string `db:"texto" json:"texto"`
	Imagen      string `db:"imagen" json:"imagen"`
	Nombre      string `db:"nombre" json:"nombre"`
	Codigo      string `db:"codigo" json:"codigo"`
	Enalmacen   int64  `db:"enalmacen" json:"enalmacen"`
}

type Noticias struct {
	Id     int64     `db:"id" json:"id"`
	Texto  string    `db:"texto" json:"texto"`
	Inicio time.Time `db:"inicio" json:"inicio"`
	Fin    time.Time `db:"fin" json:"fin"`
}

type Salidas struct {
	Id            int64     `db:"id" json:"id"`
	Idmedicamento int64     `db:"idmedicamento" json:"idmedicamento"`
	Fechahora     time.Time `db:"fechahora" json:"fechahora"`
	Cantidad      int64     `db:"cantidad" json:"cantidad"`
	Idfarmacia    int64     `db:"idfarmacia" json:"idfarmacia"`
}

func compareSalidas(a, b Salidas) bool {
	if &a == &b {
		return true
	}

	if a.Idmedicamento != b.Idmedicamento {
		return false
	}

	if a.Fechahora != b.Fechahora {
		return false
	}

	if a.Cantidad != b.Cantidad {
		return false
	}

	if a.Idfarmacia != b.Idfarmacia {
		return false
	}

	return true
}
