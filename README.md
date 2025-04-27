# 计算器应用

一个基于 Next.js 和 Go 的全栈计算器应用，支持基本的数学运算。

## 功能特性

- 支持基本的数学运算：加、减、乘、除
- 实时计算结果显示
- 响应式设计，适配各种设备
- 优雅的用户界面
- 高性能的后端服务

## 技术栈

### 前端
- Next.js 14
- TypeScript
- Tailwind CSS
- React Testing Library
- Jest

### 后端
- Go
- gRPC
- HTTP/2
- CORS 支持

## 项目结构

```
.
├── frontend/          # 前端代码
│   ├── src/          # 源代码
│   ├── public/       # 静态资源
│   └── test/         # 前端测试
├── backend/          # 后端代码
│   ├── api/          # API 定义
│   ├── internal/     # 内部实现
│   ├── config/       # 配置
│   └── test/         # 后端测试
└── README.md         # 项目文档
```

## 开发环境设置

### 前端开发

1. 进入前端目录：
```bash
cd frontend
```

2. 安装依赖：
```bash
npm install
```

3. 启动开发服务器：
```bash
npm run dev
```

### 后端开发

1. 进入后端目录：
```bash
cd backend
```

2. 安装依赖：
```bash
go mod download
```

3. 启动服务器：
```bash
go run main.go
```

## 测试

### 前端测试
```bash
cd frontend
npm test
```

### 后端测试
```bash
cd backend
go test ./...
```

## 构建和部署

### 前端构建
```bash
cd frontend
npm run build
```

### 后端构建
```bash
cd backend
go build -o calculator
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件 
