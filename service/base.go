package service

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// filterDate
func filterDate(db *gorm.DB, date string) *gorm.DB {
	dates := strings.Split(date, "-")
	if len(dates) == 2 {
		start, _ := time.Parse("2006/01/02", dates[0])
		end, _ := time.Parse("2006/01/02", dates[1])
		return db.Where("created_at BETWEEN ? AND ?", start, end)
	}
	if len(dates) == 1 {
		// { text: '今天', val: 'today' },
		// { text: '昨天', val: 'yesterday' },
		// { text: '最近7天', val: 'lately7' },
		// { text: '最近30天', val: 'lately30' },
		// { text: '本月', val: 'month' },
		// { text: '本年', val: 'year' }
		// TODO: 使用内置函数，可能造成索引失效
		switch dates[0] {
		case "today":
			return db.Where("TO_DAYS(NOW()) - TO_DAYS(created_at) < 1")
		case "yesterday":
			return db.Where("TO_DAYS(NOW()) - TO_DAYS(created_at) = 1")
		case "lately7":
			return db.Where("DATE_SUB(CURDATE(),INTERVAL 7 DAY) <= DATE(created_at)")
		case "lately30":
			return db.Where("DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(created_at)")
		case "month":
			return db.Where("DATE_FORMAT( created_at, '%Y%m' ) = DATE_FORMAT( CURDATE() , '%Y%m' )")
		case "year":
			return db.Where("YEAR(created_at)=YEAR(NOW())")
		}
	}
	return db
}
