package update_api_1c

import (
	"bytes"
	"github.com/monaco-io/request"
	"io"
)

const (
	patchesGetFilesPath = "/update-platform/patches/getFiles"
)

type PatchesGetFilesRequest struct {
	PatchUinList []string `json:"patchUinList"`
	Login        string   `json:"login"`
	Password     string   `json:"password"`
}

type PatchesGetFilesResponse struct {
	ErrorResponse
	PatchDistributionDataList []PatchDistributionData `json:"patchDistributionDataList"`
}

type PatchDistributionData struct {
	PatchUeid     string `json:"patchUeid"`
	PatchFileUrl  string `json:"patchFileUrl"`
	PatchFileName string `json:"patchFileName"`
	Size          int    `json:"size"`
	HashSum       string `json:"hashSum"`
}

type PatchDistributionFile struct {
	io.Reader
	PatchDistributionData
}

func (c PatchesGetFilesResponse) Error() error {
	if len(c.ErrorName) == 0 {
		return nil
	}
	return c.ErrorResponse
}

func (c *Client) GetPatchesFilesInfo(patchUinList ...string) (PatchesGetFilesResponse, error) {

	var result PatchesGetFilesResponse
	data := PatchesGetFilesRequest{
		patchUinList,
		c.Username,
		c.Password,
	}

	resp, err := c.doRequest(apiRequest{
		patchesGetFilesPath,
		request.POST,
		data,
	})

	if err != nil {
		return result, err
	}

	resp.Scan(&result)
	return result, result.Error()

}

func (c *Client) GetPatchDistributionData(data PatchDistributionData) (*PatchDistributionFile, error) {

	type loginParams struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	resp, err := c.doRawRequest("", apiRequest{
		data.PatchFileUrl,
		request.POST,
		loginParams{
			c.Username,
			c.Password,
		},
	})

	if err != nil {
		return nil, err
	}

	return &PatchDistributionFile{
		bytes.NewBuffer(resp.Bytes()),
		data,
	}, nil
}
