//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

import ChronoxorFbe
import Foundation
import ChronoxorProto

// Fast Binary Encoding ChronoxorTest final sender
open class FinalSender: ChronoxorFbe.SenderProtocol {
    // Imported senders
    let ProtoSender: ChronoxorProto.FinalSender

    // Sender models accessors

    public var buffer: Buffer = Buffer()
    public var final: Bool = false

    public init() {
        ProtoSender = ChronoxorProto.FinalSender(buffer: buffer)
        build(with: true)
    }

    public init(buffer: ChronoxorFbe.Buffer) {
        ProtoSender = ChronoxorProto.FinalSender(buffer: buffer)
        build(with: buffer, final: true)
    }

    public func send(obj: Any) throws -> Int {
        guard let listener = self as? ChronoxorFbe.SenderListener else { return 0 }
        return try send(obj: obj, listener: listener)
    }

    public func send(obj: Any, listener: ChronoxorFbe.SenderListener) throws -> Int {

        // Try to send using imported senders
        var result: Int = 0
        result = try ProtoSender.send(obj: obj, listener: listener)
        if result > 0 {
            return result
            }

        return 0
    }
}
