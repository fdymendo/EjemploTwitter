package routers

import (
	"net/http"

	"github.com/fdymendo/EjemploTwitter/bd"
	"github.com/fdymendo/EjemploTwitter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relacion ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
