package generation

import "fmt"

func (m *Model) View() string {

	return fmt.Sprintf("Password ditits\n%s", m.textInput.View()) + "\n"
}
