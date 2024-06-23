// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views_pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ppp3ppj/pppfav-htmx/views/variables"
import "github.com/ppp3ppj/pppfav-htmx/views"
import "fmt"
import "github.com/ppp3ppj/pppfav-htmx/views/components/themepicker"

func Index(body views_variables.BodyOpts) templ.Component {
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
		templ_7745c5c3_Err = views.Body(views_variables.BodyOpts{
			Component: index(),
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!--\n\t<div id=\"top-navbar\" class=\"navbar sticky top-0 z-20 bg-base-100 lg:bg-transparent\">\n\t\t<div class=\"flex-none\">\n\t\t\t<label hx-boost=\"false\" for=\"my-drawer\" class=\"btn btn-square btn-ghost drawer-button lg:hidden\">\n\t\t\t\t<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"inline-block w-5 h-5 stroke-current\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path></svg>\n\t\t\t</label>\n\t\t</div>\n\t\t<div class=\"lg:invisible flex-1\">\n            ppp:w\n\t\t</div>\n\t\t<div class=\"navbar-end md:px-6\">\n\t\t\t@views_themepicker.ThemePicker(views_themepicker.DefaultThemes)\n\t\t</div>\n\t</div>\n--><div class=\"fixed top-4 right-4 z-50\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = views_themepicker.ThemePicker(views_themepicker.DefaultThemes).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"px-auto relative mx-auto w-full self-center md:mx-auto md:max-w-3xl md:px-0 lg:max-w-4xl\"><div class=\"px-4 py-4\"><ul class=\"menu bg-base-200 lg:menu-horizontal rounded-box\"><li><a hx-boost=\"true\" href=\"/dashboard\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6\"></path></svg> Dashboard <span class=\"badge badge-sm\">99+</span></a></li><li><a><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg> Updates <span class=\"badge badge-sm badge-warning\">NEW</span></a></li><li><a>Stats <span class=\"badge badge-xs badge-info\"></span></a></li></ul><h2 class=\"py-4 font-semibold text-lg\">Latest Persons </h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range 3 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(i))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/pages/index/index.templ`, Line: 62, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div class=\"card lg:card-side bg-base-100 shadow-xl\"><figure><img src=\"https://img.daisyui.com/images/stock/photo-1494232410401-ad00d5433cfa.jpg\" alt=\"Album\"></figure><div class=\"card-body\"><h2 class=\"card-title\">New album is released!</h2><p>Click the button to listen on Spotiwhy app.</p><div class=\"card-actions justify-end\"><button class=\"btn btn-primary\">Listen</button></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
