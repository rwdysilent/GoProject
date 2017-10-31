# coding: utf-8
"""
__author__ = 'wupengfei'
__time__ = '2017/10/27 16:40'
__function__ = ""
"""

from Crypto.Cipher import AES
import base64
import requests
import json

headers = {
    'Cookie': 'appver=1.5.0.75771;',
    'Referer': 'http://music.163.com/'
}

#first_param = "{rid:\"\", offset:\"0\", total:\"true\", limit:\"2\", csrf_token:\"\"}"
first_param = "exampleplaintext"
second_param = "010001"
first_key = forth_param = "0CoJUm6Qyw8W8jud"


def get_params():
    iv = "0102030405060708"
    #first_key = forth_param
    second_key = 16 * 'F'
    h_encText = AES_encrypt(first_param, first_key, iv)
    #h_encText = AES_encrypt(h_encText, second_key, iv)
    return h_encText


def AES_encrypt(text, key, iv):
    pad = 16 - len(text) % 16
    text = text + pad * chr(pad)
    encryptor = AES.new(key, AES.MODE_CBC, iv)
    encrypt_text = encryptor.encrypt(text)
    encrypt_text = base64.b64encode(encrypt_text)
    return encrypt_text


def get_params_test(value=first_param, key="0CoJUm6Qyw8W8jud", block_segments=False):
    # iv = Random.new().read(AES.block_size)
    iv = "0102030405060708"
    if block_segments:
        print(1)
        # See comment in decrypt for information.
        remainder = len(value) % 16
        # padded_value = value + '\0' * (16 - remainder) if remainder else value
        # cipher = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
        # value = cipher.encrypt(padded_value)[:len(value)]
        padded_value = value + (16 - remainder) * chr(16 - remainder)
        # pad = 16 - len(value) % 16
        # padded_value = value + pad * chr(pad)
        cipher = AES.new(key, AES.MODE_CBC, iv)
        value = cipher.encrypt(padded_value)
    else:
        value = AES.new(key, AES.MODE_CFB, iv).encrypt(value)
    # The returned value has its padding stripped to avoid query string issues.
    # return base64.b64encode(iv + value, '-_').rstrip('=')
    return base64.b64encode(value)


def get_encSecKey():
    encSecKey = "257348aecb5e556c066de214e531faadd1c55d814f9be95fd06d6bff9f4c7a41f831f6394d5a3fd2e3881736d94a02ca919d952872e7d0a50ebfa1769a7a62d512f5f1ca21aec60bc3819a9c3ffca5eca9a0dba6d6f7249b06f5965ecfff3695b54e1c28f3f624750ed39e7de08fc8493242e26dbc4484a01c76f739e135637c"
    return encSecKey


def get_json(url, params, encSecKey):
    data = {
        "params": params,
        "encSecKey": encSecKey
    }
    response = requests.post(url, headers=headers, data=data)
    return response.content


if __name__ == "__main__":
    url = "http://music.163.com/weapi/v1/resource/comments/R_SO_4_30953009/?csrf_token="
    params = get_params()
    print(params)

    params_test = get_params_test(value=first_param, key="0CoJUm6Qyw8W8jud", block_segments=True)
    print(params_test)

    #encSecKey = get_encSecKey()
    #json_text = get_json(url, params, encSecKey)
    #json_dict = json.loads(json_text)
    #print json_dict['total']
    #for item in json_dict['comments']:
    #    print item['content'].encode('utf-8', 'ignore')