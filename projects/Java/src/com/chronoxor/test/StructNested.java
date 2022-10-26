//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.test;

public class StructNested extends StructOptional
{
    public EnumSimple f1000 = new EnumSimple();
    public EnumSimple f1001 = null;
    public EnumTyped f1002 = EnumTyped.ENUM_VALUE_2;
    public EnumTyped f1003 = null;
    public FlagsSimple f1004 = new FlagsSimple();
    public FlagsSimple f1005 = null;
    public FlagsTyped f1006 = FlagsTyped.fromSet(java.util.EnumSet.of(FlagsTyped.FLAG_VALUE_2.getEnum(), FlagsTyped.FLAG_VALUE_4.getEnum(), FlagsTyped.FLAG_VALUE_6.getEnum()));
    public FlagsTyped f1007 = null;
    public StructSimple f1008 = new StructSimple();
    public StructSimple f1009 = null;
    public StructOptional f1010 = new StructOptional();
    public StructOptional f1011 = null;

    public static final long fbeTypeConst = 112;
    public long fbeType() { return fbeTypeConst; }

    public StructNested() {}

    public StructNested(StructOptional parent, EnumSimple f1000, EnumSimple f1001, EnumTyped f1002, EnumTyped f1003, FlagsSimple f1004, FlagsSimple f1005, FlagsTyped f1006, FlagsTyped f1007, StructSimple f1008, StructSimple f1009, StructOptional f1010, StructOptional f1011)
    {
        super(parent);
        this.f1000 = f1000;
        this.f1001 = f1001;
        this.f1002 = f1002;
        this.f1003 = f1003;
        this.f1004 = f1004;
        this.f1005 = f1005;
        this.f1006 = f1006;
        this.f1007 = f1007;
        this.f1008 = f1008;
        this.f1009 = f1009;
        this.f1010 = f1010;
        this.f1011 = f1011;
    }

    public StructNested(StructNested other)
    {
        super(other);
        this.f1000 = other.f1000;
        this.f1001 = other.f1001;
        this.f1002 = other.f1002;
        this.f1003 = other.f1003;
        this.f1004 = other.f1004;
        this.f1005 = other.f1005;
        this.f1006 = other.f1006;
        this.f1007 = other.f1007;
        this.f1008 = other.f1008;
        this.f1009 = other.f1009;
        this.f1010 = other.f1010;
        this.f1011 = other.f1011;
    }

    public StructNested clone()
    {
        // Serialize the struct to the FBE stream
        var writer = new com.chronoxor.test.fbe.StructNestedModel();
        writer.serialize(this);

        // Deserialize the struct from the FBE stream
        var reader = new com.chronoxor.test.fbe.StructNestedModel();
        reader.attach(writer.getBuffer());
        return reader.deserialize();
    }

    @Override
    public int compareTo(Object other)
    {
        if (other == null)
            return -1;

        if (!StructNested.class.isAssignableFrom(other.getClass()))
            return -1;

        final StructNested obj = (StructNested)other;

        int result = 0;
        result = super.compareTo(obj);
        if (result != 0)
            return result;
        return result;
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!StructNested.class.isAssignableFrom(other.getClass()))
            return false;

        final StructNested obj = (StructNested)other;

        if (!super.equals(obj))
            return false;
        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        hash = hash * 31 + super.hashCode();
        return hash;
    }

    @Override
    public String toString()
    {
        var sb = new StringBuilder();
        sb.append("StructNested(");
        sb.append(super.toString());
        sb.append(",f1000="); sb.append(f1000);
        sb.append(",f1001="); if (f1001 != null) sb.append(f1001); else sb.append("null");
        sb.append(",f1002="); sb.append(f1002);
        sb.append(",f1003="); if (f1003 != null) sb.append(f1003); else sb.append("null");
        sb.append(",f1004="); sb.append(f1004);
        sb.append(",f1005="); if (f1005 != null) sb.append(f1005); else sb.append("null");
        sb.append(",f1006="); sb.append(f1006);
        sb.append(",f1007="); if (f1007 != null) sb.append(f1007); else sb.append("null");
        sb.append(",f1008="); sb.append(f1008);
        sb.append(",f1009="); if (f1009 != null) sb.append(f1009); else sb.append("null");
        sb.append(",f1010="); sb.append(f1010);
        sb.append(",f1011="); if (f1011 != null) sb.append(f1011); else sb.append("null");
        sb.append(")");
        return sb.toString();
    }

    public String toJson() { return com.chronoxor.test.fbe.Json.getEngine().toJson(this); }
    public static StructNested fromJson(String json) { return com.chronoxor.test.fbe.Json.getEngine().fromJson(json, StructNested.class); }
}
