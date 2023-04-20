package service

import (
	"encoding/base64"
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strings"
)

var Images = &imagesSrv{}

type imagesSrv struct{}

func (s *imagesSrv) SaveImageWithUrl(base, url string) (err *exception.Exception) {
	pic := base
	idx := strings.LastIndex(url, "/")

	if idx < 0 {
		pic += "/" + url
	} else {
		pic += "/" + url[idx+1:]
	}

	v, e := http.Get(url)
	if e != nil {
		common.Logger.Warningf("imagesSrv SaveImageWithUrl get url fail. | base: %s | url: %s | err: %s", base, url, e)
		err = exception.New(exception.CodeCallHttpFailed)
		return err
	}
	defer v.Body.Close()

	content, e := io.ReadAll(v.Body)
	if e != nil {
		common.Logger.Warningf("imagesSrv SaveImageWithUrl ReadAll fail. | base: %s | url: %s | err: %s", base, url, e)
		err = exception.New(exception.CodeInternalError)
		return err
	}

	pic += "." + s.getImageSuffix(content)

	e = os.WriteFile(pic, content, 0666)
	if e != nil {
		common.Logger.Warningf("imagesSrv SaveImageWithUrl WriteFile fail. | base: %s | url: %s | err: %s", base, url, e)
		err = exception.New(exception.CodeInternalError)
		return err
	}
	return nil
}

func (s *imagesSrv) SaveImageWithBase64Json(base, imageJson string) (err *exception.Exception) {
	bytes, e := base64.StdEncoding.DecodeString(imageJson)

	if e != nil {
		common.Logger.Warningf("imagesSrv SaveImageWithBase64Json DecodeString fail. | imageJson: %s | err: %s", imageJson, e)
		err = exception.New(exception.CodeInternalError)
		return err
	}

	fileName := base + "/" + uuid.New().String() + "." + s.getImageSuffix(bytes)

	e = os.WriteFile(fileName, bytes, 0666)
	if e != nil {
		common.Logger.Warningf("imagesSrv SaveImageWithUrl WriteFile fail. | base: %s | err: %s", base, e)
		err = exception.New(exception.CodeInternalError)
		return err
	}

	return
}

func (s *imagesSrv) getImageSuffix(bytes []byte) (suffix string) {
	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		suffix = "jpg"
	case "image/png":
		suffix = "png"
	case "image/gif":
		suffix = "gif"
	}

	return
}
