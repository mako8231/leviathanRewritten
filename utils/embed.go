package utils

import (
	"github.com/bwmarrin/discordgo"
)

type Embed struct{
	*discordgo.MessageEmbed
}

func NewEmbed()(*Embed){
	return &Embed{&discordgo.MessageEmbed{}}
}

func (m *Embed) SetColor(color int) (*Embed){
	m.MessageEmbed.Color = color
	return m
}

func (m *Embed) SetTitle(title string) (*Embed){
	m.MessageEmbed.Title = title
	return m 
}


func (m *Embed) SetDescription(desc string) (*Embed){
	m.MessageEmbed.Description = desc
	return m
}

func (m *Embed) SetImage(url string) (*Embed){
	m.MessageEmbed.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}

	return m 
}

func (m *Embed) SetAuthor(imgUrl, username string)(*Embed){
	m.MessageEmbed.Author = &discordgo.MessageEmbedAuthor{
		Name: username,
		IconURL: imgUrl,
	}

	return m
}

func (m *Embed) SetThumbnail(url string)(*Embed){
	m.MessageEmbed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}

	return m 
}

func (m *Embed) SetFooter(txt string) (*Embed){
	m.MessageEmbed.Footer = &discordgo.MessageEmbedFooter{
		Text: txt,
	}

	return m
}


func (m *Embed) AddField(inline bool, title, value string) (*Embed){
	m.MessageEmbed.Fields = append(m.Fields, &discordgo.MessageEmbedField{
		Value: value,
		Inline: inline,
		Name: title,
	})

	return m 
}

