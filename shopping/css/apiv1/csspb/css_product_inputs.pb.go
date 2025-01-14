// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.2
// source: google/shopping/css/v1/css_product_inputs.proto

package csspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	typepb "cloud.google.com/go/shopping/type/typepb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This resource represents input data you submit for a CSS Product, not
// the processed CSS Product that you see in CSS Center, in Shopping Ads, or
// across Google surfaces.
type CssProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the CSS Product input.
	// Format:
	// `accounts/{account}/cssProductInputs/{css_product_input}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The name of the processed CSS Product.
	// Format:
	// `accounts/{account}/cssProducts/{css_product}`
	// "
	FinalName string `protobuf:"bytes,2,opt,name=final_name,json=finalName,proto3" json:"final_name,omitempty"`
	// Required. Your unique identifier for the CSS Product. This is the same for
	// the CSS Product input and processed CSS Product. We only allow ids with
	// alphanumerics, underscores and dashes. See the [products feed
	// specification](https://support.google.com/merchants/answer/188494#id) for
	// details.
	RawProvidedId string `protobuf:"bytes,3,opt,name=raw_provided_id,json=rawProvidedId,proto3" json:"raw_provided_id,omitempty"`
	// Required. The two-letter [ISO
	// 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code for the CSS
	// Product.
	ContentLanguage string `protobuf:"bytes,4,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
	// Required. The [feed
	// label](https://developers.google.com/shopping-content/guides/products/feed-labels)
	// for the CSS Product.
	// Feed Label is synonymous to "target country" and hence should always be a
	// valid region code. For example: 'DE' for Germany, 'FR' for France.
	FeedLabel string `protobuf:"bytes,5,opt,name=feed_label,json=feedLabel,proto3" json:"feed_label,omitempty"`
	// Represents the existing version (freshness) of the CSS Product, which
	// can be used to preserve the right order when multiple updates are done at
	// the same time.
	//
	// This field must not be set to the future time.
	//
	// If set, the update is prevented if a newer version of the item already
	// exists in our system (that is the last update time of the existing
	// CSS products is later than the freshness time set in the update). If
	// the update happens, the last update time is then set to this freshness
	// time.
	//
	// If not set, the update will not be prevented and the last update time will
	// default to when this request was received by the CSS API.
	//
	// If the operation is prevented, the aborted exception will be
	// thrown.
	FreshnessTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=freshness_time,json=freshnessTime,proto3" json:"freshness_time,omitempty"`
	// A list of CSS Product attributes.
	Attributes *Attributes `protobuf:"bytes,7,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// A list of custom (CSS-provided) attributes. It can also be used for
	// submitting any attribute of the feed specification in its generic
	// form (for example:
	// `{ "name": "size type", "value": "regular" }`).
	// This is useful for submitting attributes not explicitly exposed by the
	// API, such as additional attributes used for Buy on Google.
	CustomAttributes []*typepb.CustomAttribute `protobuf:"bytes,8,rep,name=custom_attributes,json=customAttributes,proto3" json:"custom_attributes,omitempty"`
}

func (x *CssProductInput) Reset() {
	*x = CssProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CssProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CssProductInput) ProtoMessage() {}

func (x *CssProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CssProductInput.ProtoReflect.Descriptor instead.
func (*CssProductInput) Descriptor() ([]byte, []int) {
	return file_google_shopping_css_v1_css_product_inputs_proto_rawDescGZIP(), []int{0}
}

func (x *CssProductInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CssProductInput) GetFinalName() string {
	if x != nil {
		return x.FinalName
	}
	return ""
}

func (x *CssProductInput) GetRawProvidedId() string {
	if x != nil {
		return x.RawProvidedId
	}
	return ""
}

func (x *CssProductInput) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

func (x *CssProductInput) GetFeedLabel() string {
	if x != nil {
		return x.FeedLabel
	}
	return ""
}

func (x *CssProductInput) GetFreshnessTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FreshnessTime
	}
	return nil
}

func (x *CssProductInput) GetAttributes() *Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CssProductInput) GetCustomAttributes() []*typepb.CustomAttribute {
	if x != nil {
		return x.CustomAttributes
	}
	return nil
}

// Request message for the InsertCssProductInput method.
type InsertCssProductInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The account where this CSS Product will be inserted.
	// Format: accounts/{account}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The CSS Product Input to insert.
	CssProductInput *CssProductInput `protobuf:"bytes,2,opt,name=css_product_input,json=cssProductInput,proto3" json:"css_product_input,omitempty"`
	// Required. The primary or supplemental feed id. If CSS Product already
	// exists and feed id provided is different, then the CSS Product will be
	// moved to a new feed. Note: For now, CSSs do not need to provide feed ids as
	// we create feeds on the fly. We do not have supplemental feed support for
	// CSS Products yet.
	FeedId int64 `protobuf:"varint,3,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
}

func (x *InsertCssProductInputRequest) Reset() {
	*x = InsertCssProductInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertCssProductInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertCssProductInputRequest) ProtoMessage() {}

func (x *InsertCssProductInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertCssProductInputRequest.ProtoReflect.Descriptor instead.
func (*InsertCssProductInputRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_css_v1_css_product_inputs_proto_rawDescGZIP(), []int{1}
}

func (x *InsertCssProductInputRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *InsertCssProductInputRequest) GetCssProductInput() *CssProductInput {
	if x != nil {
		return x.CssProductInput
	}
	return nil
}

func (x *InsertCssProductInputRequest) GetFeedId() int64 {
	if x != nil {
		return x.FeedId
	}
	return 0
}

// Request message for the DeleteCssProductInput method.
type DeleteCssProductInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the CSS product input resource to delete.
	// Format: accounts/{account}/cssProductInputs/{css_product_input}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Content API Supplemental Feed ID.
	// The field must not be set if the action applies to a primary feed.
	// If the field is set, then product action applies to a supplemental feed
	// instead of primary Content API feed.
	SupplementalFeedId *int64 `protobuf:"varint,2,opt,name=supplemental_feed_id,json=supplementalFeedId,proto3,oneof" json:"supplemental_feed_id,omitempty"`
}

func (x *DeleteCssProductInputRequest) Reset() {
	*x = DeleteCssProductInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCssProductInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCssProductInputRequest) ProtoMessage() {}

func (x *DeleteCssProductInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCssProductInputRequest.ProtoReflect.Descriptor instead.
func (*DeleteCssProductInputRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_css_v1_css_product_inputs_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCssProductInputRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteCssProductInputRequest) GetSupplementalFeedId() int64 {
	if x != nil && x.SupplementalFeedId != nil {
		return *x.SupplementalFeedId
	}
	return 0
}

var File_google_shopping_css_v1_css_product_inputs_proto protoreflect.FileDescriptor

var file_google_shopping_css_v1_css_product_inputs_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x73, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x04,
	0x0a, 0x0f, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0f, 0x72, 0x61, 0x77,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x64, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x52, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x60, 0xea, 0x41, 0x5d, 0x0a, 0x22, 0x63, 0x73, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x37,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x7d, 0x2f, 0x63, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x7d, 0x22, 0xda, 0x01, 0x0a, 0x1c, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24,
	0x12, 0x22, 0x63, 0x73, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x11,
	0x63, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x63, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x65,
	0x65, 0x64, 0x49, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x63, 0x73, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x69, 0x64, 0x32, 0xc0, 0x03, 0x0a, 0x17, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc2, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x34, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x44, 0x3a, 0x11, 0x63, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x2f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x3a,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x37,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x2a, 0x28, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x3f, 0xca, 0x41, 0x12, 0x63, 0x73, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41,
	0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0xb7, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x63, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x43, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x63,
	0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x73, 0x73, 0x70, 0x62, 0x3b, 0x63,
	0x73, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x73, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5c,
	0x43, 0x73, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x43, 0x73, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_css_v1_css_product_inputs_proto_rawDescOnce sync.Once
	file_google_shopping_css_v1_css_product_inputs_proto_rawDescData = file_google_shopping_css_v1_css_product_inputs_proto_rawDesc
)

func file_google_shopping_css_v1_css_product_inputs_proto_rawDescGZIP() []byte {
	file_google_shopping_css_v1_css_product_inputs_proto_rawDescOnce.Do(func() {
		file_google_shopping_css_v1_css_product_inputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_css_v1_css_product_inputs_proto_rawDescData)
	})
	return file_google_shopping_css_v1_css_product_inputs_proto_rawDescData
}

var file_google_shopping_css_v1_css_product_inputs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_shopping_css_v1_css_product_inputs_proto_goTypes = []interface{}{
	(*CssProductInput)(nil),              // 0: google.shopping.css.v1.CssProductInput
	(*InsertCssProductInputRequest)(nil), // 1: google.shopping.css.v1.InsertCssProductInputRequest
	(*DeleteCssProductInputRequest)(nil), // 2: google.shopping.css.v1.DeleteCssProductInputRequest
	(*timestamppb.Timestamp)(nil),        // 3: google.protobuf.Timestamp
	(*Attributes)(nil),                   // 4: google.shopping.css.v1.Attributes
	(*typepb.CustomAttribute)(nil),       // 5: google.shopping.type.CustomAttribute
	(*emptypb.Empty)(nil),                // 6: google.protobuf.Empty
}
var file_google_shopping_css_v1_css_product_inputs_proto_depIdxs = []int32{
	3, // 0: google.shopping.css.v1.CssProductInput.freshness_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.shopping.css.v1.CssProductInput.attributes:type_name -> google.shopping.css.v1.Attributes
	5, // 2: google.shopping.css.v1.CssProductInput.custom_attributes:type_name -> google.shopping.type.CustomAttribute
	0, // 3: google.shopping.css.v1.InsertCssProductInputRequest.css_product_input:type_name -> google.shopping.css.v1.CssProductInput
	1, // 4: google.shopping.css.v1.CssProductInputsService.InsertCssProductInput:input_type -> google.shopping.css.v1.InsertCssProductInputRequest
	2, // 5: google.shopping.css.v1.CssProductInputsService.DeleteCssProductInput:input_type -> google.shopping.css.v1.DeleteCssProductInputRequest
	0, // 6: google.shopping.css.v1.CssProductInputsService.InsertCssProductInput:output_type -> google.shopping.css.v1.CssProductInput
	6, // 7: google.shopping.css.v1.CssProductInputsService.DeleteCssProductInput:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_shopping_css_v1_css_product_inputs_proto_init() }
func file_google_shopping_css_v1_css_product_inputs_proto_init() {
	if File_google_shopping_css_v1_css_product_inputs_proto != nil {
		return
	}
	file_google_shopping_css_v1_css_product_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CssProductInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertCssProductInputRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCssProductInputRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_shopping_css_v1_css_product_inputs_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_css_v1_css_product_inputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_css_v1_css_product_inputs_proto_goTypes,
		DependencyIndexes: file_google_shopping_css_v1_css_product_inputs_proto_depIdxs,
		MessageInfos:      file_google_shopping_css_v1_css_product_inputs_proto_msgTypes,
	}.Build()
	File_google_shopping_css_v1_css_product_inputs_proto = out.File
	file_google_shopping_css_v1_css_product_inputs_proto_rawDesc = nil
	file_google_shopping_css_v1_css_product_inputs_proto_goTypes = nil
	file_google_shopping_css_v1_css_product_inputs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CssProductInputsServiceClient is the client API for CssProductInputsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CssProductInputsServiceClient interface {
	// Uploads a CssProductInput to your CSS Center account. If an
	// input with the same contentLanguage, identity, feedLabel and feedId already
	// exists, this method replaces that entry.
	//
	// After inserting, updating, or deleting a CSS Product input, it may
	// take several minutes before the processed CSS Product can be retrieved.
	InsertCssProductInput(ctx context.Context, in *InsertCssProductInputRequest, opts ...grpc.CallOption) (*CssProductInput, error)
	// Deletes a CSS Product input from your CSS Center account.
	//
	// After a delete it may take several minutes until the input is no longer
	// available.
	DeleteCssProductInput(ctx context.Context, in *DeleteCssProductInputRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cssProductInputsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCssProductInputsServiceClient(cc grpc.ClientConnInterface) CssProductInputsServiceClient {
	return &cssProductInputsServiceClient{cc}
}

func (c *cssProductInputsServiceClient) InsertCssProductInput(ctx context.Context, in *InsertCssProductInputRequest, opts ...grpc.CallOption) (*CssProductInput, error) {
	out := new(CssProductInput)
	err := c.cc.Invoke(ctx, "/google.shopping.css.v1.CssProductInputsService/InsertCssProductInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cssProductInputsServiceClient) DeleteCssProductInput(ctx context.Context, in *DeleteCssProductInputRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.shopping.css.v1.CssProductInputsService/DeleteCssProductInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CssProductInputsServiceServer is the server API for CssProductInputsService service.
type CssProductInputsServiceServer interface {
	// Uploads a CssProductInput to your CSS Center account. If an
	// input with the same contentLanguage, identity, feedLabel and feedId already
	// exists, this method replaces that entry.
	//
	// After inserting, updating, or deleting a CSS Product input, it may
	// take several minutes before the processed CSS Product can be retrieved.
	InsertCssProductInput(context.Context, *InsertCssProductInputRequest) (*CssProductInput, error)
	// Deletes a CSS Product input from your CSS Center account.
	//
	// After a delete it may take several minutes until the input is no longer
	// available.
	DeleteCssProductInput(context.Context, *DeleteCssProductInputRequest) (*emptypb.Empty, error)
}

// UnimplementedCssProductInputsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCssProductInputsServiceServer struct {
}

func (*UnimplementedCssProductInputsServiceServer) InsertCssProductInput(context.Context, *InsertCssProductInputRequest) (*CssProductInput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCssProductInput not implemented")
}
func (*UnimplementedCssProductInputsServiceServer) DeleteCssProductInput(context.Context, *DeleteCssProductInputRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCssProductInput not implemented")
}

func RegisterCssProductInputsServiceServer(s *grpc.Server, srv CssProductInputsServiceServer) {
	s.RegisterService(&_CssProductInputsService_serviceDesc, srv)
}

func _CssProductInputsService_InsertCssProductInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertCssProductInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CssProductInputsServiceServer).InsertCssProductInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.css.v1.CssProductInputsService/InsertCssProductInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CssProductInputsServiceServer).InsertCssProductInput(ctx, req.(*InsertCssProductInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CssProductInputsService_DeleteCssProductInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCssProductInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CssProductInputsServiceServer).DeleteCssProductInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.css.v1.CssProductInputsService/DeleteCssProductInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CssProductInputsServiceServer).DeleteCssProductInput(ctx, req.(*DeleteCssProductInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CssProductInputsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.css.v1.CssProductInputsService",
	HandlerType: (*CssProductInputsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertCssProductInput",
			Handler:    _CssProductInputsService_InsertCssProductInput_Handler,
		},
		{
			MethodName: "DeleteCssProductInput",
			Handler:    _CssProductInputsService_DeleteCssProductInput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/css/v1/css_product_inputs.proto",
}
