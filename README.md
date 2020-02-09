# event-store-go [![Build Status](https://circleci.com/gh/im-tollu/event-store-go.svg?style=svg)](https://app.circleci.com/github/im-tollu/event-store-go/pipelines)
Event store for learning, demo and prototyping purposes

## Motivation

This project is created in the context of learning CQRS/ES. It is intentionally separated
from other logic. First, to highlight the role of an event store itself. Second, to be able to
reuse it in different business contexts.

## Target outcomes

* The project should be usable as a library in a CQRS/ES project. It should act as a database 
façade. It should provide basic features of event store like storing events in event streams,
subscribing to events and preserving the events order and consistency within a stream.
* It also should be demonstrable as a part of presentation on internal workings of event-sourcing 
system.
* It should show that event-sourcing is simple.
* It should show an example of using Go for implementing backend systems in comparison to 
JVM-based solutions