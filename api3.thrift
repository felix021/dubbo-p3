namespace go echo3

include "java.thrift"

struct EchoRequest {
    1: required string message,
    2: optional java.Object obj,
}(JavaClassName="kitex.echo3.EchoRequest")

struct EchoResponse {
    1: required string message,
}(JavaClassName="kitex.echo3.EchoResponse")

service TestService3 {
    EchoResponse Echo(1: EchoRequest req)
} (JavaClassName="kitex.echo3.TestService3")
