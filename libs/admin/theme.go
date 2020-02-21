package admin

// ThemeInterface theme interface
type ThemeInterface interface {
	GetName() string
	GetViewPaths() []string
	ConfigAdminTheme(*Resource)
}

// Theme resource theme struct
type Theme struct {
	Name      string
	ViewPaths []string
}

// GetName get name from theme
func (theme Theme) GetName() string {
	return theme.Name
}

// GetViewPaths get view paths from theme
func (theme Theme) GetViewPaths() []string {
	return theme.ViewPaths
}

// ConfigAdminTheme config theme for admin resource
func (theme Theme) ConfigAdminTheme(*Resource) {
	return
}
