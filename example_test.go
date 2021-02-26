package updateApiClient_test

import (
	apiClient "github.com/v8platform/updateApiClient"

	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ExampleNewClient() {

	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		apiClient.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(updateInfo)

}

func ExampleClient_GetUpdateInfo() {

	client := apiClient.NewClient("", "")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		apiClient.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(updateInfo)

}

func ExampleClient_GetUpdate() {

	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		apiClient.NewProgramOrRedactionUpdateType, "8.3.15.2107")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(updateInfo)

	updateData, err := client.GetUpdate(updateInfo.ConfigurationUpdate.ProgramVersionUin, updateInfo.ConfigurationUpdate.UpgradeSequence)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(updateData)

}

func ExampleClient_GetConfigurationUpdateData() {
	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	updateInfo, err := client.GetUpdateInfo("Accounting",
		"3.0.88.22",
		apiClient.NewProgramOrRedactionUpdateType, "8.3.15.2107")

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

		log.Println("Download:", updateDataFile.UpdateFileUrl)

		distPath := strings.ReplaceAll(updateDataFile.TemplatePath, "\\", string(os.PathSeparator))
		distPath = filepath.Join(".", distPath)
		log.Println("Path:", distPath)

		err = os.MkdirAll(distPath, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		f, err := ioutil.TempFile("", "."+updateDataFile.UpdateFileFormat)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, updateDataFile)

		f.Close()
		updateDataFile.Close()

		err = apiClient.UnzipFile(f.Name(), distPath)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func ExampleClient_GetPatchesInfo() {

	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	patchesInfo, err := client.GetPatchesInfo("Accounting",
		"3.0.88.22")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Get paches info for %d", len(patchesInfo.PatchUpdateList))

}

func ExampleClient_GetPatchesFilesInfo() {

	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	patchesInfo, err := client.GetPatchesInfo("Accounting",
		"3.0.88.22")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Get paches info for %d", len(patchesInfo.PatchUpdateList))

	var patchesList []string
	for _, update := range patchesInfo.PatchUpdateList {
		patchesList = append(patchesList, update.Ueid)
	}

	patchesFilesInfo, err := client.GetPatchesFilesInfo(patchesList...)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Get paches files info for %d", len(patchesFilesInfo.PatchDistributionDataList))

}

func ExampleClient_GetPatchDistributionData() {

	client := apiClient.NewClient("ITS_USER", "ITS_PASSWORD")

	patchesInfo, err := client.GetPatchesInfo("Accounting",
		"3.0.88.22")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Get paches info for %d", len(patchesInfo.PatchUpdateList))

	var patchesList []string
	patchesNames := make(map[string]string)
	for _, update := range patchesInfo.PatchUpdateList {
		patchesList = append(patchesList, update.Ueid)
		patchesNames[update.Ueid] = update.Name
	}

	patchesFilesInfo, err := client.GetPatchesFilesInfo(patchesList...)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Get paches files info for %d", len(patchesFilesInfo.PatchDistributionDataList))

	for _, data := range patchesFilesInfo.PatchDistributionDataList {
		fileInfo, err := client.GetPatchDistributionData(data)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Download:", fileInfo.PatchFileUrl)

		distPath := strings.ReplaceAll(fileInfo.PatchFileName, "\\", string(os.PathSeparator))
		distPath = filepath.Join(".", "patches", patchesNames[fileInfo.PatchUeid])
		log.Println("Path:", distPath)

		err = os.MkdirAll(distPath, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		f, err := ioutil.TempFile("", ".pzip")
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, fileInfo)

		f.Close()

		err = apiClient.UnzipFile(f.Name(), distPath)
		if err != nil {
			log.Fatal(err)
		}

	}

}
