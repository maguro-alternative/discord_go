// handlers.go
package handlers

import "github.com/bwmarrin/discordgo"

func RegisterHandlers(s *discordgo.Session) {
    s.AddHandler(OnMessageCreate)
}
