package uadmin

type MinioConfig struct {
	endpoint        string
	accessKeyId     string
	secretAccessKey string
	useSSl          bool
	bucketName      string
	location        string
	policy          string
	isHttps         bool
}

func NewMinioConfig(endpoint string, accessKeyId string, secretAccessKey string, useSSl bool, bucketName string,
	location string, policy string, isHttps bool) *MinioConfig {
	return &MinioConfig{
		endpoint:        endpoint,
		accessKeyId:     accessKeyId,
		secretAccessKey: secretAccessKey,
		useSSl:          useSSl,
		bucketName:      bucketName,
		location:        location,
		policy:          policy,
		isHttps:         isHttps,
	}
}
