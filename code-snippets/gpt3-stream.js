// 导入axios库
import axios from 'axios';

// 设置OpenAI的API密钥
const API_KEY = 'your_api_key';

// 准备请求数据
const requestData = {
    prompt: 'Hello,',
    max_tokens: 5,
    temperature: 0.5
};

// 创建一个新的axios实例
const instance = axios.create({
    baseURL: 'https://api.openai.com/v1',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${API_KEY}`
    },
});

// 发送POST请求
instance.post('/engines/davinci-codex/completions', requestData)
    .then(function (response) {
        // 处理接收到的响应内容
        console.log(response.data.choices[0].text);
    })
    .catch(function (error) {
        console.error(error);
    });
