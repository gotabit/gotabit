package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gotabit/gotabit/x/mint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/minting/parameters",
		queryParamsHandlerFn(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		"/minting/epoch-provisions",
		queryEpochProvisionsHandlerFn(clientCtx),
	).Methods("GET")
}

func queryParamsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)

		clientCtx, ok := ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, nil)
		if CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		PostProcessResponse(w, clientCtx, res)
	}
}

// ResponseWithHeight defines a response object type that wraps an original
// response with a height.
type ResponseWithHeight struct {
	Height int64           `json:"height"`
	Result json.RawMessage `json:"result"`
}

// NewResponseWithHeight creates a new ResponseWithHeight instance
func NewResponseWithHeight(height int64, result json.RawMessage) ResponseWithHeight {
	return ResponseWithHeight{
		Height: height,
		Result: result,
	}
}

// PostProcessResponse performs post processing for a REST response. The result
// returned to clients will contain two fields, the height at which the resource
// was queried at and the original result.
func PostProcessResponse(w http.ResponseWriter, ctx client.Context, resp interface{}) {
	var (
		result []byte
		err    error
	)

	if ctx.Height < 0 {
		WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("negative height in response").Error())
		return
	}

	// LegacyAmino used intentionally for REST
	marshaler := ctx.LegacyAmino

	switch res := resp.(type) {
	case []byte:
		result = res

	default:
		result, err = marshaler.MarshalJSON(resp)
		if CheckInternalServerError(w, err) {
			return
		}
	}

	wrappedResp := NewResponseWithHeight(ctx.Height, result)

	output, err := marshaler.MarshalJSON(wrappedResp)
	if CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func queryEpochProvisionsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryEpochProvisions)

		clientCtx, ok := ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, nil)
		if CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		PostProcessResponse(w, clientCtx, res)
	}
}

// CheckInternalServerError attaches an error message to an HTTP 500 INTERNAL SERVER ERROR response.
// Returns false when err is nil; it returns true otherwise.
func CheckInternalServerError(w http.ResponseWriter, err error) bool {
	return CheckError(w, http.StatusInternalServerError, err)
}

// CheckError takes care of writing an error response if err is not nil.
// Returns false when err is nil; it returns true otherwise.
func CheckError(w http.ResponseWriter, status int, err error) bool {
	if err != nil {
		WriteErrorResponse(w, status, err.Error())
		return true
	}

	return false
}

// ErrorResponse defines the attributes of a JSON error response.
type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(code int, err string) ErrorResponse {
	return ErrorResponse{Code: code, Error: err}
}

// WriteErrorResponse prepares and writes a HTTP error
// given a status code and an error message.
func WriteErrorResponse(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(legacy.Cdc.MustMarshalJSON(NewErrorResponse(0, err)))
}

// CheckBadRequestError attaches an error message to an HTTP 400 BAD REQUEST response.
// Returns false when err is nil; it returns true otherwise.
func CheckBadRequestError(w http.ResponseWriter, err error) bool {
	return CheckError(w, http.StatusBadRequest, err)
}

// ParseQueryHeightOrReturnBadRequest sets the height to execute a query if set by the http request.
// It returns false if there was an error parsing the height.
func ParseQueryHeightOrReturnBadRequest(w http.ResponseWriter, clientCtx client.Context, r *http.Request) (client.Context, bool) {
	heightStr := r.FormValue("height")
	if heightStr != "" {
		height, err := strconv.ParseInt(heightStr, 10, 64)
		if CheckBadRequestError(w, err) {
			return clientCtx, false
		}

		if height < 0 {
			WriteErrorResponse(w, http.StatusBadRequest, "height must be equal or greater than zero")
			return clientCtx, false
		}

		if height > 0 {
			clientCtx = clientCtx.WithHeight(height)
		}
	} else {
		clientCtx = clientCtx.WithHeight(0)
	}

	return clientCtx, true
}
