# LightNovelNotify
## 概要
Slackに設定日発売のラノベを通知するだけのプログラム。情報元は
[ここ](https://calendar.gameiroiro.com/litenovel.php)
。

## 設定

### 環境変数

```sh
export SLACK_TOKEN=xoxb-***
export SLACK_CHANNEL=C******
```


### 設定ファイル

```json
[
    {
        "Delta": 0, // 当日からの何日後のデータか
        "Hour": 6, // 何時に通知するか
        "Minute": 31 // 何分に通知するか
    }
]
```
