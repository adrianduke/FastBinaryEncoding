//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.enums.fbe;;

public final class EnumInt8Json implements com.google.gson.JsonSerializer<com.chronoxor.enums.EnumInt8>, com.google.gson.JsonDeserializer<com.chronoxor.enums.EnumInt8>
{
    @Override
    public com.google.gson.JsonElement serialize(com.chronoxor.enums.EnumInt8 src, java.lang.reflect.Type typeOfSrc, com.google.gson.JsonSerializationContext context)
    {
        return new com.google.gson.JsonPrimitive(src.getRaw());
    }

    @Override
    public com.chronoxor.enums.EnumInt8 deserialize(com.google.gson.JsonElement json, java.lang.reflect.Type type, com.google.gson.JsonDeserializationContext context) throws com.google.gson.JsonParseException
    {
        return new com.chronoxor.enums.EnumInt8(json.getAsJsonPrimitive().getAsByte());
    }
}
