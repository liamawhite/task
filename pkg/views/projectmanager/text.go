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

package projectmanager

import (
	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/notedownorg/task/pkg/context"
	"github.com/notedownorg/task/pkg/model"
)

type Text struct {
	base model.Base

	ctx *context.ProgramContext

	ti      textinput.Model
	focused bool
	IsValid bool
}

func NewText(ctx *context.ProgramContext, name string) *Text {
	ti := textinput.New()
	ti.Prompt = ""
	ti.SetValue(name)
	return &Text{
		ctx: ctx,
		ti:  ti,
	}
}

func (t *Text) Init() (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	t.ti, cmd = t.ti.Update(msg)
	return t, cmd
}

func (t *Text) Focus() *Text {
	t.ti.Focus()
	t.focused = true
	return t
}

func (t *Text) Blur() *Text {
	t.ti.Blur()
	t.focused = false
	return t
}

func (t Text) AtBeginning() bool {
	return t.ti.Position() == 0
}

func (t *Text) Width(i int) *Text {
	t.ti.Width = i
	return t
}

func (t Text) Value() string {
	return t.ti.Value()
}

func (t *Text) SetValue(s string) *Text {
	t.ti.SetValue(s)
	return t
}

func (t Text) Cursor() int {
	return t.ti.Position()
}

func (t *Text) SetCursor(i int) {
	t.ti.SetCursor(i)
}

func (s *Text) View() string {
	return s.ti.View()
}