package shortcut

import (
	"github.com/wilsongp/go-api/routing"
)

var controller = &Controller{Repository: Repository{}}

//Routes with associated controllers
var Routes = routing.Routes{
	{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controller.Index,
	},
	{
		Name:        "AddShortcut",
		Method:      "POST",
		Pattern:     "/",
		HandlerFunc: controller.AddShortcut,
	},
	{
		Name:        "UpdateShortcut",
		Method:      "PUT",
		Pattern:     "/",
		HandlerFunc: controller.UpdateShortcut,
	},
	{
		Name:        "DeleteShortcut",
		Method:      "DELETE",
		Pattern:     "/",
		HandlerFunc: controller.DeleteShortcut,
	},
}
