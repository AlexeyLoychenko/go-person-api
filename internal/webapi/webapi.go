package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AlexeyLoychenko/person_api/internal/model"
)

type WebApi struct {
	requestTimeout time.Duration
}

func New(timeout time.Duration) *WebApi {
	return &WebApi{requestTimeout: timeout}
}

func (api *WebApi) CollectPersonData(name string) (model.GroupedApiResponse, error) {
	requests := map[string]string{
		"age":         "https://api.agify.io/?name=",
		"gender":      "https://api.genderize.io/?name=",
		"nationality": "https://api.nationalize.io/?name=",
	}
	var wg sync.WaitGroup
	res := make(chan model.GeneralApiResponse, len(requests))
	wg.Add(len(requests))
	for k, v := range requests {
		go api.doRequest(k, v+name, res, &wg)
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	var groupedResp model.GroupedApiResponse
	for resp := range res {
		if resp.Error != nil {
			return groupedResp, fmt.Errorf("webapi - CollectPersonData:%w", resp.Error)
		} else {
			switch resp.ApiKey {
			case "age":
				if v, ok := resp.Payload["age"]; ok {
					if age, ok := v.(float64); ok {
						groupedResp.Age = int(age)
					} else {
						return groupedResp, fmt.Errorf("webapi - CollectPersonData - age:%w", resp.Error)
					}
				}
			case "gender":
				gender := resp.Payload["gender"].(string)
				groupedResp.Gender = gender
			case "nationality":
				c, ok := resp.Payload["country"].([]interface{})
				if !ok {
					return groupedResp, fmt.Errorf("webapi - CollectPersonData - nationality:%w", resp.Error)
				}
				var best float64
				for i := range c {
					cc, ok := c[i].(map[string]interface{})
					if !ok {
						return groupedResp, fmt.Errorf("webapi - CollectPersonData - nationality:%w", resp.Error)
					}
					country, ok2 := cc["country_id"].(string)
					prob, ok3 := cc["probability"].(float64)
					if !ok2 || !ok3 {
						return groupedResp, fmt.Errorf("webapi - CollectPersonData - nationality:%w", resp.Error)
					}
					if best < prob {
						groupedResp.Nationality = country
						best = prob
					}
				}
			}
		}

	}
	return groupedResp, nil
}

func (api *WebApi) doRequest(key string, url string, res chan<- model.GeneralApiResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{
		Timeout: api.requestTimeout * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		res <- model.GeneralApiResponse{
			ApiKey: key,
			Error:  fmt.Errorf("client.Do() error:%w", err),
		}
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		res <- model.GeneralApiResponse{
			ApiKey: key,
			Error:  fmt.Errorf("resp.StatusCode error: api responded with status %d", resp.StatusCode),
		}
		return
	}

	var apiResp map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		res <- model.GeneralApiResponse{
			ApiKey: key,
			Error:  fmt.Errorf("json.Decode() error: %w", err),
		}
		return
	}
	res <- model.GeneralApiResponse{
		ApiKey:  key,
		Payload: apiResp,
	}
}
