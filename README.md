# IM Bot Webhook Adapter

This package provides a webhook adapter for sending messages (text or markdown) to multiple IM bots via the official webhook API.

## Features

- **Send Messages to Multiple Bots**: Easily send text or markdown messages to all registered bots.
- **Bot Management**: Register multiple bots with the `BotManager` to send messages across various channels.

## Supported IM webhook providers

* WeCom
* FeiShu

## Installation

To use this package in your Go project, you can install it by running:

```bash
go get github.com/gmodx/im-bot-webhook-adapter
```

## Usage

### Step 1: Register Bots

You can register multiple bots using their respective webhook keys. For example:

```go
import "github.com/gmodx/im-bot-webhook-adapter"

weComBot := imbotwebhookadapter.NewWeComBot("your_wecom_webhook_key")
imbotwebhookadapter.Register("bot_name", weComBot)
```

### Step 2: Sending Messages

Once the bots are registered, you can send messages to all registered bots.

```go
imbotwebhookadapter.Send(imbotwebhookadapter.MsgType_Text, "This is a text message")
```

You can also send markdown messages:

```go
imbotwebhookadapter.Send(imbotwebhookadapter.MsgType_Markdown, "This is a **markdown** message")
```

## Message Types

- **Text**: Sends a plain text message.
- **Markdown**: Sends a message formatted with markdown syntax.

## Bot Interface

The `Bot` interface allows you to implement any type of bot that supports sending messages:

```go
type Bot interface {
    Send(msgType MsgType, content string) error
}
```

Bots can be registered and managed using the `BotManager`.

## Example

Here's a complete example of registering bots and sending a message:

```go
package main

import "github.com/gmodx/im-bot-webhook-adapter"

func main() {
    weComBot := imbotwebhookadapter.NewWeComBot("your_wecom_webhook_key")
    imbotwebhookadapter.Register("WeComBot", weComBot)

    // Send a text message to all registered bots
    imbotwebhookadapter.Send(imbotwebhookadapter.MsgType_Text, "Hello from WeCom Bot!")
}
```

