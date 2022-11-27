package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/sdbox"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/vault"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/single"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Vaults    *vault.Vaults
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Gow Coin", single.Bank.AllVaults()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		// sdboxText := r.Form.Get("sdbox")
		single.Bank.AddOwnerVault(&sdbox.Sdboxes{})
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port := fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.tmpl"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.tmpl"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Println("Explorer server listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
