# nbot

## use

    cw_api := nbot.NewChatWorkApi("{{TOKEN}}")
    cw_api.SendMessage("{{ROOM_ID}}", "hello chatwork")

    sk_api := nbot.NewSlackApi("{{TOKEN}}", "{{USER_NAME}}")
    sk_api.SendMessage("{{CHANNEL_ID}}", "hello slack")

    // https://hooks.slack.com/services/{{T-Param}}/{{B-Param}}/{{TOKEN}}
    sk_api := nbot.NewSlackIncomHookApi("{{T-Param}}", "{{B-Param}}", "{{TOKEN}}")
    sk_api.SendMessage("{{CHANNEL_ID}}", "hello slack")