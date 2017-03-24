# nbot

## use

    cw_api := nbot.NewChatWorkApi("{{TOKEN}}")
    cw_api.SendMessage("{{ROOM_ID}}", "hello chatwork")

    sk_api := nbot.NewSlackApi("{{TOKEN}}", "{{USER_NMAE}}")
    sk_api.SendMessage("{{CHANNEL_ID}}", "hello slack")
