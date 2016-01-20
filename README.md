# HipChabot

Post random messages to a HipChat room.

## Create HTTP Webhook

```sh
curl "https://api.hipchat.com/v2/room/FIXME/webhook?auth_token=FIXME" -X "Post" -H │
"Content-Type: application/json" -d @request.json
```

