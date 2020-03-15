# Stream should be created explicitly

Before appending or reading events, a stream should be created explicitly.
It is not possible to create a stream with the same key twice.
When created, stream contains a single `StreamCreated` event.