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

A `Msg` is a message object which stores `id`, `from`, `to` and `message` fields.

```protobuf
message Msg {
  uint64 id = 1;
  string from = 2 [
    (gogoproto.moretags) = "yaml:\"from_address\""
  ];
  string to = 3 [
    (gogoproto.moretags) = "yaml:\"to_address\""
  ];
  string message = 4 [
    (gogoproto.moretags) = "yaml:\"msg\""
  ];
}
```

- Msg: `0x01 | format(id) -> Msg`
- MsgBySender: `0x02 | format(sender) | format(id) -> Msg`
- MsgByReceiver: `0x03 | format(receiver) | format(id) -> Msg`
- LastMsgId `0x04 -> id`

## Messages

### MsgSend

`MsgSend` is a message to be used to send a new message.
At the time of message execution, it creates a new message object.

```protobuf
message MsgSend {
  string sender = 1;
  string from = 2 [
    (gogoproto.moretags) = "yaml:\"from_address\""
  ];
  string to = 3 [
    (gogoproto.moretags) = "yaml:\"to_address\""
  ];
  string message = 4 [
    (gogoproto.moretags) = "yaml:\"msg\""
  ];
}
message MsgSendResponse {
  uint64 id = 1;
  string from = 2;
  string to = 3;
  string message = 4;
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

Epochs keeper module provides utility functions to manage epochs.

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
