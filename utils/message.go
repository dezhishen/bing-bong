package utils

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
)

// getTemplate 消息模板
func getTemplate() string {
	result := viper.GetString("messageTemplate")
	if result != "" {
		return result
	}
	return "《%s》更新了《%s》\n链接：%s"
}

func BuildMessage(rssTitle string, item *gofeed.Item) string {
	return fmt.Sprintf(getTemplate(), rssTitle, item.Title, item.Link)
}
