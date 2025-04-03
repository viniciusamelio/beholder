package bucket

type Bucket interface {
	DownloadFile(objectKey string, fileName string) error
	UploadFile(objectKey string, filename string, filetype string) error
}
