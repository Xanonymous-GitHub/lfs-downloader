package lfs

import "github.com/Xanonymous-GitHub/lfs-downloader/src/common"

// Worker is an implementation according to https://gist.github.com/fkraeutli/66fa741d9a8c2a6a238a01d17ed0edc5.
type Worker interface {
	RetrieveRawContent() (*common.LFSDownloadableObj, error)
}
