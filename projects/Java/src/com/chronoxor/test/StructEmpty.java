//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.test;

public class StructEmpty implements Comparable<Object>
{
    public static final long fbeTypeConst = 143;
    public long fbeType() { return fbeTypeConst; }

    public StructEmpty() {}

    public StructEmpty(StructEmpty other)
    {
    }

    public StructEmpty clone()
    {
        // Serialize the struct to the FBE stream
        var writer = new com.chronoxor.test.fbe.StructEmptyModel();
        writer.serialize(this);

        // Deserialize the struct from the FBE stream
        var reader = new com.chronoxor.test.fbe.StructEmptyModel();
        reader.attach(writer.getBuffer());
        return reader.deserialize();
    }

    @Override
    public int compareTo(Object other)
    {
        if (other == null)
            return -1;

        if (!StructEmpty.class.isAssignableFrom(other.getClass()))
            return -1;

        final StructEmpty obj = (StructEmpty)other;

        int result = 0;
        return result;
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!StructEmpty.class.isAssignableFrom(other.getClass()))
            return false;

        final StructEmpty obj = (StructEmpty)other;

        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        return hash;
    }

    @Override
    public String toString()
    {
        var sb = new StringBuilder();
        sb.append("StructEmpty(");
        sb.append(")");
        return sb.toString();
    }

    public String toJson() { return com.chronoxor.test.fbe.Json.getEngine().toJson(this); }
    public static StructEmpty fromJson(String json) { return com.chronoxor.test.fbe.Json.getEngine().fromJson(json, StructEmpty.class); }
}
