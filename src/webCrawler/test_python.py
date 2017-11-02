# coding: utf-8
"""
__author__ = 'wupengfei'
__time__ = '2017/11/2 11:25'
__function__ = ""
"""
from Crypto.Cipher import AES

import md5


def pad(data):
    length = 16 - (len(data) % 16)
    data += chr(length)*length
    return data


def encrypt(plainText, workingKey):
    iv = '\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f'
    plainText = pad(plainText)
    encDigest = md5.new()
    encDigest.update(workingKey)
    enc_cipher = AES.new(encDigest.digest(), AES.MODE_CBC, iv)
    encryptedText = enc_cipher.encrypt(plainText).encode('hex')
    return encryptedText


if __name__ == "__main__":
    plainText = "it is a test string"
    workingKey = "0CoJUm6Qyw8W8jud"
    s = encrypt(plainText, workingKey)
    print(s)
