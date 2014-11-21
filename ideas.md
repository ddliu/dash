## Programming Language

nodejs vs golang

## Concepts

A central server.

Security layer.

Remote control with any devices(through API).

## core

## modules

Notify.

Sync.

## Plan

core => personal toolkit => ...

## Useful packages

https://github.com/codegangsta/cli

## Code Play


```
dash.Serve(ip, port)
dash.AddNode(node1)
dash.AddNode(node2)

node1 := NewNode(server)
node2 := NewNode(server)

node1.broadcast("hello world")
node1.talk("hello", node2)
```