// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "beholder-api/internal/application/models"
import "fmt"

func Sidebar(envs *[]*models.Environment) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<aside class=\"sidebar w-64 h-screen sticky top-0 left-0 bg-card border-r-2 border-border flex flex-col text-foreground\"><div class=\"p-4 border-b border-border\"><div class=\"dropdown w-full\"><button class=\"w-full px-3 py-2 text-sm text-left rounded-lg hover:bg-secondary flex items-center\"><i class=\"ri-command-line mr-2\"></i> Acme Inc <i class=\"ri-arrow-down-s-line ml-auto\"></i></button></div></div><!-- Sidebar Content --><nav class=\"flex-1 overflow-y-auto p-4\"><div class=\"space-y-4\"><!-- Personal Section --><div><h3 class=\"text-xs font-semibold text-gray-500 mb-2\">Personal</h3><div class=\"space-y-1\"><a href=\"#\" class=\"sidebar-item flex items-center text-sm px-3 py-2 rounded-lg\"><i class=\"ri-folders-line mr-2\"></i> Calls</a> <a href=\"#\" class=\"sidebar-item flex items-center text-sm px-3 py-2 rounded-lg\"><i class=\"ri-history-line mr-2\"></i> Replay history</a></div></div><div><h3 class=\"text-xs font-semibold text-gray-500 mb-2\">Environments</h3><div class=\"space-y-1\"><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, env := range *envs {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<li><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/env/%d", env.ID))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" class=\"sidebar-item flex items-center text-sm px-3 py-2 rounded-lg\"><span class=\"w-2 h-2 rounded-full bg-blue-500 mr-2\"></span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(env.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/resources/components/sidebar.templ`, Line: 59, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</ul></div></div></div></nav><!-- Sidebar Footer --></aside>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
