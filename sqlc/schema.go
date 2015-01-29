package sqlc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func sqlc_tmpl_fields_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x58,
		0x5f, 0x6f, 0xdb, 0x36, 0x10, 0x7f, 0xd7, 0xa7, 0x38, 0x18, 0x45, 0x20,
		0x07, 0x8a, 0x12, 0x6c, 0x45, 0x1f, 0x0c, 0xe4, 0xc1, 0x69, 0x9c, 0xd5,
		0x83, 0xeb, 0x04, 0xb1, 0xb2, 0x60, 0x4f, 0x03, 0x23, 0x53, 0x29, 0x37,
		0x95, 0xf2, 0x24, 0x3a, 0x4b, 0x61, 0xe8, 0xbb, 0xef, 0x8e, 0xa4, 0x64,
		0x4a, 0x96, 0x1c, 0x7b, 0xc3, 0x8a, 0xe5, 0xa1, 0x96, 0x8e, 0xf7, 0xe7,
		0x77, 0x77, 0xbf, 0x23, 0xa9, 0x9e, 0x9f, 0x43, 0xf4, 0x69, 0xba, 0x80,
		0x9b, 0xe9, 0x6c, 0x02, 0x8f, 0xe3, 0x05, 0x8c, 0x1f, 0xa2, 0xdb, 0x9f,
		0x26, 0xf3, 0xc9, 0xfd, 0x38, 0x9a, 0x5c, 0xc3, 0x19, 0x8c, 0xe7, 0xbf,
		0xc2, 0xe4, 0x7a, 0x1a, 0x2d, 0x20, 0xba, 0x35, 0xaa, 0x8f, 0xd3, 0xd9,
		0x0c, 0xae, 0x26, 0x30, 0xbb, 0x5d, 0x44, 0xf0, 0xf8, 0x69, 0x32, 0x87,
		0x69, 0x04, 0x28, 0xbf, 0x9f, 0xd4, 0x76, 0x9e, 0xb7, 0xd9, 0xc0, 0xbb,
		0x55, 0xce, 0x97, 0x05, 0x8c, 0x2e, 0x21, 0xa4, 0x27, 0x11, 0x33, 0xc5,
		0x0b, 0x28, 0x4b, 0xbd, 0x96, 0xac, 0x65, 0x6c, 0xd6, 0xe8, 0x49, 0x89,
		0x4c, 0xea, 0x25, 0x6f, 0xc5, 0xe2, 0x3f, 0xd8, 0x33, 0x87, 0xe2, 0xcf,
		0x34, 0xf6, 0x3c, 0xf1, 0x75, 0x95, 0xe5, 0x0a, 0x7c, 0x0f, 0x60, 0xb0,
		0x64, 0x8a, 0x3d, 0xb1, 0x82, 0x9f, 0xe3, 0xd2, 0x80, 0x04, 0x39, 0x4f,
		0x52, 0x1e, 0x2b, 0xfd, 0xac, 0xc4, 0x57, 0x3e, 0xf0, 0x86, 0x9e, 0xf7,
		0xc2, 0x72, 0xad, 0xae, 0xbe, 0xad, 0xf8, 0x55, 0x96, 0xa5, 0x70, 0x09,
		0x56, 0x2f, 0x8c, 0x50, 0x74, 0x9b, 0xf8, 0x09, 0x4b, 0x0b, 0x3e, 0xb4,
		0x2a, 0x37, 0x69, 0xc6, 0xd4, 0x8f, 0x3f, 0x74, 0x68, 0x99, 0x05, 0xff,
		0x22, 0x1c, 0x36, 0x74, 0x3f, 0xbc, 0xef, 0xd1, 0xfd, 0xf0, 0xde, 0xd5,
		0x9d, 0x4a, 0xb5, 0xab, 0x27, 0xa4, 0xf2, 0x2f, 0x5c, 0x95, 0x2e, 0x67,
		0x42, 0x6a, 0x57, 0xb5, 0xda, 0x7c, 0x9d, 0xa6, 0xdd, 0x89, 0x60, 0x1d,
		0xc2, 0x6a, 0x75, 0x53, 0xba, 0xfa, 0xbd, 0x59, 0x55, 0x26, 0x36, 0x95,
		0x0e, 0xab, 0x2e, 0x48, 0x7b, 0xad, 0x3a, 0x33, 0xad, 0x2c, 0x74, 0x8e,
		0x3b, 0xfa, 0xfb, 0x62, 0x74, 0x58, 0x2c, 0x54, 0x2e, 0xe4, 0x73, 0xbf,
		0x89, 0x59, 0x6f, 0xda, 0x44, 0xc8, 0x87, 0x5d, 0x0b, 0x5a, 0x61, 0x4f,
		0x29, 0xa7, 0xd5, 0xad, 0x7e, 0x9f, 0xff, 0xc1, 0xa0, 0xd2, 0xe8, 0xf6,
		0x46, 0x9c, 0x0b, 0x1f, 0xa4, 0x78, 0xf5, 0x2f, 0x02, 0xd8, 0x36, 0xec,
		0x1a, 0x69, 0x7e, 0x80, 0x32, 0x72, 0x95, 0xb4, 0x61, 0x2a, 0x0b, 0x9e,
		0xab, 0x05, 0x57, 0x0b, 0xc5, 0x57, 0x80, 0xdd, 0xe7, 0x79, 0xc2, 0x62,
		0x0e, 0x1b, 0x74, 0x87, 0x52, 0x3f, 0x22, 0xc0, 0x37, 0x82, 0xa7, 0xcb,
		0x60, 0xbb, 0x8a, 0xe0, 0xb7, 0x86, 0x9f, 0xb3, 0x9c, 0x93, 0x31, 0x1a,
		0xe0, 0x64, 0xe5, 0x4c, 0xe2, 0xf8, 0xbc, 0xfb, 0x2d, 0x80, 0x77, 0x4a,
		0xcf, 0x17, 0x45, 0xd1, 0xb3, 0xa5, 0xfd, 0xd1, 0xec, 0xa9, 0xf0, 0x0e,
		0xd1, 0x89, 0x57, 0x14, 0xfa, 0xad, 0x77, 0x1b, 0xc8, 0x48, 0x67, 0x02,
		0xc3, 0xb1, 0x14, 0xc5, 0xbd, 0xd1, 0xb8, 0x5c, 0x92, 0xeb, 0xd2, 0x26,
		0xf3, 0xb0, 0xc2, 0x29, 0xe5, 0xff, 0x20, 0x99, 0xda, 0xf0, 0xbb, 0x24,
		0xd3, 0x13, 0x6d, 0x9b, 0xcc, 0xde, 0xd0, 0xb4, 0x61, 0x81, 0x2f, 0xe0,
		0x54, 0xe8, 0x9a, 0x0c, 0xbb, 0x90, 0x24, 0xd0, 0x8d, 0xe5, 0xe5, 0x90,
		0xd2, 0xea, 0x72, 0xe5, 0x5c, 0xad, 0x73, 0x09, 0x22, 0xa4, 0xba, 0x25,
		0x68, 0x39, 0xf4, 0xf4, 0xce, 0x69, 0x41, 0x1e, 0x02, 0x71, 0x0d, 0xa7,
		0x6b, 0x9d, 0xe9, 0xbf, 0x86, 0xb8, 0x53, 0x30, 0x17, 0xe2, 0xba, 0x0f,
		0xe2, 0x39, 0xfd, 0x59, 0x66, 0xdc, 0x9b, 0x71, 0xa0, 0xfe, 0x37, 0x78,
		0xf1, 0x46, 0x93, 0xbb, 0x20, 0xfa, 0x92, 0xe1, 0x38, 0x16, 0x7a, 0x6a,
		0x87, 0x9d, 0x1a, 0xde, 0x2e, 0x33, 0x6f, 0xec, 0x29, 0x83, 0x09, 0xf5,
		0x87, 0x4f, 0x28, 0xbc, 0x3d, 0x9a, 0xea, 0xf0, 0x49, 0x38, 0xa7, 0x78,
		0x86, 0x5e, 0x85, 0x78, 0x96, 0x22, 0x11, 0x3c, 0x27, 0x65, 0xaa, 0x4c,
		0x47, 0xbc, 0x03, 0x3a, 0x53, 0xc0, 0x69, 0xc1, 0xa9, 0x1e, 0x88, 0xa8,
		0x3b, 0x85, 0xb7, 0x93, 0x74, 0x5b, 0x70, 0x82, 0x0a, 0x2a, 0x9b, 0x65,
		0x7f, 0x11, 0xb0, 0xb6, 0xe2, 0x86, 0x5c, 0x8d, 0x80, 0xfe, 0x25, 0x7c,
		0x06, 0x81, 0x02, 0xdd, 0x8b, 0xef, 0x19, 0x3c, 0x80, 0x3a, 0xe7, 0x11,
		0xa8, 0xb2, 0x93, 0x2b, 0x7b, 0x6b, 0x67, 0x3a, 0xb9, 0x2f, 0x1a, 0x61,
		0x5e, 0xc7, 0x4a, 0xa3, 0x73, 0x72, 0xc0, 0xb7, 0x3a, 0x34, 0x8e, 0x42,
		0x45, 0x44, 0x14, 0xb3, 0x54, 0xb0, 0x62, 0xab, 0x85, 0xb5, 0x31, 0x1d,
		0xad, 0xe8, 0xe2, 0x39, 0x51, 0x77, 0xa3, 0x35, 0x37, 0xb8, 0xed, 0xe6,
		0xb6, 0xb3, 0x79, 0xad, 0x34, 0xb1, 0xcc, 0x7d, 0xa8, 0x26, 0xd6, 0x2a,
		0x6c, 0x84, 0x22, 0x86, 0xbd, 0xb0, 0x74, 0xcd, 0x3b, 0xe6, 0xef, 0x63,
		0x26, 0x97, 0x42, 0xe3, 0xa9, 0x4c, 0x7f, 0xce, 0x84, 0xec, 0xb3, 0x6c,
		0xa2, 0x1c, 0x02, 0xe9, 0xb6, 0x3c, 0x6c, 0xd9, 0x6a, 0xe8, 0x10, 0xc3,
		0xe9, 0xbe, 0xba, 0x0e, 0xeb, 0xf9, 0xf1, 0x87, 0xcd, 0x02, 0xb9, 0x44,
		0x68, 0x2c, 0x90, 0x1c, 0x60, 0xae, 0xbb, 0x0f, 0x31, 0xdd, 0xf2, 0xf4,
		0x1c, 0x05, 0x5a, 0x3c, 0x79, 0x5d, 0xe5, 0xb5, 0x98, 0x5e, 0x8c, 0x78,
		0x9c, 0x3f, 0x17, 0xb5, 0x98, 0x5e, 0x8c, 0xf8, 0xe3, 0x17, 0x91, 0x2e,
		0x47, 0x56, 0xac, 0x5f, 0x48, 0x7e, 0x0c, 0xfa, 0x24, 0xc6, 0x0d, 0x6a,
		0x2d, 0x03, 0xe0, 0x18, 0xcb, 0xb6, 0x3b, 0x00, 0x86, 0x11, 0x20, 0x0c,
		0xc3, 0xc6, 0x49, 0xb4, 0xa5, 0xb7, 0x48, 0xe0, 0x44, 0xc7, 0x84, 0xcb,
		0x4b, 0x90, 0x22, 0x05, 0x93, 0xd2, 0x41, 0xac, 0xd7, 0x9a, 0x86, 0x83,
		0x23, 0xf3, 0x18, 0x87, 0xb2, 0x4e, 0x1f, 0xdc, 0x49, 0x88, 0xc3, 0xfa,
		0xa5, 0x5a, 0xc5, 0x98, 0xd6, 0xaa, 0x55, 0x53, 0x5d, 0x4e, 0x9d, 0x88,
		0xae, 0x20, 0x65, 0x13, 0xd8, 0xaa, 0x51, 0x32, 0xa5, 0x71, 0x40, 0x04,
		0x2b, 0x81, 0xe3, 0x35, 0xf7, 0x7f, 0x00, 0xd9, 0xae, 0xd7, 0x54, 0x20,
		0xf4, 0xb5, 0xcc, 0xf2, 0x80, 0xd7, 0x0c, 0x70, 0x58, 0xc0, 0xea, 0xfe,
		0x3b, 0x1c, 0x38, 0xe9, 0x71, 0xde, 0xcb, 0xb4, 0x46, 0x9c, 0x36, 0xdf,
		0x1a, 0xf1, 0xda, 0xac, 0x6b, 0xc4, 0x6d, 0x71, 0xcf, 0xfc, 0x95, 0xd5,
		0xa3, 0x5b, 0xf7, 0x23, 0x58, 0x39, 0x2e, 0x7c, 0x77, 0xff, 0x71, 0xb9,
		0x77, 0x78, 0xc7, 0x4c, 0xbf, 0xdc, 0x56, 0xed, 0x6b, 0x94, 0x8e, 0x37,
		0x32, 0x3f, 0x46, 0xa2, 0x1b, 0xd7, 0x59, 0xd5, 0xde, 0x8a, 0xf6, 0x56,
		0xb3, 0xb7, 0x92, 0x7d, 0x55, 0x2c, 0x8f, 0x1d, 0xe4, 0x31, 0x01, 0xc7,
		0x3d, 0xc8, 0x54, 0xcc, 0x2d, 0x55, 0x1c, 0xea, 0xa4, 0x8e, 0xf0, 0xf5,
		0x99, 0x7d, 0x7b, 0xe2, 0x1d, 0x0e, 0x71, 0xee, 0xad, 0x33, 0x1a, 0xfc,
		0xc1, 0xa0, 0x39, 0x44, 0xa6, 0xd2, 0x3d, 0x13, 0x56, 0x81, 0x38, 0x2e,
		0x29, 0xaa, 0x6d, 0x4f, 0x4e, 0x3a, 0xd6, 0xe1, 0x9e, 0xe8, 0x63, 0x03,
		0x3d, 0xb9, 0xdf, 0x1e, 0xae, 0x3f, 0x3a, 0xc9, 0x5a, 0x47, 0xc4, 0x11,
		0xce, 0xef, 0x58, 0xce, 0xf1, 0x0b, 0x76, 0xe8, 0x9c, 0xa0, 0x4d, 0xb0,
		0x35, 0xd7, 0x3c, 0x7d, 0xa0, 0xc3, 0xd9, 0x59, 0xfb, 0x40, 0x6f, 0x9d,
		0x83, 0x87, 0x46, 0xee, 0x39, 0x2d, 0xc9, 0xcf, 0xbe, 0xc3, 0xd2, 0x45,
		0x57, 0x0b, 0x37, 0x57, 0x02, 0x9f, 0xe4, 0xb3, 0x25, 0xbd, 0x7d, 0xdb,
		0xfc, 0x42, 0xe7, 0xe7, 0x08, 0xc8, 0x65, 0x60, 0x56, 0x90, 0xae, 0x65,
		0x00, 0x77, 0xd5, 0xff, 0x5c, 0x8c, 0x2c, 0x8a, 0x5a, 0x80, 0xb1, 0x8e,
		0xa9, 0x5e, 0xf7, 0xb1, 0xed, 0xa4, 0xb0, 0xf7, 0xd4, 0x76, 0x53, 0x69,
		0x2c, 0x6c, 0x66, 0x5f, 0x70, 0xe4, 0xe2, 0x00, 0xee, 0xe9, 0xd7, 0xc0,
		0x7f, 0x1b, 0x73, 0xe3, 0xe2, 0xa5, 0xfb, 0xa4, 0xb3, 0x68, 0x7f, 0x22,
		0x14, 0x4e, 0xab, 0x03, 0xf8, 0xef, 0x2f, 0x85, 0x45, 0xe9, 0x55, 0x57,
		0xc1, 0x9d, 0xbb, 0x60, 0xeb, 0x6e, 0x7e, 0x44, 0xd5, 0x0f, 0xb8, 0xc0,
		0x37, 0x79, 0x4c, 0xb7, 0x85, 0x41, 0xc3, 0x72, 0x10, 0x80, 0x15, 0xd0,
		0x76, 0x47, 0x02, 0x7c, 0x13, 0xf2, 0x77, 0x02, 0xee, 0x38, 0xdb, 0xf9,
		0x42, 0xb3, 0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x7e, 0x5f,
		0x1a, 0x65, 0x13, 0x00, 0x00,
	},
		"sqlc/tmpl/fields.tmpl",
	)
}

func sqlc_tmpl_schema_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x5b, 0x6f, 0x9b, 0x30, 0x14, 0x7e, 0xe7, 0x57, 0x1c, 0xa1, 0x6a, 0x82,
		0x29, 0x23, 0xef, 0x91, 0xf2, 0xc0, 0x54, 0xda, 0x22, 0x31, 0xb2, 0x15,
		0xda, 0x6a, 0x9a, 0xa6, 0xca, 0x21, 0x26, 0x61, 0xe5, 0x36, 0x6c, 0xba,
		0x45, 0x11, 0xff, 0x7d, 0xbe, 0x40, 0x42, 0x8a, 0x69, 0xd0, 0xfa, 0x50,
		0x11, 0xfb, 0x9c, 0xef, 0x76, 0x6c, 0xcf, 0xe7, 0x10, 0xde, 0xb9, 0x01,
		0xdc, 0xb8, 0x9e, 0x03, 0x4f, 0x76, 0x00, 0xf6, 0x43, 0xb8, 0xba, 0x75,
		0x7c, 0xe7, 0xde, 0x0e, 0x9d, 0x6b, 0xf8, 0x04, 0xb6, 0xff, 0x1d, 0x9c,
		0x6b, 0x37, 0x0c, 0x20, 0x5c, 0xc9, 0xd2, 0x27, 0xd7, 0xf3, 0xe0, 0xb3,
		0x03, 0xde, 0x2a, 0x08, 0xe1, 0xe9, 0xce, 0xf1, 0xc1, 0x0d, 0x81, 0xad,
		0xdf, 0x3b, 0xc7, 0x3e, 0x8d, 0xc1, 0xda, 0x21, 0x1c, 0x0e, 0x60, 0x7d,
		0xad, 0x8a, 0x57, 0x9c, 0xa3, 0x3c, 0xc2, 0x56, 0x98, 0x64, 0x98, 0x50,
		0x94, 0x95, 0xd0, 0x34, 0xf0, 0x10, 0xb8, 0xfe, 0x2d, 0x90, 0xdf, 0x69,
		0x04, 0x8f, 0xce, 0x7d, 0xe0, 0xae, 0xfc, 0xb7, 0xe5, 0x8f, 0xb8, 0x22,
		0x49, 0x91, 0xb3, 0x62, 0x4d, 0x63, 0x5b, 0x57, 0x74, 0x5f, 0x62, 0x02,
		0x8b, 0x25, 0x58, 0xa1, 0xf8, 0xe2, 0xeb, 0x25, 0x8a, 0x5e, 0xd0, 0x16,
		0xcb, 0xd6, 0xf6, 0x9b, 0xaf, 0x27, 0x59, 0x59, 0x54, 0x14, 0x0c, 0x0d,
		0x40, 0xdf, 0x26, 0x74, 0x57, 0xaf, 0xad, 0xa8, 0xc8, 0xe6, 0x64, 0x57,
		0x53, 0xfc, 0x6b, 0xce, 0x59, 0xc5, 0x3f, 0x9d, 0xef, 0x13, 0x5a, 0x25,
		0xf9, 0x96, 0xe8, 0x9a, 0xa9, 0x69, 0x51, 0x91, 0x13, 0x0a, 0x41, 0xb4,
		0xc3, 0x19, 0x82, 0x25, 0xe8, 0x1c, 0xb7, 0xfd, 0xd5, 0x34, 0xba, 0xd0,
		0x51, 0xa1, 0x9c, 0x91, 0x5c, 0x3d, 0xcf, 0x98, 0x22, 0xa9, 0x06, 0xad,
		0xd3, 0x56, 0x0e, 0x97, 0xc8, 0xb5, 0xd0, 0xc2, 0x2b, 0xfe, 0xe0, 0x8a,
		0x55, 0x58, 0x3e, 0xca, 0xb8, 0x24, 0x60, 0x2c, 0x75, 0x44, 0xe1, 0xc0,
		0x18, 0xcf, 0x41, 0x62, 0x0e, 0xc2, 0x0a, 0x6f, 0x12, 0x9c, 0x6e, 0x04,
		0x0c, 0x48, 0x88, 0x87, 0xb2, 0xe4, 0x10, 0xf1, 0x09, 0x82, 0x29, 0xb6,
		0x78, 0x12, 0xb1, 0x48, 0x80, 0x2d, 0x89, 0x1e, 0x59, 0x8f, 0xf3, 0x8d,
		0xec, 0x45, 0x69, 0x82, 0x08, 0x48, 0x53, 0x1a, 0xd3, 0x14, 0xd7, 0x79,
		0x04, 0x06, 0x85, 0x8f, 0x4a, 0x5d, 0x26, 0xb8, 0x24, 0xc0, 0x29, 0x8e,
		0x28, 0x77, 0x61, 0x98, 0x70, 0x98, 0xd0, 0x22, 0x13, 0x61, 0xc5, 0x92,
		0x45, 0x98, 0xaa, 0x30, 0xad, 0xab, 0xbc, 0xdd, 0x9a, 0xc2, 0xcb, 0x3f,
		0x94, 0x10, 0xc2, 0xe6, 0xb7, 0x9a, 0xf9, 0x88, 0x13, 0xbc, 0x31, 0xa8,
		0xd5, 0xd1, 0xcd, 0xc4, 0x3c, 0x86, 0x58, 0xba, 0x39, 0x85, 0xcf, 0x26,
		0x06, 0x6a, 0xd9, 0x4c, 0xc9, 0x71, 0xf2, 0xdd, 0xa7, 0xff, 0xa0, 0xec,
		0xe7, 0x05, 0x53, 0x26, 0x37, 0x32, 0xbb, 0x05, 0x50, 0x4b, 0xb9, 0x31,
		0xeb, 0x7a, 0x8e, 0xf3, 0x6b, 0x27, 0xb8, 0x40, 0x7c, 0xab, 0x99, 0xe4,
		0x8c, 0xd7, 0x2b, 0xa3, 0xa4, 0x96, 0xc0, 0x9a, 0x02, 0xf2, 0x05, 0xed,
		0xd7, 0x58, 0x81, 0x94, 0xc4, 0x1d, 0x0a, 0x2c, 0xd9, 0x85, 0xd0, 0x41,
		0x26, 0xd1, 0x12, 0x8c, 0x0c, 0x84, 0x2b, 0x07, 0x9c, 0x12, 0x7c, 0x5e,
		0xdd, 0xc9, 0x69, 0x7d, 0xcd, 0xf9, 0xdf, 0xe0, 0x4e, 0xed, 0x65, 0xaa,
		0xdd, 0x15, 0xbf, 0xa4, 0x5b, 0xbe, 0x0c, 0xec, 0xdd, 0xc0, 0x71, 0xf2,
		0xb7, 0xbb, 0x12, 0x46, 0xce, 0xb7, 0xcf, 0x86, 0xad, 0xac, 0x1b, 0x1c,
		0xbb, 0xb7, 0x55, 0x46, 0x7f, 0x68, 0x47, 0xd6, 0x59, 0x0b, 0x4d, 0xac,
		0x50, 0xee, 0x09, 0x3e, 0x93, 0x1f, 0xc3, 0xd3, 0x24, 0x3b, 0x7b, 0x97,
		0x0c, 0xc8, 0xd3, 0xc3, 0x42, 0xff, 0xf1, 0x53, 0x48, 0x18, 0x08, 0xeb,
		0xaf, 0x4f, 0x3f, 0x86, 0xc3, 0xc7, 0x62, 0xcc, 0x8b, 0xae, 0x3c, 0x98,
		0xba, 0x39, 0x3c, 0x9a, 0x62, 0x6a, 0x3d, 0x87, 0xef, 0xbf, 0x86, 0xaf,
		0xa8, 0x82, 0xe7, 0x67, 0xf5, 0x6b, 0xb8, 0x1c, 0xbb, 0x66, 0xb2, 0x4d,
		0xa9, 0x74, 0xb4, 0x89, 0x85, 0xa5, 0xfd, 0xe7, 0x9b, 0xba, 0x50, 0xe5,
		0x34, 0x22, 0xfa, 0xbd, 0xa4, 0x7a, 0xa1, 0x9c, 0x45, 0xc4, 0xcd, 0x88,
		0x44, 0xbc, 0xe4, 0x85, 0xa5, 0xb2, 0xec, 0x86, 0x79, 0x5c, 0x3b, 0x5c,
		0x08, 0x11, 0xd4, 0x59, 0xf4, 0x19, 0x1b, 0xed, 0x5f, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xd4, 0xa6, 0xda, 0x3b, 0xc5, 0x07, 0x00, 0x00,
	},
		"sqlc/tmpl/schema.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"sqlc/tmpl/fields.tmpl": sqlc_tmpl_fields_tmpl,
	"sqlc/tmpl/schema.tmpl": sqlc_tmpl_schema_tmpl,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"sqlc": &_bintree_t{nil, map[string]*_bintree_t{
		"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
			"fields.tmpl": &_bintree_t{sqlc_tmpl_fields_tmpl, map[string]*_bintree_t{
			}},
			"schema.tmpl": &_bintree_t{sqlc_tmpl_schema_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}
