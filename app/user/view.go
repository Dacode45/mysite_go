package user

import(
  "net/http"
  "html/template"

  "github.com/Dacode45/mysite_go/app/common"

  "github.com/gorilla/mux"
)

func GetViewPage(rw http.ResponseWriter, req *http.Request){
  type Page struct{
    Title string
    UserId string
  }

  params := mux.Vars(req)
  userId := params["id"]
  //get user Id
  p := Page{
    Title: "user_view",
    UserId: userId,
  }

  common.Templates = template.Must(template.ParseFiles("templates/user/view.html", common.LayoutPath))
  err := common.Templates.ExecuteTemplate(rw, "base", p)
  common.CheckError(err, common.Error)
}
