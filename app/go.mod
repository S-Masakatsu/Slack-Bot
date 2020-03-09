module main

go 1.14

require (
	github.com/nlopes/slack v0.6.0
	github.com/slack-go/slack v0.6.2 // indirect
	golang.org/x/exp v0.0.0-20200228211341-fcea875c7e85 // indirect
	handlers/handlers v0.5.0
	slack/config v0.5.0
	slack/messages v0.5.0
	slack/todo v0.5.0
)

replace (
	handlers/handlers => /app/src/handlers
	slack/config => /app/src/slack/config
	slack/messages => /app/src/slack/messages
	slack/todo => /app/src/slack/todo
)
