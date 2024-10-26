// Copyright 2024 Notedown Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package taskeditor

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/notedownorg/notedown/pkg/ast"
	"github.com/notedownorg/notedown/pkg/workspace/documents/writer"
)

func (m *Model) submit() (tea.Model, tea.Cmd) {
	// Create is intentionally run syncronously to prevent losing progress on error
	if err := m.tasks.Create("README.md", m.text.Value(), m.status.Value(), ast.WithLine(writer.AtEnd)); err != nil {
		slog.Error("failed to create task", "error", err)

		// TODO: We should probably show an error message to the user
		return m, nil
	}

	// If we've successfully created the task, we can navigate back to the previous view
	return m.ctx.Back(), nil
}
