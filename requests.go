package update_api_1c

type PatchesInfoRequest struct {
	ProgramVersion
	ProgramVersionList   []ProgramVersion `json:"programVersionList"`
	InstalledPatchesList []string         `json:"installedPatchesList"`
}
