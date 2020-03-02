# Slack Bot App

## ■ Docker container environment

- `~/docker/golang`
  - `Dockerfile_dep` => production(dep)
  - `Dockerfile_dev` => development(dev)

Change the value of the environment variable `DOCKER_APP_ENV` in the `~/.env file`.

## ■ App development environment

``` yml
build:
  context: .
  dockerfile: Dockerfile_dev
```

App(Golang) development is directly under `~/app`.  
Go into the container and execute the `go run` command.

``` shell
$ docker exec -it slackbot-app /bin/ash

/app # go run main.go
```

## ■ Read Token

- Bot User OAuth Access Token

  ``` shell
  # cat $BOT_USER_OAUTH_ACCESS_TOKEN_PATH
  ```

- Verification Token

  ``` shell
  # cat $VERIFICATION_TOKEN_PATH
  ```

- Workspace Token

  ``` shell
  # cat $WORKSPACE_TOKEN_PATH
  ```

