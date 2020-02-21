package activity

import (
	"net/http"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/responder"
)

type controller struct {
	ActivityResource *admin.Resource
}

func (ctrl controller) GetActivity(context *admin.Context) {
	var (
		activities       []QorActivity
		activityResource = ctrl.ActivityResource
		result, err      = context.FindOne()
	)

	if err == nil {
		activities, err = GetActivities(context, result, "-tag")
	}

	context.AddError(err)

	if context.HasError() {
		responder.With("json", func() {
			context.Encode("edit", map[string]interface{}{"errors": context.GetErrors()})
		}).Respond(context.Request)
	} else {
		responder.With("json", func() {
			context.NewResourceContext(activityResource).Encode("index", activities)
		}).Respond(context.Request)
	}
}

func (ctrl controller) CreateActivity(context *admin.Context) {
	var (
		activityResource = ctrl.ActivityResource
		newActivity      = &QorActivity{}
		result, err      = context.FindOne()
	)

	if err == nil {
		if context.AddError(activityResource.Decode(context.Context, newActivity)); !context.HasError() {
			context.AddError(CreateActivity(context, newActivity, result))
		}
	}

	context.AddError(err)

	redirectTo := context.Request.Referer()
	if context.HasError() {
		responder.With("html", func() {
			context.Flash(context.Error(), "error")
			http.Redirect(context.Writer, context.Request, redirectTo, http.StatusFound)
		}).With("json", func() {
			context.Encode("edit", map[string]interface{}{"errors": context.GetErrors()})
		}).Respond(context.Request)
	} else {
		responder.With("html", func() {
			context.Flash(string(context.Admin.T(context.Context, "activity.successfully_created", "Activity was successfully created")), "success")
			http.Redirect(context.Writer, context.Request, redirectTo, http.StatusFound)
		}).With("json", func() {
			context.NewResourceContext(activityResource).Encode("show", newActivity)
		}).Respond(context.Request)
	}
}

func (ctrl controller) UpdateActivity(context *admin.Context) {
	c := context.Admin.NewContext(context.Writer, context.Request)
	c.ResourceID = ctrl.ActivityResource.GetPrimaryValue(context.Request)
	c.Resource = ctrl.ActivityResource
	c.Searcher = &admin.Searcher{Context: c}
	result, err := c.FindOne()

	context.AddError(err)
	if !context.HasError() {
		if context.AddError(c.Resource.Decode(c.Context, result)); !context.HasError() {
			context.AddError(context.Resource.CallSave(result, c.Context))
		}
	}

	redirectTo := context.Request.Referer()
	if context.HasError() {
		context.Writer.WriteHeader(admin.HTTPUnprocessableEntity)
		responder.With("html", func() {
			http.Redirect(context.Writer, context.Request, redirectTo, http.StatusFound)
		}).With("json", func() {
			context.Encode("edit", map[string]interface{}{"errors": context.GetErrors()})
		}).Respond(context.Request)
	} else {
		responder.With("html", func() {
			context.Flash(string(context.Admin.T(context.Context, "activity.successfully_updated", "Activity was successfully updated")), "success")
			http.Redirect(context.Writer, context.Request, redirectTo, http.StatusFound)
		}).With("json", func() {
			c.Encode("show", result)
		}).Respond(context.Request)
	}
}
