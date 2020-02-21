package help

import "GoTenancy/libs/admin"

type controller struct {
}

func (ctr controller) Index(context *admin.Context) {
	helpResource := context.Resource
	results := helpResource.NewSlice()

	context.Execute("help/index", map[string]interface{}{
		"HelpResults":  results,
		"HelpResource": helpResource,
	})
}

func (ctr controller) New(context *admin.Context) {
	context.Execute("help/new", context.Resource.NewStruct())
}
