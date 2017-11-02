# coding: utf-8
"""
__author__ = 'wupengfei'
__time__ = '2017/11/2 11:25'
__function__ = ""
"""

from Crypto.Cipher import AES
import base64


def get_params():
    iv = "0102030405060708"
    # first_key = forth_param
    second_key = 16 * 'F'
    h_encText = AES_encrypt(first_param, first_key, iv)
    h_encText = AES_encrypt(h_encText, second_key, iv)
    return h_encText


def AES_encrypt(text, key, iv):
    pad = 16 - len(text) % 16
    text = text + pad * chr(pad)
    print("length text: %s" % len(text))
    encryptor = AES.new(key, AES.MODE_CBC, iv)
    encrypt_text = encryptor.encrypt(text)
    encrypt_text = base64.b64encode(encrypt_text)
    return encrypt_text


if __name__ == '__main__':
    first_param = "{rid:\"\", offset:\"0\", total:\"true\", limit:\"2\", csrf_token:\"\"}"
    second_param = "010001"
    first_key = "0CoJUm6Qyw8W8jud"
    my_pass = get_params()
    print(my_pass)
