//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.protoex.fbe;;

public final class OrderTypeJson implements com.google.gson.JsonSerializer<com.chronoxor.protoex.OrderType>, com.google.gson.JsonDeserializer<com.chronoxor.protoex.OrderType>
{
    @Override
    public com.google.gson.JsonElement serialize(com.chronoxor.protoex.OrderType src, java.lang.reflect.Type typeOfSrc, com.google.gson.JsonSerializationContext context)
    {
        return new com.google.gson.JsonPrimitive(src.getRaw());
    }

    @Override
    public com.chronoxor.protoex.OrderType deserialize(com.google.gson.JsonElement json, java.lang.reflect.Type type, com.google.gson.JsonDeserializationContext context) throws com.google.gson.JsonParseException
    {
        return new com.chronoxor.protoex.OrderType(json.getAsJsonPrimitive().getAsByte());
    }
}
