package models

import (
	"drink-api/services"
	"drink-api/utils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	ID          int32     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Channel     string    `json:"channel" binding:"required"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" sql:"default:now()"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	DrinkShopID int32     `json:"drink_shop"`
	UserID      int32     `json:"user" pg:"fk:user_id"`
	Orders      []*Order  `json:"orders"`
	ImageURL    string    `json:image_url`
}

// GetEndTime return endtime unix
func (menu Menu) GetEndTime() int64 {
	return menu.EndTime.Unix()
}

func format(date time.Time, formatter string) string {
	return fmt.Sprintf(formatter, date.Month(), date.Day(), date.Hour(), date.Minute())
}

func (menu *Menu) AfterSave() (err error) {
	channel := menu.Channel
	if menu.Channel == "" {
		channel = "#frontend-underground"
	}

	formatStr := "%02d-%02d %02d:%02d:%02d"
	// Duration
	remainTime := menu.EndTime.Sub(menu.CreatedAt)

	msg := fmt.Sprintf(`
	已經發起了訂飲料活動 %s\n	
	*店家名稱*: %s
	*開始時間*: %s
	*結束時間*: %s
	*訂餐連結*: %s
	*剩餘時間*: %s
	*訂單圖片*: %s
	`,
		menu.Name,
		format(menu.CreatedAt, formatStr),
		format(menu.EndTime, formatStr),
		utils.GetURLByID("menus", menu.ID),
		fmt.Sprintf("%.0f分%.0f秒", remainTime.Minutes(), remainTime.Seconds()),
		menu.ImageURL,
	)

	return services.SendMessage(msg, channel)
}
