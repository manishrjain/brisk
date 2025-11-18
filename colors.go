package main

import "github.com/charmbracelet/lipgloss"

// Monokai Pro color scheme
var (
	MonokaiPink   = lipgloss.Color("#FF6188") // Primary accent - titles, focused elements
	MonokaiOrange = lipgloss.Color("#FC9867") // Secondary accent - group headers
	MonokaiCyan   = lipgloss.Color("81")      // Tertiary accent - labels, table headers
	MonokaiBorder = lipgloss.Color("238")     // Borders

	// Adaptive colors for different terminal backgrounds
	MonokaiAdaptiveText = lipgloss.AdaptiveColor{
		Light: "240", // Dark gray for light backgrounds
		Dark:  "255", // Bright white for dark backgrounds
	}
	MonokaiGrey = lipgloss.AdaptiveColor{
		Light: "243", // Medium gray for light backgrounds (readable)
		Dark:  "250", // Light gray for dark backgrounds
	}
)
