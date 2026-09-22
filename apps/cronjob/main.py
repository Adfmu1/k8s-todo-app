import requests

wikiURL = 'https://en.wikipedia.org/wiki/Special:Random'

headers={
    'sec-ch-ua': '"Chromium";v="152", "Not?A_Brand";v="24", "Brave";v="152"',
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/152.0.0.0 Safari/537.36'
}

getReq = requests.get(url=wikiURL, headers=headers)

text = f'Read: {getReq.url}'

response = requests.post('http://todo-backend-svc:8990/todos', json={"newTodo": text})
response.raise_for_status()