package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	baseURL      = "http://127.0.0.1:8888/usercenter/v1/user"
	registerPath = "/register"
	loginPath    = "/login"
	csvFilePath  = "tokens.csv"
)

type RegisterRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken string `json:"accessToken"`
	} `json:"data"`
}

func registerUser(mobile, password string) error {
	reqBody := RegisterRequest{
		Mobile:   mobile,
		Password: password,
	}

	reqBytes, _ := json.Marshal(reqBody)
	resp, err := http.Post(baseURL+registerPath, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register failed for mobile %s, status code: %d", mobile, resp.StatusCode)
	}
	return nil
}

func loginUser(mobile, password string) (string, error) {
	reqBody := LoginRequest{
		Mobile:   mobile,
		Password: password,
	}

	reqBytes, _ := json.Marshal(reqBody)
	resp, err := http.Post(baseURL+loginPath, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return "", err
	}

	if loginResp.Code != 200 {
		return "", fmt.Errorf("login failed for mobile %s, message: %s", mobile, loginResp.Msg)
	}

	return loginResp.Data.AccessToken, nil
}

func writeTokensToCSV(tokens []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"token"})
	for _, token := range tokens {
		writer.Write([]string{token})
	}

	return nil
}

func main() {
	mobilePrefix := "17142560"
	password := "zZizwGGf8H3nF47"
	tokens := make([]string, 0, 1000)

	for i := 0; i < 1000; i++ {
		mobile := fmt.Sprintf("%s%03d", mobilePrefix, i)

		// 注册用户
		err := registerUser(mobile, password)
		if err != nil {
			fmt.Printf("Error registering user %s: %v\n", mobile, err)
			continue
		}

		// 登录用户
		token, err := loginUser(mobile, password)
		if err != nil {
			fmt.Printf("Error logging in user %s: %v\n", mobile, err)
			continue
		}

		fmt.Printf("User %s registered and logged in successfully. Token: %s\n", mobile, token)
		tokens = append(tokens, token)
	}

	// 写入 CSV 文件
	err := writeTokensToCSV(tokens, csvFilePath)
	if err != nil {
		fmt.Printf("Error writing tokens to CSV: %v\n", err)
		return
	}

	fmt.Println("Tokens successfully written to", csvFilePath)
}
