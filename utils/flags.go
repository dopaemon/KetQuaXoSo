package utils

import (
	"context"
	"errors"
	_ "fmt"
	"os"

	_ "KetQuaXoSo/internal/configs"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func GenFlags() string {
	var runCLI, runGUI bool
	var result string

	cmd := &cobra.Command{
		Use:   "KetQuaXoSo",
		Short: "Xem kết quả xổ số kiến thiết (CLI/GUI)",
		Long:  `Chạy chế độ GUI (mặc định) hoặc CLI (--cli).`,
		Example: `
  # GUI mặc định:
  KetQuaXoSo

  # Chạy CLI:
  KetQuaXoSo --cli

  # Chạy GUI rõ ràng:
  KetQuaXoSo --gui
		`,
		RunE: func(c *cobra.Command, args []string) error {
			if runCLI && runGUI {
				return errors.New("không thể dùng đồng thời --cli và --gui")
			}
			switch {
			case runCLI:
				result = "cli"
			case runGUI:
				result = "gui"
			default:
				result = "gui"
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&runCLI, "cli", false, "Chạy chế độ CLI")
	cmd.Flags().BoolVar(&runGUI, "gui", false, "Chạy chế độ GUI")

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.Version = ""
	cmd.SetVersionTemplate("")

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithNotifySignal(os.Interrupt, os.Kill),
	); err != nil {
		os.Exit(1)
	}

	return result
}
