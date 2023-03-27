# 智齡科技小型專題

請做一個簡單的 List 呈現 Patients，並於點擊後跳出 Dialog 呈現該 Patient 的 Order(醫囑)，
於 Dialog 右上增加可新增 Order 按鈕，並提供編輯回存功能。

<https://docs.google.com/document/d/1E5frZ_bw80f163XclcMHwkyRZAzMs3ubEstNhlqh3kI/edit>

## service start and stop

```bash
# start service in container
bash project.sh -c

# stop service
bash project.sh -s
```

## backend api

```bash
curl -X GET --location "http://localhost:8888/v1/api/patients" \
-H "Accept: application/json"

curl -X GET --location "http://localhost:8888/v1/api/orders" \
-H "Accept: application/json"

curl -X GET --location "http://localhost:8888/v1/api/orders?patient_id=01GWFC9277K77PP1XX57BNHBFS" \
-H "Accept: application/json"

curl -X POST --location "http://localhost:8888/v1/api/orders" \
-H "Accept: application/json" \
-d '{ "message": "病人精神狀態不佳, 要注意睡眠", "patient_id": "01GWFC9277K77PP1XX57BNHBFS" }'

curl -X GET --location "http://localhost:8888/v1/api/orders/01GWFCJQAY4QCSXJ1W1SF3ACJG" \
-H "Accept: application/json"

curl -X PATCH --location "http://localhost:8888/v1/api/orders/01GWFCJQAY4QCSXJ1W1SF3ACJG" \
-H "Accept: application/json" \
-d '{ "message": "血糖超過120，建議施打 10 單位胰島素" }'

```

## 使用技術

- backend:  
	golang gin gorm  
	wire testtify go-playground  
	ulid PostgreSQL  

- frontend:  
	react Next.js axios  
	MUI(material-ui)  

# Dockerfile

```bash
# backend
docker build -f Dockerfile-backend -t x246libra/jubo-homework:v1.0 . && \
    docker rmi `docker images --filter label=stage=builder -q`
```

