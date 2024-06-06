// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"PurchaseOrderSystem/models"
	"PurchaseOrderSystem/views/components"
)

func GetLogin() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Head(models.HeadInputs{Title: "Login", CssFile: "/static/GetLogin.css"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><section class=\"vh-100\"><div class=\"container-fluid\"><div class=\"row\"><div class=\"col-sm-6 text-black\"><div class=\"d-flex align-items-center h-custom-2 px-5 ms-xl-4 mt-5 pt-5 pt-xl-0 mt-xl-n5\"><form action=\"/login\" method=\"post\" style=\"width: 23rem;\"><h3 class=\"fw-normal mb-3 pb-3\" style=\"letter-spacing: 1px;\">Log in</h3><div data-mdb-input-init class=\"form-outline mb-4\"><input type=\"text\" id=\"username\" name=\"username\" class=\"form-control form-control-lg\"> <label class=\"form-label\" for=\"username\">Username</label></div><div class=\"form-outline mb-4\"><input type=\"password\" id=\"password\" name=\"password\" class=\"form-control form-control-lg\"> <label class=\"form-label\" for=\"password\">Password</label></div><div class=\"pt-1 mb-4\"><button class=\"btn btn-info btn-lg btn-block\" type=\"submit\">Login</button></div><p class=\"small mb-5 pb-lg-2\"><a class=\"text-muted\" href=\"#!\">Forgot password?</a></p><p>Don't have an account? <a href=\"#!\" class=\"link-info\">Go fuck yourself</a></p></form></div></div><div class=\"col-sm-6 px-0 d-none d-sm-block\"><img src=\"/static/jesus.jpg\" alt=\"Login image\" class=\"w-100 vh-100\" style=\"object-fit: cover; object-position: left;\"></div></div></div></section><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz\" crossorigin=\"anonymous\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
