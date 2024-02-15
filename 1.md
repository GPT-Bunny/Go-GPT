package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CORS 中间件函数
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置 CORS 头部
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// 设置中转接口地址和 API Key
	apiURL := "https://api.nextapi.fun"
	apiKey := "ak-yPIAfSUvEUdhjD8UbAtKqlTRdDf3L7dPxT5ZqfOpgO7XCAnm"

	// 创建路由处理器
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		// 从前端接收 JSON 请求
		var requestData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
			return
		}

		// 构建请求数据
		data := map[string]interface{}{
			"model":    "gpt-3.5-turbo",
			"messages": requestData["messages"],
		}

		// 将请求数据转换为 JSON
		payload, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "JSON marshal error", http.StatusInternalServerError)
			return
		}

		// 创建 HTTP 客户端
		client := &http.Client{}

		// 创建 POST 请求到 ChatGPT API
		req, err := http.NewRequest("POST", apiURL+"/v1/chat/completions", bytes.NewBuffer(payload))
		if err != nil {
			http.Error(w, "HTTP request error", http.StatusInternalServerError)
			return
		}

		// 添加 ChatGPT API Key 到请求头
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "HTTP request error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 解析响应数据
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			http.Error(w, "JSON unmarshal error", http.StatusInternalServerError)
			return
		}

		// 提取生成的回复
		choices, ok := response["choices"].([]interface{})
		if !ok || len(choices) == 0 {
			http.Error(w, "Invalid response from ChatGPT API", http.StatusInternalServerError)
			return
		}

		message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
		if !ok {
			http.Error(w, "Invalid response format from ChatGPT API", http.StatusInternalServerError)
			return
		}

		reply, ok := message["content"].(string)
		if !ok {
			http.Error(w, "Invalid reply format from ChatGPT API", http.StatusInternalServerError)
			return
		}

		// 返回回复给前端
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"reply": reply})
	})

	// 使用中间件处理 CORS 头部
	handler := corsMiddleware(http.DefaultServeMux)

	// 启动服务器
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
