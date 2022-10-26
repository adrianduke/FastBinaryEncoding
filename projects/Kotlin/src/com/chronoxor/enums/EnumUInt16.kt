//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
class EnumUInt16 : Comparable<EnumUInt16>
{
    companion object
    {
        val ENUM_VALUE_0 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_0)
        val ENUM_VALUE_1 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_1)
        val ENUM_VALUE_2 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_2)
        val ENUM_VALUE_3 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_3)
        val ENUM_VALUE_4 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_4)
        val ENUM_VALUE_5 = EnumUInt16(EnumUInt16Enum.ENUM_VALUE_5)
    }

    var value: EnumUInt16Enum? = EnumUInt16Enum.values()[0]
        private set

    val raw: UShort
        get() = value!!.raw

    constructor()
    constructor(value: UShort) { setEnum(value) }
    constructor(value: EnumUInt16Enum) { setEnum(value) }
    constructor(value: EnumUInt16) { setEnum(value) }

    fun setDefault() { setEnum(0.toUShort()) }

    fun setEnum(value: UShort) { this.value = EnumUInt16Enum.mapValue(value) }
    fun setEnum(value: EnumUInt16Enum) { this.value = value }
    fun setEnum(value: EnumUInt16) { this.value = value.value }

    override fun compareTo(other: EnumUInt16): Int
    {
        if (value == null)
            return -1
        if (other.value == null)
            return 1
        return (value!!.raw - other.value!!.raw).toInt()
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!EnumUInt16::class.java.isAssignableFrom(other.javaClass))
            return false

        val enm = other as EnumUInt16? ?: return false

        if (enm.value == null)
            return false
        if (value!!.raw != enm.value!!.raw)
            return false
        return true
    }

    override fun hashCode(): Int
    {
        var hash = 17
        hash = hash * 31 + if (value != null) value!!.hashCode() else 0
        return hash
    }

    override fun toString(): String
    {
        return if (value != null) value!!.toString() else "<unknown>"
    }
}
