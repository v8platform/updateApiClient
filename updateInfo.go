package update_api_1c

import (
	"github.com/monaco-io/request"
)

const (
	getUpdateInfoPath = "update-platform/programs/update/info"

	NewProgramOrRedactionUpdateType = "NewProgramOrRedaction"
)

type UpdateInfoRequest struct {
	ProgramVersion
	UpdateType      string `json:"updateType"`
	PlatformVersion string `json:"platformVersion"`
}

func (c *Client) GetUpdateInfo(programName, version string, updateTypeAndPlatformVersion ...string) (UpdateInfoResponse, error) {

	var result UpdateInfoResponse

	updateType := NewProgramOrRedactionUpdateType
	platformVersion := ""

	if len(updateTypeAndPlatformVersion) == 1 {
		updateType = updateTypeAndPlatformVersion[0]
	}

	if len(updateTypeAndPlatformVersion) == 2 {
		platformVersion = updateTypeAndPlatformVersion[1]
	}

	data := UpdateInfoRequest{
		ProgramVersion{programName,
			version},
		updateType,
		platformVersion,
	}

	resp, err := c.doRequest(apiRequest{
		getUpdateInfoPath,
		request.POST,
		data,
	})

	if err != nil {
		return result, err
	}

	resp.Scan(&result)
	return result, result.Error()

}

type UpdateInfoResponse struct {
	*ErrorResponse
	ConfigurationUpdateResponse ConfigurationUpdateResponse `json:"configurationUpdateResponse"`
	PlatformUpdateResponse      PlatformUpdateResponse      `json:"platformUpdateResponse"`

	AdditionalParameters map[string]string `json:"additionalParameters"`
}

func (c UpdateInfoResponse) Error() error { return c.ErrorResponse }

type ConfigurationUpdateResponse struct {
	ConfigurationVersion string   `json:"configurationVersion"`
	Size                 int      `json:"size"`
	PlatformVersion      string   `json:"platformVersion"`
	UpdateInfoUrl        string   `json:"updateInfoUrl"`
	HowToUpdateInfoUrl   string   `json:"howToUpdateInfoUrl"`
	UpgradeSequence      []string `json:"upgradeSequence"`
	ProgramVersionUin    string   `json:"programVersionUin"`
}

type PlatformUpdateResponse struct {
	PlatformVersion   string `json:"platformVersion"`
	TransitionInfoUrl string `json:"transitionInfoUrl"`
	ReleaseUrl        string `json:"releaseUrl"`
	DistributionUin   string `json:"distributionUin"`
	Size              int    `json:"size"`
	Recommended       bool   `json:"recommended"`
}
