package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method 'Get' not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main2() {
	flag.Parse()
	room := newRoom()
	go room.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("LisenAndServe: ", err)
	}
}

type sessionState int

const (
	mainView sessionState = iota
	chatView
)

type mainModel struct {
	state     sessionState
	menuitems []string
	cursor    int
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.menuitems)-1 {
				m.cursor++
			}
		case "enter", " ":
		}
	}
    switch m.state {
    case mainView:

    }
	return m, nil
}

func (m mainModel) View() string {
	s := "Go Websocket Chat\n"
	for i, menuitem := range m.menuitems {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, menuitem)
	}
	s += "\nPress q to quit.\n"
	return s
}

func initialModel() mainModel {
	return mainModel{
		menuitems: []string{"Join room", "Create room"},
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running bubbletea program: %v", err)
		os.Exit(1)
	}
}
