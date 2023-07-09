/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/schimini/chad-generator/pkg/gen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chad",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		image, _ := cmd.Flags().GetString("image")
		text, _ := cmd.Flags().GetString("text")
		font, _ := cmd.Flags().GetString("font")
		out, _ := cmd.Flags().GetString("out")
		generator := gen.NewGenerator(image, font)
		generator.Generate(text, out)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chad.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("image", "i", "images/5.jpeg", "Background image path")
	rootCmd.Flags().StringP("text", "t", "Alpha males don't wait for opportunities; they create them. Like lions, they lead, relentless in their pursuit of success. Mediocrity is for sheep; excellence, the hallmark of the alpha.", "Overlay text")
	rootCmd.Flags().StringP("font", "f", "fonts/DancingScript-Regular.ttf", "Path of the font to be used")
	rootCmd.Flags().StringP("out", "o", "out.png", "Path of the output image")
}
