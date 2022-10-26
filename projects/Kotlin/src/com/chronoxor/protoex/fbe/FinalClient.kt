//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex.fbe

// Fast Binary Encoding com.chronoxor.protoex final client
@Suppress("MemberVisibilityCanBePrivate", "PropertyName")
open class FinalClient : com.chronoxor.fbe.Client, IFinalClientListener
{
    // Imported clients
    val protoClient: com.chronoxor.proto.fbe.FinalClient

    // Client sender models accessors
    val OrderMessageSenderModel: OrderMessageFinalModel
    val BalanceMessageSenderModel: BalanceMessageFinalModel
    val AccountMessageSenderModel: AccountMessageFinalModel

    // Client receiver values accessors
    private val OrderMessageReceiverValue: com.chronoxor.protoex.OrderMessage
    private val BalanceMessageReceiverValue: com.chronoxor.protoex.BalanceMessage
    private val AccountMessageReceiverValue: com.chronoxor.protoex.AccountMessage

    // Client receiver models accessors
    private val OrderMessageReceiverModel: OrderMessageFinalModel
    private val BalanceMessageReceiverModel: BalanceMessageFinalModel
    private val AccountMessageReceiverModel: AccountMessageFinalModel

    constructor() : super(true)
    {
        protoClient = com.chronoxor.proto.fbe.FinalClient(sendBuffer, receiveBuffer)
        OrderMessageSenderModel = OrderMessageFinalModel(sendBuffer)
        OrderMessageReceiverValue = com.chronoxor.protoex.OrderMessage()
        OrderMessageReceiverModel = OrderMessageFinalModel()
        BalanceMessageSenderModel = BalanceMessageFinalModel(sendBuffer)
        BalanceMessageReceiverValue = com.chronoxor.protoex.BalanceMessage()
        BalanceMessageReceiverModel = BalanceMessageFinalModel()
        AccountMessageSenderModel = AccountMessageFinalModel(sendBuffer)
        AccountMessageReceiverValue = com.chronoxor.protoex.AccountMessage()
        AccountMessageReceiverModel = AccountMessageFinalModel()
    }

    constructor(sendBuffer: com.chronoxor.fbe.Buffer, receiveBuffer: com.chronoxor.fbe.Buffer) : super(sendBuffer, receiveBuffer, true)
    {
        protoClient = com.chronoxor.proto.fbe.FinalClient(sendBuffer, receiveBuffer)
        OrderMessageSenderModel = OrderMessageFinalModel(sendBuffer)
        OrderMessageReceiverValue = com.chronoxor.protoex.OrderMessage()
        OrderMessageReceiverModel = OrderMessageFinalModel()
        BalanceMessageSenderModel = BalanceMessageFinalModel(sendBuffer)
        BalanceMessageReceiverValue = com.chronoxor.protoex.BalanceMessage()
        BalanceMessageReceiverModel = BalanceMessageFinalModel()
        AccountMessageSenderModel = AccountMessageFinalModel(sendBuffer)
        AccountMessageReceiverValue = com.chronoxor.protoex.AccountMessage()
        AccountMessageReceiverModel = AccountMessageFinalModel()
    }

    fun send(obj: Any): Long
    {
        return sendListener(this, obj)
    }

    @Suppress("JoinDeclarationAndAssignment", "UNUSED_PARAMETER")
    fun sendListener(listener: IFinalClientListener, obj: Any): Long
    {
        when (obj)
        {
            is com.chronoxor.protoex.OrderMessage -> if (obj.fbeType == OrderMessageSenderModel.fbeType) return sendListener(listener, obj)
            is com.chronoxor.protoex.BalanceMessage -> if (obj.fbeType == BalanceMessageSenderModel.fbeType) return sendListener(listener, obj)
            is com.chronoxor.protoex.AccountMessage -> if (obj.fbeType == AccountMessageSenderModel.fbeType) return sendListener(listener, obj)
        }

        // Try to send using imported clients
        @Suppress("CanBeVal")
        var result: Long
        result = protoClient.sendListener(listener, obj)
        if (result > 0)
            return result

        return 0
    }

    fun send(value: com.chronoxor.protoex.OrderMessage): Long
    {
        return sendListener(this, value)
    }

    fun sendListener(listener: IFinalClientListener, value: com.chronoxor.protoex.OrderMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = OrderMessageSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.protoex.OrderMessage serialization failed!" }
        assert(OrderMessageSenderModel.verify()) { "com.chronoxor.protoex.OrderMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            listener.onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(listener, serialized)
    }
    fun send(value: com.chronoxor.protoex.BalanceMessage): Long
    {
        return sendListener(this, value)
    }

    fun sendListener(listener: IFinalClientListener, value: com.chronoxor.protoex.BalanceMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = BalanceMessageSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.protoex.BalanceMessage serialization failed!" }
        assert(BalanceMessageSenderModel.verify()) { "com.chronoxor.protoex.BalanceMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            listener.onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(listener, serialized)
    }
    fun send(value: com.chronoxor.protoex.AccountMessage): Long
    {
        return sendListener(this, value)
    }

    fun sendListener(listener: IFinalClientListener, value: com.chronoxor.protoex.AccountMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = AccountMessageSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.protoex.AccountMessage serialization failed!" }
        assert(AccountMessageSenderModel.verify()) { "com.chronoxor.protoex.AccountMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            listener.onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(listener, serialized)
    }

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: IFinalClientListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            com.chronoxor.protoex.fbe.OrderMessageFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                OrderMessageReceiverModel.attach(buffer, offset)
                assert(OrderMessageReceiverModel.verify()) { "com.chronoxor.protoex.OrderMessage validation failed!" }
                val deserialized = OrderMessageReceiverModel.deserialize(OrderMessageReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.OrderMessage deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = OrderMessageReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(OrderMessageReceiverValue)
                return true
            }
            com.chronoxor.protoex.fbe.BalanceMessageFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                BalanceMessageReceiverModel.attach(buffer, offset)
                assert(BalanceMessageReceiverModel.verify()) { "com.chronoxor.protoex.BalanceMessage validation failed!" }
                val deserialized = BalanceMessageReceiverModel.deserialize(BalanceMessageReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.BalanceMessage deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = BalanceMessageReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(BalanceMessageReceiverValue)
                return true
            }
            com.chronoxor.protoex.fbe.AccountMessageFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                AccountMessageReceiverModel.attach(buffer, offset)
                assert(AccountMessageReceiverModel.verify()) { "com.chronoxor.protoex.AccountMessage validation failed!" }
                val deserialized = AccountMessageReceiverModel.deserialize(AccountMessageReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.AccountMessage deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = AccountMessageReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(AccountMessageReceiverValue)
                return true
            }
        }

        if (protoClient.onReceiveListener(listener, type, buffer, offset, size))
            return true

        return false
    }
}
