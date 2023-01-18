package common

import "time"

type RepositoryFileInfo struct {
	// Repository name.
	Name string `json:"name,omitempty"`

	// The repository owner's identifier.
	OrgName string `json:"orgName,omitempty"`

	// The checkout point.
	// It may be a branch name or a sha string.
	Head string `json:"head,omitempty"`

	// The absolute path of the file from the repository's root.
	Path string `json:"path,omitempty"`

	// If provided, it means the special git remote domain provided by user instead of `github.com`.
	SpecificDomain *string `json:"specificDomain,omitempty"`
}

type LFSDownloadReq struct {
	Operation string               `json:"operation,omitempty"`
	Transfer  []string             `json:"transfer,omitempty"`
	Objects   []LFSDownloadableObj `json:"objects,omitempty"`
}

type LFSDownloadableObj struct {
	Oid     string               `json:"oid,omitempty"`
	Size    string               `json:"size,omitempty"`
	Actions map[string]LFSAction `json:"actions,omitempty"`
}

type LFSAction struct {
	Href      string    `json:"href,omitempty"`
	ExpiresAt time.Time `json:"expires_at"`
	ExpiresIn int       `json:"expires_in,omitempty"`
}
