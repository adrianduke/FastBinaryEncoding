//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.protoex;

public final class StateEx implements Comparable<StateEx>
{
    public static final StateEx unknown = new StateEx(StateExEnum.unknown);
    public static final StateEx invalid = new StateEx(StateExEnum.invalid);
    public static final StateEx initialized = new StateEx(StateExEnum.initialized);
    public static final StateEx calculated = new StateEx(StateExEnum.calculated);
    public static final StateEx broken = new StateEx(StateExEnum.broken);
    public static final StateEx happy = new StateEx(StateExEnum.happy);
    public static final StateEx sad = new StateEx(StateExEnum.sad);
    public static final StateEx good = new StateEx(StateExEnum.good);
    public static final StateEx bad = new StateEx(StateExEnum.bad);

    private StateExEnum value = StateExEnum.values()[0];
    private byte flags = value.getRaw();

    public StateEx() {}
    public StateEx(byte value) { setEnum(value); }
    public StateEx(StateExEnum value) { setEnum(value); }
    public StateEx(java.util.EnumSet<StateExEnum> value) { setEnum(value); }
    public StateEx(StateEx value) { setEnum(value); }

    public StateExEnum getEnum() { return value; }
    public byte getRaw() { return flags; }

    public void setDefault() { setEnum((byte)0); }

    public void setEnum(byte value) { this.flags = value; this.value = StateExEnum.mapValue(value); }
    public void setEnum(StateExEnum value) { this.value = value; this.flags = value.getRaw(); }
    public void setEnum(java.util.EnumSet<StateExEnum> value) { setEnum(StateEx.fromSet(value)); }
    public void setEnum(StateEx value) { this.value = value.value; this.flags = value.flags; }

    public boolean hasFlags(byte flags) { return (((this.flags & flags) != 0) && ((this.flags & flags) == flags)); }
    public boolean hasFlags(StateExEnum flags) { return hasFlags(flags.getRaw()); }
    public boolean hasFlags(StateEx flags) { return hasFlags(flags.flags); }

    public StateEx setFlags(byte flags) { setEnum((byte)(this.flags | flags)); return this; }
    public StateEx setFlags(StateExEnum flags) { setFlags(flags.getRaw()); return this; }
    public StateEx setFlags(StateEx flags) { setFlags(flags.flags); return this; }

    public StateEx removeFlags(byte flags) { setEnum((byte)(this.flags & ~flags)); return this; }
    public StateEx removeFlags(StateExEnum flags) { removeFlags(flags.getRaw()); return this; }
    public StateEx removeFlags(StateEx flags) { removeFlags(flags.flags); return this; }

    public java.util.EnumSet<StateExEnum> getAllSet() { return value.getAllSet(); }
    public java.util.EnumSet<StateExEnum> getNoneSet() { return value.getNoneSet(); }
    public java.util.EnumSet<StateExEnum> getCurrentSet() { return value.getCurrentSet(); }

    public static StateEx fromSet(java.util.EnumSet<StateExEnum> set)
    {
        byte result = 0;
        if (set.contains(unknown.getEnum()))
        {
            result |= unknown.flags;
        }
        if (set.contains(invalid.getEnum()))
        {
            result |= invalid.flags;
        }
        if (set.contains(initialized.getEnum()))
        {
            result |= initialized.flags;
        }
        if (set.contains(calculated.getEnum()))
        {
            result |= calculated.flags;
        }
        if (set.contains(broken.getEnum()))
        {
            result |= broken.flags;
        }
        if (set.contains(happy.getEnum()))
        {
            result |= happy.flags;
        }
        if (set.contains(sad.getEnum()))
        {
            result |= sad.flags;
        }
        if (set.contains(good.getEnum()))
        {
            result |= good.flags;
        }
        if (set.contains(bad.getEnum()))
        {
            result |= bad.flags;
        }
        return new StateEx(result);
    }

    @Override
    public int compareTo(StateEx other)
    {
        return (int)(flags - other.flags);
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!StateEx.class.isAssignableFrom(other.getClass()))
            return false;

        final StateEx flg = (StateEx)other;

        if (flags != flg.flags)
            return false;
        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        hash = hash * 31 + (int)flags;
        return hash;
    }

    @Override
    public String toString()
    {
        var sb = new StringBuilder();
        boolean first = true;
        if (hasFlags(unknown.flags))
        {
            sb.append(first ? "" : "|").append("unknown");
            first = false;
        }
        if (hasFlags(invalid.flags))
        {
            sb.append(first ? "" : "|").append("invalid");
            first = false;
        }
        if (hasFlags(initialized.flags))
        {
            sb.append(first ? "" : "|").append("initialized");
            first = false;
        }
        if (hasFlags(calculated.flags))
        {
            sb.append(first ? "" : "|").append("calculated");
            first = false;
        }
        if (hasFlags(broken.flags))
        {
            sb.append(first ? "" : "|").append("broken");
            first = false;
        }
        if (hasFlags(happy.flags))
        {
            sb.append(first ? "" : "|").append("happy");
            first = false;
        }
        if (hasFlags(sad.flags))
        {
            sb.append(first ? "" : "|").append("sad");
            first = false;
        }
        if (hasFlags(good.flags))
        {
            sb.append(first ? "" : "|").append("good");
            first = false;
        }
        if (hasFlags(bad.flags))
        {
            sb.append(first ? "" : "|").append("bad");
            first = false;
        }
        return sb.toString();
    }
}
