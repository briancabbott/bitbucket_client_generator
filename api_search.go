/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type SearchApi interface {

	/*
	 * SearchAccount Search for code in the repositories of the specified team.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param username The account to search in; either the username or the UUID in curly braces
	 * @return SearchApiApiSearchAccountRequest
	 */
	SearchAccount(ctx _context.Context, username string) SearchApiApiSearchAccountRequest

	/*
	 * SearchAccountExecute executes the request
	 * @return SearchResultPage
	 */
	SearchAccountExecute(r SearchApiApiSearchAccountRequest) (SearchResultPage, *_nethttp.Response, error)
}

// SearchApiService SearchApi service
type SearchApiService service

type SearchApiApiSearchAccountRequest struct {
	ctx _context.Context
	ApiService SearchApi
	username string
	searchQuery *string
	page *int32
	pagelen *int32
}

func (r SearchApiApiSearchAccountRequest) SearchQuery(searchQuery string) SearchApiApiSearchAccountRequest {
	r.searchQuery = &searchQuery
	return r
}
func (r SearchApiApiSearchAccountRequest) Page(page int32) SearchApiApiSearchAccountRequest {
	r.page = &page
	return r
}
func (r SearchApiApiSearchAccountRequest) Pagelen(pagelen int32) SearchApiApiSearchAccountRequest {
	r.pagelen = &pagelen
	return r
}

func (r SearchApiApiSearchAccountRequest) Execute() (SearchResultPage, *_nethttp.Response, error) {
	return r.ApiService.SearchAccountExecute(r)
}

/*
 * SearchAccount Search for code in the repositories of the specified team.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param username The account to search in; either the username or the UUID in curly braces
 * @return SearchApiApiSearchAccountRequest
 */
func (a *SearchApiService) SearchAccount(ctx _context.Context, username string) SearchApiApiSearchAccountRequest {
	return SearchApiApiSearchAccountRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

/*
 * Execute executes the request
 * @return SearchResultPage
 */
func (a *SearchApiService) SearchAccountExecute(r SearchApiApiSearchAccountRequest) (SearchResultPage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResultPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.SearchAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{username}/search/code"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", _neturl.PathEscape(parameterToString(r.username, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.searchQuery == nil {
		return localVarReturnValue, nil, reportError("searchQuery is required and must be specified")
	}

	localVarQueryParams.Add("search_query", parameterToString(*r.searchQuery, ""))
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pagelen != nil {
		localVarQueryParams.Add("pagelen", parameterToString(*r.pagelen, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}