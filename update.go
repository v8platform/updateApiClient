package updateApiClient

import (
	"github.com/monaco-io/request"
	"io"
)

const (
	updatePath = "/update-platform/programs/update/"
)

type ConfigurationUpdateData struct {
	TemplatePath         string `json:"templatePath"`
	ExecuteUpdateProcess bool   `json:"executeUpdateProcess"`
	UpdateFileUrl        string `json:"updateFileUrl"`
	UpdateFileName       string `json:"updateFileName"`
	UpdateFileFormat     string `json:"updateFileFormat"`
	Size                 int    `json:"size"`
	HashSum              string `json:"hashSum"`
}

type ConfigurationUpdateFile struct {
	io.ReadCloser
	ConfigurationUpdateData
}

type PlatformUpdateFile struct {
	io.ReadCloser
}

type UpdateRequest struct {
	ProgramVersionUin       string   `json:"programVersionUin"`
	UpgradeSequence         []string `json:"upgradeSequence"`
	PlatformDistributionUin string   `json:"platformDistributionUin"`
	Login                   string   `json:"login"`
	Password                string   `json:"password"`
}

func (c *Client) GetUpdate(programVersionUin string, UpgradeSequence []string) (UpdateResponse, error) {

	var result UpdateResponse
	data := UpdateRequest{
		programVersionUin,
		UpgradeSequence,
		"",
		c.Username,
		c.Password,
	}

	resp, err := c.doRequest(apiRequest{
		updatePath,
		request.POST,
		data,
	})

	if err != nil {
		return result, err
	}

	resp.Scan(&result)
	return result, result.Error()

}

func (c *Client) GetPlatformUpdate(platformDistributionUin string) (UpdateResponse, error) {

	var result UpdateResponse
	data := UpdateRequest{
		"",
		[]string{},
		platformDistributionUin,
		c.Username,
		c.Password,
	}

	resp, err := c.doRequest(apiRequest{
		updatePath,
		request.POST,
		data,
	})

	if err != nil {
		return result, err
	}

	resp.Scan(&result)
	return result, result.Error()

}

func (c *Client) GetConfigurationUpdateData(data ConfigurationUpdateData) (*ConfigurationUpdateFile, error) {

	resp, err := c.doFileRequest(data.UpdateFileUrl)

	if err != nil {
		return nil, err
	}

	return &ConfigurationUpdateFile{
		resp,
		data,
	}, nil
}

func (c *Client) GetPlatformUpdateData(data UpdateResponse) (*PlatformUpdateFile, error) {

	resp, err := c.doFileRequest(data.PlatformDistributionUrl)

	if err != nil {
		return nil, err
	}

	return &PlatformUpdateFile{
		resp,
	}, nil
}

type UpdateResponse struct {
	ErrorResponse
	ConfigurationUpdateDataList []ConfigurationUpdateData `json:"configurationUpdateDataList"`
	PlatformDistributionUrl     string                    `json:"platformDistributionUrl"`
	AdditionalParameters        map[string]string         `json:"additionalParameters"`
}

func (c UpdateResponse) Error() error {
	if len(c.ErrorName) == 0 {
		return nil
	}
	return c.ErrorResponse
}
