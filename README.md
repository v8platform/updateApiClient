# update-api-1c
Client for update-api.1c.ru

[![go.dev][pkg-img]][pkg] [![goreport][report-img]][report] [![build][build-img]][build] [![coverage][cov-img]][cov] ![stability-stable][stability-img]


## How to use

### Quick start

```go
package main

import (
	apiClient "github.com/v8platform/updateApiClient"

	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	
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

```

[pkg-img]: http://img.shields.io/badge/godoc-reference-5272B4.svg
[pkg]: https://godoc.org/github.com/v8platform/updateApiClient
[report-img]: https://goreportcard.com/badge/github.com/v8platform/updateApiClient
[report]: https://goreportcard.com/report/github.com/v8platform/updateApiClient
[build-img]: https://github.com/v8platform/updateApiClient/workflows/goreleaser/badge.svg
[build]: https://github.com/v8platform/updateApiClient/actions
[cov-img]: http://gocover.io/_badge/github.com/v8platform/updateApiClient
[cov]: https://gocover.io/github.com/v8platform/updateApiClient
[stability-img]: https://img.shields.io/badge/stability-stable-green.svg

