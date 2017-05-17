# nbot

## use

    cwApi := nbot.NewChatWorkApi("{{TOKEN}}")
    cwApi.SendMessage("{{ROOM_ID}}", "hello chatwork")

    skApi := nbot.NewSlackApi("{{TOKEN}}", "{{USER_NAME}}")
    skApi.SendMessage("{{CHANNEL_ID}}", "hello slack")

    // https://hooks.slack.com/services/{{T-Param}}/{{B-Param}}/{{TOKEN}}
    skApi := nbot.NewSlackIncomHookApi("{{T-Param}}", "{{B-Param}}", "{{TOKEN}}")
    skApi.SendMessage("{{CHANNEL_ID}}", "hello slack")