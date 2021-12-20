package main

import (
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
	//...
)

func action_guild(eventinfo HMC_Bot.MessageGuild) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGuildMsg("Hello World", eventinfo.GuildID, eventinfo.ChannelID)
	}
}
func action_private(eventinfo HMC_Bot.MessagePrivate) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendPrivateMsg("Hello World", eventinfo.UserID)
	}
}
func action_group(eventinfo HMC_Bot.MessageGroup) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGroupMsg("Hello World", eventinfo.GroupID)
	}
}

func main() {
	//绑定频道消息处理函数
	HMC_Bot.Listeners.OnGuildMsg = append(HMC_Bot.Listeners.OnGuildMsg, action_guild)
	HMC_Bot.Listeners.OnPrivateMsg = append(HMC_Bot.Listeners.OnPrivateMsg, action_private)
	HMC_Bot.Listeners.OnGroupMsg = append(HMC_Bot.Listeners.OnGroupMsg, action_group)

	Bot := HMC_Bot.NewBot()
	Bot.Config = HMC_Bot.Config{
		Loglvl:   HMC_Bot.LOGGER_LEVEL_INFO,
		Host:     "0.0.0.0:6700",
		MasterQQ: 1234567890,
		Path:     "/",
	}
	Bot.Run()
}
