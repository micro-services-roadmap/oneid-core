package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/micro-services-roadmap/oneid-core/model"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
)

const (
	OneidCaptchaEndpoint = "/oneid/captcha"
	OneidTokenEndpoint   = "/oneid/token"
)

var (
	TokenUrl    = "/base/token"
	CaptchaUrl  = "/base/captcha"
	UpdateUrl   = "/access-key/:id"
	RegisterUrl = "/access-keys"

	client = &http.Client{}
)

func DoReq(method, url string, body []byte) (*model.Response, error) {

	if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
		url = "http://" + url
	}

	if request, err := http.NewRequest(method, url, bytes.NewBuffer(body)); err != nil {
		return nil, err
	} else {
		request.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("请求OneId服务失败(" + resp.Status + ")")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		result := &model.Response{}
		if err = json.Unmarshal(body, result); err != nil {
			return nil, err
		}

		return result, nil
	}
}

func GenerateJwt(aacshost string, body *model.JwtReq) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	if bts, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		return DoReq(http.MethodPost, path.Join(aacshost, TokenUrl), bts)
	}
}

func Register(aacshost string, body *model.AccessKeyReq) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	if bts, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		return DoReq(http.MethodPost, path.Join(aacshost, RegisterUrl), bts)
	}
}

func UpdateAccessKey(aacshost string, body *model.AccessKeyUpdateReq, ID int64) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	UpdateUrl = strings.Replace(UpdateUrl, ":id", strconv.FormatInt(ID, 10), 1)
	if bts, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		return DoReq(http.MethodPut, path.Join(aacshost, UpdateUrl), bts)
	}
}

func GetCaptcha(aacshost string) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	return DoReq(http.MethodGet, path.Join(aacshost, CaptchaUrl), nil)
}
