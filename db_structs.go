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

type Entradas struct {
	Id            int64     `db:"id" json:"id"`
	Idmedicamento int64     `db:"idmedicamento" json:"idmedicamento,string"`
	Cantidad      int64     `db:"cantidad" json:"cantidad,string"`
	Fechahora     time.Time `db:"fechahora" json:"fechahora"`
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
