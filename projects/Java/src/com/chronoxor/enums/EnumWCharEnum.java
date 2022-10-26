//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.enums;

public enum EnumWCharEnum
{
    ENUM_VALUE_0((char)0 + 0)
    , ENUM_VALUE_1((char)0x0444 + 0)
    , ENUM_VALUE_2((char)0x0444 + 1)
    , ENUM_VALUE_3((char)0x0555 + 0)
    , ENUM_VALUE_4((char)0x0555 + 1)
    , ENUM_VALUE_5(ENUM_VALUE_3.getRaw())
    ;

    private int value;

    EnumWCharEnum(int value) { this.value = (int)value; }
    EnumWCharEnum(EnumWCharEnum value) { this.value = value.value; }

    public int getRaw() { return value; }

    public static EnumWCharEnum mapValue(int value) { return mapping.get(value); }

    @Override
    public String toString()
    {
        if (this == ENUM_VALUE_0) return "ENUM_VALUE_0";
        if (this == ENUM_VALUE_1) return "ENUM_VALUE_1";
        if (this == ENUM_VALUE_2) return "ENUM_VALUE_2";
        if (this == ENUM_VALUE_3) return "ENUM_VALUE_3";
        if (this == ENUM_VALUE_4) return "ENUM_VALUE_4";
        if (this == ENUM_VALUE_5) return "ENUM_VALUE_5";
        return "<unknown>";
    }

    private static final java.util.Map<Integer, EnumWCharEnum> mapping = new java.util.HashMap<>();
    static
    {
        for (var value : EnumWCharEnum.values())
            mapping.put(value.value, value);
    }
}
