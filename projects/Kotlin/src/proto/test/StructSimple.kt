// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.2.0.0

@file:Suppress("UnusedImport", "unused")

package test

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import proto.*

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
open class StructSimple : Comparable<Any?>
{
    var id: Int = 0
    var f1: Boolean = false
    var f2: Boolean = true
    var f3: Byte = 0.toByte()
    var f4: Byte = UByte.MAX_VALUE
    var f5: Char = '\u0000'
    var f6: Char = '!'.toChar()
    var f7: Char = '\u0000'
    var f8: Char = 0x0444.toChar()
    var f9: Byte = 0.toByte()
    var f10: Byte = Byte.MAX_VALUE
    var f11: UByte = UByte.MIN_VALUE
    var f12: UByte = UByte.MAX_VALUE
    var f13: Short = 0.toShort()
    var f14: Short = Short.MAX_VALUE
    var f15: UShort = UShort.MIN_VALUE
    var f16: UShort = UShort.MAX_VALUE
    var f17: Int = 0
    var f18: Int = Int.MAX_VALUE
    var f19: UInt = UInt.MIN_VALUE
    var f20: UInt = UInt.MAX_VALUE
    var f21: Long = 0L
    var f22: Long = Long.MAX_VALUE
    var f23: ULong = ULong.MIN_VALUE
    var f24: ULong = ULong.MAX_VALUE
    var f25: Float = 0.0f
    var f26: Float = 123.456f
    var f27: Double = 0.0
    var f28: Double = -123.456e+123
    var f29: BigDecimal = BigDecimal.valueOf(0L)
    var f30: BigDecimal = BigDecimal.valueOf(123456.123456)
    var f31: String = ""
    var f32: String = "Initial string!"
    var f33: Instant = Instant.EPOCH
    var f34: Instant = Instant.EPOCH
    var f35: Instant = Instant.now()
    var f36: UUID = fbe.UUIDGenerator.nil()
    var f37: UUID = fbe.UUIDGenerator.sequential()
    var f38: UUID = UUID.fromString("123e4567-e89b-12d3-a456-426655440000")
    var f39: proto.OrderSide = proto.OrderSide()
    var f40: proto.OrderType = proto.OrderType()
    var f41: proto.Order = proto.Order()
    var f42: proto.Balance = proto.Balance()
    var f43: proto.State = proto.State()
    var f44: proto.Account = proto.Account()

    constructor()

    constructor(id: Int, f1: Boolean, f2: Boolean, f3: Byte, f4: Byte, f5: Char, f6: Char, f7: Char, f8: Char, f9: Byte, f10: Byte, f11: UByte, f12: UByte, f13: Short, f14: Short, f15: UShort, f16: UShort, f17: Int, f18: Int, f19: UInt, f20: UInt, f21: Long, f22: Long, f23: ULong, f24: ULong, f25: Float, f26: Float, f27: Double, f28: Double, f29: BigDecimal, f30: BigDecimal, f31: String, f32: String, f33: Instant, f34: Instant, f35: Instant, f36: UUID, f37: UUID, f38: UUID, f39: proto.OrderSide, f40: proto.OrderType, f41: proto.Order, f42: proto.Balance, f43: proto.State, f44: proto.Account)
    {
        this.id = id
        this.f1 = f1
        this.f2 = f2
        this.f3 = f3
        this.f4 = f4
        this.f5 = f5
        this.f6 = f6
        this.f7 = f7
        this.f8 = f8
        this.f9 = f9
        this.f10 = f10
        this.f11 = f11
        this.f12 = f12
        this.f13 = f13
        this.f14 = f14
        this.f15 = f15
        this.f16 = f16
        this.f17 = f17
        this.f18 = f18
        this.f19 = f19
        this.f20 = f20
        this.f21 = f21
        this.f22 = f22
        this.f23 = f23
        this.f24 = f24
        this.f25 = f25
        this.f26 = f26
        this.f27 = f27
        this.f28 = f28
        this.f29 = f29
        this.f30 = f30
        this.f31 = f31
        this.f32 = f32
        this.f33 = f33
        this.f34 = f34
        this.f35 = f35
        this.f36 = f36
        this.f37 = f37
        this.f38 = f38
        this.f39 = f39
        this.f40 = f40
        this.f41 = f41
        this.f42 = f42
        this.f43 = f43
        this.f44 = f44
    }

    @Suppress("UNUSED_PARAMETER")
    constructor(other: StructSimple)
    {
        this.id = other.id
        this.f1 = other.f1
        this.f2 = other.f2
        this.f3 = other.f3
        this.f4 = other.f4
        this.f5 = other.f5
        this.f6 = other.f6
        this.f7 = other.f7
        this.f8 = other.f8
        this.f9 = other.f9
        this.f10 = other.f10
        this.f11 = other.f11
        this.f12 = other.f12
        this.f13 = other.f13
        this.f14 = other.f14
        this.f15 = other.f15
        this.f16 = other.f16
        this.f17 = other.f17
        this.f18 = other.f18
        this.f19 = other.f19
        this.f20 = other.f20
        this.f21 = other.f21
        this.f22 = other.f22
        this.f23 = other.f23
        this.f24 = other.f24
        this.f25 = other.f25
        this.f26 = other.f26
        this.f27 = other.f27
        this.f28 = other.f28
        this.f29 = other.f29
        this.f30 = other.f30
        this.f31 = other.f31
        this.f32 = other.f32
        this.f33 = other.f33
        this.f34 = other.f34
        this.f35 = other.f35
        this.f36 = other.f36
        this.f37 = other.f37
        this.f38 = other.f38
        this.f39 = other.f39
        this.f40 = other.f40
        this.f41 = other.f41
        this.f42 = other.f42
        this.f43 = other.f43
        this.f44 = other.f44
    }

    open fun clone(): StructSimple
    {
        // Serialize the struct to the FBE stream
        val writer = test.fbe.StructSimpleModel()
        writer.serialize(this)

        // Deserialize the struct from the FBE stream
        val reader = test.fbe.StructSimpleModel()
        reader.attach(writer.buffer)
        return reader.deserialize()
    }

    override fun compareTo(other: Any?): Int
    {
        if (other == null)
            return -1

        if (!StructSimple::class.java.isAssignableFrom(other.javaClass))
            return -1

        @Suppress("UNUSED_VARIABLE")
        val obj = other as StructSimple? ?: return -1

        @Suppress("VARIABLE_WITH_REDUNDANT_INITIALIZER", "CanBeVal", "UnnecessaryVariable")
        var result = 0
        result = id.compareTo(obj.id)
        if (result != 0)
            return result
        return result
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!StructSimple::class.java.isAssignableFrom(other.javaClass))
            return false

        @Suppress("UNUSED_VARIABLE")
        val obj = other as StructSimple? ?: return false

        if (id != obj.id)
            return false
        return true
    }

    override fun hashCode(): Int
    {
        @Suppress("CanBeVal", "UnnecessaryVariable")
        var hash = 17
        hash = hash * 31 + id.hashCode()
        return hash
    }

    override fun toString(): String
    {
        val sb = StringBuilder()
        sb.append("StructSimple(")
        sb.append("id="); sb.append(id)
        sb.append(",f1="); sb.append(if (f1) "true" else "false")
        sb.append(",f2="); sb.append(if (f2) "true" else "false")
        sb.append(",f3="); sb.append(f3)
        sb.append(",f4="); sb.append(f4)
        sb.append(",f5="); sb.append("'").append(f5).append("'")
        sb.append(",f6="); sb.append("'").append(f6).append("'")
        sb.append(",f7="); sb.append("'").append(f7).append("'")
        sb.append(",f8="); sb.append("'").append(f8).append("'")
        sb.append(",f9="); sb.append(f9)
        sb.append(",f10="); sb.append(f10)
        sb.append(",f11="); sb.append(f11)
        sb.append(",f12="); sb.append(f12)
        sb.append(",f13="); sb.append(f13)
        sb.append(",f14="); sb.append(f14)
        sb.append(",f15="); sb.append(f15)
        sb.append(",f16="); sb.append(f16)
        sb.append(",f17="); sb.append(f17)
        sb.append(",f18="); sb.append(f18)
        sb.append(",f19="); sb.append(f19)
        sb.append(",f20="); sb.append(f20)
        sb.append(",f21="); sb.append(f21)
        sb.append(",f22="); sb.append(f22)
        sb.append(",f23="); sb.append(f23)
        sb.append(",f24="); sb.append(f24)
        sb.append(",f25="); sb.append(f25)
        sb.append(",f26="); sb.append(f26)
        sb.append(",f27="); sb.append(f27)
        sb.append(",f28="); sb.append(f28)
        sb.append(",f29="); sb.append(f29)
        sb.append(",f30="); sb.append(f30)
        sb.append(",f31="); sb.append("\"").append(f31).append("\"")
        sb.append(",f32="); sb.append("\"").append(f32).append("\"")
        sb.append(",f33="); sb.append(f33.epochSecond * 1000000000 + f33.nano)
        sb.append(",f34="); sb.append(f34.epochSecond * 1000000000 + f34.nano)
        sb.append(",f35="); sb.append(f35.epochSecond * 1000000000 + f35.nano)
        sb.append(",f36="); sb.append("\"").append(f36).append("\"")
        sb.append(",f37="); sb.append("\"").append(f37).append("\"")
        sb.append(",f38="); sb.append("\"").append(f38).append("\"")
        sb.append(",f39="); sb.append(f39)
        sb.append(",f40="); sb.append(f40)
        sb.append(",f41="); sb.append(f41)
        sb.append(",f42="); sb.append(f42)
        sb.append(",f43="); sb.append(f43)
        sb.append(",f44="); sb.append(f44)
        sb.append(")")
        return sb.toString()
    }

    open fun toJson(): String = test.fbe.Json.engine.toJson(this)
    companion object
    {
        fun fromJson(json: String): StructSimple = test.fbe.Json.engine.fromJson(json, StructSimple::class.java)
    }
}
