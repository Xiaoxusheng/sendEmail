# 发送生日祝福邮件

## 1.配置文件在config.json

```txt
group中添加祝福信息
```

```json
{
  "username": "邮箱",
  "password": "邮箱验证码",
  "group": [
    {
      "email": "对方email",
      "title": "标题",
      "message": "生日祝福",
      "birthday": "2024-12-19",
      "nickname": "昵称"
    }
  ]
}
```
