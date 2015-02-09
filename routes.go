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
			"getAllCategorias",
			"GET",
			"/categorias",
			getAllCategorias,
		},
		HandlerFuncRoute{
			"getCategoria",
			"GET",
			"/categorias/{id}",
			getCategoria,
		},
		HandlerFuncRoute{
			"postCategoria",
			"POST",
			"/categorias",
			postCategoria,
		},
		HandlerFuncRoute{
			"putCategoria",
			"PUT",
			"/categorias/{id}",
			putCategoria,
		},
		HandlerFuncRoute{
			"deleteCategoria",
			"DELETE",
			"/categorias/{id}",
			deleteCategoria,
		},

		// Entradas
		HandlerFuncRoute{
			"getAllEntradas",
			"GET",
			"/entradas",
			getAllEntradas,
		},
		HandlerFuncRoute{
			"getEntradas",
			"GET",
			"/entradas/{fromDate}/{toDate}",
			getEntradas,
		},
		HandlerFuncRoute{
			"postEntrada",
			"POST",
			"/entradas",
			postEntrada,
		},

		// Salidas
		HandlerFuncRoute{
			"getAllSalidas",
			"GET",
			"/salidas",
			getAllSalidas,
		},
		HandlerFuncRoute{
			"getSalidas",
			"GET",
			"/salidas/{fromDate}/{toDate}",
			getSalidas,
		},
		HandlerFuncRoute{
			"postSalida",
			"POST",
			"/salidas",
			postSalida,
		},

		// Farmacias
		HandlerFuncRoute{
			"getAllFarmacias",
			"GET",
			"/farmacias",
			getAllFarmacias,
		},
		HandlerFuncRoute{
			"getFarmacia",
			"GET",
			"/farmacias/{nik}",
			getFarmacia,
		},
		HandlerFuncRoute{
			"postFarmacia",
			"POST",
			"/farmacias",
			postFarmacia,
		},
		HandlerFuncRoute{
			"putFarmacia",
			"PUT",
			"/farmacias/{nik}",
			putFarmacia,
		},
		HandlerFuncRoute{
			"deleteFarmacia",
			"DELETE",
			"/farmacias/{nik}",
			deleteFarmacia,
		},

		// Medicamentos
		HandlerFuncRoute{
			"getAllMedicamentos",
			"GET",
			"/medicamentos",
			getAllMedicamentos,
		},
		HandlerFuncRoute{
			"getMedicamento",
			"GET",
			"/medicamentos/{id}",
			getMedicamento,
		},
		HandlerFuncRoute{
			"postMedicamento",
			"POST",
			"/medicamentos",
			postMedicamento,
		},
		HandlerFuncRoute{
			"putMedicamento",
			"PUT",
			"/medicamentos/{id}",
			putMedicamento,
		},
		HandlerFuncRoute{
			"deleteMedicamento",
			"DELETE",
			"/medicamentos/{id}",
			deleteMedicamento,
		},
		HandlerFuncRoute{
			"sumaEnAlmancen",
			"PUT",
			"/medicamentos/{id}/stock/{quantity}",
			sumaEnAlmancen,
		},

		// Noticias
		HandlerFuncRoute{
			"getAllNoticias",
			"GET",
			"/noticias",
			getAllNoticias,
		},
		HandlerFuncRoute{
			"getNoticia",
			"GET",
			"/noticias/{id}",
			getNoticia,
		},
		HandlerFuncRoute{
			"getNoticiasFromDate",
			"GET",
			"/noticias/date/{date}",
			getNoticiasFromDate,
		},
		HandlerFuncRoute{
			"postNoticia",
			"POST",
			"/noticias",
			postNoticia,
		},
		HandlerFuncRoute{
			"putNoticia",
			"PUT",
			"/noticias/{id}",
			putNoticia,
		},

		// File upload
		HandlerFuncRoute{
			"uploadHandler",
			"POST",
			"/upload",
			uploadHandler,
		},

		// Login
		HandlerFuncRoute{
			"rememberMeHandler",
			"POST",
			"/rememberLogin",
			rememberMeHandler,
		},
		HandlerFuncRoute{
			"loginHandler",
			"POST",
			"/login",
			loginHandler,
		},
		HandlerFuncRoute{
			"logoutHandler",
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
