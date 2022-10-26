//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

import Foundation
import ChronoxorFbe

// Fast Binary Encoding ChronoxorProto final receiver
open class FinalReceiver: ChronoxorFbe.ReceiverProtocol {
    // Receiver values accessors
    private var OrderMessageValue: ChronoxorProto.OrderMessage
    private var BalanceMessageValue: ChronoxorProto.BalanceMessage
    private var AccountMessageValue: ChronoxorProto.AccountMessage

    // Receiver models accessors
    private var OrderMessageModel: OrderMessageFinalModel
    private var BalanceMessageModel: BalanceMessageFinalModel
    private var AccountMessageModel: AccountMessageFinalModel

    public var buffer: Buffer = Buffer()
    public var final: Bool = false

    public init() {
        OrderMessageValue = ChronoxorProto.OrderMessage()
        OrderMessageModel = ChronoxorProto.OrderMessageFinalModel()
        BalanceMessageValue = ChronoxorProto.BalanceMessage()
        BalanceMessageModel = ChronoxorProto.BalanceMessageFinalModel()
        AccountMessageValue = ChronoxorProto.AccountMessage()
        AccountMessageModel = ChronoxorProto.AccountMessageFinalModel()
        build(final: true)
    }

    public init(buffer: ChronoxorFbe.Buffer) {
        OrderMessageValue = ChronoxorProto.OrderMessage()
        OrderMessageModel = ChronoxorProto.OrderMessageFinalModel()
        BalanceMessageValue = ChronoxorProto.BalanceMessage()
        BalanceMessageModel = ChronoxorProto.BalanceMessageFinalModel()
        AccountMessageValue = ChronoxorProto.AccountMessage()
        AccountMessageModel = ChronoxorProto.AccountMessageFinalModel()
        build(with: buffer, final: true)
    }

    open func onReceive(type: Int, buffer: Data, offset: Int, size: Int) -> Bool {
        guard let listener = self as? FinalReceiverListener else { return false }
        return onReceiveListener(listener: listener, type: type, buffer: buffer, offset: offset, size: size)
    }

    open func onReceiveListener(listener: FinalReceiverListener, type: Int, buffer: Data, offset: Int, size: Int) -> Bool {
        switch type {
        case ChronoxorProto.OrderMessageFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            OrderMessageModel.attach(buffer: buffer, offset: offset)
            assert(OrderMessageModel.verify(), "Proto.OrderMessage validation failed!")
            let deserialized = OrderMessageModel.deserialize(value: &OrderMessageValue)
            assert(deserialized > 0, "Proto.OrderMessage deserialization failed!")

            // Log the value
            if listener.logging {
                let message = OrderMessageValue.description
                listener.onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: OrderMessageValue)
            return true
        case ChronoxorProto.BalanceMessageFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            BalanceMessageModel.attach(buffer: buffer, offset: offset)
            assert(BalanceMessageModel.verify(), "Proto.BalanceMessage validation failed!")
            let deserialized = BalanceMessageModel.deserialize(value: &BalanceMessageValue)
            assert(deserialized > 0, "Proto.BalanceMessage deserialization failed!")

            // Log the value
            if listener.logging {
                let message = BalanceMessageValue.description
                listener.onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: BalanceMessageValue)
            return true
        case ChronoxorProto.AccountMessageFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            AccountMessageModel.attach(buffer: buffer, offset: offset)
            assert(AccountMessageModel.verify(), "Proto.AccountMessage validation failed!")
            let deserialized = AccountMessageModel.deserialize(value: &AccountMessageValue)
            assert(deserialized > 0, "Proto.AccountMessage deserialization failed!")

            // Log the value
            if listener.logging {
                let message = AccountMessageValue.description
                listener.onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: AccountMessageValue)
            return true
        default: break
        }

        return false
    }
}
