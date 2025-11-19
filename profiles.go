package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const profilesDir = ".rentobuy_profiles"

// ensureProfilesDir creates the profiles directory if it doesn't exist
func ensureProfilesDir() error {
	return os.MkdirAll(profilesDir, 0755)
}

// listProfiles returns a sorted list of available profile names
func listProfiles() ([]string, error) {
	if err := ensureProfilesDir(); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(profilesDir)
	if err != nil {
		return nil, err
	}

	var profiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			// Remove .json extension
			name := strings.TrimSuffix(entry.Name(), ".json")
			profiles = append(profiles, name)
		}
	}

	sort.Strings(profiles)
	return profiles, nil
}

// loadProfile loads inputs from a named profile
func loadProfile(name string) (map[string]string, error) {
	profilePath := filepath.Join(profilesDir, name+".json")
	data, err := os.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}

	var inputs map[string]string
	err = json.Unmarshal(data, &inputs)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// saveProfile saves inputs to a named profile
func saveProfile(name string, inputs map[string]string) error {
	if err := ensureProfilesDir(); err != nil {
		return err
	}

	profilePath := filepath.Join(profilesDir, name+".json")
	data, err := json.MarshalIndent(inputs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(profilePath, data, 0644)
}

// deleteProfile deletes a named profile
func deleteProfile(name string) error {
	profilePath := filepath.Join(profilesDir, name+".json")
	return os.Remove(profilePath)
}

// showProfileMenu displays the profile selection menu and returns the selected profile name or empty for new
func showProfileMenu() (string, bool, error) {
	re := lipgloss.NewRenderer(os.Stdout)
	titleStyle := re.NewStyle().Foreground(MonokaiPink).Bold(true)
	labelStyle := re.NewStyle().Foreground(MonokaiCyan)
	itemStyle := re.NewStyle().Foreground(MonokaiAdaptiveText)

	fmt.Println()
	fmt.Println(titleStyle.Render("RENT vs BUY CALCULATOR"))
	fmt.Println()

	profiles, err := listProfiles()
	if err != nil {
		return "", false, err
	}

	if len(profiles) == 0 {
		fmt.Println(labelStyle.Render("No saved profiles found. Starting new calculation..."))
		fmt.Println()
		return "", false, nil
	}

	fmt.Println(labelStyle.Render("Available profiles:"))
	for i, profile := range profiles {
		fmt.Printf("  %s %d. %s\n", itemStyle.Render(""), i+1, profile)
	}
	fmt.Println()

	fmt.Println(labelStyle.Render("Options:"))
	fmt.Printf("  %s Enter profile number to load\n", itemStyle.Render(""))
	fmt.Printf("  %s Press 'n' for new calculation\n", itemStyle.Render(""))
	fmt.Printf("  %s Press 'd' to delete a profile\n", itemStyle.Render(""))
	fmt.Printf("  %s Press 'q' to quit\n", itemStyle.Render(""))
	fmt.Println()

	fmt.Print(labelStyle.Render("Your choice: "))
	var choice string
	fmt.Scanln(&choice)

	choice = strings.TrimSpace(strings.ToLower(choice))

	if choice == "q" {
		return "", true, nil
	}

	if choice == "n" || choice == "" {
		return "", false, nil
	}

	if choice == "d" {
		return handleDeleteProfile(profiles)
	}

	// Try to parse as number
	var num int
	_, err = fmt.Sscanf(choice, "%d", &num)
	if err != nil || num < 1 || num > len(profiles) {
		fmt.Println("Invalid choice. Starting new calculation...")
		return "", false, nil
	}

	return profiles[num-1], false, nil
}

// handleDeleteProfile handles the profile deletion flow
func handleDeleteProfile(profiles []string) (string, bool, error) {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(MonokaiCyan)

	fmt.Println()
	fmt.Print(labelStyle.Render("Enter profile number to delete (or 'c' to cancel): "))
	var choice string
	fmt.Scanln(&choice)

	choice = strings.TrimSpace(strings.ToLower(choice))
	if choice == "c" || choice == "" {
		return showProfileMenu() // Return to main menu
	}

	var num int
	_, err := fmt.Sscanf(choice, "%d", &num)
	if err != nil || num < 1 || num > len(profiles) {
		fmt.Println("Invalid choice. Returning to menu...")
		return showProfileMenu()
	}

	profileName := profiles[num-1]
	fmt.Printf("Delete profile '%s'? (y/n): ", profileName)
	var confirm string
	fmt.Scanln(&confirm)

	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
		if err := deleteProfile(profileName); err != nil {
			fmt.Printf("Error deleting profile: %v\n", err)
		} else {
			fmt.Printf("Profile '%s' deleted.\n", profileName)
		}
	}

	return showProfileMenu() // Return to main menu
}

// promptSaveProfile prompts the user to save the current configuration
func promptSaveProfile(inputs map[string]string) {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(MonokaiCyan)

	fmt.Println()
	fmt.Print(labelStyle.Render("Save this configuration? (y/n): "))
	var save string
	fmt.Scanln(&save)

	if strings.ToLower(strings.TrimSpace(save)) != "y" {
		return
	}

	fmt.Print(labelStyle.Render("Enter profile name: "))
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Profile name cannot be empty. Configuration not saved.")
		return
	}

	// Check if profile already exists
	profiles, _ := listProfiles()
	for _, profile := range profiles {
		if profile == name {
			fmt.Printf("Profile '%s' already exists. Overwrite? (y/n): ", name)
			var overwrite string
			fmt.Scanln(&overwrite)
			if strings.ToLower(strings.TrimSpace(overwrite)) != "y" {
				fmt.Println("Configuration not saved.")
				return
			}
			break
		}
	}

	if err := saveProfile(name, inputs); err != nil {
		fmt.Printf("Error saving profile: %v\n", err)
	} else {
		fmt.Printf("Configuration saved as '%s'\n", name)
	}
}
