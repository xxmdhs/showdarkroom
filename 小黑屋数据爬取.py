import requests
import time
print('输入起始cid')
cid = input()
print('输入结束cid')
end = input()
cid = int(cid)
end = int(end)
while cid > end:
    cids = str(cid)
    r = requests.get(
        'https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom&cid=' + cids + '&ajaxdata=json')
    f = open('./' + cids + '.json', 'w', encoding='utf-8')
    f.write(r.text)
    f.close()
    a = r.text
    b = a.find('"cid":"')
    b = b + 7
    c = a.find('"},"data"')
    cids = a[b:c]
    cid = int(cids)
    print(cid)
    time.sleep(3)