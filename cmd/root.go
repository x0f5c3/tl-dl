package cmd

import (
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/x0f5c3/tl-dl/pkg/products"
	"os"
	"os/signal"
)

var rootCmd = &cobra.Command{
	Use:     "toolbox-download [DOWNLOAD_DIR]",
	Short:   "A tool and a library to download the jetbrains-toolbox",
	Version: "v0.0.7", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	RunE: runFunc,
	Args: cobra.ExactArgs(1),
}

func runFunc(_ *cobra.Command, args []string) error {
	b, err := products.DownloadNative()
	if err != nil {
		return err
	}
	return b.Save(args[0])
}

func handleUpdate() {
	err := pcli.CheckForUpdates()
	if err != nil {
		pterm.Error.Printf("Update check failed: %s\n", err)
		os.Exit(1)
	}
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
		handleUpdate()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		handleUpdate()
		os.Exit(1)
	}

	handleUpdate()
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "n", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	_ = pcli.SetRepo("x0f5c3/tl-dl")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
