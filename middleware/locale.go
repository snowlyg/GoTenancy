package middleware

import (
	"context"

	"GoTenancy/config/db"
	"GoTenancy/libs/publish2"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"github.com/kataras/iris/v12"
)

func Locale(ctx iris.Context) {
	var (
		tx         = db.DB
		qorContext = &qor.Context{Request: ctx.Request(), Writer: ctx.ResponseWriter()}
	)

	if locale := utils.GetLocale(qorContext); locale != "" {
		tx = tx.Set("l10n:locale", locale)
	}

	r := ctx.Request()
	stdCtx := context.WithValue(r.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
	ctx.ResetRequest(r.WithContext(stdCtx))

	ctx.Next()
}
