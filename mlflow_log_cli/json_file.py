import json
from typing import List

from loguru import logger
from mlflow.tracking import MlflowClient
from mlflow.tracking.fluent import _get_or_start_run


def load_json(f: str) -> dict:
    with open(f, 'r') as fp:
        return json.load(fp)


def log_params(client: MlflowClient, run_id: str, params: List[dict]) -> None:
    if params is None:
        return

    logger.info(f'log parameters to run ID: {run_id}')
    for param in params:
        client.log_param(run_id, **param)


def log_metrics(client: MlflowClient, run_id: str, metrics: List[dict]) -> None:
    if metrics is None:
        return

    logger.info(f'log metrics to run ID: {run_id}')
    for metric in metrics:
        client.log_metric(run_id, **metric)


def set_tags(client: MlflowClient, run_id: str, tags: List[dict]) -> None:
    if tags is None:
        return

    logger.info(f'set tags to run ID: {run_id}')
    for tag in tags:
        client.set_tag(run_id, **tag)


def log_json_file(f:str) -> None:
    data = load_json(f)
    logger.info('load json file: {}'.format(f))

    run_id = _get_or_start_run().info.run_id
    logger.info('run ID: {}'.format(run_id))

    client = MlflowClient()
    log_params(client, run_id, data.get('params'))
    log_metrics(client, run_id, data.get('metrics'))
    set_tags(client, run_id, data.get('tags'))
