package cases

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/coalalib/coalago"
)

func getService(cid, gumAddr, service string) (srvAddr string, err error) {
	resp, err := coalago.NewClient().GET(
		fmt.Sprintf("coaps://%s/get?peer_cid=%s&cid=%s", gumAddr, service, cid),
	)
	if err != nil {
		return "", fmt.Errorf("get request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return "", fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	return string(resp.Body), nil
}

func getServiceInfo(serviceAddr string) (info string, err error) {
	resp, err := coalago.NewClient().GET("coaps://" + serviceAddr + "/info")
	if err != nil {
		return "", fmt.Errorf("get request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return "", fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	return string(resp.Body), nil
}

func downloadData(addr string) (duration int64, err error) {
	now := time.Now()
	resp, err := coalago.NewClient().GET(
		fmt.Sprintf("coaps://%s/tests/large?size=%d", addr, 512*1024),
	)
	if err != nil {
		return 0, fmt.Errorf("get request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	duration = time.Since(now).Milliseconds()

	return duration, nil
}

func sendData(addr string) (duration int64, err error) {
	data := make([]byte, 512*1024)
	_, err = rand.Read(data)
	if err != nil {
		return 0, fmt.Errorf("generate payload: %s", err)
	}

	now := time.Now()
	resp, err := coalago.NewClient().POST(data,
		fmt.Sprintf("coaps://%s/tests/large", addr),
	)
	if err != nil {
		return 0, fmt.Errorf("post request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	duration = time.Since(now).Milliseconds()

	return duration, nil
}

func sendMirror(addr string) (duration int64, err error) {
	data := make([]byte, 512*1024)
	_, err = rand.Read(data)
	if err != nil {
		return 0, fmt.Errorf("generate payload: %s", err)
	}

	now := time.Now()
	resp, err := coalago.NewClient().POST(data,
		fmt.Sprintf("coaps://%s/tests/mirror", addr),
	)
	if err != nil {
		return 0, fmt.Errorf("post request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	duration = time.Since(now).Milliseconds()

	return duration, nil
}
