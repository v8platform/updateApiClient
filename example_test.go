package update_api_1c_test

import (
	"github.com/k0kubun/pp"
	update_api_1c "github.com/v8platform/update-api-1c"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ExampleNewClient() {

	client := update_api_1c.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		update_api_1c.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(updateInfo)

}

func ExampleClient_GetUpdateInfo() {

	client := update_api_1c.NewClient("", "")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		update_api_1c.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(updateInfo)

}

func ExampleClient_GetUpdate() {

	client := update_api_1c.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		update_api_1c.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(updateInfo)

	updateData, err := client.GetUpdate(updateInfo.ConfigurationUpdate.ProgramVersionUin, updateInfo.ConfigurationUpdate.UpgradeSequence)

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(updateData)

}

func ExampleClient_GetConfigurationUpdateData() {
	client := update_api_1c.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		update_api_1c.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	updateData, err := client.GetUpdate(updateInfo.ConfigurationUpdate.ProgramVersionUin, updateInfo.ConfigurationUpdate.UpgradeSequence)

	if err != nil {
		log.Fatal(err)
	}

	for _, data := range updateData.ConfigurationUpdateDataList {

		updateDataFile, err := client.GetConfigurationUpdateData(data)

		if err != nil {
			log.Fatal(err)
		}

		distPath := strings.ReplaceAll(updateDataFile.TemplatePath, "\\", string(os.PathSeparator))

		f, err := os.Create(filepath.Join(".", distPath, updateDataFile.UpdateFileName))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		updateDataFile.Close()

		_, err = io.Copy(f, updateDataFile)
	}

}