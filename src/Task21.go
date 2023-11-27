package main

import "fmt"

/* Задание: Реализовать паттерн «адаптер» на любом примере. */

// Интерфейс для медиаплеера
type IMediaPlayer interface {
	Play(filename string)
}

// MP4-медиаплеер
type Mp4Player struct{}

func (self *Mp4Player) PlayMp4(filename string) {
	fmt.Printf("I'm playing an MP4 file '%s'!\n", filename)
}

// Адаптер для MP4-формата
type Mp4PlayerAdapter struct {
	*Mp4Player
}

func (self *Mp4PlayerAdapter) Play(filename string) {
	self.PlayMp4(filename)
}

// VLC-медиаплеер
type VlcPlayer struct{}

func (self *VlcPlayer) PlayVLC(filename string) {
	fmt.Printf("I'm playing a VLC file '%s'!\n", filename)
}

// Адаптер для VLC-формата
type VlcPlayerAdapter struct {
	*VlcPlayer
}

func (self *VlcPlayerAdapter) Play(filename string) {
	self.PlayVLC(filename)
}

// Конструктор для создания адаптера MP4
func NewMp4Adapter(player *Mp4Player) IMediaPlayer {
	return &Mp4PlayerAdapter{player}
}

// Конструктор для создания адаптера VLC
func NewVLCAdapter(player *VlcPlayer) IMediaPlayer {
	return &VlcPlayerAdapter{player}
}

func main() {
	adapter1 := NewMp4Adapter(&Mp4Player{})
	adapter1.Play("MyMp4File.mp4")

	adapter2 := NewVLCAdapter(&VlcPlayer{})
	adapter2.Play("MyVlcFile.mp4")
}
