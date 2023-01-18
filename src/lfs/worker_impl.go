package lfs

import (
	"fmt"
	"github.com/Xanonymous-GitHub/lfs-downloader/pkg/utils"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/api"
	"github.com/Xanonymous-GitHub/lfs-downloader/src/common"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

type workerImpl struct {
	logger   *logrus.Logger
	client   api.HttpClient
	fileInfo *common.RepositoryFileInfo
}

func CreateWorker(
	logger *logrus.Logger,
	client api.HttpClient,
	fileInfo common.RepositoryFileInfo,
) Worker {
	return &workerImpl{
		logger:   logger,
		client:   client,
		fileInfo: &fileInfo,
	}
}

func (w *workerImpl) RetrieveRawContent() (*common.LFSDownloadableObj, error) {
	reqUrl := w.getRawFileReqUrlString()
	w.logger.Infoln(reqUrl)
	resp, err := w.client.Get(reqUrl)

	if err != nil {
		w.logger.Errorln(err)
		return nil, err
	}

	return w.parseLFSReferenceFile(resp.String())
}

func (w *workerImpl) parseLFSReferenceFile(raw string) (*common.LFSDownloadableObj, error) {
	oid := utils.ExtractSpecificStringSliceFrom(raw, "sha256:", "\n")
	if oid == nil {
		return nil, fmt.Errorf("oid parse failed")
	}

	size := utils.ExtractSpecificStringSliceFrom(raw, "size ", "\n")
	if size == nil {
		return nil, fmt.Errorf("size parse failed")
	}

	return &common.LFSDownloadableObj{
		Oid:  *oid,
		Size: *size,
	}, nil
}

func (w *workerImpl) getRawFileReqUrlString() string {
	var gitRemoteDomain *string
	gitRemoteDomain = w.fileInfo.SpecificDomain
	if gitRemoteDomain == nil {
		gitRemoteDomain = utils.ToPointerType("github.com")
	}

	// refer to https://docs.github.com/en/repositories/working-with-files/using-files/viewing-a-file.
	path := fmt.Sprintf("/%s/%s/raw/%s/%s", w.fileInfo.OrgName, w.fileInfo.Name, w.fileInfo.Head, w.fileInfo.Path)

	return utils.ToPointerType(
		url.URL{
			Scheme: "https",
			Host:   *gitRemoteDomain,
			Path:   strings.ReplaceAll(path, "//", "/"),
		},
	).String()
}
