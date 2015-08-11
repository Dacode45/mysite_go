package common

import (
  "html/template"
)

//pointer to template so we don't constantly create new ones?

type BasePage struct{
  Title string
  Nav []string
}
var Templates *template.Template
const LayoutPath string = "templates/layout.html"
