# slack-to-telegram
Slack incoming webhook to telegram

![slack-to-telegram-image](docs/image.png)

Send messages from slack incoming webhook integrations to telegram

```Tested on Gitlab Slack Notifications```

## Configuration
```yaml
server:
  # Server address
  address: 0.0.0.0:3000
telegram:
  # Telegram bot token
  token: bot_token
  # Telegram default chat to which all messages will be sent, if no chat find in mapChats
  defaultChat: 000000000
  # Map slack channels to telegram chats, add channel without #
  mapChats:
    general: 111111111
  # Disable notifications in telegram
  disableNotification: true
# Path to template to file
template: ./assets/default.tmpl
```

## Usage
Docker: 
```bash
docker run -v $PWD/config:/etc/slack-to-telegram timmiles/slack-to-telegram --config /etc/slack-to-telegram/config.yaml
```

Kubernetes:
```bash
kubectl apply -f deploy/k8s
```

To change message sent to telegram, modify ```assets/default.tmpl``` and mount it to container, because default.tmpl is 
baked into docker image
