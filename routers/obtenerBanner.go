package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/fdymendo/EjemploTwitter/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado ", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Imagen al copiar la imagen", http.StatusBadRequest)
		return
	}

}
