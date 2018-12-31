package cmd

import (
    "github.com/spf13/cobra"
)

func NewFoobarCommand() *cobra.Command {
    var (
        flagVerbose bool
    )

    cmd := &cobra.Command{
        Use: "foobar",
    }

    cmd.AddCommand(
        newInitCommand(),
    )

    // `Persistent` ってつくやつはサブコマンドまで影響が及ぶやつ
    cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")

    return cmd
}