
> https://go.dev/wiki/WebAssembly


### 使用 tinygo 编译 Go 代码为 WebAssembly 需要安装 wasm-opt

```bash
npm i wasm-opt -g
```

### 使用 wasmedge 运行 WebAssembly

#### 安装 wasmedge

+ Windows 安装命令
```bash
winget install wasmedge
```

#### 运行 WebAssembly

```bash
make run-wasi
```