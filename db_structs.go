package main

type Categorias struct {
	Id     int64  `db:"id" json:"id"`
	Nombre string `db:"nombre" json:"nombre"`
	Texto  string `db:"texto" json:"texto"`
	Imagen string `db:"imagen" json:"imagen"`
}

type Entradas struct {
	Id            int64  `db:"id" json:"id"`
	Idmedicamento int64  `db:"idmedicamento" json:"idmedicamento,string"`
	Cantidad      int64  `db:"cantidad" json:"cantidad,string"`
	Fechahora     string `db:"fechahora" json:"fechahora"`
}

type EntradasJoin struct {
	Id                int64  `db:"id" json:"id"`
	Idmedicamento     int64  `db:"idmedicamento" json:"idmedicamento,string"`
	Nombremedicamento string `db:"nombremedicamento" json:"nombremedicamento"`
	Cantidad          int64  `db:"cantidad" json:"cantidad,string"`
	Fechahora         string `db:"fechahora" json:"fechahora"`
}

type Farmacias struct {
	Id    int64  `db:"id" json:"id"`
	Nik   string `db:"nik" json:"nik"`
	Pass  string `db:"pass" json:"pass"`
	Nivel int64  `db:"nivel" json:"nivel,string"`
}

type Medicamentos struct {
	Id          int64  `db:"id" json:"id"`
	Idcategoria int64  `db:"idcategoria" json:"idcategoria,string"`
	Texto       string `db:"texto" json:"texto"`
	Imagen      string `db:"imagen" json:"imagen"`
	Nombre      string `db:"nombre" json:"nombre"`
	Codigo      string `db:"codigo" json:"codigo"`
	Enalmacen   int64  `db:"enalmacen" json:"enalmacen,string"`
}

type Noticias struct {
	Id     int64  `db:"id" json:"id"`
	Texto  string `db:"texto" json:"texto"`
	Inicio string `db:"inicio" json:"inicio"`
	Fin    string `db:"fin" json:"fin"`
}

type Salidas struct {
	Id            int64  `db:"id" json:"id"`
	Idmedicamento int64  `db:"idmedicamento" json:"idmedicamento,string"`
	Fechahora     string `db:"fechahora" json:"fechahora"`
	Cantidad      int64  `db:"cantidad" json:"cantidad,string"`
	Idfarmacia    int64  `db:"idfarmacia" json:"idfarmacia,string"`
}

type SalidasJoin struct {
	Id                int64  `db:"id" json:"id"`
	Idmedicamento     int64  `db:"idmedicamento" json:"idmedicamento,string"`
	Nombremedicamento string `db:"nombremedicamento" json:"nombremedicamento"`
	Fechahora         string `db:"fechahora" json:"fechahora"`
	Cantidad          int64  `db:"cantidad" json:"cantidad,string"`
	Idfarmacia        int64  `db:"idfarmacia" json:"idfarmacia,string"`
	Nombrefarmacia    string `db:"nombrefarmacia" json:"nombrefarmacia"`
}
