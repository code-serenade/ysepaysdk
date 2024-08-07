package api

// func TestNotifyPayload_VerifySign(t *testing.T) {
// 	payload := &NotifyPayload{
// 		TimeStamp:  "2023-01-01T00:00:00Z",
// 		Src:        "test_src",
// 		ReqID:      "test_req_id",
// 		Charset:    "UTF-8",
// 		Sign:       "test_sign",
// 		BizContent: "{\"orderID\": \"123456\"}",
// 	}

// 	key := []byte("example_public_key")

// 	err := payload.VerifySign(key)
// 	assert.NotNil(t, err, "expected error due to incorrect key/sign")
// }

// func TestReceiveAndVerifyPayload(t *testing.T) {
// 	jsonPayload := `{
// 		"timeStamp": "2023-01-01T00:00:00Z",
// 		"src": "test_src",
// 		"reqId": "test_req_id",
// 		"charset": "UTF-8",
// 		"sign": "test_sign",
// 		"bizContent": "{\"orderID\": \"123456\"}"
// 	}`
// 	req := httptest.NewRequest(http.MethodPost, "/v1/ysepay/notifyCust", bytes.NewBuffer([]byte(jsonPayload)))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()

// 	session := web.NewSession(resp, req)

// 	payload, err := receiveAndVerifyPayload(session)
// 	assert.Nil(t, payload)
// 	assert.NotNil(t, err, "expected error due to incorrect key/sign")
// }

// func TestHandleNotify(t *testing.T) {
// 	jsonPayload := `{
// 		"timeStamp": "2023-01-01T00:00:00Z",
// 		"src": "test_src",
// 		"reqId": "test_req_id",
// 		"charset": "UTF-8",
// 		"sign": "test_sign",
// 		"bizContent": "{\"orderID\": \"123456\"}"
// 	}`
// 	req := httptest.NewRequest(http.MethodPost, "/v1/ysepay/notifyCust", bytes.NewBuffer([]byte(jsonPayload)))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()

// 	session := web.NewSession(resp, req)
// 	handler := handleNotify(NotifyCustDone)
// 	handler(session)

// 	result := resp.Result()
// 	defer result.Body.Close()

// 	assert.Equal(t, http.StatusOK, result.StatusCode, "expected status OK")
// }

// func TestParseBizContent(t *testing.T) {
// 	bizContent := `{"orderID": "123456"}`
// 	result, err := parseBizContent(bizContent)
// 	assert.Nil(t, err, "expected no error when parsing bizContent")
// 	assert.Equal(t, "123456", result.Str("orderID"), "expected orderID to be '123456'")
// }

// func TestHandle(t *testing.T) {
// 	mux := web.NewSessionMux("")
// 	Handle("/api", mux)

// 	// Verify routes are registered
// 	reqCust := httptest.NewRequest(http.MethodPost, "/api/v1/ysepay/notifyCust", nil)
// 	respCust := httptest.NewRecorder()
// 	mux.ServeHTTP(respCust, reqCust)
// 	assert.Equal(t, http.StatusOK, respCust.Code, "expected status OK for notifyCust route")

// 	reqPay := httptest.NewRequest(http.MethodPost, "/api/v1/ysepay/notifyPay", nil)
// 	respPay := httptest.NewRecorder()
// 	mux.ServeHTTP(respPay, reqPay)
// 	assert.Equal(t, http.StatusOK, respPay.Code, "expected status OK for notifyPay route")
// }
