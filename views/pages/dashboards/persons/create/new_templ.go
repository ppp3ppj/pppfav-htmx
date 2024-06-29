// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views_dashboards_persons_new

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ppp3ppj/pppfav-htmx/views/variables"
import "github.com/ppp3ppj/pppfav-htmx/pkg/models"
import "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards"

//import "fmt"
import "github.com/ppp3ppj/pppfav-htmx/utils/ternary"
import views_dashboards_persons_new_components "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards/persons/create/components"

type NewPersonVM struct {
	Opts    views_variables.DashboardOpts
	BaseURL string
	Person  *models.Person
}

func New(vm NewPersonVM) templ.Component {
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
		templ_7745c5c3_Err = views_dashboard.Dashboard(views_variables.DashboardOpts{
			Nav:               vm.Opts.Nav,
			AdditionalHeaders: vm.Opts.AdditionalHeaders,
			Comp:              new(vm.Person, vm.BaseURL),
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

func new(existingPerson *models.Person, baseURL string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full h-full\"><div class=\"flex flex-col justify-center items-center\"><form id=\"form-person-post\" class=\"w-full md:max-w-3xl lg:max-w-4xl\" hx-encoding=\"multipart/form-data\"><div class=\"text-3xl font-bold\">Create Person</div><!-- Image Field --><label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">What is your name?</span> <span class=\"label-text-alt\">Top Right label</span></div><div class=\"flex items-center space-x-6\"><div class=\"avatar shrink-0\"><div class=\"w-24 rounded-full\"><img src=\"https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/20230905_Haerin_%28NewJeans%29.jpg/455px-20230905_Haerin_%28NewJeans%29.jpg?20230905163647\"></div></div><label class=\"block\"><span class=\"sr-only\">Choose profile photo</span> <input class=\"file-input w-full max-w-xs block\" type=\"file\" id=\"inputUploadImage\" name=\"imageData\"></label></div><div class=\"label\"><span class=\"label-text-alt\">Bottom Left label</span> <span class=\"label-text-alt\">Bottom Right label</span></div></label><!-- Name Field -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = views_dashboards_persons_new_components.NameFieldLabelValidation(views_dashboards_persons_new_components.NewPersonVadidateVM{
			ElementId: "lblFieldName",
			TitleName: "What is your name",
			BasePath:  "/persons/validate/name",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!--\n            <label\n                hx-target=\"this\"\n                hx-swap=\"outerHTML\"\n                class=\"form-control\">\n              <div class=\"label\">\n                <span class=\"label-text\">What is your name?</span>\n                <span class=\"label-text-alt\">Top Right label</span>\n              </div>\n              <input\n                  name=\"name\"\n                  hx-post=\"/persons/validate/name\"\n                  type=\"text\" placeholder=\"Type here\" class=\"input input-bordered w-full md:max-w-3xl lg:max-w-4xl\" />\n              <div class=\"label\">\n                <span class=\"label-text-alt\">Bottom Left label</span>\n                <span class=\"label-text-alt\">Bottom Right label</span>\n              </div>\n            </label>\n            --><!-- Age Field --><label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">What is your name?</span> <span class=\"label-text-alt\">Top Right label</span></div><input type=\"text\" placeholder=\"Type here\" class=\"input input-bordered w-full md:max-w-3xl lg:max-w-4xl\"><div class=\"label\"><span class=\"label-text-alt\">Bottom Left label</span> <span class=\"label-text-alt\">Bottom Right label</span></div></label><!-- Birth date Field --><label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">What is your name?</span> <span class=\"label-text-alt\">Top Right label</span></div><input type=\"date\" placeholder=\"Type here\" class=\"input input-bordered w-full md:max-w-3xl lg:max-w-4xl\"><div class=\"label\"><span class=\"label-text-alt\">Bottom Left label</span> <span class=\"label-text-alt\">Bottom Right label</span></div></label><!-- Description Field --><label class=\"form-control\"><div class=\"label\"><span class=\"label-text\">Your bio</span> <span class=\"label-text-alt\">Alt label</span></div><textarea class=\"textarea textarea-bordered h-24\" placeholder=\"Bio\"></textarea><div class=\"label\"><span class=\"label-text-alt\">Your bio</span> <span class=\"label-text-alt\">Alt label</span></div></label><!-- Cute Rating Field --><div class=\"form-control\"><div class=\"label\"><span class=\"label-text\">Cute Rating</span></div><div class=\"rating gap-1\"><input type=\"radio\" name=\"rating-3\" class=\"mask mask-heart bg-green-400\" checked=\"checked\"> <input type=\"radio\" name=\"rating-3\" class=\"mask mask-heart bg-lime-400\"> <input type=\"radio\" name=\"rating-3\" class=\"mask mask-heart bg-yellow-400\"> <input type=\"radio\" name=\"rating-3\" class=\"mask mask-heart bg-orange-400\"> <input type=\"radio\" name=\"rating-3\" class=\"mask mask-heart bg-red-400\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = submitButton(ternary.Struct(existingPerson, &models.Person{}).ID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func submitButton(existingId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full flex flex-row justify-end space-x-3\"><button id=\"draft-button\" class=\"btn btn-secondary md:max-w-3xl lg:max-w-4xl w-24 md:w-32 lg:w-40\" hx-get=\"/dashboard/persons\" hx-target=\"#admin-root\" hx-push-url=\"true\" hx-swap=\"innerHTML\">Cancel</button> <button id=\"publish-button\" class=\"btn btn-primary md:max-w-3xl lg:max-w-4xl w-32 md:w-40 lg:w-48\" hx-post=\"/dashboard/persons/push\" hx-target=\"#admin-root\" hx-swap=\"outerHTML\" hx-indicator=\"#global-progress\">Save</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
