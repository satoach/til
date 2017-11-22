# -*- coding: utf-8 -*-
import sys
import urllib.request
import json
import pandas as pd


def get_myissues(url):
    with urllib.request.urlopen(url+"/issues.json?assigned_to_id=me") as resp:
        return json.loads(resp.read().decode('utf-8'))

def issue2df(issue):
    return pd.DataFrame.from_dict(issue["issues"])

class DfConv:
    def __init__(self, key, fltr=""):
        self._key = key
        self._fltr = fltr

    def conv(self, d, *clmname):
        if type(d) == dict and self._key in d:
            if self._fltr == "" or "".join(clmname) == self._fltr:
                return d[self._key]
        return d

def main():
    if len(sys.argv) < 2:
        print("need url")
        return 1

    # 自分のチケットをDataFrame形式で保持
    issue = get_myissues(sys.argv[1])
    df_org = issue2df(issue)

    # 不要な列削除
    droplist = ["assigned_to", "author", "created_on", "estimated_hours",
                "done_ratio", "start_date", "updated_on", "description",
                "parent", "fixed_version"]
    df = df_org.drop(droplist, axis=1)

    # ステータスで絞り込む
    df[df.apply(lambda x: x.apply(DfConv("id", "status").conv, args=(x.name)))["status"] < 17]
    print(df.applymap(DfConv("name").conv).sort_values(by="status"))

    return 0


if __name__ == '__main__':
    main()
