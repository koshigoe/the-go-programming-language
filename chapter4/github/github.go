// パッケージ github は、 GitHub のイシュートラッカーに対する Go の API を提供します。
// https://developer.github.com/v3/search/#search-issues を参照のこと。
package github

import "time"
import "encoding/json"
import "net/url"
import "net/http"
import "strings"
import "fmt"

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues は GitHub のイシュートラッカーに問い合わせます。
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// すべての実行パスで resp.Body を閉じなければなりません。
	// (この処理を簡単にする `defer` を第5章で説明しています。)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
