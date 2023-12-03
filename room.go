package main

type Room struct {
    clients map[*Client]bool
}
