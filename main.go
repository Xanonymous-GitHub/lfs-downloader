package main

import (
	"github.com/Xanonymous-GitHub/lfs-downloader/pkg/logrux"
	"github.com/Xanonymous-GitHub/lfs-downloader/pkg/utils"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/api"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/cmd"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/common"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/lfs"
	"github.com/sirupsen/logrus"
	"strings"
)

type userInteractionHandler struct {
	logger       *logrus.Logger
	parsedResult *common.RepositoryFileInfo
}

func (h *userInteractionHandler) handleProvidedCliParams(args []string) {
	length := len(args)
	if length < 4 {
		h.logger.Warningln("You should at least provide four values includes " +
			"'Repository Name', " +
			"'Repository owner's identifier', " +
			"'Checkout point', " +
			"'Absolute path of the file'.")
		return
	}

	trimmedArgs := utils.Map(args, strings.TrimSpace)

	var specificDomain *string

	if length >= 5 {
		specificDomain = utils.ToPointerType(trimmedArgs[4])
	}

	h.parsedResult = &common.RepositoryFileInfo{
		Name:           trimmedArgs[0],
		OrgName:        trimmedArgs[1],
		Head:           trimmedArgs[2],
		Path:           trimmedArgs[3],
		SpecificDomain: specificDomain,
	}
}

func main() {
	logger := logrux.NewLogger()
	handler := userInteractionHandler{logger: logger}
	cliCmd := cmd.CreateRootCmd(logger, handler.handleProvidedCliParams)

	if err := cliCmd.Execute(); err != nil {
		logger.Warningln(err)
		return
	}

	client := api.CreateHttpClient(logger)
	lfsWorker := lfs.CreateWorker(logger, client, *handler.parsedResult)

	_, err := lfsWorker.RetrieveRawContent()
	if err != nil {
		logger.Warningln(err)
		return
	}
}
