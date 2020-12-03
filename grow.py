import requests
import os


base_url = "http://127.0.0.1:9093"

def run():
    cmd = "./grow"
    print(os.popen(cmd))

def add_keyword(word):
    return requests.get(f"{base_url}/add?keyword={word}").text

def build():
    return requests.get(f"{base_url}/build").text

def extract(text):
    return requests.get(f"{base_url}/extract?text={text}").text

if __name__ == "__main__":
    run()
    add_keyword("基金")
    add_keyword("阿里巴巴")
    build()
    rst = extract("阿里巴巴asdfasd阿里巴巴I love Big Apple and Bay Area.基金基金")
    print(rst)