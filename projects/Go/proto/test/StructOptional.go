// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.2.0.0

package test

import "fmt"
import "strconv"
import "strings"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// StructOptional key
type StructOptionalKey struct {
    StructSimpleKey
}

// Convert StructOptional flags key to string
func (k *StructOptionalKey) String() string {
    var sb strings.Builder
    sb.WriteString("StructOptionalKey(")
    sb.WriteString(k.StructSimpleKey.String())
    sb.WriteString(")")
    return sb.String()
}

// StructOptional struct
type StructOptional struct {
    *StructSimple
    F100 *bool `json:"f100"`
    F101 *bool `json:"f101"`
    F102 *bool `json:"f102"`
    F103 *byte `json:"f103"`
    F104 *byte `json:"f104"`
    F105 *byte `json:"f105"`
    F106 *rune `json:"f106"`
    F107 *rune `json:"f107"`
    F108 *rune `json:"f108"`
    F109 *rune `json:"f109"`
    F110 *rune `json:"f110"`
    F111 *rune `json:"f111"`
    F112 *int8 `json:"f112"`
    F113 *int8 `json:"f113"`
    F114 *int8 `json:"f114"`
    F115 *uint8 `json:"f115"`
    F116 *uint8 `json:"f116"`
    F117 *uint8 `json:"f117"`
    F118 *int16 `json:"f118"`
    F119 *int16 `json:"f119"`
    F120 *int16 `json:"f120"`
    F121 *uint16 `json:"f121"`
    F122 *uint16 `json:"f122"`
    F123 *uint16 `json:"f123"`
    F124 *int32 `json:"f124"`
    F125 *int32 `json:"f125"`
    F126 *int32 `json:"f126"`
    F127 *uint32 `json:"f127"`
    F128 *uint32 `json:"f128"`
    F129 *uint32 `json:"f129"`
    F130 *int64 `json:"f130"`
    F131 *int64 `json:"f131"`
    F132 *int64 `json:"f132"`
    F133 *uint64 `json:"f133"`
    F134 *uint64 `json:"f134"`
    F135 *uint64 `json:"f135"`
    F136 *float32 `json:"f136"`
    F137 *float32 `json:"f137"`
    F138 *float32 `json:"f138"`
    F139 *float64 `json:"f139"`
    F140 *float64 `json:"f140"`
    F141 *float64 `json:"f141"`
    F142 *fbe.Decimal `json:"f142"`
    F143 *fbe.Decimal `json:"f143"`
    F144 *fbe.Decimal `json:"f144"`
    F145 *string `json:"f145"`
    F146 *string `json:"f146"`
    F147 *string `json:"f147"`
    F148 *fbe.Timestamp `json:"f148"`
    F149 *fbe.Timestamp `json:"f149"`
    F150 *fbe.Timestamp `json:"f150"`
    F151 *fbe.UUID `json:"f151"`
    F152 *fbe.UUID `json:"f152"`
    F153 *fbe.UUID `json:"f153"`
    F154 *proto.OrderSide `json:"f154"`
    F155 *proto.OrderSide `json:"f155"`
    F156 *proto.OrderType `json:"f156"`
    F157 *proto.OrderType `json:"f157"`
    F158 *proto.Order `json:"f158"`
    F159 *proto.Order `json:"f159"`
    F160 *proto.Balance `json:"f160"`
    F161 *proto.Balance `json:"f161"`
    F162 *proto.State `json:"f162"`
    F163 *proto.State `json:"f163"`
    F164 *proto.Account `json:"f164"`
    F165 *proto.Account `json:"f165"`
}

// Create a new StructOptional struct
func NewStructOptional() *StructOptional {
    return &StructOptional{
        StructSimple: NewStructSimple(),
        F100: nil,
        F101: fbe.OptionalBool(true),
        F102: nil,
        F103: nil,
        F104: fbe.OptionalByte(255),
        F105: nil,
        F106: nil,
        F107: fbe.OptionalRune('!'),
        F108: nil,
        F109: nil,
        F110: fbe.OptionalRune(rune(0x0444)),
        F111: nil,
        F112: nil,
        F113: fbe.OptionalInt8(127),
        F114: nil,
        F115: nil,
        F116: fbe.OptionalUInt8(255),
        F117: nil,
        F118: nil,
        F119: fbe.OptionalInt16(32767),
        F120: nil,
        F121: nil,
        F122: fbe.OptionalUInt16(65535),
        F123: nil,
        F124: nil,
        F125: fbe.OptionalInt32(2147483647),
        F126: nil,
        F127: nil,
        F128: fbe.OptionalUInt32(4294967295),
        F129: nil,
        F130: nil,
        F131: fbe.OptionalInt64(9223372036854775807),
        F132: nil,
        F133: nil,
        F134: fbe.OptionalUInt64(18446744073709551615),
        F135: nil,
        F136: nil,
        F137: fbe.OptionalFloat32(float32(123.456)),
        F138: nil,
        F139: nil,
        F140: fbe.OptionalFloat64(float64(-123.456e+123)),
        F141: nil,
        F142: nil,
        F143: fbe.OptionalDecimal(fbe.DecimalFromString("123456.123456")),
        F144: nil,
        F145: nil,
        F146: fbe.OptionalString("Initial string!"),
        F147: nil,
        F148: nil,
        F149: fbe.OptionalTimestamp(fbe.TimestampUTC()),
        F150: nil,
        F151: nil,
        F152: fbe.OptionalUUID(fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000")),
        F153: nil,
        F154: nil,
        F155: nil,
        F156: nil,
        F157: nil,
        F158: nil,
        F159: nil,
        F160: nil,
        F161: nil,
        F162: nil,
        F163: nil,
        F164: nil,
        F165: nil,
    }
}

// Create a new StructOptional struct from the given field values
func NewStructOptionalFromFieldValues(Parent *StructSimple, F100 *bool, F101 *bool, F102 *bool, F103 *byte, F104 *byte, F105 *byte, F106 *rune, F107 *rune, F108 *rune, F109 *rune, F110 *rune, F111 *rune, F112 *int8, F113 *int8, F114 *int8, F115 *uint8, F116 *uint8, F117 *uint8, F118 *int16, F119 *int16, F120 *int16, F121 *uint16, F122 *uint16, F123 *uint16, F124 *int32, F125 *int32, F126 *int32, F127 *uint32, F128 *uint32, F129 *uint32, F130 *int64, F131 *int64, F132 *int64, F133 *uint64, F134 *uint64, F135 *uint64, F136 *float32, F137 *float32, F138 *float32, F139 *float64, F140 *float64, F141 *float64, F142 *fbe.Decimal, F143 *fbe.Decimal, F144 *fbe.Decimal, F145 *string, F146 *string, F147 *string, F148 *fbe.Timestamp, F149 *fbe.Timestamp, F150 *fbe.Timestamp, F151 *fbe.UUID, F152 *fbe.UUID, F153 *fbe.UUID, F154 *proto.OrderSide, F155 *proto.OrderSide, F156 *proto.OrderType, F157 *proto.OrderType, F158 *proto.Order, F159 *proto.Order, F160 *proto.Balance, F161 *proto.Balance, F162 *proto.State, F163 *proto.State, F164 *proto.Account, F165 *proto.Account) *StructOptional {
    return &StructOptional{Parent, F100, F101, F102, F103, F104, F105, F106, F107, F108, F109, F110, F111, F112, F113, F114, F115, F116, F117, F118, F119, F120, F121, F122, F123, F124, F125, F126, F127, F128, F129, F130, F131, F132, F133, F134, F135, F136, F137, F138, F139, F140, F141, F142, F143, F144, F145, F146, F147, F148, F149, F150, F151, F152, F153, F154, F155, F156, F157, F158, F159, F160, F161, F162, F163, F164, F165}
}

// Create a new StructOptional struct from JSON
func NewStructOptionalFromJSON(buffer []byte) (*StructOptional, error) {
    result := *NewStructOptional()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *StructOptional) Copy() *StructOptional {
    var result = *s
    return &result
}

// Struct deep clone
func (s *StructOptional) Clone() *StructOptional {
    // Serialize the struct to the FBE stream
    writer := NewStructOptionalModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewStructOptionalModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *StructOptional) Key() StructOptionalKey {
    return StructOptionalKey{
        StructSimpleKey: s.StructSimple.Key(),
    }
}

// Convert struct to optional
func (s *StructOptional) Optional() *StructOptional {
    return s
}

// Convert struct to string
func (s *StructOptional) String() string {
    var sb strings.Builder
    sb.WriteString("StructOptional(")
    sb.WriteString(s.StructSimple.String())
    sb.WriteString(",f100=")
    if s.F100 != nil { 
        sb.WriteString(strconv.FormatBool(*s.F100))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f101=")
    if s.F101 != nil { 
        sb.WriteString(strconv.FormatBool(*s.F101))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f102=")
    if s.F102 != nil { 
        sb.WriteString(strconv.FormatBool(*s.F102))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f103=")
    if s.F103 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F103), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f104=")
    if s.F104 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F104), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f105=")
    if s.F105 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F105), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f106=")
    if s.F106 != nil { 
        sb.WriteString("'" + string(*s.F106) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f107=")
    if s.F107 != nil { 
        sb.WriteString("'" + string(*s.F107) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f108=")
    if s.F108 != nil { 
        sb.WriteString("'" + string(*s.F108) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f109=")
    if s.F109 != nil { 
        sb.WriteString("'" + string(*s.F109) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f110=")
    if s.F110 != nil { 
        sb.WriteString("'" + string(*s.F110) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f111=")
    if s.F111 != nil { 
        sb.WriteString("'" + string(*s.F111) + "'")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f112=")
    if s.F112 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F112), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f113=")
    if s.F113 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F113), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f114=")
    if s.F114 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F114), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f115=")
    if s.F115 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F115), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f116=")
    if s.F116 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F116), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f117=")
    if s.F117 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F117), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f118=")
    if s.F118 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F118), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f119=")
    if s.F119 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F119), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f120=")
    if s.F120 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F120), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f121=")
    if s.F121 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F121), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f122=")
    if s.F122 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F122), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f123=")
    if s.F123 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F123), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f124=")
    if s.F124 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F124), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f125=")
    if s.F125 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F125), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f126=")
    if s.F126 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F126), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f127=")
    if s.F127 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F127), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f128=")
    if s.F128 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F128), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f129=")
    if s.F129 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F129), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f130=")
    if s.F130 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F130), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f131=")
    if s.F131 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F131), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f132=")
    if s.F132 != nil { 
        sb.WriteString(strconv.FormatInt(int64(*s.F132), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f133=")
    if s.F133 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F133), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f134=")
    if s.F134 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F134), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f135=")
    if s.F135 != nil { 
        sb.WriteString(strconv.FormatUint(uint64(*s.F135), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f136=")
    if s.F136 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F136), 'g', -1, 32))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f137=")
    if s.F137 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F137), 'g', -1, 32))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f138=")
    if s.F138 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F138), 'g', -1, 32))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f139=")
    if s.F139 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F139), 'g', -1, 64))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f140=")
    if s.F140 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F140), 'g', -1, 64))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f141=")
    if s.F141 != nil { 
        sb.WriteString(strconv.FormatFloat(float64(*s.F141), 'g', -1, 64))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f142=")
    if s.F142 != nil { 
        sb.WriteString((*s.F142).String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f143=")
    if s.F143 != nil { 
        sb.WriteString((*s.F143).String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f144=")
    if s.F144 != nil { 
        sb.WriteString((*s.F144).String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f145=")
    if s.F145 != nil { 
        sb.WriteString("\"" + *s.F145 + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f146=")
    if s.F146 != nil { 
        sb.WriteString("\"" + *s.F146 + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f147=")
    if s.F147 != nil { 
        sb.WriteString("\"" + *s.F147 + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f148=")
    if s.F148 != nil { 
        sb.WriteString(strconv.FormatInt((*s.F148).UnixNano(), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f149=")
    if s.F149 != nil { 
        sb.WriteString(strconv.FormatInt((*s.F149).UnixNano(), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f150=")
    if s.F150 != nil { 
        sb.WriteString(strconv.FormatInt((*s.F150).UnixNano(), 10))
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f151=")
    if s.F151 != nil { 
        sb.WriteString("\"" + (*s.F151).String() + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f152=")
    if s.F152 != nil { 
        sb.WriteString("\"" + (*s.F152).String() + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f153=")
    if s.F153 != nil { 
        sb.WriteString("\"" + (*s.F153).String() + "\"")
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f154=")
    if s.F154 != nil { 
        sb.WriteString(s.F154.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f155=")
    if s.F155 != nil { 
        sb.WriteString(s.F155.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f156=")
    if s.F156 != nil { 
        sb.WriteString(s.F156.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f157=")
    if s.F157 != nil { 
        sb.WriteString(s.F157.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f158=")
    if s.F158 != nil { 
        sb.WriteString(s.F158.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f159=")
    if s.F159 != nil { 
        sb.WriteString(s.F159.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f160=")
    if s.F160 != nil { 
        sb.WriteString(s.F160.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f161=")
    if s.F161 != nil { 
        sb.WriteString(s.F161.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f162=")
    if s.F162 != nil { 
        sb.WriteString(s.F162.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f163=")
    if s.F163 != nil { 
        sb.WriteString(s.F163.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f164=")
    if s.F164 != nil { 
        sb.WriteString(s.F164.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(",f165=")
    if s.F165 != nil { 
        sb.WriteString(s.F165.String())
    } else {
        sb.WriteString("null")
    }
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *StructOptional) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}
