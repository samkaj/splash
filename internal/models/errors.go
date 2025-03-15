package models

import "errors"

var ErrInvalidHexString = errors.New("invalid hex string");
var ErrInvalidHexPrefix = errors.New("hex strings must begin with '#' or '0x'");
var ErrInvalidHexLength = errors.New("invalid hex length");
var ErrHexConversionFailure = errors.New("failed to parse hexadecimal");
