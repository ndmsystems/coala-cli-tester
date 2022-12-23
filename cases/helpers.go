package cases

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
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

func getServiceInfoViaProxy(s SessionData) (info string, err error) {
	msg := coalago.NewCoAPMessage(coalago.CON, coalago.GET)
	msg.SetURIPath("/info")
	msg.SetSchemeCOAPS()

	resp, err := sendViaProxy(msg, s.Address, s.Proxy)
	if err != nil {
		return "", fmt.Errorf("get request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return "", fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	return string(resp.Body), nil
}

func sendViaProxy(message *coalago.CoAPMessage, addr, proxy string) (resp *coalago.Response, err error) {
	escapeQueryMessage(message)
	message.SetProxy(message.GetSchemeString(), addr)
	return coalago.NewClient().Send(message, proxy)
}

func escapeQueryMessage(message *coalago.CoAPMessage) {
	queries := message.GetURIQueryArray()
	message.RemoveOptions(coalago.OptionURIQuery)
	for _, v := range queries {
		kv := strings.SplitN(v, "=", 2)
		message.SetURIQuery(kv[0], kv[1])
	}
}

func downloadData(addr string, size int) (duration int64, err error) {
	now := time.Now()
	resp, err := coalago.NewClient().GET(
		fmt.Sprintf("coaps://%s/tests/large?size=%d", addr, size),
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

func downloadDataViaProxy(s SessionData, size int) (duration int64, err error) {
	msg := coalago.NewCoAPMessage(coalago.CON, coalago.GET)
	msg.SetURIPath("/tests/large")
	msg.SetURIQuery("size", fmt.Sprint(size))
	msg.SetSchemeCOAPS()

	now := time.Now()
	resp, err := sendViaProxy(msg, s.Address, s.Proxy)
	if err != nil {
		return 0, fmt.Errorf("get request by coala: %s", err)
	}
	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}
	duration = time.Since(now).Milliseconds()

	return duration, nil
}

func sendData(addr string, size int) (duration int64, err error) {
	data := make([]byte, size)
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

func sendDataViaProxy(s SessionData, size int) (duration int64, err error) {
	data := make([]byte, size)
	_, err = rand.Read(data)
	if err != nil {
		return 0, fmt.Errorf("generate payload: %s", err)
	}

	msg := coalago.NewCoAPMessage(coalago.CON, coalago.POST)
	msg.SetURIPath("/tests/large")
	msg.Payload = coalago.NewBytesPayload(data)
	msg.SetSchemeCOAPS()

	now := time.Now()
	resp, err := sendViaProxy(msg, s.Address, s.Proxy)
	if err != nil {
		return 0, fmt.Errorf("post request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	duration = time.Since(now).Milliseconds()

	return duration, nil
}

func sendMirror(addr string, size int) (duration int64, err error) {
	data := make([]byte, size)
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

func sendMirrorViaProxy(s SessionData, size int) (duration int64, err error) {
	data := make([]byte, size)
	_, err = rand.Read(data)
	if err != nil {
		return 0, fmt.Errorf("generate payload: %s", err)
	}

	msg := coalago.NewCoAPMessage(coalago.CON, coalago.POST)
	msg.SetURIPath("/tests/mirror")
	msg.Payload = coalago.NewBytesPayload(data)
	msg.SetSchemeCOAPS()

	now := time.Now()
	resp, err := sendViaProxy(msg, s.Address, s.Proxy)
	if err != nil {
		return 0, fmt.Errorf("post request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return 0, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	duration = time.Since(now).Milliseconds()

	return duration, nil
}

func gumSession(cid, peerCid, gumAddr string) (s SessionData, err error) {
	resp, err := coalago.NewClient().GET(
		fmt.Sprintf("coaps://%s/session?cid=%s&peer_cid=%s", gumAddr, cid, peerCid),
	)
	if err != nil {
		return s, fmt.Errorf("get request by coala: %s", err)
	}

	if resp.Code != coalago.CoapCodeContent {
		return s, fmt.Errorf("invalid response code %s with payload: %s", resp.Code.String(), string(resp.Body))
	}

	err = json.Unmarshal(resp.Body, &s)
	if err != nil {
		return s, fmt.Errorf("unmarshal json: %s with payload: %s", err, string(resp.Body))
	}

	return s, nil
}

func ByteCountBinary(b int64) string {
	b *= 8
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
