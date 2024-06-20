package views_themepicker

type ThemeOption struct {
    Label string
    Theme string
    IsDark bool
}

    var DefaultThemes = []ThemeOption{
	{
		Label:  "Cupcake",
		Theme:  "cupcake",
		IsDark: false,
	},
	{
		Label:  "Cyber Punk",
		Theme:  "cyberpunk",
		IsDark: false,
	},
	{
		Label:  "Valentine",
		Theme:  "valentine",
		IsDark: false,
	},
	{
		Label:  "Nord",
		Theme:  "nord",
		IsDark: false,
	},
	{
		Label:  "Synth Wave",
		Theme:  "synthwave",
		IsDark: true,
	},
	{
		Label:  "Dracula",
		Theme:  "dracula",
		IsDark: true,
	},
	{
		Label:  "Sunset",
		Theme:  "sunset",
		IsDark: true,
	},
    {
        Label: "Dim",
        Theme: "dim",
        IsDark: true,
    },
}
