// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// AIAdGenerateService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdGenerateService] method instead.
type AIAdGenerateService struct {
	Options []option.RequestOption
}

// NewAIAdGenerateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdGenerateService(opts ...option.RequestOption) (r *AIAdGenerateService) {
	r = &AIAdGenerateService{}
	r.Options = opts
	return
}

// Submits a highly customized request to generate a video ad, allowing
// fine-grained control over artistic style, aspect ratio, voiceover, background
// music, target audience, and call-to-action elements for professional-grade
// productions.
func (r *AIAdGenerateService) Advanced(ctx context.Context, body AIAdGenerateAdvancedParams, opts ...option.RequestOption) (res *AIAdGenerateAdvancedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate/advanced"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits a request to generate a high-quality video ad using the advanced Veo 2.0
// generative AI model. This is an asynchronous operation, suitable for standard ad
// content creation.
func (r *AIAdGenerateService) Standard(ctx context.Context, body AIAdGenerateStandardParams, opts ...option.RequestOption) (res *AIAdGenerateStandardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GenerateVideoParam struct {
	Prompt        param.Field[string]   `json:"prompt,required"`
	AspectRatio   param.Field[string]   `json:"aspectRatio"`
	BrandColors   param.Field[[]string] `json:"brandColors"`
	LengthSeconds param.Field[int64]    `json:"lengthSeconds"`
	Style         param.Field[string]   `json:"style"`
}

func (r GenerateVideoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdGenerateAdvancedResponse struct {
	EstimatedCompletionTimeSeconds int64                            `json:"estimatedCompletionTimeSeconds"`
	OperationID                    string                           `json:"operationId"`
	JSON                           aiAdGenerateAdvancedResponseJSON `json:"-"`
}

// aiAdGenerateAdvancedResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateAdvancedResponse]
type aiAdGenerateAdvancedResponseJSON struct {
	EstimatedCompletionTimeSeconds apijson.Field
	OperationID                    apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AIAdGenerateAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateStandardResponse struct {
	EstimatedCompletionTimeSeconds int64                            `json:"estimatedCompletionTimeSeconds"`
	OperationID                    string                           `json:"operationId"`
	JSON                           aiAdGenerateStandardResponseJSON `json:"-"`
}

// aiAdGenerateStandardResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateStandardResponse]
type aiAdGenerateStandardResponseJSON struct {
	EstimatedCompletionTimeSeconds apijson.Field
	OperationID                    apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AIAdGenerateStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateStandardResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateAdvancedParams struct {
	Prompt         param.Field[string]                                 `json:"prompt,required"`
	AspectRatio    param.Field[string]                                 `json:"aspectRatio"`
	AudienceTarget param.Field[string]                                 `json:"audienceTarget"`
	BrandAssets    param.Field[[]string]                               `json:"brandAssets" format:"uri"`
	BrandColors    param.Field[[]string]                               `json:"brandColors"`
	CallToAction   param.Field[AIAdGenerateAdvancedParamsCallToAction] `json:"callToAction"`
	LengthSeconds  param.Field[int64]                                  `json:"lengthSeconds"`
	Style          param.Field[string]                                 `json:"style"`
	VoiceoverStyle param.Field[string]                                 `json:"voiceoverStyle"`
	VoiceoverText  param.Field[string]                                 `json:"voiceoverText"`
}

func (r AIAdGenerateAdvancedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdGenerateAdvancedParamsCallToAction struct {
	DisplayTimeSeconds param.Field[int64]  `json:"displayTimeSeconds"`
	Text               param.Field[string] `json:"text"`
	URL                param.Field[string] `json:"url" format:"uri"`
}

func (r AIAdGenerateAdvancedParamsCallToAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdGenerateStandardParams struct {
	GenerateVideo GenerateVideoParam `json:"generate_video,required"`
}

func (r AIAdGenerateStandardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GenerateVideo)
}
