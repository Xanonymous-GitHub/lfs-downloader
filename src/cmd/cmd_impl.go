package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cobraCmdWrapperImpl struct {
	rawCobraRootCmd *cobra.Command
	logger          *logrus.Logger
}

func CreateRootCmd(logger *logrus.Logger, voidCallback func(args []string)) Cmd {
	rootCmd := &cobra.Command{
		Run: func(_ *cobra.Command, args []string) {
			voidCallback(args)
		},
	}

	return &cobraCmdWrapperImpl{
		rawCobraRootCmd: rootCmd,
		logger:          logger,
	}
}

func (wrapper cobraCmdWrapperImpl) Execute() error {
	err := wrapper.rawCobraRootCmd.Execute()

	if err != nil {
		wrapper.logger.Warningln(err)
	}

	return err
}
