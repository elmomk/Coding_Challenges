#!/usr/bin/env python3


import sys
import typer
from typing import Optional
from typing_extensions import Annotated

def count_bytes(name: str):
    file = open(name, "r")
    bytes_size = len(file.read().encode("utf-8"))
    file.close()
    return bytes_size

def count_lines(name: str):
    file = open(name, "r")
    lines = len(file.readlines())
    file.close()
    return lines

def count_words(name: str):
    file = open(name, "r")
    words = len(file.read().split())
    file.close()
    return words


# flag --bytes/-c to count bytes
# flag --lines/-l to count count_lines
# flag --words/-w to count words
# iterate through arguments
def main(names: Annotated[Optional[list[str]], typer.Argument()] = None,
         bytes: bool = typer.Option(False, "--bytes", "-c"),
         lines: bool = typer.Option(False, "--lines", "-l"),
         words: bool = typer.Option(False, "--words", "-w")):

    # check if stdin is being piped
    if not sys.stdin.isatty():
        if bytes:
            result = len(sys.stdin.read().encode("utf-8"))
            typer.echo(f"{result}")
        elif lines:
            result = len(sys.stdin.readlines())
            typer.echo(f"{result}")
        elif words:
            result = len(sys.stdin.read().split())
            typer.echo(f"{result}")
        else:
            result_lines = len(sys.stdin.readlines())
            result_words = len(sys.stdin.read().split())
            result_bytes = len(sys.stdin.read().encode("utf-8"))
            typer.echo(f"{result_lines} {result_words} {result_bytes}")

    else:
        for name in names:
            if bytes:
                result = count_bytes(name)
                typer.echo(f"{result} {name}")
            elif lines:
                result = count_lines(name)
                typer.echo(f"{result} {name}")
            elif words:
                result = count_words(name)
                typer.echo(f"{result} {name}")
            else:
                result_bytes = count_bytes(name)
                result_lines = count_lines(name)
                result_words = count_words(name)
                typer.echo(f"{result_lines} {result_words} {result_bytes} {name}")



if __name__ == "__main__":
    typer.run(main)
