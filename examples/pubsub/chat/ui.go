package main

import (
	"io"

	"github.com/rivo/tview"
)

type ChatUI struct {
	cr        *ChatRoom
	app       *tview.Application
	peersList *tview.TextView

	msgW    io.Writer
	inputCh chan string
	doneCh  chan struct{}
}

func NewChatUI(cr *ChatRoom) *ChatUI { _ = "STUB: not implemented"; return nil }

func (ui *ChatUI) Run() error { _ = "STUB: not implemented"; return nil }

func (ui *ChatUI) end() { _ = "STUB: not implemented"; return }

func (ui *ChatUI) refreshPeers() { _ = "STUB: not implemented"; return }

func (ui *ChatUI) displayChatMessage(cm *ChatMessage) { _ = "STUB: not implemented"; return }

func (ui *ChatUI) displaySelfMessage(msg string) { _ = "STUB: not implemented"; return }

func (ui *ChatUI) handleEvents() { _ = "STUB: not implemented"; return }

func withColor(color, msg string) string { _ = "STUB: not implemented"; return "" }
