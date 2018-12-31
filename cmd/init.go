package cmd

import (
    "github.com/spf13/cobra"
)

func newInitCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "init",
        Args: cobra.ExactArgs(1), // 引数のバリデーションもできる
        RunE: func(_ *cobra.Command, args []string) error {
            // snip.
            return nil
        },
    }

    return cmd
}