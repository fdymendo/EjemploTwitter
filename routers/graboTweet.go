package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fdymendo/EjemploTwitter/bd"
	"github.com/fdymendo/EjemploTwitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Ocurrio un error el decodificar el body "+err.Error(), http.StatusBadRequest)
		return
	}
	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Nose  ha logrado insertar el tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
