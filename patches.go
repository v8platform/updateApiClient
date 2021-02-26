package update_api_1c

import "github.com/monaco-io/request"

const (
	patchesGetInfoPath = "/update-platform/patches/getInfo"
)

type PatchesInfoRequest struct {
	ProgramVersionList   []ProgramVersion `json:"programVersionList"`
	InstalledPatchesList []string         `json:"installedPatchesList"`
}

type PatchesInfoResponse struct {
	*ErrorResponse
	PatchUpdateList []PatchUpdate `json:"patchUpdateList"`
}

func (c PatchesInfoResponse) Error() error {
	if len(c.ErrorName) == 0 {
		return nil
	}
	return c.ErrorResponse
}

type PatchUpdate struct {
	Ueid                string `json:"ueid"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	BuildDate           int64  `json:"buildDate"`
	ModificatedMetadata string `json:"modificatedMetadata"`
	Status              string `json:"status"`
	Size                int    `json:"size"`
	ApplyToVersion      []ProgramVersion `json:"applyToVersion"`
}

func (c *Client) GetPatchesInfo(programName, programVersion string, InstalledPatchesList... string) (PatchesInfoResponse, error) {

	return c.GetPatchesInfoRequest([]ProgramVersion{
		{programName,programVersion},
	}, InstalledPatchesList)

}

func (c *Client) GetPatchesInfoRequest(programVersionList []ProgramVersion, installedPatchesList []string) (PatchesInfoResponse, error) {

	var result PatchesInfoResponse
	data := PatchesInfoRequest{
		programVersionList,
		installedPatchesList,
	}

	resp, err := c.doRequest(apiRequest{
		patchesGetInfoPath,
		request.POST,
		data,
	})

	if err != nil {
		return result, err
	}

	resp.Scan(&result)
	return result, result.Error()

}
