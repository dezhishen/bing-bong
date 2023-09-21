package client

import "github.com/spf13/viper"

func getSubscribeCommands() []string {
	commands := viper.GetStringSlice("commands.subscribe")
	if commands == nil {
		return []string{"订阅", "subscribe"}
	}
	return commands
}

func getUnsubscribeCommands() []string {
	commands := viper.GetStringSlice("commands.unsubscribe")
	if commands == nil {
		return []string{"取消订阅", "unsubscribe"}
	}
	return commands
}

func getsearchSubscriptionCommands() []string {
	commands := viper.GetStringSlice("commands.searchSubscription")
	if commands == nil {
		return []string{"查询订阅", "searchSubscription"}
	}
	return commands
}
