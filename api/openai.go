package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func IsValidModel(modelName string) bool {
	// 在这里添加您支持的有效模型列表
	validModels := []string{"gpt-3.5-turbo", "gpt-4.0", "gpt-3.5-turbo-1106"}

	// 检查所选择的模型是否在有效模型列表中
	for _, validModel := range validModels {
		if modelName == validModel {
			return true
		}
	}

	return false
}

func SetModelHandler(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	// 获取模型名称
	modelName, ok := requestData["model"].(string)
	if !ok || modelName == "" {
		http.Error(w, "Invalid or missing model name", http.StatusBadRequest)
		return
	}

	// 验证模型的有效性
	if !IsValidModel(modelName) {
		http.Error(w, "Invalid model", http.StatusBadRequest)
		return
	}

	// 处理模型切换逻辑，可以根据需要更新相关状态
	fmt.Printf("Selected model: %s\n", modelName)

	// 返回响应给前端，表示成功
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Model switched successfully"})
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	// 从前端接收 JSON 请求
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	// 从 URL 参数中获取模型名称，默认为 "gpt-3.5-turbo"
	modelName := r.URL.Query().Get("model")
	if modelName == "" {
		modelName = "gpt-3.5-turbo"
	}

	// 根据模型名称构建请求数据
	data := map[string]interface{}{
		"model":    modelName,
		"messages": requestData["messages"],
	}

	// 将请求数据转换为 JSON
	payload, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}

	// 使用函数参数或在函数内部声明这些变量
	apiURL := "https://api.nextapi.fun"
	apiKey := "ak-yPIAfSUvEUdhjD8UbAtKqlTRdDf3L7dPxT5ZqfOpgO7XCAnm"

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 POST 请求到 ChatGPT API
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/chat/completions", apiURL), bytes.NewBuffer(payload))
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
}
