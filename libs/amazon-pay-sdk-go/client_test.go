package amazonpay

import (
	"fmt"
	"testing"
)

var Client = New(&Config{
	MerchantID: "test",
	AccessKey:  "test",
	SecretKey:  "test",
	Sandbox:    true,
	Region:     "jp",
})

func TestSetOrderReferenceDetails(t *testing.T) {
	data, err := Client.SetOrderReferenceDetails("orderReferenceID", OrderReferenceAttributes{
		OrderTotal: OrderTotal{
			CurrencyCode: "usd",
			Amount:       "100",
		},
		PlatformID: "test",
		SellerNote: "test",
		SellerOrderAttributes: SellerOrderAttributes{
			SellerOrderID:     "orderid",
			StoreName:         "test",
			CustomInformation: "test",
		},
	})

	fmt.Printf("%#v\n", data)
	fmt.Println(err)
}

func TestGetProfile(t *testing.T) {
	t.Skip()

	// The token is got from https://demo.getqor.com/cart
	profile, err := Client.GetProfile("Atza|IwEBIB7_PgcZFL8P1b4Pz1NJ6gfzSSGleqKE59p-56jqFjY7dMjbN5OnBtdi4YQQJnG-PwZpnCjlHGlgLgk98UqOc6zFvoCOU7sNoD_5HZ2KGfwnqrjrO-lg6oAyR3iIAIM_m1gYikVXEa0M65al4tbr3Vauqbhtg3B4D2k72mxwByiyjhUJcO4hau4FCGYShXulWkHO5NwOV3NsYk8LEKqS9pmnljL-l9I1BAAIBdNn2Zy5az4bF_z74KSBkqL2lEd44RaIaVuN0bY_J6E4pd_ujo3T2AfVvKuHg-k3heW0DkS8ZPdYok3fmlwtlBRfIJOB2OADUzRFXN7JzQI2reNWngkZ4PRhmGNh6YF7xezQUfRUVTvaoE9vNeobQF-AKHIaJ9LpYoWg5eb67m823JA4OMLRBiJgNHPEPmGykZHEa3VxBGrf2lprKYgsq-UmJJjV_PZXITkjnOlO1gYfGbLZFDplRhV15Ua2gfHWtG-TWqvXtTieEK2_dBXJmNVd3IoXwHRLHtgbwKiJCTLQOGn0zHleN98r3W3HNwIpXFcGuMMnXpU17MkHiP_n4cPDoDmPuqPaGo1xgA1y2XUpipLf1FkSeIO61uPkfMEyqqE_TuTcJA")
	if err != nil {
		t.Errorf("No error should happen when get profile, but got %v", err)
	}

	if profile.Name == "" {
		t.Errorf("Profile should has value")
	}
}
