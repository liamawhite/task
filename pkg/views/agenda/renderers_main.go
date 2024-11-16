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

package agenda

import (
	"log/slog"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/notedownorg/notedown/pkg/providers/tasks"
	"github.com/notedownorg/task/pkg/components/groupedlist"
	"github.com/notedownorg/task/pkg/themes"
)

func mainRendererFuncs(theme themes.Theme, dateRetriever func() time.Time) groupedlist.Renderers[tasks.Task] {
	paddingHorizontal := 2

	return groupedlist.Renderers[tasks.Task]{
		Header: func(name string, width int) string {
			bg := func(s string) lipgloss.Color {
				switch s {
				case "Doing":
					return theme.Green
				case "Todo":
					return theme.Text
				case "Blocked":
					return theme.Red
				default:
					slog.Warn("unexpected task status", "status", s)
					return theme.Text
				}
			}(name)

			return lipgloss.JoinVertical(lipgloss.Top,
				s().Margin(0, 0, 1, 0).
					Background(bg).
					Foreground(theme.TextCursor).
					Bold(true).
					Padding(0, paddingHorizontal).
					Render(strings.ToUpper(name)),
				s().Width(width).Background(theme.Panel).
					Render(""),
			)
		},

		Footer: func(name string, width int) string {
			return lipgloss.JoinVertical(lipgloss.Bottom,
				s().Width(width).Background(theme.Panel).
					Render(""),
				"",
			)
		},

		Item: func(task tasks.Task, width int) string {
			bg, fg, err := colors(theme, task)
			if err != nil {
				slog.Warn("unexpected task status", "status", task.Status())
				return ""
			}

			right := buildRight(false)(theme, task, dateRetriever, bg)

			remainingSpace := width - w(right) - 2*paddingHorizontal
			left := buildLeft(false)(theme, task, remainingSpace, bg)

			middlePadding := width - w(left) - w(right) - 2*paddingHorizontal
			middle := s().Background(bg).PaddingRight(middlePadding).Render("") // fill out the rest of the space

			return s().Width(width).Padding(0, paddingHorizontal).Background(bg).Foreground(fg).Render(left + middle + right)
		},

		Selected: func(task tasks.Task, width int) string {
			bg, fg, err := selectedColors(theme, task)
			if err != nil {
				slog.Warn("unexpected task status", "status", task.Status())
				return ""
			}

			right := buildRight(true)(theme, task, dateRetriever, bg)

			remainingSpace := width - w(right) - 2*paddingHorizontal
			left := buildLeft(true)(theme, task, remainingSpace, bg)

			middlePadding := width - w(left) - w(right) - 2*paddingHorizontal
			middle := s().Background(bg).PaddingRight(middlePadding).Render("") // fill out the rest of the space

			return s().Width(width).Padding(0, paddingHorizontal).Background(bg).Foreground(fg).Render(left + middle + right)
		},
	}
}