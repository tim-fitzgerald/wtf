package spoke

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget for Spoke
type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	result   *RequestsArray
	settings *Settings
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	requestArray, err := widget.listRequests()
	widget.err = err
	widget.result = requestArray
	widget.SetItemCount(widget.result.Total)
	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s (%d)", widget.CommonSettings().Title, widget.result.Total)
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	items := widget.result.Results
	if len(items) == 0 {
		return title, "No unassigned tickets in queue - woop!!", false
	}

	str := ""
	for idx, data := range items {
		str += widget.format(data, idx)
	}

	return title, str, false
}

func (widget *Widget) format(request Request, idx int) string {
	textColor := widget.settings.common.Colors.Background
	if idx == widget.GetSelected() {
		textColor = widget.settings.common.Colors.BorderTheme.Focused
	}

	str := fmt.Sprintf(" [%s:]%s\n %s\n\n", textColor, request.Subject, request.Permalink)
	return str
}

func (widget *Widget) openRequest() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Results) {
		request := &widget.result.Results[sel]
		requestURL := request.Permalink
		utils.OpenFile(requestURL)
	}
}
