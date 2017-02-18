package insta_controller

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var ClientID = "702843ad7cb74e8e90cf9de804fe8f3a"
var ClientSecret = "8e29598519cc4053986c5826eab26a07"
var RedirectURI = "https://localhost:8080/acces_token"

func permiso(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://api.instagram.com/oauth/authorize/?client_id=702843ad7cb74e8e90cf9de804fe8f3a&redirect_uri=https://localhost:8080/acces_token&response_type=code", nil)

}

func accestoken(res http.ResponseWriter, req *http.Request) {
	// Después de que un usuario permite le da permisos a nuestra
	// aplicación regresa con un código de autenticación y lo guardamos
	code := req.FormValue("code")

	// Verificamos que el codigo no este vacio ( tamaño = 0)
	if len(code) != 0 {

		// Enviamos el codigo a la api, a cambio de un access token
		formResponse, err := http.PostForm("https://api.instagram.com/oauth/access_token", url.Values{"client_id": {ClientID}, "client_secret": {ClientSecret}, "grant_type": {"authorization_code"}, "redirect_uri": {RedirectURI}, "code": {code}})
		if err != nil {
			http.NotFound(res, req)
			return
		}
		defer formResponse.Body.Close()

		// Nos aseguramos que la API retorne un 200
		if formResponse.StatusCode == 200 {

			// Convertir el cuerpo en [] byte
			body, _ := ioutil.ReadAll(formResponse.Body)

			// Enviamos el access token a el usuario
			res.Write(body)
			return
		}

		// Si esto falla obtenemos un 404
		http.NotFound(res, req)
	}
}
