package views

import (
	"context"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/KalinduGandara/erp-system/internal/app/services"
)

type MainWindow struct {
	window      fyne.Window
	userService *services.UserService
}

func NewMainWindow(userService *services.UserService) *MainWindow {
	mainApp := app.New()
	window := mainApp.NewWindow("ERP System")

	return &MainWindow{
		window:      window,
		userService: userService,
	}
}

func (m *MainWindow) createLoginScreen() fyne.CanvasObject {
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: username},
			{Text: "Password", Widget: password},
		},
		OnSubmit: func() {
			if m.userService.ValidateUser(context.Background(), username.Text, password.Text) {
				m.showMainContent()
			} else {
				dialog.ShowError(errors.New("Invalid credentials"), m.window)
			}
		},
	}
	return form
}

func (m *MainWindow) showMainContent() {
	tabs := container.NewAppTabs(
		container.NewTabItem("Dashboard", m.createDashboardTab()),
		container.NewTabItem("Users", m.createUsersTab()),
		// container.NewTabItem("Customers", m.createCustomersTab()),
		// container.NewTabItem("Products", m.createProductsTab()),
	)

	m.window.SetContent(tabs)
}

func (m *MainWindow) createDashboardTab() fyne.CanvasObject {
	return widget.NewLabel("Welcome to ERP System")
}

func (m *MainWindow) createUsersTab() fyne.CanvasObject {
	// Basic user management UI
	return container.NewVBox(
		widget.NewLabel("Users Management"),
		widget.NewButton("Add New User", func() {
			// TODO: Implement add user dialog
		}),
	)
}

func (m *MainWindow) Show() {
	m.window.SetContent(m.createLoginScreen())
	m.window.Resize(fyne.NewSize(800, 600))
	m.window.ShowAndRun()
}
