package plugs

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	var MSG string = `
*Hello, I am BORZ 𝘼𝙣𝙩𝙞 𝘾𝙝𝙖𝙣𝙣𝙚𝙡 spam
detector bot*.
I can ban the channels which
spams your chat!

*(c) @muharibun07*
	`
	if ctx.EffectiveChat.Type != "private" {
		ctx.EffectiveMessage.Reply(
			bot,
			"Bot is Alive (:",
			&gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)
	} else {
		ctx.EffectiveMessage.Reply(
			bot,
			MSG,
			&gotgbot.SendMessageOpts{
				ParseMode: "markdown",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Octave_Updates", Url: "https://t.me/IsIam07"},
					}},
				},
			},
		)
	}
	return ext.EndGroups
}
