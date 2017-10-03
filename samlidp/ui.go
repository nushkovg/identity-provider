package samlidp

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/crewjam/saml"
)

// LoginForm produces a form which requests a username and password and directs the user
// back to the IDP authorize URL to restart the SAML login flow, this time establishing a
// session based on the credentials that were provided.
func LoginForm(w http.ResponseWriter, r *http.Request, req *saml.IdpAuthnRequest, message string) {
	data := map[string]interface{}{
		"Error":       message,
		"URL":         req.IDP.SSOURL.String(),
		"SAMLRequest": base64.StdEncoding.EncodeToString(req.RequestBuffer),
		"RelayState":  req.RelayState,
	}

	renderTemplate("public/login/login-form.html", 200, data, w, r)
}

// ErrorForm shows error message if something went wrong.
func ErrorForm(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	data := map[string]interface{}{
		"Message": message,
	}

	renderTemplate("public/error.html", statusCode, data, w, r)
}

// BadRequestForm shows bad request message if SAML request is not valid
func BadRequestForm(w http.ResponseWriter, r *http.Request, message string) {
	data := map[string]interface{}{
		"Message": message,
	}

	renderTemplate("public/bad-request.html", 400, data, w, r)

}

// renderTemplate renders html template file
func renderTemplate(templateFile string, statusCode int, data map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	tplContent, err := loadTemplateFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(templateFile).Parse(tplContent)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")

	if err := tmpl.Execute(w, data); err != nil {
		ErrorForm(w, r, err.Error(), 500)
	}
}

// renderTemplate reads template file
func loadTemplateFile(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(b), nil
}