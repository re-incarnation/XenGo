package xfrequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/re-incarnation/XenGo/xftypes"
)

func Auth(conf xftypes.Config) {
	method := "POST"
	AuthUrl := conf.URL + "auth"
	client := &http.Client{}
	req, err := http.NewRequest(method, AuthUrl, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("XF-Api-Key", conf.Key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}

func UserCreate(user xftypes.UserUpdateRequest, config xftypes.Config) (string, xftypes.User) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	body := xftypes.UserUpdateRequest{
		Option:  user.Option,
		Profile: user.Profile,
		Privacy: user.Privacy,
		Dob:     user.Dob,
		User:    user.User,
	}

	jsonData, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", config.URL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("XF-Api-Key", config.Key)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("Status" + resp.Status)

	return resp.Status, body.User
}
