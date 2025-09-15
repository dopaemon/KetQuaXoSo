package utils

import (
	"context"
	"errors"
	_ "fmt"
	"os"

	_ "github.com/dopaemon/KetQuaXoSo/internal/configs"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var (
	runCLI	bool
	runGUI	bool
	runAPI	bool
	result	string
)

func GenFlags() string {
	cmd := &cobra.Command{
		Use:   "KetQuaXoSo",
		Short: "Xem kết quả xổ số kiến thiết (CLI/GUI/API)",
		Long:  `Chạy chế độ GUI (mặc định) hoặc CLI (--cli).`,
		Example: `
  # GUI mặc định:
  KetQuaXoSo

  # Chạy CLI:
  KetQuaXoSo --cli

  # Chạy GUI:
  KetQuaXoSo --gui

  # Chạy API:
  KetQuaXoSo --api
		`,
		RunE: func(c *cobra.Command, args []string) error {
			if runCLI && runGUI && runAPI {
				return errors.New("không thể dùng đồng thời --cli và --gui và --api")
			}
			switch {
			case runCLI:
				result = "cli"
			case runGUI:
				result = "gui"
			case runAPI:
				result = "api"
			default:
				result = "gui"
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&runCLI, "cli", false, "Chạy chế độ CLI")
	cmd.Flags().BoolVar(&runGUI, "gui", false, "Chạy chế độ GUI")
	cmd.Flags().BoolVar(&runAPI, "api", false, "Chạy chế độ API")

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
