# random-sticker-bot

A Telegram bot that sends a sticker randomly from a sticker set when specific command is received.


## Deployment
Build with `go build` or use pre-built container from:

 https://github.com/users/star-39/packages/container/package/random-sticker-bot

 Run the container with:
```
podman run -dt --rm -e BOT_TOKEN=your_token -e COMMAND=your_command -e STICKER_SET=sticker_set_name ghcr.io/star-39/random-sticker-bot
```
Of course you can use `docker` as well.

## License
The GPL v3 License
