'''
JJMail的链接工具
用于发送信息到JJMail服务

任务介绍：
task1 发送博客订阅
task2 发送自动回复
task3 发送验证码
task4 发送mgek订阅
'''

# 引入sys path
# 否则会找不到库
import sys
import json

jjmail_root = "/home/apps"
sys.path.append(jjmail_root)
sys.path.append(jjmail_root + "/jjmail")

from jjmail.task1 import send_blog_mail
from jjmail.task2 import send_reply_mail
from jjmail.task3 import send_code_mail
from jjmail.task4 import send_mgek_mail


def lib_blog_mail(mail_address, posts):
    res = json.loads(posts, encoding="utf-8")
    send_blog_mail.apply_async([mail_address, res],
                               countdown=60,
                               retry=False)


def lib_reply_mail(mail_address):
    send_reply_mail.apply_async([mail_address],
                                countdown=30,
                                retry=False)


def lib_mgek_mail(mail_address):
    send_mgek_mail.apply_async([mail_address],
                               countdown=30,
                               retry=False)


if __name__ == '__main__':
    args = sys.argv
    length = len(args)
    if length <= 2:
        print("args not enough")
    else:
        # 默认参数格式 address service
        address = args[1]
        service = args[2]
        if service == "blog":
            posts = args[3]
            lib_blog_mail(address, posts)
        elif service == "mgek":
            lib_mgek_mail(address)
        elif service == "reply":
            lib_reply_mail(address)
        else:
            print("wrong service")
