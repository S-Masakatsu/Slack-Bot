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

## ■ ngrok

1. Connect your account

    ``` shell
    $ cd ngrok
    $ ./ngrok authtoken <Token>
    ```

1. Start Docker container for production

    Make Docker a production environment
    ``` env
    DOCKER_APP_ENV=dep
    ```

    ``` shell
    $ docker-compose up --build
    ```

1. Fire it up

    ``` shell
    $ ./ngrok http -host-header="0.0.0.0:3000" 3000
    ```
