package lazadago

import (
	// "encoding/json"
	"encoding/json"
	"log"
)

const (
	// please to go https://open.lazada.com/ to register one
	apiKey      = "app-id"
	apiSecret   = "app-secret"
	accessToken = "access-token"
)
func main() {

	//callBackUrl := "https://htbk.byteco.cn/lazada/authCb"
	appKey := "125991"
	appSecret := "TrBcWRniy1X6HPMIUBzxTlyTqUPi9XvK"
	// 授权
	// Using
	var clientOptions = ClientOptions{
		APIKey:    appKey,
		APISecret: appSecret,
		Region: "TH",
	}
	client := NewClient(&clientOptions).
		AddAPIParam("CreatedAfter",  "2020-08-01T00:00:00+08:00").
		AddAPIParam("CreatedBefore",  "2020-08-02T00:00:00+08:00").
		AddAPIParam("Status",  "delivered")

	client.SetAccessToken("50000701d27uO6YccvgKuVifIxG3gqVFh0HQRBv1fdeca53r0qsRv47tBT88Mvm")
	orderList,_ := client.GetOrders()
	orderListJson, _ := json.Marshal(orderList)
	log.Println(string(orderListJson))
}
//func TestPhotoUpload(t *testing.T) {
	// filename := "/Users/hsinhoyeh/Desktop/1.png"
	// blob, _ := ioutil.ReadFile(filename)
	// r := New(apiKey, apiSecret, ApiGatewayID).
	// 	Method(http.MethodPost).
	// 	ApiPath("/image/upload").
	// 	AddFileParam("image", blob).
	// 	AccessToken(accessToken)
	// resp, err := r.Do()
	// fmt.Printf("resp:%v, err:%v\n", string(resp.Data), err)
//}

//func TestGetProduct(t *testing.T) {
	// r := NewRequest(apiKey, apiSecret, ApiGatewayID).
	// 	Method(http.MethodGet).
	// 	ApiPath("/products/get").
	// 	AddApiParam("filter", "live").
	// 	AddApiParam("limit", "10").
	// 	AddApiParam("offset", "0").
	// 	AccessToken(accessToken)
	// resp, err := r.Do()
	// fmt.Printf("resp:%v, err:%v\n", string(resp.Data), err)
/*
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": {"name": "Maka", "short_code": "developer@jhonmike.com.br"}}`))
	}))
	defer ts.Close()*/

	/*var clientOptions = ClientOptions{
		APIKey:    apiKey,
		APISecret: apiSecret,
		ServerURL: ts.URL,
	}
	lc := NewClient(&clientOptions)
	resp, _ := lc.GetSeller()

	fmt.Println(resp.Name)*/

//}
