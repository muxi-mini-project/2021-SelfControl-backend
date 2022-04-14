package user

import (
	"SC/model"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io"
	"strconv"
	"strings"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"time"
)

var (
	accessKey, secretKey, bucketName, domainName, upToken string
)

func initOSS() {
	accessKey = viper.GetString("oss.access_key")
	secretKey = viper.GetString("oss.secret_key")
	bucketName = viper.GetString("oss.bucket_name")
	domainName = viper.GetString("oss.domain_name")
}

func getToken() {
	var maxInt uint64 = 1 << 32
	initOSS()
	putPolicy := storage.PutPolicy{
		Scope:   bucketName,
		Expires: maxInt,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
}

func getObjectName(filename string, id uint32) (string, error) {
	i := strings.LastIndex(filename, ".")
	fileType := filename[i+1:]

	timeEpochNow := time.Now().Unix()
	objectName := strconv.FormatUint(uint64(id), 10) + "-" + strconv.FormatInt(timeEpochNow, 10) + "." + fileType
	return objectName, nil
}

func uploadFile(filename string, id uint32, r io.ReaderAt, dataLen int64) (string, error) {
	if upToken == "" {
		getToken()
	}

	objectName, err := getObjectName(filename, id)
	if err != nil {
		return "", err
	}

	// 下面是七牛云的oss所需信息，objectName对应key是文件上传路径
	cfg := storage.Config{Zone: &storage.ZoneHuanan, UseHTTPS: false, UseCdnDomains: true}
	formUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{Params: map[string]string{"x:name": "self-control"}}
	err = formUploader.Put(context.Background(), &ret, upToken, objectName, r, dataLen, &putExtra)
	if err != nil {
		return "", err
	}

	url := domainName + "/" + objectName
	return url, nil
}

func UpdateAvatar(filename string, id string, r io.ReaderAt, dataLen int64) (string, error) {
	Id, err := strconv.Atoi(id)
	url, err := uploadFile(filename, uint32(Id), r, dataLen)
	if err != nil {
		return "", err
	}

	user := model.User{
		StudentID:   id,
		UserPicture: url,
	}

	err = model.UpdateUserInfo(user)
	return url, err
}
