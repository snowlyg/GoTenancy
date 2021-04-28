package initialize

import (
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
)

func Timer() {
	if g.TENANCY_CONFIG.Timer.Start {
		for _, detail := range g.TENANCY_CONFIG.Timer.Detail {
			fmt.Println(detail)
			// go func(detail config.Detail) {
			// 	g.TENANCY_Timer.AddTaskByFunc("ClearDB", g.TENANCY_CONFIG.Timer.Spec, func() {
			// 		err := utils.ClearTable(g.TENANCY_DB, detail.TableName, detail.CompareField, detail.Interval)
			// 		if err != nil {
			// 			fmt.Println("timer error:", err)
			// 		}
			// 	})
			// }(detail)
		}
	}
}
