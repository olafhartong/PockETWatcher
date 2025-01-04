# PockETWatcher

A tiny program to consume an ETW provider, for research purpose mostly.

Command line options:
  -complete
        Output complete event information
  -eventIds string
        Comma-separated list of EventIds to filter
  -output string
        File to write the trace output
  -provider string
        ETW provider name or GUID

You need to supply a provider name or GUID, the rest is optional.

By default the output is basic, the information I need most for my research.

An example event looks like this:

```
.\PockETWatcher.exe -provider Microsoft-Windows-TCPIP -eventIds 1033

PPPPPP                kk     EEEEEEE TTTTTTT WW      WW         tt           hh
PP   PP  oooo    cccc kk  kk EE        TTT   WW      WW   aa aa tt      cccc hh        eee  rr rr
PPPPPP  oo  oo cc     kkkkk  EEEEE     TTT   WW   W  WW  aa aaa tttt  cc     hhhhhh  ee   e rrr  r
PP      oo  oo cc     kk kk  EE        TTT    WW WWW WW aa  aaa tt    cc     hh   hh eeeee  rr
PP       oooo   ccccc kk  kk EEEEEEE   TTT     WW   WW   aaa aa  tttt  ccccc hh   hh  eeeee rr

2025/01/04 22:07:08 [i] PocketWatcher started with provider Microsoft-Windows-TCPIP
2025/01/04 22:07:08 [i] Started consuming events ... (Ctrl-c to end)
{
  "EventData": {
    "Compartment": "0",
    "LocalAddress": "10.211.55.21:53423",
    "LocalAddressLength": "16",
    "ProcessId": "1316",
    "ProcessStartKey": "3096224743817238",
    "RemoteAddress": "10.211.55.1:53",
    "RemoteAddressLength": "16",
    "Status": "STATUS_SUCCESS",
    "Tcb": "0xFFFFBB8E44B71AB0"
  },
  "EventID": 1033,
  "Keywords": 9223372054034646148,
  "Level": 4,
  "Opcode": 0,
  "Provider": "Microsoft-Windows-TCPIP",
  "TaskName": "TcpConnectTcbComplete",
  "Time": "2025-01-04T21:07:54.4941022Z"
}
{
  "EventData": {
    "Compartment": "0",
    "LocalAddress": "[::ffff:10.211.55.21]:53424",
    "LocalAddressLength": "28",
    "ProcessId": "1872",
    "ProcessStartKey": "3096224743818591",
    "RemoteAddress": "[::ffff:20.50.88.244]:443",
    "RemoteAddressLength": "28",
    "Status": "STATUS_SUCCESS",
    "Tcb": "0xFFFFBB8E46838530"
  },
  "EventID": 1033,
  "Keywords": 9223372054034646148,
  "Level": 4,
  "Opcode": 0,
  "Provider": "Microsoft-Windows-TCPIP",
  "TaskName": "TcpConnectTcbComplete",
  "Time": "2025-01-04T21:07:54.5255544Z"
}
2025/01/04 22:07:57 [~] Shutting down... Bye! 👋
```
