from pathlib import Path

import click
from loguru import logger

from .json_file import log_json_file


@click.command()
@click.option('-f','--file', type=click.STRING)
def cli(file):
    f = Path(file)
    if f.with_suffix('.json'):
        log_json_file(f)
    else:
        logger.info('file not supported')
