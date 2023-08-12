package huddersfield_town

import (
	"context"
	"encoding/xml"
	"fmt"
	"incrowd-backend/config"
	"incrowd-backend/domain/models"
	"incrowd-backend/internal/httpclient"
	"net/http"
	"net/url"
	"strconv"
)

const (
	listNewArticlesPath = "/api/incrowd/getnewlistinformation"
	articleDetailsPath  = "/api/incrowd/getnewsarticleinformation"
)

type HuddersfieldTownProvider struct {
	config config.Provider
	client httpclient.HTTPClient
}

func NewHuddersfieldTownProvider(config config.Provider, client httpclient.HTTPClient) *HuddersfieldTownProvider {
	return &HuddersfieldTownProvider{
		config: config,
		client: client,
	}
}

// GetNewArticlesIDs retrieves new article IDs from the Huddersfield Town API
func (p *HuddersfieldTownProvider) GetNewArticlesIDs(ctx context.Context) ([]string, error) {
	baseUrl, err := url.Parse(fmt.Sprintf("%s%s", p.config.Host, listNewArticlesPath))
	if err != nil {
		return nil, fmt.Errorf("unable to parse the URL with error %s", err)
	}

	params := url.Values{}
	params.Add("count", strconv.Itoa(p.config.Count))
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create the request with error %s", err)
	}

	response, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http GET call failed to endpoint %s with error %s", baseUrl, err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var result GetNewArticlesIDsResponse
	if err = xml.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("unable to decode response body with error %s", err)
	}

	return result.mapToIDs(), nil
}

// GetArticleInformation retrieves detailed information about an article from the Huddersfield Town API
func (p *HuddersfieldTownProvider) GetArticleInformation(ctx context.Context, id string) (*models.Article, error) {
	baseUrl, err := url.Parse(fmt.Sprintf("%s%s", p.config.Host, articleDetailsPath))
	if err != nil {
		return nil, fmt.Errorf("unable to parse the URL with error %s", err)
	}

	params := url.Values{}
	params.Add("id", id)
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create the request with error %s", err)
	}

	response, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http GET call failed to endpoint %s with error %s", baseUrl, err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var result GetArticleInformationResponse
	if err = xml.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("unable to decode response body with error %s", err)
	}

	return result.mapToArticleModel(ctx, p.config), nil
}
