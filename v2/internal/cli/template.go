package cli

const UserInput = `git commit -m "{{.Title}}" -m "{{.Body}}"`

const CommandTemplate = `{{.Instruction}} {{range .Options}} {{.Flag}} {{.Value}} {{end}}`
