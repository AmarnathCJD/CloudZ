package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StarReq struct {
	PaymentIndent string `json:"payment_indent"`
	Secret        string `json:"secret"`
}

var (
	httpx, _ = NewHttpSession()
)

const (
	PK_LIVE = "pk_live_FllziDXKdkK8AW2o2YbWUTWZ00Pn3cJglW"
)

func (r *StarReq) StepOne() error {
	headers := map[string]string{
		"authority":  "starregistration.net",
		"accept":     "application/json, text/javascript, */*; q=0.01",
		"cookie":     "frontend=st5dg364cpde0aomkdlfl87kg1; frontend_cid=WVcXnT2Abf9DfnL3; _ga=8aa9093a-ee31-4875-92bd-e9bf7076ba5b; _fbp=fb.1.1678192270012.766096246; external_no_cache=1; _gcl_au=1.1.980549279.1678192271; _gcl_aw=GCL.1678192274.CjwKCAiA3pugBhAwEiwAWFzwdSTuiYY-esEwkQ7RSAXFm77SXqtkVd1NwFNDbhba0TnF2uyMjPLoWBoCuhgQAvD_BwE; uenc=aHR0cHM6Ly9zdGFycmVnaXN0cmF0aW9uLm5ldC9leHRlbmRzL2Jsb2NrL2dldC8%2C; __zlcmid=1ElleYCVaE1x7Nw; msg_dlv=1; currency=USD; __kla_id=eyIkcmVmZXJyZXIiOnsidHMiOjE2NzgxOTIyNzEsInZhbHVlIjoiaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8iLCJmaXJzdF9wYWdlIjoiaHR0cHM6Ly9zdGFycmVnaXN0cmF0aW9uLm5ldC8/Z2NsaWQ9Q2p3S0NBaUEzcHVnQmhBd0Vpd0FXRnp3ZFNUdWlZWS1lc0V3a1E3UlNBWEZtNzdTWHF0a1ZkMU53Rk5EYmhiYTBUbkYydXlNalBMb1dCb0N1aGdRQXZEX0J3RSJ9LCIkbGFzdF9yZWZlcnJlciI6eyJ0cyI6MTY3ODE5MjM3MiwidmFsdWUiOiJodHRwczovL3d3dy5nb29nbGUuY29tLyIsImZpcnN0X3BhZ2UiOiJodHRwczovL3N0YXJyZWdpc3RyYXRpb24ubmV0Lz9nY2xpZD1DandLQ0FpQTNwdWdCaEF3RWl3QVdGendkU1R1aVlZLWVzRXdrUTdSU0FYRm03N1NYcXRrVmQxTndGTkRiaGJhMFRuRjJ1eU1qUExvV0JvQ3VoZ1FBdkRfQndFIn19",
		"origin":     "https://starregistration.net",
		"referer":    "https://starregistration.net/standard-star-registration.html",
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	}
	var h http.Header = make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	resp, err := httpx.Get("https://starregistration.net/amstripefront/payment/update/", h, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r.PaymentIndent = "pi" + strings.Split(string(body), `_secret`)[0][len(`///`):]
	r.Secret = strings.ReplaceAll(string(body), `"`, ``)
	return nil
}

func (r StarReq) StepTwo() error {
	headers := map[string]string{
		"authority":    "starregistration.net",
		"accept":       "application/json, text/javascript, */*; q=0.01",
		"origin":       "https://starregistration.net",
		"referer":      "https://starregistration.net/checkout/cart/",
		"content-type": "application/x-www-form-urlencoded; charset=UTF-8",
		"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	}
	var h http.Header = make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	if _, err := httpx.Post("https://starregistration.net/reclaim/index/saveEmail/", h, []byte(`email=roseloverx%40proton.me`), false); err != nil {
		return err
	}
	resp, err := httpx.Post("https://starregistration.net/checkout/onepage/saveBilling/", h, []byte(`billing%5Baddress_id%5D=10454134&billing%5Bemail%5D=roseloverx%40proton.me&billing%5Btelephone%5D=&billing%5Bfirstname%5D=Jenna&billing%5Blastname%5D=Ortega&billing%5Bstreet%5D%5B%5D=7775%20N%20Palm%20Ave&billing%5Bcountry_id%5D=US&billing%5Bregion_id%5D=12&billing%5Bregion%5D=&billing%5Bcity%5D=Fresno&billing%5Bpostcode%5D=93711&billing%5Baadhar%5D=&billing%5Bsave_in_address_book%5D=1&billing%5Buse_for_shipping%5D=1`), false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return nil
}

type stripeResp struct {
	Error struct {
		Charge      string `json:"charge"`
		Code        string `json:"code"`
		DeclineCode string `json:"decline_code"`
		DocURL      string `json:"doc_url"`
		Message     string `json:"message"`
	} `json:"error"`
}

func (r StarReq) StepThree(card, expMo, expYr, cvv string) (string, string, string, error) {
	headers := map[string]string{
		"authority":    "api.stripe.com",
		"accept":       "application/json",
		"origin":       "https://js.stripe.com",
		"referer":      "https://js.stripe.com",
		"content-type": "application/x-www-form-urlencoded",
		"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	}
	var h http.Header = make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	payload := url.Values{
		"payment_method_data[type]":                                  {"card"},
		"payment_method_data[billing_details][address][city]":        {"Fresno"},
		"payment_method_data[billing_details][address][country]":     {"US"},
		"payment_method_data[billing_details][address][line1]":       {"7775 N Palm Ave"},
		"payment_method_data[billing_details][address][postal_code]": {"93711"},
		"payment_method_data[billing_details][address][state]":       {"CA"},
		"payment_method_data[billing_details][name]":                 {"Jenna Ortega"},
		"payment_method_data[card][number]":                          {card},
		"payment_method_data[card][exp_month]":                       {expMo},
		"payment_method_data[card][exp_year]":                        {expYr},
		"payment_method_data[card][cvc]":                             {cvv},
		"payment_method_data[payment_user_agent]":                    {"stripe.js/b998f5daf; stripe-js-v3/b998f5daf"},
		"expected_payment_method_type":                               {"card"},
		"use_stripe_sdk":                                             {"true"},
		"key":                                                        {PK_LIVE},
		"client_secret":                                              {r.Secret},
	}
	resp, err := httpx.Post("https://api.stripe.com/v1/payment_intents/"+r.PaymentIndent+"/confirm", h, []byte(payload.Encode()), false)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}
	var a stripeResp
	if err := json.Unmarshal(body, &a); err != nil {
		return "", "", "", err
	}
	if a.Error == (stripeResp{}.Error) {
		return "", "", "", fmt.Errorf("no error")
	}
	return a.Error.Code, a.Error.DeclineCode, a.Error.Message, nil
}

func writeJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
