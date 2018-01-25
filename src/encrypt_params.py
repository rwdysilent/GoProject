# coding: utf-8
"""
__author__ = 'wupengfei'
__time__ = '2017/11/2 11:25'
__function__ = ""
"""

from Crypto.Cipher import AES
import base64


def get_params():
    plaintext = "{rid:\"\", offset:\"0\", total:\"true\", limit:\"2\", csrf_token:\"\"}"
    encrypt_key = "0CoJUm6Qyw8W8jud"
    iv = "0102030405060708"
    second_key = 16 * 'F'
    h_encText = AES_encrypt(plaintext, encrypt_key, iv)
    h_encText = AES_encrypt(h_encText, second_key, iv)
    return h_encText


def AES_encrypt(text, key, iv):
    pad = 16 - len(text) % 16
    text = text + pad * chr(pad)
    encryptor = AES.new(key, AES.MODE_CBC, iv)
    encrypt_text = encryptor.encrypt(text)
    encrypt_text = base64.b64encode(encrypt_text)
    return encrypt_text


if __name__ == '__main__':
    my_pass = get_params()
    print(my_pass)
