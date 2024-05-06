package main

var Args struct {
	Client struct {
		ServerAddr string  `help:"server address"`
		Addr       *string `help:"optional client address, defaults to first available port"`
		// Pvp        struct {
		// 	OpponentAddr string `help:"opponents address"`
		// } `default:"1" help:"play against other client"`
	} `cmd:"" help:"run as client"`
	Server struct {
		Addr string `help:"server address" default:":5000"`
	} `cmd:"" help:"run as server"`
}
