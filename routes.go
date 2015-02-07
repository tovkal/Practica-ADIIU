package main

import (
	"net/http"
)

// API routes

type HandlerFuncRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type ApiRoutes struct {
	RoutesList []HandlerFuncRoute
	Prefix     string
}

var apiRoutes = ApiRoutes{
	[]HandlerFuncRoute{
		// Categorias
		HandlerFuncRoute{
			"GetAllCategorias",
			"GET",
			"/categorias",
			getAllCategorias,
		},
		HandlerFuncRoute{
			"GetCategoriaById",
			"GET",
			"/categorias/{id}",
			getCategoria,
		},
		HandlerFuncRoute{
			"PostCategoria",
			"POST",
			"/categorias",
			postCategoria,
		},
		HandlerFuncRoute{
			"PutCategoria",
			"PUT",
			"/categorias/{id}",
			putCategoria,
		},
		HandlerFuncRoute{
			"DeleteCategoria",
			"DELETE",
			"/categorias/{id}",
			deleteCategoria,
		},

		// Entradas
		HandlerFuncRoute{
			"GetAllEntradas",
			"GET",
			"/entradas",
			getAllEntradas,
		},
		HandlerFuncRoute{
			"GetEntradasBetweenDates",
			"GET",
			"/entradas/{fromDate}/{toDate}",
			getEntradas,
		},
		HandlerFuncRoute{
			"PostEntrada",
			"POST",
			"/entradas",
			postEntrada,
		},

		// Salidas
		HandlerFuncRoute{
			"GetAllSalidas",
			"GET",
			"/salidas",
			getAllSalidas,
		},
		HandlerFuncRoute{
			"GetSalidasBetweenDates",
			"GET",
			"/salidas/{fromDate}/{toDate}",
			getSalidas,
		},
		HandlerFuncRoute{
			"PostSalida",
			"POST",
			"/salidas",
			postSalida,
		},

		// Farmacias
		HandlerFuncRoute{
			"GetAllFarmacias",
			"GET",
			"/farmacias",
			getAllFarmacias,
		},
		HandlerFuncRoute{
			"GetFarmaciaById",
			"GET",
			"/farmacias/{nik}",
			getFarmacia,
		},
		HandlerFuncRoute{
			"PostFarmacia",
			"POST",
			"/farmacias",
			postFarmacia,
		},
		HandlerFuncRoute{
			"PutFarmacia",
			"PUT",
			"/farmacias/{nik}",
			putFarmacia,
		},
		HandlerFuncRoute{
			"DeleteFarmacia",
			"DELETE",
			"/farmacias/{nik}",
			deleteFarmacia,
		},

		// Medicamentos
		HandlerFuncRoute{
			"GetAllMedicamentos",
			"GET",
			"/medicamentos",
			getAllMedicamentos,
		},
		HandlerFuncRoute{
			"GetMedicamentoById",
			"GET",
			"/medicamentos/{id}",
			getMedicamento,
		},
		HandlerFuncRoute{
			"PostMedicamento",
			"POST",
			"/medicamentos",
			postMedicamento,
		},
		HandlerFuncRoute{
			"PutMedicamento",
			"PUT",
			"/medicamentos/{id}",
			putMedicamento,
		},
		HandlerFuncRoute{
			"DeleteMedicamento",
			"DELETE",
			"/medicamentos/{id}",
			deleteMedicamento,
		},
		HandlerFuncRoute{
			"AddStock",
			"PUT",
			"/medicamentos/{id}/stock/{quantity}",
			sumaEnAlmancen,
		},

		// Noticias
		HandlerFuncRoute{
			"GetAllNoticias",
			"GET",
			"/noticias",
			getAllNoticias,
		},
		HandlerFuncRoute{
			"GetNoticiaById",
			"GET",
			"/noticias/{id}",
			getNoticia,
		},
		HandlerFuncRoute{
			"GetNoticiasForDate",
			"GET",
			"/noticias/date/{date}",
			getNoticiasFromDate,
		},
		HandlerFuncRoute{
			"PostNoticia",
			"POST",
			"/noticias",
			postNoticia,
		},
		HandlerFuncRoute{
			"PutNoticia",
			"PUT",
			"/noticias/{id}",
			putNoticia,
		},

		// File upload
		HandlerFuncRoute{
			"FileUpload",
			"POST",
			"/upload",
			uploadHandler,
		},

		// Login
		HandlerFuncRoute{
			"RememberMe",
			"POST",
			"/rememberLogin",
			rememberMeHandler,
		},
		HandlerFuncRoute{
			"Login",
			"POST",
			"/login",
			loginHandler,
		},
		HandlerFuncRoute{
			"Logout",
			"POST",
			"/logout",
			logoutHandler,
		},
	},
	"/api",
}

// Static routes

type HandlerRoute struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type StaticRoutes []HandlerRoute

var staticRoutes = StaticRoutes{
	HandlerRoute{
		"Static",
		"GET",
		"/static",
		http.StripPrefix("/static", http.FileServer(http.Dir("static/"))),
	},
}

// Templates routes

type WebRoutes []HandlerFuncRoute

var webRoutes = WebRoutes{
	HandlerFuncRoute{
		"Templates",
		"GET",
		"/{path}",
		renderTemplate,
	},
	HandlerFuncRoute{
		"Templates",
		"GET",
		"/",
		renderTemplate,
	},
}
