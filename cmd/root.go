package cmd

import (
	"fmt"
	"github.com/x0f5c3/tl-dl/pkg/products"
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "toolbox-download",
	Short:   "A tool and a library to download the jetbrains-toolbox",
	Version: "v0.0.1", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	RunE: runFunc,
	Args: cobra.ExactArgs(1),
}

func runFunc(cmd *cobra.Command, args []string) error {
	prod, err := products.GetToolbox()
	if err != nil {
		pterm.Fatal.Sprintf("Failed to retrive toolbox data: %s\n", err)
		return err
	}
	rel, err := prod.LatestRelease()
	if err != nil {
		return err
	}
	b, err := rel.Download()
	if err != nil {
		return err
	}
	err = b.Data.Save(fmt.Sprintf("%s/%s", args[0], b.Data.FileName))
	if err != nil {
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		if err := pcli.CheckForUpdates(); err != nil {
			pterm.Fatal.Sprintf("Failed to check for updates: %s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		err := pcli.CheckForUpdates()
		pterm.Fatal.Sprintf("Failed to check for updates: %s\n", err)
		os.Exit(1)
	}

	if err := pcli.CheckForUpdates(); err != nil {
		pterm.Fatal.Sprintf("Failed to check for updates: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("x0f5c3/tl-dl")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
