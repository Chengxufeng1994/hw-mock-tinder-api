# hw-mock-tinder-api

## API 開發

In this service, it will provide APIs for the client to do the following things:

* Get Recommendation Users (用戶配對列表 API)
  * GET /api/v1/users/recommendations/

* Login with facebook (用戶認證 API)
  * POST /api/v1/login/facebook

* Login with Sms (用戶認證 API)
  * POST /api/v1/login/sms

* Get Current User (用戶資料 API)
  * GET /api/v1/users/

* Update Current User (用戶資料 API)
  * PUT /api/v1/users/

## 系統水平擴展設計

## 架構概述

本系統使用微服務架構並通過 Kubernetes 進行容器化部署，支持根據負載自動調整副本數量。系統的每個服務都可以單獨擴展，確保在高併發情況下能夠高效運行。

## 高併發優化

* 使用 Redis 進行熱點數據的緩存，減少對資料庫的訪問。
* 實現批量處理機制，減少 API 請求的數量。
* 將長時間操作放入消息隊列中，使用背景處理方式非同步執行。

## 水平擴展

* 服務副本數量可通過 Kubernetes 進行自動擴展（水平擴展）。
* 使用 Kubernetes 的負載均衡來分發流量，避免單點瓶頸。
* 根據業務需求進行資料庫的分片和連接池的設置，保證資料庫的高效處理。

## 實施步驟

1. 配置並部署 Kubernetes 集群。
2. 設置服務的擴展策略。
3. 配置 Redis 和資料庫連接池。
4. 配置流量限流和熔斷機制，確保高可用性。

## 安全性考量

1. 基於 JWT 的身份驗證，用戶登錄後通過 JWT 標頭來驗證請求。

## 部署與運維

1. 提供基於容器化技術（如 Docker）的一鍵部署方案。

    創建一個 Dockerfile 來構建應用程序的容器映像, 根目錄 Dockerfile

2. 使用 CI/CD 工具（如 GitHub Actions、Jenkins）實現自動化測試與部署流程。

    * 創建 .github/workflows/deploy.yml 文件，定義自動化流程。
    * 當代碼推送到 main 分支時，工作流會觸發。
    * 步驟包括檢出代碼、安裝 Go 環境、運行測試、構建 Docker 映像、推送 Docker 映像到 Docker Hub。
