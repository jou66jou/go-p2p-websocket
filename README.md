### P2P使用websocket實作

#### 簡要：
以Websocket實作p2p（peer to peer）結構，每個節點皆建立http server，並以客戶端向指定ip發出websocket連線請求，節點收到新連線時會透過websocket發出廣播，通知其他節點向新節點建立連線。
