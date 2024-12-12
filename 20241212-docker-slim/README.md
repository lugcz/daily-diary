# SlimToolkit 使用文档

SlimToolkit 是一个用于优化和缩小 Docker 镜像的工具。它通过分析和删除不必要的文件和依赖项，帮助用户创建更小、更安全的 Docker 镜像。

## 安装

你可以通过以下命令安装 SlimToolkit：

```bash
curl -sL https://raw.githubusercontent.com/slimtoolkit/slim/master/scripts/install-slim.sh | sudo -E bash -
```

## 使用方法

### 1. 构建 Docker 镜像

首先，构建你的 Docker 镜像。例如：

```bash
docker build -t my-app .
```

### 2. 使用 SlimToolkit 优化镜像

使用 `slim` 命令优化你的 Docker 镜像：

```bash
slim build --tag my-app-slim my-app
```

### 3. 运行优化后的镜像

你可以像运行普通 Docker 镜像一样运行优化后的镜像：

```bash
docker run --rm my-app-slim
```

## 常用命令

- `slim build <image>`: 优化指定的 Docker 镜像。
- `slim profile <image>`: 分析 Docker 镜像并生成优化配置文件。
- `slim lint <image>`: 检查 Docker 镜像中的潜在问题。

## 示例

以下是一个完整的示例：

```bash
# 构建原始镜像
docker build -t my-app .

# 优化镜像
slim build --tag my-app-slim my-app

# 运行优化后的镜像
docker run --rm my-app-slim
```

## 更多信息

请访问 [SlimToolkit 官方文档](https://github.com/slimtoolkit/slim) 获取更多信息和高级用法。
