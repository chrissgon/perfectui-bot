package main

import "errors"

const CHANNEL_SERVER_NAME = "server"

var (
	ERR_UNKNOWN_CHANNEL_CONN = errors.New("unknown channel connection")
)
