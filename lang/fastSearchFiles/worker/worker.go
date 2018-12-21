package worker

import (
	"awesomeProject/lang/fastSearchFiles/models"
	"io/ioutil"
	"strings"
)

func FindFiles(workers typesF.Workers) (typesF.WorkerResult, error) {
	var filesResult typesF.WorkerResult

	dir_list, _ := ioutil.ReadDir(workers.Fliepath)
	for _, v := range dir_list {
		if strings.Contains(v.Name(), "business") {
			//fmt.Println("已找到", v.Name(), "文件，目录为：", workers.Fliepath, )
			filesResult.FindsName = append(filesResult.FindsName, workers.Fliepath+v.Name())
		}
		if v.IsDir() {
			str := workers.Fliepath + v.Name() + "/"
			filesResult.FilesName = append(filesResult.FilesName, typesF.Workers{
				Fliepath:   str,
				WorderFunc: FindFiles,
			})
		}
	}
	return filesResult, nil
}
