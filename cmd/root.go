package cmd

// TODO 考虑添加本地配置文件

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "gitf",
	Short: "gitf is also git, but for gltf",
	Long: `GitF by Herbert He @ <https://github.com/gitf-project/gitf>
A version control tool like git.
gitf is also git, but for gltf.
gitf is an open source project under GPL v2`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version for gitf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Printf("gitf @ %s for %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCommand.AddCommand(versionCmd)
	rootCommand.Version = Version
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
