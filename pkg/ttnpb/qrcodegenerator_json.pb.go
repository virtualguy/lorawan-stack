// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.0
// - protoc             v3.9.1
// source: lorawan-stack/api/qrcodegenerator.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GenerateEndDeviceQRCodeRequest message to JSON.
func (x *GenerateEndDeviceQRCodeRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.FormatId != "" || s.HasField("format_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("format_id")
		s.WriteString(x.FormatId)
	}
	if x.EndDevice != nil || s.HasField("end_device") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("end_device")
		x.EndDevice.MarshalProtoJSON(s.WithField("end_device"))
	}
	if x.Image != nil || s.HasField("image") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("image")
		// NOTE: GenerateEndDeviceQRCodeRequest_Image does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Image)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GenerateEndDeviceQRCodeRequest to JSON.
func (x GenerateEndDeviceQRCodeRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the GenerateEndDeviceQRCodeRequest message from JSON.
func (x *GenerateEndDeviceQRCodeRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "format_id", "formatId":
			s.AddField("format_id")
			x.FormatId = s.ReadString()
		case "end_device", "endDevice":
			if !s.ReadNil() {
				x.EndDevice = &EndDevice{}
				x.EndDevice.UnmarshalProtoJSON(s.WithField("end_device", true))
			}
		case "image":
			s.AddField("image")
			// NOTE: GenerateEndDeviceQRCodeRequest_Image does not seem to implement UnmarshalProtoJSON.
			var v GenerateEndDeviceQRCodeRequest_Image
			gogo.UnmarshalMessage(s, &v)
			x.Image = &v
		}
	})
}

// UnmarshalJSON unmarshals the GenerateEndDeviceQRCodeRequest from JSON.
func (x *GenerateEndDeviceQRCodeRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
