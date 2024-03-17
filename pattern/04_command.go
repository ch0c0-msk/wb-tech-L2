package pattern

// Command interface
type Command interface {
	Execute()
}

// Command implementations
type LightOn struct {
	light *Light
}

func NewLightOn(l *Light) *LightOn {
	return &LightOn{light: l}
}

func (l *LightOn) Execute() {
	l.light.lightOn()
}

type LightOff struct {
	light *Light
}

func NewLightOff(l *Light) *LightOff {
	return &LightOff{light: l}
}

func (l *LightOff) Execute() {
	l.light.lightOff()
}

// Command receiver
type Light struct {
	IsLightOn bool
}

func (l *Light) lightOn() {
	l.IsLightOn = true
}

func (l *Light) lightOff() {
	l.IsLightOn = false
}

// Commands caller
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) ExecuteCommand() {
	if len(i.commands) == 0 {
		return
	}
	i.commands[0].Execute()
	i.commands = i.commands[1:]
}
