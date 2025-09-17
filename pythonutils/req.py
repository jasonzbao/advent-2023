import os
import requests


def make_request(day: int) -> str:
    session_cookie = os.getenv("AOC_SESSION")
    if not session_cookie:
        raise ValueError("AOC_SESSION environment variable must be set")

    url = f"https://adventofcode.com/2023/day/{day}/input"

    cookies = {"session": session_cookie}
    response = requests.get(url, cookies=cookies)
    response.raise_for_status()

    return response.text


def format_input(input_str: str) -> list[str]:
    return input_str.rstrip('\n').split('\n')


def formatted_request(day: int) -> list[str]:
    return format_input(make_request(day))
