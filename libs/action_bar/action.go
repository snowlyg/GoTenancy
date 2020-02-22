package action_bar

import (
	"fmt"
	"html/template"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor/utils"
)

type ActionInterface interface {
	InlineAction() bool
	ToHTML(*admin.Context) template.HTML
}

type Action struct {
	EditModeOnly bool
	Inline       bool
	Name         string
	Link         string
}

func (action Action) InlineAction() bool {
	return action.Inline
}

func (action Action) ToHTML(context *admin.Context) template.HTML {
	if action.EditModeOnly && !isEditMode(context) {
		return template.HTML("")
	}
	name := context.Admin.T(context.Context, "qor_action_bar.action."+action.Name, action.Name)

	return toLink(string(name), action.Link, context.Admin)
}

type EditResourceAction struct {
	EditModeOnly bool
	Inline       bool
	Value        interface{}
	Resource     *admin.Resource
}

func (action EditResourceAction) InlineAction() bool {
	return action.Inline
}

func (action EditResourceAction) ToHTML(context *admin.Context) template.HTML {
	if action.EditModeOnly && !isEditMode(context) {
		return template.HTML("")
	}

	var (
		admin          = context.Admin
		resourceParams = utils.ModelType(action.Value).Name()
		resourceName   = resourceParams
		editURL, _     = utils.JoinURL(context.URLFor(action.Value, action.Resource), "edit")
	)

	if action.Resource != nil {
		resourceParams = action.Resource.ToParam()
		resourceName = string(admin.T(context.Context, fmt.Sprintf("%v.name", resourceParams), action.Resource.Name))
	}

	name := context.Admin.T(context.Context, "qor_action_bar.action.edit_"+resourceParams, fmt.Sprintf("Edit %v", resourceName))

	return toLink(string(name), editURL, context.Admin)
}

type HTMLAction struct {
	EditModeOnly bool
	HTML         template.HTML
}

func (action HTMLAction) InlineAction() bool {
	return true
}

func (action HTMLAction) ToHTML(context *admin.Context) template.HTML {
	if action.EditModeOnly && !isEditMode(context) {
		return template.HTML("")
	}

	return action.HTML
}

func toLink(name, link string, admin *admin.Admin) template.HTML {
	prefix := admin.GetRouter().Prefix

	if strings.HasPrefix(link, prefix) {
		jsURL := fmt.Sprintf("<script data-prefix=\"%v\" src=\"%v/assets/javascripts/action_bar_check.js?theme=action_bar\"></script>", prefix, prefix)
		frameURL := fmt.Sprintf("%v/action_bar/inline_edit", prefix)

		return template.HTML(fmt.Sprintf(`%v<a target="_blank" data-iframe-url="%v" data-url="%v" href="#" class="qor-actionbar-button">%v</a>`, jsURL, frameURL, link, name))
	} else {
		return template.HTML(fmt.Sprintf(`<a target="_blank" href="%v" class="qor-actionbar-button">%v</a>`, link, name))
	}
}
