# hw-mock-tinder-api

## API 開發

In this service, it will provide APIs for the client to do the following things:

* Get Recommendation Users (用戶配對列表 API)
  * GET /api/v1/users/recommendations/
  * GET /api/v1/users/preferences/recommendations/ (使用自己偏好設定)

* Login with facebook (用戶認證 API)
  * POST /api/v1/login/facebook [](./internal/application/service/authn/authn_service.go#L8)

* Login with Sms (用戶認證 API)
  * POST /api/v1/login/sms [](./internal/application/service/authn/authn_service.go#L168)

* Chat API:
  * GET /api/v1/users/chats
  * POST /api/v1/users/message

* Get Current User (查看用戶資料 API)
  * GET /api/v1/users/

* Update Current User (更新用戶資料 API)
  * PUT /api/v1/users/

## 系統水平擴展設計

## 架構概述

本系統使用微服務架構並通過 Kubernetes 進行容器化部署，支持根據負載自動調整副本數量。系統的每個服務都可以單獨擴展，確保在高併發情況下能夠高效運行。

## 高併發優化

* 使用緩存
  * 方案：引入分布式緩存系統（如 Redis、Memcached）以減少數據庫壓力。
  * 應用場景：頻繁訪問的靜態數據（如JWT Token、熱門查詢結果）。
  * 策略：設置合理的 TTL（Time-to-Live）避免緩存雪崩。

* 請求排隊與限流
  * 方案：引入限流中間件（如令牌桶算法）避免過載。
  * 應用場景：API 請求高峰期間。
  * 工具：Gin 限流中間件或 Nginx 配置。

## 水平擴展

* 確保 API 無狀態, 使用 JWT TOKEN 管理。
* 使用 Traefik, Nginx, Kubernetes 的負載均衡來分發流量，避免單點瓶頸。
* 服務副本數量可通過 Kubernetes 進行自動擴展（水平擴展）。
* 根據業務需求進行資料庫的分片和連接池的設置，保證資料庫的高效處理。
* 引入消息隊列, 解耦系統模塊, 處理異步事件。

+-----------------------+         +---------------------+
|   Load Balancer       |         | Distributed Cache  |
|  (Nginx/K8s Ingress)  |         |   (Redis)          |
+-----------------------+         +---------------------+
             |                              |
+-----------------------+         +---------------------+
|   Mocker Tinder API   |  <----> | Message Queue       |
|                       |         |   (Kafka/NATS)      |
+-----------------------+         +---------------------+
             |
+-----------------------+
|   Database Layer      |
| (Master-Slave DB)     |
+-----------------------+

## 實施步驟

1. 配置並部署 Kubernetes 集群。
2. 設置服務的擴展策略。
3. 配置 Redis 和資料庫連接池。
4. 配置流量限流和熔斷機制，確保高可用性。

## 安全性考量

1. 基於 JWT 的身份驗證，用戶登錄後通過 JWT 標頭來驗證請求。

## 即時消息處理

  1. 使用 WebSocket 或長輪詢技術實現即時聊天功能。
  2. 支持消息的已讀未讀狀態更新。

## 部署與運維

1. 提供基於容器化技術（如 Docker）的一鍵部署方案。

    * 創建一個 Dockerfile 來構建應用程序的容器映像, 根目錄 Dockerfile

2. 使用 CI/CD 工具（如 GitHub Actions、Jenkins）實現自動化測試與部署流程。

    * 創建 .github/workflows/deploy.yml 文件，定義自動化流程。
    * 當代碼推送到 main 分支時，工作流會觸發。
    * 步驟包括檢出代碼、安裝 Go 環境、運行測試、構建 Docker 映像、推送 Docker 映像到 Docker Hub。
