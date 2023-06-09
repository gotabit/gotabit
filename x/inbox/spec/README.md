# Inbox

## Abstract

This module is used to send on-chain messages between accounts.

Users can send messages and query their sent/received messages.

## Contents

1. **[State](#state)**
2. **[Messages](#messages)**
3. **[Events](#events)**
4. **[Queries](#queries)**

## State

### Msg

A `Msg` is a message object which stores `id`, `sender`, `to`, `topics` and `message` fields.

```protobuf
// Msg defines the inbox item - msg
message Msg {
  option (gogoproto.goproto_getters) = false;
  option (gogoproto.goproto_stringer) = false;

  // msg id
  uint64 id = 1;

  // msg sender address
  string sender = 2 [
    (gogoproto.moretags) = "yaml:\"sender\""
  ];

  // msg recipient address
  string to = 3 [
    (gogoproto.moretags) = "yaml:\"to_address\""
  ];

  // msg topics
  string topics = 4 [
    (gogoproto.moretags) = "yaml:\"the_message_topics\""
  ];

  // msg message
  string message = 5 [
    (gogoproto.moretags) = "yaml:\"message\""
  ];
}
```

- Msg: `0x01 | format(id) -> Msg`
- MsgBySender: `0x02 | format(sender) | format(id) -> Msg`
- MsgByReceiver: `0x03 | format(receiver) | format(id) -> Msg`
- MsgByReceiverAndTopics: `0x04 | format(receiver) | format(topics) | format(id) -> Msg`
- LastMsgId `0x05 -> id`

## Messages

### MsgSend

`MsgSend` is a message to be used to send a new message.
At the time of message execution, it creates a new message object.

```protobuf
// MsgSend defines a message for sending a message
message MsgSend {
  // msg sender address
  string sender = 1;

  // msg recipient address
  string to = 2 [
    (gogoproto.moretags) = "yaml:\"to_address\""
  ];

  // msg topics
  string topics = 3 [
    (gogoproto.moretags) = "yaml:\"the_message_topics\""
  ];

  // msg message
  string message = 4 [
    (gogoproto.moretags) = "yaml:\"message\""
  ];
}

// MsgSendResponse defines the MsgSend response type
message MsgSendResponse {
  // msg id
  uint64 id = 1;

  // msg sender address
  string sender = 2;

  // msg recipient address
  string to = 3;

  // msg topics
  string topics = 4;

  // msg message
  string message = 5;
}
```

Steps:

1. Get unique message id by using last message id
2. Store newly generated message id as last message id
3. Create message object with newly generated message id
4. Store message object on storage
5. Emit event for message creation

## Events

The `inbox` module emits the following event:

### EventMsgSend

|  Type          | Attribute Key |  Attribute Value |
|  --------------| ---------------| -----------------|
|  gotabit.inbox.v1beta1.EventMsgSend |  sender |  {sender} |
|  gotabit.inbox.v1beta1.EventMsgSend |  receiver |  {receiver} |
|  gotabit.inbox.v1beta1.EventMsgSend |  id |  {id} |

## Keepers

### Keeper functions

Inbox keeper module provides utility functions to manage inbox.

```go
// Keeper is the interface for lockup module keeper
type Keeper interface {
  // GetLastMsgId returns last msg id
  GetLastMsgId(ctx sdk.Context) uint64
  // SetLastMsgId set last msg id
  SetLastMsgId(ctx sdk.Context, id uint64)
  // GetMsgById returns msg by id
  GetMsgById(ctx sdk.Context, id uint64) (*types.Msg, error)
  // GetMsgsBySender returns all msgs by sender
  GetMsgsBySender(ctx sdk.Context, sender string) []*types.Msg
  // GetMsgsByReceiver returns all msgs by receiver
  GetMsgsByReceiver(ctx sdk.Context, receiver string) []*types.Msg
  // SetMsg stores msg to store
  SetMsg(ctx sdk.Context, msg *types.Msg)
}
```

## Queries

`inbox` module is providing below queries to check the module's state.

```protobuf
service Query {
  // SentMessages returns messages sent from user
  rpc SentMessages(SentMessagesRequest) returns (SentMessagesResponse) {}
  // ReceivedMessages returns messages received by user
  rpc ReceivedMessages(ReceivedMessagesRequest) returns (ReceivedMessagesResponse) {}
}
```

### Sent Messages

Query the sent messages

```sh
gotabitd query inbox sent --sender=SENDER_ADDRESS
```

### Received Messages

Query the received messages

```sh
gotabitd query inbox received --receiver=RECEIVER_ADDRESS
```
