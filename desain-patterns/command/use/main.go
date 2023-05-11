package main

type Command interface {
	Execute()
}

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) Execute() {
	for _, c := range m.commands {
		c.Execute()
	}
}

func (m *MacroCommand) Append(c Command) {
	m.commands = append(m.commands, c)
}

func (m *MacroCommand) Undo() {
	if len(m.commands) != 0 {
		m.commands = m.commands[:len(m.commands)-1]
	}
}

func (m *MacroCommand) Clear() {
	m.commands = []Command{}
}

func (m *MacroCommand) History() []Command {
	return m.commands
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{}
}
