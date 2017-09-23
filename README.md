# dosukoi

## Description
Command to read and parse HTTP-Request packets


## Option
- -i  Interface to get packets from (default: "en0")
- -p  Read from pcap-file, overrides -i (default: "")
- -f  BPF filter for pcap (default: "tcp")
- -l  Log file path (default: "")
- -a  Display only information including Basic-Auth (default: false)
- -s  Display information in short format (default: false)
- -S  SnapLen for pcap packet capture (default: 1600)


## Examples
### Standard output
```
$ ./dosukoi -i en0
2017/09/24 06:01:57 Starting capture on interface "en0"
2017/09/24 06:01:57 reading in packets
-----------------------------------------------------
Date:         2017/09/24 06:02:00
Method:       GET
URL:          http://example.com/
Basic-Auth:   user:pass
Referer:      http://example.com
UserAgent:    curl/7.54.0
Source:       192.168.0.1:63990 (ab:cd:12:34:ef:gh)
Destination:  93.184.26.34:80 (hg:fg:11:23:ab:cd)
-----------------------------------------------------
Date:         2017/09/24 06:02:25
Method:       GET
URL:          http://h0ge.net/
UserAgent:    curl/7.54.0
Source:       192.168.0.1:63992 (fa:ke:ma:ca:dd:rs)
Destination:  104.198.14.52:80 (aa:ii:uu:ee:oo:kk)

```

### Log
```
{"Date":"2017/09/24 06:02:00","HTTP":{"URL":"http://example.com/","Method":"GET","Version":"HTTP/1.1","BasicAuth":"user:pass","UseProxy":false,"ProxyAuth":"","Referer":"http://example.com","UserAgent":"curl/7.54.0"},"SrcPort":"63990","DstPort":"80","SrcIP":"192.168.0.1","DstIP":"93.184.216.34","SrcMAC":"ab:cd:12:34:ef:gh","DstMAC":"hg:fg:11:23:ab:cd"}
{"Date":"2017/09/24 06:02:25","HTTP":{"URL":"http://h0ge.net/","Method":"GET","Version":"HTTP/1.1","BasicAuth":"","UseProxy":false,"ProxyAuth":"","Referer":"","UserAgent":"curl/7.54.0"},"SrcPort":"63992","DstPort":"80","SrcIP":"192.168.0.1","DstIP":"104.198.14.52","SrcMAC":"ka:ke:ma:ca:dd:rs","DstMAC":"aa:ii:uu:ee:oo:kk"}

```


## Notes
- HTTPリクエストパケットのデフラグは行っていないため、第一パケットから漏れた情報は捨てられます
- ログファイルは1行につき1リクエストずつJson形式で書き込まれます


## Licence
MIT
