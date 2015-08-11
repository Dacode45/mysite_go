package home

import(
  "net/http"
  "html/template"

  "github.com/Dacode45/mysite_go/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request){


  p := common.BasePage{
    Title: "David Ayeke",
    Nav: []string{"discover"},
  }

  common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
  err := common.Templates.ExecuteTemplate(rw, "base", p)
  common.CheckError(err, common.Error)
}
