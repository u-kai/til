# 各種イベントに劇団員を派遣して演劇のパフォーマンスを行う会社を考える
# 顧客がいくつかの劇を選択し，席数や演じた劇の種類に応じて請求する
# この会社は演劇に関するデータを単純なJSONファイルで保持している



"""plays.json
{
    "hamlet": {"name": "Hamlet", "type": "tragedy"},
    "as-like": {"name": "As You Like It", "type": "comedy"},
    "othello": {"name": "Othello", "type": "tragedy"}
}
"""
"""invoice.json
[
    {
        "customer": "BigCo",
        "performances": [
            {
                "playID": "hamlet",
                "audience": 55
            },
            {
                "playID": "as-like",
                "audience": 35
            },
            {
                "playID": "othello",
                "audience": 40
            }
    }
]
"""


# 請求書を作成するコード


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f"Statement for {invoice['customer']}\n"
    format = "  {:<30} {:>10} {:>10}\n"

    for perf in invoice["performances"]:
        play = plays[perf["playID"]]
        this_amount = 0

        if play["type"] == "tragedy":
            this_amount = 40000
            if perf["audience"] > 30:
                this_amount += 1000 * (perf["audience"] - 30)
                break
        elif play["type"] == "comedy":
            this_amount = 30000
            if perf["audience"] > 20:
                this_amount += 10000 + 500 * (perf["audience"] - 20)
            this_amount += 300 * perf["audience"]
        else:
            Exception(f"unknown type: {play['type']}")
        volume_credits += max(perf["audience"] - 30, 0)
        if "comedy" == play["type"]:
            volume_credits += perf["audience"] // 5
        result += format.format(play["name"], this_amount // 100, perf["audience"])
        total_amount += this_amount

    result += f"Amount owed is {total_amount // 100}\n"
    result += f"You earned {volume_credits} credits\n"
    return result


