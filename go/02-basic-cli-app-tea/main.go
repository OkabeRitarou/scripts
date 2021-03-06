package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    count int
}

func (m *Model) Init() tea.Cmd {
    return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg: 
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "up" :
            m.count++
        case "down" :
            m.count--
        }
    }

    return m, nil
}

func (m *Model) View() string {
    return fmt.Sprintf("count is: %d\n\n ↑: increse    ↓: decrese", m.count)
}


func main() {
    //err := tea.NewProgram(&Model{}, tea.WithAltScreen()).Start()
    err := tea.NewProgram(&Model{}).Start()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
