/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows all of the tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your ToDo's for today")
		showPrettyTodoList()
		fmt.Println("show called")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showPrettyTodoList() {
	pinkHeader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFB6C1")).
		Bold(true)

	purpleId := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E6E6FA"))

	mintActive := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#98FB98")).
		Bold(true)

	rosePause := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFC0CB"))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{
		pinkHeader.Render("#ID"),
		pinkHeader.Render("Name"),
		pinkHeader.Render("Status"),
	})

	t.AppendRows([]table.Row{
		{purpleId.Render("01"), "Cry Yourself to Sleep", mintActive.Render("● done")},
		{purpleId.Render("02"), "Build this CLI Toolp", mintActive.Render("○ in Progress")},
		{purpleId.Render("03"), "Sleep", rosePause.Render("○ not Started")},
	})

	cuteStyle := table.StyleRounded

	frameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#D8BFD8"))

	cuteStyle.Box.Left = frameStyle.Render("│")
	cuteStyle.Box.Right = frameStyle.Render("│")
	cuteStyle.Box.TopLeft = frameStyle.Render("─")
	cuteStyle.Box.TopRight = frameStyle.Render("─")
	cuteStyle.Box.BottomLeft = frameStyle.Render("─")
	cuteStyle.Box.BottomRight = frameStyle.Render("─")
	cuteStyle.Box.MiddleHorizontal = frameStyle.Render("─")
	cuteStyle.Box.MiddleVertical = frameStyle.Render("│")
	cuteStyle.Box.TopLeft = frameStyle.Render("╭")
	cuteStyle.Box.TopRight = frameStyle.Render("╮")
	cuteStyle.Box.BottomLeft = frameStyle.Render("╰")
	cuteStyle.Box.BottomRight = frameStyle.Render("╯")
	cuteStyle.Box.TopSeparator = frameStyle.Render("┬")
	cuteStyle.Box.BottomSeparator = frameStyle.Render("┴")
	cuteStyle.Box.LeftSeparator = frameStyle.Render("├")
	cuteStyle.Box.RightSeparator = frameStyle.Render("┤")
	cuteStyle.Box.PageSeparator = frameStyle.Render("┼")

	t.SetStyle(cuteStyle)

	t.Render()
}
