package updateApiClient

import (
	"reflect"
	"testing"
)

func TestClient_GetUpdateInfo(t *testing.T) {
	type fields struct {
		BaseURL  string
		Username string
		Password string
	}
	type args struct {
		programName                  string
		version                      string
		updateTypeAndPlatformVersion []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UpdateInfoResponse
		wantErr bool
	}{
		{
			"test",
			fields{
				BaseURL: "https://update-api.1c.ru",
			},
			args{
				"Accounting",
				"3.0.88.22",
				[]string{
					"NewProgramOrRedaction",
					"8.3.15.2107",
				},
			},
			UpdateInfoResponse{
				ErrorResponse{},
				ConfigurationUpdateInfo{
					ConfigurationVersion: "3.0.89.43",
					Size:                 122421513,
					PlatformVersion:      "8.3.15.2107",
					UpdateInfoUrl:        "https://dl03.1c.ru/content//AutoUpdatesFiles/Accounting/3_0_89_43/82/News1cv8.htm",
					HowToUpdateInfoUrl:   "",
					UpgradeSequence: []string{
						"2ade0a5e-6469-4d48-99b5-da940e04bd41",
						"7e3b30ae-112c-452f-9068-e401e985a70e",
					},
					ProgramVersionUin: "f78f38d1-cdc3-49d8-9c7d-99fb3f992b51",
				},
				PlatformUpdateInfo{
					PlatformVersion:   "8.3.17.1851",
					TransitionInfoUrl: "https://dl03.1c.ru/content//AutoUpdatesFiles/Platform/8_3_17_1851/V8Update.htm",
					ReleaseUrl:        "https://releases.1c.ru/version_files?nick=Platform83&ver=8.3.17.1851&needAccessToken=true",
					DistributionUin:   "f44be38c-754e-4a76-90db-af06f2f381ac",
					Size:              595401719,
					Recommended:       true,
				},
				nil,
			},
			false,
		},

		{
			"error",
			fields{
				BaseURL: "https://update-api.1c.ru",
			},
			args{
				"Accounting",
				"3.0.88.22",
				[]string{NewProgramOrRedactionUpdateType, ""},
			},
			UpdateInfoResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				BaseURL:  tt.fields.BaseURL,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			got, err := c.GetUpdateInfo(tt.args.programName, tt.args.version, tt.args.updateTypeAndPlatformVersion...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpdateInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdateInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUpdate(t *testing.T) {
	type fields struct {
		BaseURL  string
		Username string
		Password string
	}
	type args struct {
		programVersionUin string
		UpgradeSequence   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UpdateResponse
		wantErr bool
	}{
		{
			"simpe",
			fields{
				BaseURL: "https://update-api.1c.ru",
			},
			args{
				"f78f38d1-cdc3-49d8-9c7d-99fb3f992b51",
				[]string{
					"2ade0a5e-6469-4d48-99b5-da940e04bd41",
					"7e3b30ae-112c-452f-9068-e401e985a70e",
				},
			},
			UpdateResponse{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				BaseURL:  tt.fields.BaseURL,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			got, err := c.GetUpdate(tt.args.programVersionUin, tt.args.UpgradeSequence)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
