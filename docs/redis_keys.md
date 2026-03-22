| key | 说明 |
|-----|------|
| im:token:{uid} | Token，redis string，登录态 |
| im:unread:{uid}:{sessionid} | 未读计数 |
| im:friends:{uid} | 好友列表缓存 |
| im:sessions:{uid} | 最近活跃会话ZSet |
| im:history:{sessionid} | 最近N条消息list(优先缓存) |