package updateApiClient

import (
	"github.com/monaco-io/request"
)

const (
	updateInfoPath = "/update-platform/programs/update/info"

	NewConfigurationAndOrPlatformUpdateType = "NewConfigurationAndOrPlatform" // - РабочееОбновление
	NewProgramOrRedactionUpdateType         = "NewProgramOrRedaction"         // - ПереходНаДругуюПрограммуИлиРедакцию
	NewPlatformUpdateType                   = "NewPlatform"                   // Новая платформа

	defaultPlatformVersion = "8.3.3.641"
)

type UpdateInfoRequest struct {
	ProgramVersion
	UpdateType      string `json:"updateType"`
	PlatformVersion string `json:"platformVersion"`
}

func (c *Client) GetUpdateInfo(programName, version string, updateTypeAndPlatformVersion ...string) (UpdateInfoResponse, error) {

	var result UpdateInfoResponse

	updateType := NewProgramOrRedactionUpdateType
	platformVersion := defaultPlatformVersion

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
		updateInfoPath,
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
	ErrorResponse
	ConfigurationUpdate  ConfigurationUpdateInfo `json:"configurationUpdateResponse"`
	PlatformUpdate       PlatformUpdateInfo      `json:"platformUpdateResponse"`
	AdditionalParameters map[string]string       `json:"additionalParameters"`
}

func (c UpdateInfoResponse) Error() error {

	if len(c.ErrorName) == 0 {
		return nil
	}

	return c.ErrorResponse
}

type ConfigurationUpdateInfo struct {
	ConfigurationVersion string   `json:"configurationVersion"`
	Size                 int      `json:"size"`
	PlatformVersion      string   `json:"platformVersion"`
	UpdateInfoUrl        string   `json:"updateInfoUrl"`
	HowToUpdateInfoUrl   string   `json:"howToUpdateInfoUrl"`
	UpgradeSequence      []string `json:"upgradeSequence"`
	ProgramVersionUin    string   `json:"programVersionUin"`
}

type PlatformUpdateInfo struct {
	PlatformVersion   string `json:"platformVersion"`
	TransitionInfoUrl string `json:"transitionInfoUrl"`
	ReleaseUrl        string `json:"releaseUrl"`
	DistributionUin   string `json:"distributionUin"`
	Size              int    `json:"size"`
	Recommended       bool   `json:"recommended"`
}
