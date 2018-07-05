
markdown = """
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects.
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions.
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes.
    * [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy.
    * [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language.
    * [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing.
    * [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces.
    * [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter.    
    * [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests.
"""

import requests
import re
import time

re_url = re.compile(r'\]\(https?://([^)]+)')
re_last_link = re.compile(r'<([^?]+\?page=)(\d+)>; rel="last"')


def list_repos(markdown):
    res = re_url.findall(markdown)
    for url in res:
        if url.endswith('/'):
            url = url[:-1]
        if url.startswith('github.com'):
            yield url[len('github.com/'):]


def extract_last_link(link_header):
    links = link_header.split(',')
    res = re_last_link.findall(link_header)
    assert len(res) == 1, f'len(res) is {len(res)}'
    return(res[0])


extract_last_link('<https://api.github.com/repositories/96570421/stargazers?page=2>; rel="next", <https://api.github.com/repositories/96570421/stargazers?page=20>; rel="last"')

up = ('ramalho', '?')

def count_stars(repo):
    url = f'https://api.github.com/repos/{repo}/stargazers'
    r = requests.get(url, auth=up)
    assert r.status_code == 200, f'{url} status_code is {r.status_code}'
    links = r.headers.get('Link')
    count = len(r.json())
    if links:
        url, pages = extract_last_link(r.headers['Link'])
        url = url + pages
        r = requests.get(url, auth=up)
        assert r.status_code == 200, f'{url} status_code is {r.status_code}'
        count = count * (int(pages) - 1) + len(r.json())
    return count


for repo in list_repos(markdown):
    print(repo, end=' ')
    stars = count_stars(repo)
    print(stars)

"""
cavaliercoder/badio 7
h2non/baloo 502
fulldump/biff 2
marioidival/bro 20
bradleyjkemp/cupaloy 49
khaiql/dbcleaner 37
viant/dsunit 17
viant/endly 35
verdverm/frisby 228
msoap/go-carpet 179
zimmski/go-mutesting 146
dnaeon/go-vcr 206
franela/goblin 512
smartystreets/goconvey 3685
corbym/gocrest 6
DATA-DOG/godog 404
appleboy/gofight 170
corbym/gogiven 6
orfjackal/gospec 109
stesla/gospecify 49
pavlo/gosuite 5
rdrdr/hamcrest 23
gavv/httpexpect 884
yookoala/restit 45
go-testfixtures/testfixtures 180
stretchr/testify 5251
vcaesar/tt 1
posener/wstest 37
go-check/check 387
onsi/gomega 709

"""
