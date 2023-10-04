There is no special interface for middleware. 

Solidity multiple inheritance is used instead.



Run of example design via `forge test -vvvv --match-test test_example_appstack` gives:

```
[⠰] Compiling...
No files changed, compilation skipped

Running 1 test for test/Middleware.t.sol:MiddlewareTest
[PASS] test_app() (gas: 645363)
Logs:
  App: sendTransfer called
  FeeMiddlewarePacketSender: send packet/hello
  HookMiddlewarePacketSender: send packet/hello_hooked
  Base: send packet via IBC Handler/hello_hooked
  =================
  HookMiddlewarePacketReceiver: recv packet/hello_hooked
  FeeMiddlewarePacketReceiver: recv packet/hello_hooked
  App: recv packet/hello_hooked
  FeeMiddlewarePacketReceiver: after recv packet
  HookMiddlewarePacketReceiver: after recv packet

Traces:
  [645363] MiddlewareTest::test_app()
    ├─ [592610] → new HookFeeMiddlewaredApp@0x5615dEB798BB3E4dFa0139dFa1b3D433Cc23b72f
    │   └─ ← 2844 bytes of code
    ├─ [8422] HookFeeMiddlewaredApp::sendTransfer(hello)
    │   ├─ [0] console::log(App: sendTransfer called) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(FeeMiddlewarePacketSender: send packet/hello) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(HookMiddlewarePacketSender: send packet/hello_hooked) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(Base: send packet via IBC Handler/hello_hooked) [staticcall]
    │   │   └─ ← ()
    │   └─ ← ()
    ├─ [0] console::log(=================) [staticcall]
    │   └─ ← ()
    ├─ [9338] HookFeeMiddlewaredApp::onRecvPacket((0x68656c6c6f5f686f6f6b6564))
    │   ├─ [0] console::log(HookMiddlewarePacketReceiver: recv packet/hello_hooked) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(FeeMiddlewarePacketReceiver: recv packet/hello_hooked) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(App: recv packet/hello_hooked) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(FeeMiddlewarePacketReceiver: after recv packet) [staticcall]
    │   │   └─ ← ()
    │   ├─ [0] console::log(HookMiddlewarePacketReceiver: after recv packet) [staticcall]
    │   │   └─ ← ()
    │   └─ ← 1
    └─ ← ()

Test result: ok. 1 passed; 0 failed; 0 skipped; finished in 1.13ms

Ran 1 test suites: 1 tests passed, 0 failed, 0 skipped (1 total tests)
```