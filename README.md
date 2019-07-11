# P2P使用websocket實作範例

### 簡要：

以Websocket實作p2p（peer to peer）結構，每個節點皆建立http server，並以客戶端向指定ip發出websocket連線請求，指定節點收到新連線時會透過websocket發出廣播，通知其他節點向新節點建立連線。

### 執行：

1. `go run main.go -p {本地port} -seed {種子ip}`

**參數說明：**

* `-p` 本機建立http服務監聽埠號（預設8080)
* `-seed` 連向已建立節點的ip   (預設127.0.0.1:8080)

**注意：** 第一次建立請將`-seed`設為127.0.0.1，並與`-p`設定相同的port。

2. 節點建立後可透過 `http://$URI/peers` 可查看已建立雙向連線的節點位置。

### 檔案結構

```
.
├── README.md
├── common
│   └── const.go    // route name 設定
├── handler
│   └── handler.go  // http請求處理
├── main.go         // 入口
├── p2p
│   ├── p2p.go      // 節點控制器
│   └── peer.go     // 節點實作
└── router
    └── route.go    // http配置

```