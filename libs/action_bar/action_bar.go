package action_bar

import (
	"fmt"
	"html/template"
	"net/http"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor/utils"
)

func init() {
	admin.RegisterViewPath("libs/action_bar/views")
}

// ActionBar stores configuration about a action bar.
type ActionBar struct {
	Admin         *admin.Admin
	GlobalActions []ActionInterface
	actions       []ActionInterface
}

// New will create an ActionBar object
func New(admin *admin.Admin) *ActionBar {
	bar := &ActionBar{Admin: admin}
	ctr := &controller{ActionBar: bar}
	admin.GetRouter().Get("action_bar/switch_mode", ctr.SwitchMode)
	admin.GetRouter().Get("action_bar/inline_edit", ctr.InlineEdit)
	return bar
}

// RegisterAction register global action
func (bar *ActionBar) RegisterAction(action ActionInterface) {
	bar.GlobalActions = append(bar.GlobalActions, action)
	bar.actions = bar.GlobalActions
}

// Actions register actions
func (bar *ActionBar) Actions(actions ...ActionInterface) *ActionBar {
	newBar := &ActionBar{Admin: bar.Admin, actions: bar.GlobalActions}
	newBar.actions = append(newBar.actions, actions...)
	return newBar
}

// Render will return the HTML of the bar, used this function to render the bar in frontend page's template or layout
func (bar *ActionBar) Render(w http.ResponseWriter, r *http.Request) template.HTML {
	var (
		actions, inlineActions []ActionInterface
		context                = bar.Admin.NewContext(w, r)
	)
	for _, action := range bar.actions {
		if action.InlineAction() {
			inlineActions = append(inlineActions, action)
		} else {
			actions = append(actions, action)
		}
	}
	context.Context.CurrentUser = bar.Admin.Auth.GetCurrentUser(context)
	result := map[string]interface{}{
		"EditMode":      bar.EditMode(w, r),
		"Auth":          bar.Admin.Auth,
		"CurrentUser":   context.Context.CurrentUser,
		"Actions":       actions,
		"InlineActions": inlineActions,
		"RouterPrefix":  bar.Admin.GetRouter().Prefix,
	}
	return context.Render("action_bar/action_bar", result)
}

// FuncMap will return helper to render inline edit button
func (bar *ActionBar) FuncMap(w http.ResponseWriter, r *http.Request) template.FuncMap {
	funcMap := template.FuncMap{}

	funcMap["render_edit_button"] = func(value interface{}, resources ...*admin.Resource) template.HTML {
		return bar.RenderEditButtonWithResource(w, r, value, resources...)
	}

	return funcMap
}

// EditMode return whether current mode is `Preview` or `Edit`
func (bar *ActionBar) EditMode(w http.ResponseWriter, r *http.Request) bool {
	return isEditMode(bar.Admin.NewContext(w, r))
}

func isEditMode(context *admin.Context) bool {
	if auth := context.Admin.Auth; auth != nil {
		if auth.GetCurrentUser(context) == nil {
			return false
		}
	}
	if cookie, err := context.Request.Cookie("qor-action-bar"); err == nil {
		return cookie.Value == "true"
	}
	return false
}

func (bar *ActionBar) RenderEditButtonWithResource(w http.ResponseWriter, r *http.Request, value interface{}, resources ...*admin.Resource) template.HTML {
	var (
		admin        = bar.Admin
		context      = admin.NewContext(w, r)
		resourceName = utils.ModelType(value).String()
		editURL, _   = utils.JoinURL(context.URLFor(value, resources...), "edit")
	)

	if res := admin.GetResource(resourceName); res != nil {
		resourceName = string(admin.T(context.Context, fmt.Sprintf("%v.name", res.ToParam()), res.Name))
	}

	for _, res := range resources {
		resourceName = string(admin.T(context.Context, fmt.Sprintf("%v.name", res.ToParam()), res.Name))
	}

	title := string(admin.T(context.Context, "qor_action_bar.action.edit_resource", "Edit {{$1}}", resourceName))
	return bar.RenderEditButton(w, r, title, editURL)
}

func (bar *ActionBar) RenderEditButton(w http.ResponseWriter, r *http.Request, title string, link string) template.HTML {
	if bar.EditMode(w, r) {
		var (
			prefix   = bar.Admin.GetRouter().Prefix
			jsURL    = fmt.Sprintf("<script data-prefix=\"%v\" src=\"%v/assets/javascripts/action_bar_check.js?theme=action_bar\"></script>", prefix, prefix)
			frameURL = fmt.Sprintf("%v/action_bar/inline_edit", prefix)
		)

		return template.HTML(fmt.Sprintf(`%v<a target="blank" data-iframe-url="%v" data-url="%v" href="#" class="qor-actionbar-button">%v</a>`, jsURL, frameURL, link, title))
	}
	return template.HTML("")
}
