package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro-services-roadmap/oneid-core/model"
	"io"
	"net/http"
	"path"
	"strings"
)

var (
	tokenUrl    = "/auth/token"
	updateUrl   = "/access-key/%d"
	registerUrl = "/access-keys"

	client = &http.Client{}
)

func DoReq(method, url string, body []byte) (*model.Response, error) {

	if !strings.HasPrefix(url, "http") || !strings.HasPrefix(url, "https") {
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
			return nil, errors.New("请求OneId服务失败")
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
		return DoReq(http.MethodPost, path.Join(aacshost, tokenUrl), bts)
	}
}

func Register(aacshost string, body *model.AccessKeyReq) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	if bts, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		return DoReq(http.MethodPost, path.Join(aacshost, registerUrl), bts)
	}
}

func UpdateAccessKey(aacshost string, body *model.AccessKeyUpdateReq, ID int64) (*model.Response, error) {

	if len(aacshost) == 0 {
		return nil, errors.New("OneidSvc is empty")
	}

	if bts, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		return DoReq(http.MethodPost, path.Join(aacshost, fmt.Sprintf(updateUrl, ID)), bts)
	}
}
