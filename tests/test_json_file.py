import pytest
from mlflow_log_cli.json_file import log_json_file


def test_log_json_file():
    log_json_file('tests/test.json')
