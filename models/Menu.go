package models

import (
	"drink-api/model"
	"drink-api/services"
	"drink-api/utils"
	"fmt"
	"time"
)

type Menu struct {
	model.Model
	Name        string    `json:"name" binding:"required"`
	Channel     string    `json:"channel" binding:"required"`
	IsActive    bool      `json:"is_active"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	DrinkShopID int32     `json:"-"`
	DrinkShop   DrinkShop `json:"drink_shop"`
	Orders      []Order   `json:"orders"`
	User        User      `json:"user"`
	UserID      uint      `json:"user_id"`
}

// GetEndTime return endtime unix
func (menu Menu) GetEndTime() int64 {
	return menu.EndTime.Unix()
}

func (menu Menu) GetRemainTime() (sec float64, isEnded bool) {
	sec = menu.EndTime.Sub(time.Now()).Seconds()
	if sec > 0 {
		return sec, false
	}

	return sec, true
}

func format(date time.Time, formatter string) string {
	return fmt.Sprintf(
		formatter,
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
	)
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
		menu.DrinkShop.ImageURL,
	)

	return services.SendMessage(msg, channel)
}
