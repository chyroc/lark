import json
import os
import requests
import base64
import datetime
import libnacl.sealed
import libnacl.public
import libnacl.secret

username = 'chyroc'
repo = 'lark'


def libnacl_encrypt(key, msg):
    if isinstance(msg, str):
        msg = msg.encode('utf-8')
    key = base64.standard_b64decode(key)
    keypair = libnacl.public.PublicKey(key)
    box = libnacl.sealed.SealedBox(keypair)
    encrypt_res = box.encrypt(msg)
    res = base64.standard_b64encode(encrypt_res)
    if isinstance(res, bytes):
        res = res.decode('utf-8')
    return res


def get_public_key(token):
    url = f'https://api.github.com/repos/{username}/{repo}/actions/secrets/public-key'
    r = requests.get(url, headers={
        'Accept': 'application/vnd.github.v3+json',
        'Authorization': 'token ' + token
    }, verify=False)
    r.raise_for_status()
    return r.json()['key_id'], r.json()['key']


def set_secret(token, name, val):
    url = f'https://api.github.com/repos/{username}/{repo}/actions/secrets/{name}'
    key_id, key = get_public_key(token=token)
    encrypted_value = libnacl_encrypt(key, val)
    r = requests.put(url, json={
        'key_id': key_id,
        'encrypted_value': encrypted_value,
    }, headers={
        'Accept': 'application/vnd.github.v3+json',
        'Authorization': 'token ' + token
    })
    r.raise_for_status()
    print(f'set secret:{name} success')


def get_plain_text(s):
    if not s:
        return s
    res = []
    for i in s:
        res.append(i)
        res.append('-')
    return ''.join(res)


def read_token_from_pre_file():
    with open('./.github/internal/token.json', 'r') as f:
        data = json.load(f)

    return data.get('access_token'), data.get('refresh_token')


if __name__ == '__main__':
    github_token = os.getenv('INTERNAL_GH_TOKEN') or ''
    key_internal_secret_just_for_fun = 'INTERNAL_SECRETS_JUST_FOR_FUN'
    key_access_token = 'LARK_ACCESS_TOKEN_ALL_PERMISSION_APP'
    key_refresh_token = 'LARK_REFRESH_TOKEN_ALL_PERMISSION_APP'
    internal_secret_just_for_fun_val = os.getenv(key_internal_secret_just_for_fun)

    print("hi, just-for-fun env:", get_plain_text(internal_secret_just_for_fun_val))
    access_token, refresh_token = read_token_from_pre_file()

    set_secret(token=github_token, name=key_internal_secret_just_for_fun, val=str(datetime.datetime.now().timestamp()))
    if access_token:
        set_secret(token=github_token, name=key_access_token, val=str(access_token))
    if refresh_token:
        set_secret(token=github_token, name=key_refresh_token, val=str(refresh_token))
